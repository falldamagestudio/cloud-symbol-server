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
                "list-stores",
                "--service-url", Helpers.GetAdminAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
            });

            Assert.Equal("", result.Stderr);
            Assert.DoesNotContain(Helpers.TestStore, result.Stdout);
            Assert.Equal(0, result.ExitCode);
        }

        {
            await Helpers.EnsureTestStoreExists();

            Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
                "list-stores",
                "--service-url", Helpers.GetAdminAPIEndpoint(),
                "--email", Helpers.GetTestEmail(),
                "--pat", Helpers.GetTestPAT(),
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
            "create-store",
            Helpers.TestStore,
            "--service-url", Helpers.GetAdminAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
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
            "create-store",
            Helpers.TestStore,
            "--service-url", Helpers.GetAdminAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
        });

        Assert.NotEqual("", result.Stderr);
        Assert.Equal("", result.Stdout);
        Assert.Equal(1, result.ExitCode);
    }

    [Fact]
    public async Task DeleteStoreSucceedsIfStoreAlreadyExists()
    {
        await Helpers.EnsureTestStoreExists();

        Helpers.CLICommandResult result = await Helpers.RunCLICommand(new string[]{
            "delete-store",
            Helpers.TestStore,
            "--service-url", Helpers.GetAdminAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
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
            "delete-store",
            Helpers.TestStore,
            "--service-url", Helpers.GetAdminAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
        });

        Assert.NotEqual("", result.Stderr);
        Assert.Equal("", result.Stdout);
        Assert.Equal(1, result.ExitCode);
    }
}