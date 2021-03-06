using Xunit;
using System.IO;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task TestUploadFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "--service-url", Helpers.GetAdminAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
                "uploads",
                "create",
                "--description", "testupload",
                "--build-id", "build 432",
                Helpers.TestStore,
                "../../../../testdata/*.pdb",
                "../../../../testdata/*.exe",
            });

            Assert.NotEqual("", result.Stderr);
            Assert.DoesNotContain("Exception", result.Stderr);
            Assert.Equal(1, result.ExitCode);
        }
    }

    [Fact]
    public async Task TestUploadSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "--service-url", Helpers.GetAdminAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
                "uploads",
                "create",
                "--description", "testupload",
                "--build-id", "build 432",
                Helpers.TestStore,
                "../../../../testdata/*.pdb",
                "../../../../testdata/*.exe",
            });

            Assert.Equal("", result.Stderr);
            Assert.Equal(0, result.ExitCode);
            byte[] content = await Helpers.DownloadFile("example.pdb", "7F416863ABF34C3E894BAD1739BAA5571");
            byte[] desiredContent = File.ReadAllBytes("../../../../testdata/example.pdb");
            Assert.NotNull(content);
            Assert.Equal(desiredContent, content);
        }
    }

    [Fact]
    public async Task TestListUploadsFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "--service-url", Helpers.GetAdminAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
                "uploads",
                "list",
                Helpers.TestStore,
            });

            Assert.NotEqual("", result.Stderr);
            Assert.DoesNotContain("Exception", result.Stderr);
            Assert.Equal(1, result.ExitCode);
        }
    }

    [Fact]
    public async Task TestListUploadsSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "--service-url", Helpers.GetAdminAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
                "uploads",
                "list",
                Helpers.TestStore,
            });

            Assert.Equal("", result.Stderr);
            Assert.NotEqual("", result.Stdout);
            Assert.Equal(0, result.ExitCode);
        }
    }
}