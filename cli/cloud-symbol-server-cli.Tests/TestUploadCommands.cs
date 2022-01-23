using Xunit;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task TestUpload()
    {
        {
            await Helpers.EnsureTestStoreExists();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "upload",
                "--service-url", Helpers.GetAdminAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
                "--description", "testupload",
                "--build-id", "build 432",
                Helpers.TestStore,
                "../../../../testdata/*.pdb",
                "../../../../testdata/*.exe",
            });

            Assert.Equal("", result.Stderr);
            Assert.Equal(0, result.ExitCode);
        }
    }
}