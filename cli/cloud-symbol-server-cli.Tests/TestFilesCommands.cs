using Xunit;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task TestListFilesFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            await TestSpecRunner.RunSpecCommand("../../../../testspecs/TestListFilesFailsIfStoreDoesNotExist", output);
        }
    }

    [Fact]
    public async Task TestListFilesSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();
            await Helpers.PopulateTestStore();

            await TestSpecRunner.RunSpecCommand("../../../../testspecs/TestListFilesSucceedsIfStoreExists", output);
        }
    }
}