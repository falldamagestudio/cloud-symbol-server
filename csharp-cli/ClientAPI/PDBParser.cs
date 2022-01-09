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

        }
    }
}