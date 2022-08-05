using System;
using System.IO;
using Xunit;
using ClientAPI;
using System.Linq;
using System.Threading.Tasks;
using System.Text;

namespace ClientAPI.Tests
{

    public class PEParserTests
    {
        [Fact]
        public void GetHashSucceedsForExe()
        {
            string hash = PEParser.GetHash("../../../../testdata/example.exe");
            string expectedHash = "61C0D4547000";
            Assert.Equal(expectedHash, hash);
        }

        [Fact]
        public void GetHashFailsForNonExe()
        {
            string hash = PEParser.GetHash("../../../../testdata/example.pdb");
            Assert.Null(hash);
        }
    }
}