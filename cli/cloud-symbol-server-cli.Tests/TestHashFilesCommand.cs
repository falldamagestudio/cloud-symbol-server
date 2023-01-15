using Xunit;
using System.IO;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{

    [Fact]
    public async Task HashFilesSucceedsIfFilesExist()
    {
        await TestSpecRunner.RunSpecCommand("../../../../testspecs/HashFilesSucceedsIfFilesExist", output);
    }
}