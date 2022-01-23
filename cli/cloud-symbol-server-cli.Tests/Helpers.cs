using System;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public static class Helpers
{
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