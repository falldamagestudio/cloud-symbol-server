using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class ListUploads
    {
        public static async Task<int> DoListUploads(ListUploadsOptions options)
        {
            IEnumerable<ClientAPI.ListUploads.StoreUpload> uploads = await ClientAPI.ListUploads.DoListUploads(options.ServiceURL, options.Email, options.PAT, options.Store);
            Console.WriteLine($"Uploads in store {options.Store}:");
            foreach (ClientAPI.ListUploads.StoreUpload upload in uploads) {
                Console.WriteLine($"  Upload {upload.UploadId}:");
                Console.WriteLine($"    Description: {upload.Upload.Description}");
                Console.WriteLine($"    Build ID: {upload.Upload.BuildId}");
                foreach (var uploadFile in upload.Upload.Files) {
                    Console.WriteLine($"      FileName: {uploadFile.FileName}, Hash: {uploadFile.Hash}");
                }
            }

            return 0;
        }
    }
}