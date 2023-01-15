using Xunit;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task ListFilesFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            await SpecRunner.RunSpecCommand("../../../../testspecs/ListFilesFailsIfStoreDoesNotExist", output);
        }
    }

    [Fact]
    public async Task ListFilesSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();
            await Helpers.PopulateTestStore();

            await SpecRunner.RunSpecCommand("../../../../testspecs/ListFilesSucceedsIfStoreExists", output);
        }
    }
}