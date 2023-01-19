namespace CLI
{
    public static class ListBlobs
    {
        public static async Task<int> DoListBlobs(GlobalOptions globalOptions, string store, string file)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            try {
                Console.WriteLine($"Blobs of file {file} in store {store}:");

                // Request file info in batches, since the backend API limits max response size
                const int limit = 100;

                for (int offset = 0; ; offset += limit) {
                    BackendAPI.Model.GetStoreFileBlobsResponse blobsResponse = await ClientAPI.ListBlobs.DoListBlobs(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, file, offset, limit);
                    for (int batchOffset = 0; batchOffset < blobsResponse.Blobs.Count; batchOffset++) {
                        BackendAPI.Model.GetStoreFileBlobResponse blob = blobsResponse.Blobs[batchOffset];
                        Console.WriteLine($"  BlobIdentifier {blob.BlobIdentifier}:");
                        Console.WriteLine($"    Type: {blob.Type}");
                        Console.WriteLine($"    Size {blob.Size}");
                        Console.WriteLine($"    Content SHA256 Hash: {blob.ContentHash}");
                        Console.WriteLine($"    Status: {blob.Status}");
                    }

                    if (offset + blobsResponse.Blobs.Count >= blobsResponse.Pagination.Total)
                        break;
                }
            } catch (ClientAPI.ClientAPIException exception) {
                Console.Error.WriteLine($"Error while listing blobs of files in store {store} / file {file}: {exception.Message}");
                return 1;
            }

            return 0;
        }
    }
}