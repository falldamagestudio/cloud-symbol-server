using System;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public static class Helpers
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

    public static async Task<CLICommandResult> RunCLICommand(string[] args)
    {
        using (var consoleStdout = new CaptureStdout()) {
            using (var consoleStderr = new CaptureStderr()) {
                
                int exitCode = await CLI.Program.Main(args);

                return new CLICommandResult(exitCode, consoleStdout.GetOutput(), consoleStderr.GetError());
            }
        }
    }

    public const string TestStore = "example";

    public static string GetAdminAPIEndpoint()
    {
        return Environment.GetEnvironmentVariable("ADMIN_API_ENDPOINT");
    }

    public static string GetTestEmail()
    {
        return Environment.GetEnvironmentVariable("TEST_EMAIL");
    }

    public static string GetTestPAT()
    {
        return Environment.GetEnvironmentVariable("TEST_PAT");
    }

    public static async Task DeleteTestStore(bool ignoreIfNotExists)
    {
        bool deleted = await ClientAPI.DeleteStore.DoDeleteStore(GetAdminAPIEndpoint(), GetTestEmail(), GetTestPAT(), TestStore);
        if (!deleted && !ignoreIfNotExists)
            throw new ApplicationException("Test store did not exist");
    }

    public static async Task CreateTestStore(bool ignoreIfAlreadyExists)
    {
        bool created = await ClientAPI.CreateStore.DoCreateStore(GetAdminAPIEndpoint(), GetTestEmail(), GetTestPAT(), TestStore);
        if (!created && !ignoreIfAlreadyExists)
            throw new ApplicationException("Test store alredady existed");
    }

    public static async Task EnsureTestStoreDoesNotExist()
    {
        await DeleteTestStore(true);
    }

    public static async Task EnsureTestStoreExists()
    {
        await DeleteTestStore(true);
        await CreateTestStore(false);
    }
}