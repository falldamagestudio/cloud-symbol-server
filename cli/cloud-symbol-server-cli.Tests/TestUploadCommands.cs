using Xunit;
using System.IO;
using System.Threading.Tasks;

namespace cloud_symbol_server_cli.Tests;

public partial class TestCommands
{
    [Fact]
    public async Task CreateUploadFailsIfStoreDoesNotExist()
    {
        await Helpers.EnsureTestStoreDoesNotExist();
        await SpecRunner.RunSpecCommand("../../../../testspecs/CreateUploadFailsIfStoreDoesNotExist", output);
    }

    [Fact]
    public async Task CreateUploadSucceedsIfStoreExists()
    {
        await Helpers.EnsureTestStoreExists();
        await SpecRunner.RunSpecCommand("../../../../testspecs/CreateUploadSucceedsIfStoreExists", output);

        byte[] content = await Helpers.DownloadFile("example.pdb", "7F416863ABF34C3E894BAD1739BAA5571");
        byte[] desiredContent = File.ReadAllBytes("../../../../testdata/example.pdb");
        Assert.NotNull(content);
        Assert.Equal(desiredContent, content);
    }

    [Fact]
    public async Task ListUploadsFailsIfStoreDoesNotExist()
    {
        await Helpers.EnsureTestStoreDoesNotExist();
        await SpecRunner.RunSpecCommand("../../../../testspecs/ListUploadsFailsIfStoreDoesNotExist", output);
    }

    [Fact]
    public async Task ListUploadsSucceedsIfStoreExists()
    {
        await Helpers.EnsureTestStoreExists();
        await SpecRunner.RunSpecCommand("../../../../testspecs/ListUploadsSucceedsIfStoreExists", output);
    }

    [Fact]
    public async Task ExpireUploadSucceedsIfUploadExists()
    {
        await Helpers.EnsureTestStoreExists();

        {
            // Upload build
            await SpecRunner.RunSpecCommand("../../../../testspecs/ExpireUploadSucceedsIfUploadExists/1. UploadBuild", output);

            byte[] content = await Helpers.DownloadFile("example.pdb", "7F416863ABF34C3E894BAD1739BAA5571");
            byte[] desiredContent = File.ReadAllBytes("../../../../testdata/example.pdb");
            Assert.NotNull(content);
            Assert.Equal(desiredContent, content);
        }

        string uploadId = "0";

        // Ensure that the upload is not expired

        await SpecRunner.RunSpecCommand("../../../../testspecs/ExpireUploadSucceedsIfUploadExists/2. EnsureUploadIsNotExpired", output);

        // Expire upload

        await SpecRunner.RunSpecCommand("../../../../testspecs/ExpireUploadSucceedsIfUploadExists/3. ExpireUpload", output);

        // Ensure that the upload is expired

        await SpecRunner.RunSpecCommand("../../../../testspecs/ExpireUploadSucceedsIfUploadExists/4. EnsureUploadIsExpired", output);
    }


}