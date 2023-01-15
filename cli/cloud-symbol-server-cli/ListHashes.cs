namespace CLI
{
    public static class ListHashes
    {
        public static async Task<int> DoListHashes(GlobalOptions globalOptions, string store, string file)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            try {
                Console.WriteLine($"Hashes of file {file} in store {store}:");

                // Request file info in batches, since the backend API limits max response size
                const int limit = 100;

                for (int offset = 0; ; offset += limit) {
                    BackendAPI.Model.GetStoreFileHashesResponse hashesResponse = await ClientAPI.ListHashes.DoListHashes(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, file, offset, limit);
                    for (int batchOffset = 0; batchOffset < hashesResponse.Hashes.Count; batchOffset++) {
                        BackendAPI.Model.GetStoreFileHashResponse hash = hashesResponse.Hashes[batchOffset];
                        Console.WriteLine($"  Hash {hash.Hash}:");
                        Console.WriteLine($"    Status: {hash.Status}");
                    }

                    if (offset + hashesResponse.Hashes.Count >= hashesResponse.Pagination.Total)
                        break;
                }
            } catch (ClientAPI.ClientAPIException exception) {
                Console.Error.WriteLine($"Error while listing hashes of files in store {store} / file {file}: {exception.Message}");
                return 1;
            }

            return 0;
        }
    }
}