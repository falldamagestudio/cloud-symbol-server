using Xunit;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task ListBlobsFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            await SpecRunner.RunSpecCommand("../../../../testspecs/ListBlobsFailsIfStoreDoesNotExist", output);
        }
    }

    [Fact]
    public async Task ListBlobsSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();
            await Helpers.PopulateTestStore();

            await SpecRunner.RunSpecCommand("../../../../testspecs/ListBlobsSucceedsIfStoreExists", output);
        }
    }

    [Fact]
    public async Task ComputeHashesSucceedsIfFilesExist()
    {
        await SpecRunner.RunSpecCommand("../../../../testspecs/ComputeHashesSucceedsIfFilesExist", output);
    }

    [Fact]
    public async Task DownloadBlobFailsIfStoreDoesNotExist()
    {
        {
            await Helpers.EnsureTestStoreDoesNotExist();

            await SpecRunner.RunSpecCommand("../../../../testspecs/DownloadBlobFailsIfStoreDoesNotExist", output);
        }
    }

    [Fact]
    public async Task DownloadBlobSucceedsIfStoreExists()
    {
        {
            await Helpers.EnsureTestStoreExists();
            await Helpers.PopulateTestStore();

            const string downloadedFileName = "example.pdb";

            if (System.IO.File.Exists(downloadedFileName)) {
                System.IO.File.Delete(downloadedFileName);
            }

            await SpecRunner.RunSpecCommand("../../../../testspecs/DownloadBlobSucceedsIfStoreExists", output);
            Assert.True(System.IO.File.Exists(downloadedFileName));
        }
    }
}