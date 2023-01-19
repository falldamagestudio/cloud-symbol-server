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
        public void GetCodeIdentifierSucceedsForExe()
        {
            string hash = PEParser.GetCodeIdentifier("../../../../testdata/example.exe");
            string expectedCodeIdentifier = "61C0D4547000";
            Assert.Equal(expectedCodeIdentifier, hash);
        }

        [Fact]
        public void GetCodeIdentifierFailsForNonExe()
        {
            string hash = PEParser.GetCodeIdentifier("../../../../testdata/example.pdb");
            Assert.Null(hash);
        }
    }
}