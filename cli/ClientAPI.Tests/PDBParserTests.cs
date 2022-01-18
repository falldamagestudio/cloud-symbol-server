using System.IO;
using Xunit;
using ClientAPI;
using System.Linq;
using System.Threading.Tasks;
using System.Text;

namespace ClientAPI.Tests
{

    public class PDBParserTests
    {
        private PDBParser.PagedStream GetPagedStream()
        {
            uint[] pageMap = new uint[] {1, 0, 2};
            byte[] content = new byte[] { 4, 5, 6, 7, 0, 1, 2, 3, 8, 9, 10, 11 };
            uint pageSize = 4;

            MemoryStream memoryStream = new MemoryStream(content);

            PDBParser.PagedStream pagedStream = new PDBParser.PagedStream(memoryStream, pageMap, pageSize);
            return pagedStream;
        }

        [Fact]
        public void PagedStreamLengthMatchesOriginalStreamLength() {
            PDBParser.PagedStream pagedStream = GetPagedStream();
            Assert.Equal(pagedStream.OriginalStream.Length, pagedStream.Length);
        }

        [Fact]
        public void PagedStreamAllowsReads() {
            PDBParser.PagedStream pagedStream = GetPagedStream();
            Assert.True(pagedStream.CanRead);
            byte[] result = new byte[1];
            pagedStream.Read(result, 0, 1);
            Assert.Equal(0, result[0]);
        }

        [Fact]
        public void PagedStreamSupportsSeekFromBeginning() {
            PDBParser.PagedStream pagedStream = GetPagedStream();
            Assert.True(pagedStream.CanSeek);
            long newPosition = pagedStream.Seek(7, SeekOrigin.Begin);
            Assert.Equal(7, newPosition);
            Assert.Equal(7, pagedStream.Position);
        }

        [Fact]
        public void PagedStreamSupportsSeekFromCurrent() {
            PDBParser.PagedStream pagedStream = GetPagedStream();
            Assert.True(pagedStream.CanSeek);
            long newPosition = pagedStream.Seek(7, SeekOrigin.Begin);
            Assert.Equal(7, newPosition);
            Assert.Equal(7, pagedStream.Position);
            long newPosition2 = pagedStream.Seek(-2, SeekOrigin.Current);
            Assert.Equal(5, newPosition2);
            Assert.Equal(5, pagedStream.Position);
        }

        [Fact]
        public void PagedStreamSupportsSeekFromEnd() {
            PDBParser.PagedStream pagedStream = GetPagedStream();
            Assert.True(pagedStream.CanSeek);
            long newPosition = pagedStream.Seek(2, SeekOrigin.End);
            Assert.Equal(10, newPosition);
            Assert.Equal(10, pagedStream.Position);
        }

        [Fact]
        public async Task PagedStreamDoesNotSupportWrites() {
            PDBParser.PagedStream pagedStream = GetPagedStream();
            Assert.False(pagedStream.CanWrite);
            byte[] result = new byte[1] { 0 };
            var exception = await Record.ExceptionAsync(() => Task.Run(() => pagedStream.Write(result, 0, 1)));
            Assert.NotNull(exception);
        }

        [Fact]
        public async Task PagedStreamDoesNotSupportSetLength() {
            PDBParser.PagedStream pagedStream = GetPagedStream();
            byte[] result = new byte[1] { 0 };
            var exception = await Record.ExceptionAsync(() => Task.Run(() => pagedStream.SetLength(10)));
            Assert.NotNull(exception);
        }

        [Fact]
        public void PagedStreamRemapWorks()
        {
            PDBParser.PagedStream pagedStream = GetPagedStream();

            // Read a section that covers three pages; part of the first, all of the second, and part of the third

            byte[] data = new byte[9];

            pagedStream.Seek(1, SeekOrigin.Begin);
            int bytesRead = pagedStream.Read(data, 0, data.Length);

            Assert.Equal(data.Length, bytesRead);

            byte[] expectedData = new byte[] {1, 2, 3, 4, 5, 6, 7, 8, 9};

            Assert.True(expectedData.SequenceEqual(data), $"Read data expected to be {{{string.Join(", ", expectedData)}}} but was {{{string.Join(", ", data)}}}");
        }

        [Fact]
        public void IsMSF7ValidTestSucceedsForPDB()
        {
            using (FileStream fileStream = new FileStream("../../../../testdata/example.pdb", FileMode.Open))
            {
                Assert.True(PDBParser.MSF7Parser.IsValid(fileStream));
            }
        }

       [Fact]
        public void IsMSF7ValidTestFailsForEmptyFile()
        {
            MemoryStream memoryStream = new MemoryStream(new byte[] { });
            Assert.False(PDBParser.MSF7Parser.IsValid(memoryStream));
        }

       [Fact]
        public void IsMSF7ValidTestFailsForNonPDBFile()
        {
            MemoryStream memoryStream = new MemoryStream(Encoding.ASCII.GetBytes(
                "abcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefgh"));
            Assert.False(PDBParser.MSF7Parser.IsValid(memoryStream));
        }

       [Fact]
        public void GetHashSucceedsForMSF7File()
        {
            string hash = PDBParser.GetHash("../../../../testdata/example.pdb");
            string expectedHash = "7F416863ABF34C3E894BAD1739BAA5571";
            Assert.Equal(expectedHash, hash);
        }

       [Fact]
        public void GetHashFailsForNonMSF7File()
        {
            string hash = PDBParser.GetHash("../../../../testdata/example.exe");
            Assert.Null(hash);
        }
    }
}