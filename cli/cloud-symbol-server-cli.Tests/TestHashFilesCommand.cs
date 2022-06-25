using Xunit;
using System.IO;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{

    [Fact]
    public async Task TestHashFilesSucceedsIfFilesExist()
    {
        Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
            "hash",
            "../../../../testdata/*.pdb",
            "../../../../testdata/*.exe",
        });

        Assert.Equal(0, result.ExitCode);
        Assert.Contains("example.pdb", result.Stdout);
        Assert.Contains("7F416863ABF34C3E894BAD1739BAA5571", result.Stdout);
        Assert.Contains("example.exe", result.Stdout);
        Assert.Contains("61C0D4547000", result.Stdout);
    }
}