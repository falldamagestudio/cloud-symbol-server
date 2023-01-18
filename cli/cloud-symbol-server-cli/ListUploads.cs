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
                Console.WriteLine($"Uploads in store {store}:");

                // Request upload info in batches, since the backend API limits max response size
                const int limit = 100;

                for (int offset = 0; ; offset += limit) {
                    BackendAPI.Model.GetStoreUploadsResponse uploadsResponse = await ClientAPI.ListUploads.DoListUploads(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, offset, limit);
                    for (int batchOffset = 0; batchOffset < uploadsResponse.Uploads.Count; batchOffset++) {
                        BackendAPI.Model.GetStoreUploadResponse upload = uploadsResponse.Uploads[batchOffset];
                        Console.WriteLine($"  Upload {offset + batchOffset}:");
                        Console.WriteLine($"    Status: {upload.Status}");
                        Console.WriteLine($"    Description: {upload.Description}");
                        Console.WriteLine($"    Build ID: {upload.BuildId}");
                        foreach (var uploadFile in upload.Files) {
                            Console.WriteLine($"      FileName: {uploadFile.FileName}, BlobIdentifier: {uploadFile.BlobIdentifier}, Status: {uploadFile.Status}");
                        }
                    }

                    if (offset + uploadsResponse.Uploads.Count >= uploadsResponse.Pagination.Total)
                        break;
                }
            } catch (ClientAPI.ClientAPIException exception) {
                Console.Error.WriteLine($"Error while listing uploads in store {store}: {exception.Message}");
                return 1;
            }

            return 0;
        }
    }
}