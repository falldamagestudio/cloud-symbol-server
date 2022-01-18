using System;
using System.IO;
using System.Linq;
using System.Text;

namespace ClientAPI
{
    public class PDBParser
    {
        public class PagedStream : System.IO.Stream
        {
        	public Stream OriginalStream;

        	public uint[] PageMap;
            public uint PageSize;

            public override long Position { get; set; }
 
            private static long RoundDown(long value, long granularity) {
                return value / granularity * granularity;
            }

            public PagedStream(Stream originalStream, uint[] pageMap, uint pageSize) {
                OriginalStream = originalStream;
                PageMap = pageMap;
                PageSize = pageSize;
            }

            public override bool CanRead => true;
            public override bool CanWrite => false;
            public override bool CanSeek => true;

            public override long Length { get { return OriginalStream.Length; }}

            private int ReadPage(byte[] buffer, long pageStartPosition, long pageEndPosition, int writeOffset) {
                uint pageIndexVirtual = (uint)(pageStartPosition / (long)PageSize);
                uint pageIndexPhysical = PageMap[pageIndexVirtual];
                long pageOffset = pageStartPosition % (long)PageSize;
                long readStart = (long)pageIndexPhysical * (long)PageSize + pageOffset;
                int readLength = (int)(pageEndPosition - pageStartPosition);

                OriginalStream.Seek(readStart, SeekOrigin.Begin);

                return OriginalStream.Read(buffer, writeOffset, readLength);
            }

            public override int Read(byte[] buffer, int offset, int count) {

                long startPosition = Position;
                long endPosition = Position + count;

                int bytesRead = 0;

                long currentPosition = startPosition;
                int currentWriteOffset = offset;

                while (currentPosition != endPosition) {
                    long pageStartPosition = currentPosition;
                    long pageEndPosition = Math.Min(RoundDown(currentPosition + PageSize, PageSize), endPosition);
                    int pageBytesToRead = (int)(pageEndPosition - pageStartPosition);

                    int pageBytesRead = ReadPage(buffer, pageStartPosition, pageEndPosition, currentWriteOffset);

                    bytesRead += pageBytesRead;
                    currentPosition = pageEndPosition;
                    currentWriteOffset += pageBytesToRead;
                }

                Position = endPosition;

                return bytesRead;
            }

            public override long Seek(long offset, SeekOrigin origin) {
                switch (origin) {
                    case SeekOrigin.Begin:
                        Position = offset;
                        break;
                    case SeekOrigin.Current:
                        Position += offset;
                        break;
                    case SeekOrigin.End:
                        Position = OriginalStream.Length - offset;
                        break;
                }

                return Position;
            }

            public override void Flush() { }

            public override void SetLength(long value) {
                throw new NotSupportedException();
            }

            public override void Write(byte[] buffer, int offset, int count)
            {
                throw new NotSupportedException();
            }
        }

        public class MSF7Parser
        {
            private static readonly byte[] msf7Signature = Encoding.ASCII.GetBytes("Microsoft C/C++ MSF 7.00\r\n\x001ADS\x0000\x0000\x0000");

            public static bool IsValid(Stream stream) {
                stream.Seek(0, SeekOrigin.Begin);
                byte[] buffer = new byte[msf7Signature.Length];
                try {
                    if (stream.Read(buffer, 0, buffer.Length) != buffer.Length)
                        return false;

                    if (!buffer.SequenceEqual(msf7Signature))
                        return false;
                } catch (Exception) {
                    return false;
                }

                return true;
            }

            public class SuperBlock {
                // byte{32]FileSignature
                public uint BlockSize;
                public uint FreeBlockMapBlock;
                public uint NumBlocks;
                public uint NumDirectoryBytes;
                //uint Unknown;
                public uint BlockMapAddr;
            }

