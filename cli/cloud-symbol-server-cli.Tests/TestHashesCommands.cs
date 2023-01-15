using Xunit;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task ListHashesFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            await SpecRunner.RunSpecCommand("../../../../testspecs/ListHashesFailsIfStoreDoesNotExist", output);
        }
    }

    [Fact]
    public async Task ListHashesSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();
            await Helpers.PopulateTestStore();

            await SpecRunner.RunSpecCommand("../../../../testspecs/ListHashesSucceedsIfStoreExists", output);
        }
    }
}