using Xunit;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public class TestStoreCommands
{
    public class CLICommandResult
    {
        public readonly int ExitCode;
        public readonly string Stdout;
        public readonly string Stderr;

        public CLICommandResult(int exitCode, string stdout, string stderr)
        {
            ExitCode = exitCode;
            Stdout = stdout;
            Stderr = stderr;
        }
    }

    public async Task<CLICommandResult> RunCLICommand(string[] args)
    {
        using (var consoleStdout = new CaptureStdout()) {
            using (var consoleStderr = new CaptureStderr()) {
                
                int exitCode = await CLI.Program.Main(args);

                return new CLICommandResult(exitCode, consoleStdout.GetOutput(), consoleStderr.GetError());
            }
        }
    }

    [Fact]
    public async Task TestListStores()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            CLICommandResult result = await RunCLICommand(new string[]{
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

            CLICommandResult result = await RunCLICommand(new string[]{
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

        CLICommandResult result = await RunCLICommand(new string[]{
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

        CLICommandResult result = await RunCLICommand(new string[]{
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

        CLICommandResult result = await RunCLICommand(new string[]{
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

        CLICommandResult result = await RunCLICommand(new string[]{
            "delete-store",
            Helpers.TestStore,
            "--service-url", Helpers.GetAdminAPIEndpoint(),
            "--email", Helpers.GetTestEmail(),
            "--pat", Helpers.GetTestPAT(),
        });

        // Assert.NotEqual("", result.Stderr);
        Assert.Equal("", result.Stdout);
        Assert.Equal(1, result.ExitCode);
    }
}