            public static SuperBlock ReadSuperBlock(Stream stream)
            {
                stream.Seek(msf7Signature.Length, SeekOrigin.Begin);
                BinaryReader binaryReader = new BinaryReader(stream);

                SuperBlock superBlock = new SuperBlock();
                superBlock.BlockSize = binaryReader.ReadUInt32();
                superBlock.FreeBlockMapBlock = binaryReader.ReadUInt32();
                superBlock.NumBlocks = binaryReader.ReadUInt32();
                superBlock.NumDirectoryBytes = binaryReader.ReadUInt32();
                binaryReader.ReadUInt32(); // Skip unknown value
                superBlock.BlockMapAddr = binaryReader.ReadUInt32();
                return superBlock;
            }

            public static byte[] ReadBlock(Stream stream, uint blockIndex, uint blockSize)
            {
                byte[] block = new byte[blockSize];
                stream.Seek((long)blockIndex * (long)blockSize, SeekOrigin.Begin);
                stream.Read(block, 0, (int)blockSize);
                return block;
            }

            public static uint[] ReadPageMap(Stream stream, uint blockMapAddr, uint blockSize, uint numDirectoryBytes)
            {
                byte[] block = ReadBlock(stream, blockMapAddr, blockSize);
            	uint numBlockMapEntries = (numDirectoryBytes + blockSize - 1) / blockSize;

                BinaryReader binaryReader = new BinaryReader(new MemoryStream(block));

                uint[] blockMap = new uint[numBlockMapEntries];
                for (uint blockMapIndex = 0; blockMapIndex < numBlockMapEntries; blockMapIndex++)
                    blockMap[blockMapIndex] = binaryReader.ReadUInt32();

                return blockMap;
            }


            public static uint[] ReadStreamDirectoryBlocks(Stream stream, uint[] blockMap, uint blockSize, uint numDirectoryBytes)
            {

                uint numStreamDirectoryBlocks = (numDirectoryBytes + blockSize - 1) / blockSize;
                uint[] streamDirectoryBlocks = new uint[numStreamDirectoryBlocks];

                uint streamDirectoryEntriesPerBlock = blockSize / 4;

                for (uint blockMapBlockId = 0; blockMapBlockId < blockMap.Length; blockMapBlockId++)
                {
                    uint blockMapBlockValue = blockMap[blockMapBlockId];
                    byte[] block = ReadBlock(stream, blockMapBlockValue, blockSize);

                    BinaryReader binaryReader = new BinaryReader(new MemoryStream(block));

                    for (uint streamDirectoryEntryWithinBlockId = 0; streamDirectoryEntryWithinBlockId < streamDirectoryEntriesPerBlock; streamDirectoryEntryWithinBlockId++)
                        streamDirectoryBlocks[blockMapBlockId * streamDirectoryEntriesPerBlock + streamDirectoryEntryWithinBlockId] = binaryReader.ReadUInt32();
                }

                return streamDirectoryBlocks;
            }

            public class StreamDirectoryHeader {
                public uint NumStreams;
                public uint[] StreamSizes;
                public uint[] StreamPageMapOffsets;
                public uint[] StreamPageMapCounts;
            }

            public static uint[] ReadStreamPageMap(PagedStream pagedStreamDirectory, StreamDirectoryHeader streamDirectoryHeader, int streamId)
            {
                if (streamId < 0 || streamId >= (int)streamDirectoryHeader.NumStreams) {
                    throw new ApplicationException($"StreamId out of bounds: id {streamId}, but there are only {streamDirectoryHeader.NumStreams} streams");
                }

                uint[] pageMap = new uint[streamDirectoryHeader.StreamPageMapCounts[streamId]];
                pagedStreamDirectory.Seek((long)(streamDirectoryHeader.StreamPageMapOffsets[streamId]), 0);
                BinaryReader binaryReader = new BinaryReader(pagedStreamDirectory);
                for (uint pageMapEntryIndex = 0; pageMapEntryIndex < pageMap.Length; pageMapEntryIndex++)
                    pageMap[pageMapEntryIndex] = binaryReader.ReadUInt32();

                return pageMap;
            }

