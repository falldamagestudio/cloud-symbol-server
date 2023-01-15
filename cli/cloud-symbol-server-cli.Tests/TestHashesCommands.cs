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

    [Fact]
    public async Task ComputeHashesSucceedsIfFilesExist()
    {
        await SpecRunner.RunSpecCommand("../../../../testspecs/ComputeHashesSucceedsIfFilesExist", output);
    }

    [Fact]
    public async Task DownloadHashFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            await SpecRunner.RunSpecCommand("../../../../testspecs/DownloadHashFailsIfStoreDoesNotExist", output);
        }
    }

    [Fact]
    public async Task DownloadHashSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();
            await Helpers.PopulateTestStore();

            const string downloadedFileName = "example.pdb";

            if (System.IO.File.Exists(downloadedFileName)) {
                System.IO.File.Delete(downloadedFileName);
            }

            await SpecRunner.RunSpecCommand("../../../../testspecs/DownloadHashSucceedsIfStoreExists", output);
            Assert.True(System.IO.File.Exists(downloadedFileName));
        }
    }
}