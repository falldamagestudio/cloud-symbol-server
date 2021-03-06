namespace CLI
{
    public static class ListUploads
    {
        public static async Task<int> DoListUploads(GlobalOptions globalOptions, string store)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            try {
                IEnumerable<ClientAPI.ListUploads.StoreUpload> uploads = await ClientAPI.ListUploads.DoListUploads(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store);
                Console.WriteLine($"Uploads in store {store}:");
                foreach (ClientAPI.ListUploads.StoreUpload upload in uploads) {
                    Console.WriteLine($"  Upload {upload.UploadId}:");
                    Console.WriteLine($"    Description: {upload.Upload.Description}");
                    Console.WriteLine($"    Build ID: {upload.Upload.BuildId}");
                    foreach (var uploadFile in upload.Upload.Files) {
                        Console.WriteLine($"      FileName: {uploadFile.FileName}, Hash: {uploadFile.Hash}");
                    }
                }
            } catch (ClientAPI.ListUploads.ListUploadsException listUploadsException) {
                Console.Error.WriteLine($"Error while listing uploads in store {store}: {listUploadsException.Message}");
                return 1;
            }

            return 0;
        }
    }
}