            public static StreamDirectoryHeader ReadStreamDirectoryHeader(PagedStream pagedStreamDirectory)
            {
                StreamDirectoryHeader streamDirectoryHeader = new StreamDirectoryHeader();

                BinaryReader binaryReader = new BinaryReader(pagedStreamDirectory);
                streamDirectoryHeader.NumStreams = binaryReader.ReadUInt32();

                streamDirectoryHeader.StreamSizes = new uint[streamDirectoryHeader.NumStreams];
                streamDirectoryHeader.StreamPageMapOffsets = new uint[streamDirectoryHeader.NumStreams];
                streamDirectoryHeader.StreamPageMapCounts = new uint[streamDirectoryHeader.NumStreams];

                for (uint streamIndex = 0; streamIndex < streamDirectoryHeader.StreamSizes.Length; streamIndex++)
                    streamDirectoryHeader.StreamSizes[streamIndex] = binaryReader.ReadUInt32();

                uint offset = (1 + streamDirectoryHeader.NumStreams) * 4;

                for (uint streamIndex = 0; streamIndex < streamDirectoryHeader.StreamSizes.Length; streamIndex++)
                {
                    uint streamSize = streamDirectoryHeader.StreamSizes[streamIndex];

                    streamDirectoryHeader.StreamPageMapOffsets[streamIndex] = offset;
                    uint numPages = (streamSize + pagedStreamDirectory.PageSize - 1) / pagedStreamDirectory.PageSize;
                    streamDirectoryHeader.StreamPageMapCounts[streamIndex] = numPages;
                    offset += numPages * 4;
                }

                return streamDirectoryHeader;
            }

            public class PdbStreamHeader
            {
                public uint Version;
                public uint Signature;
                public uint Age;
                public Guid UniqueId;
            }

            public const int PdbStreamIndex = 1;

            public static PdbStreamHeader ReadPdbStreamHeader(PagedStream pagedStreamDirectory, StreamDirectoryHeader streamDirectoryHeader)
            {
                uint[] pageMap = ReadStreamPageMap(pagedStreamDirectory, streamDirectoryHeader, PdbStreamIndex);

                PagedStream pagedPdbStream = new PagedStream(
                    pagedStreamDirectory.OriginalStream,
                    pageMap,
                    pagedStreamDirectory.PageSize);

                PdbStreamHeader pdbStreamHeader = new PdbStreamHeader();

                BinaryReader binaryReader = new BinaryReader(pagedPdbStream);
                pdbStreamHeader.Version = binaryReader.ReadUInt32();
                pdbStreamHeader.Signature = binaryReader.ReadUInt32();
                pdbStreamHeader.Age = binaryReader.ReadUInt32();
                byte[] guidBytes = new byte[16];
                pagedPdbStream.Read(guidBytes, 0, guidBytes.Length);
                pdbStreamHeader.UniqueId = new Guid(guidBytes);

                return pdbStreamHeader;
            }

            public static string GetHash(Stream stream)
            {
                SuperBlock superBlock = ReadSuperBlock(stream);

                uint[] streamDirectoryPageMap = ReadPageMap(stream, superBlock.BlockMapAddr, superBlock.BlockSize, superBlock.NumDirectoryBytes);

                PagedStream pagedStreamDirectory = new PagedStream(
                    stream,
                    streamDirectoryPageMap,
                    superBlock.BlockSize
                );

                StreamDirectoryHeader streamDirectoryHeader = ReadStreamDirectoryHeader(pagedStreamDirectory);

                PdbStreamHeader pdbStreamHeader = ReadPdbStreamHeader(pagedStreamDirectory, streamDirectoryHeader);

                string hash = string.Format("{0}{1}", pdbStreamHeader.UniqueId.ToString("N").ToUpper(), pdbStreamHeader.Age);

                return hash;
            }

        }

        public static string GetHash(string pdbPath)
        {
            using (Stream stream = new FileStream(pdbPath, FileMode.Open))
            {
                if (MSF7Parser.IsValid(stream)) {
                    return MSF7Parser.GetHash(stream);
                } else {
                    return null;
                }
            }
        }

    }
}