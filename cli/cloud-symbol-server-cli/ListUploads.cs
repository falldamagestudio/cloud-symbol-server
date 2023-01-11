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
                int offset = 0;
                int limit = 100;
                BackendAPI.Model.GetStoreUploadsResponse uploadsResponse = await ClientAPI.ListUploads.DoListUploads(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, offset, limit);
                Console.WriteLine($"Uploads in store {store}:");
                int i = offset;
                foreach (BackendAPI.Model.GetStoreUploadResponse upload in uploadsResponse.Uploads) {
                    Console.WriteLine($"  Upload {i}:");
                    Console.WriteLine($"    Status: {upload.Status}");
                    Console.WriteLine($"    Description: {upload.Description}");
                    Console.WriteLine($"    Build ID: {upload.BuildId}");
                    foreach (var uploadFile in upload.Files) {
                        Console.WriteLine($"      FileName: {uploadFile.FileName}, Hash: {uploadFile.Hash}, Status: {uploadFile.Status}");
                    }
                    i++;
                }
            } catch (ClientAPI.ClientAPIException exception) {
                Console.Error.WriteLine($"Error while listing uploads in store {store}: {exception.Message}");
                return 1;
            }

            return 0;
        }
    }
}