using Xunit;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task TestListStores()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "--service-url", Helpers.GetBackendAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
                "stores",
                "list"
            });

            Assert.Equal("", result.Stderr);
            Assert.DoesNotContain(Helpers.TestStore, result.Stdout);
            Assert.Equal(0, result.ExitCode);
        }

        {
            await Helpers.EnsureTestStoreExists();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "--service-url", Helpers.GetBackendAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
                "stores",
                "list"
            });

            Assert.Equal("", result.Stderr);
            Assert.Contains(Helpers.TestStore, result.Stdout);
            Assert.Equal(0, result.ExitCode);
        }
    }

    [Fact]
    public async Task CreateStoreSucceedsIfStoreDoesNotAlreadyExist()
    {
        await Helpers.EnsureTestStoreDoesNotExist();

        Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
            "--service-url", Helpers.GetBackendAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
            "stores",
            "create",
            Helpers.TestStore,
        });

        Assert.Equal("", result.Stderr);
        Assert.NotEqual("", result.Stdout);
        Assert.Equal(0, result.ExitCode);
    }

    [Fact]
    public async Task CreateStoreFailsIfStoreAlreadyExists()
    {
        await Helpers.EnsureTestStoreExists();

        Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
            "--service-url", Helpers.GetBackendAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
            "stores",
            "create",
            Helpers.TestStore,
        });

        Assert.NotEqual("", result.Stderr);
        Assert.DoesNotContain("Exception", result.Stderr);
        Assert.Equal("", result.Stdout);
        Assert.Equal(1, result.ExitCode);
    }

    [Fact]
    public async Task DeleteStoreSucceedsIfStoreAlreadyExists()
    {
        await Helpers.EnsureTestStoreExists();

        Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
            "--service-url", Helpers.GetBackendAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
            "stores",
            "delete",
            Helpers.TestStore,
        });

        Assert.Equal("", result.Stderr);
        Assert.NotEqual("", result.Stdout);
        Assert.Equal(0, result.ExitCode);
    }

    [Fact]
    public async Task DeleteStoreFailsIfStoreDoesNotAlreadyExist()
    {
        await Helpers.EnsureTestStoreDoesNotExist();

        Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
            "--service-url", Helpers.GetBackendAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
            "stores",
            "delete",
            Helpers.TestStore,
        });

        Assert.NotEqual("", result.Stderr);
        Assert.DoesNotContain("Exception", result.Stderr);
        Assert.Equal("", result.Stdout);
        Assert.Equal(1, result.ExitCode);
    }
}