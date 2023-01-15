namespace CLI
{
    public static class ListFiles
    {
        public static async Task<int> DoListFiles(GlobalOptions globalOptions, string store)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            try {
                Console.WriteLine($"Files in store {store}:");

                // Request file info in batches, since the backend API limits max response size
                const int limit = 100;

                for (int offset = 0; ; offset += limit) {
                    BackendAPI.Model.GetStoreFilesResponse filesResponse = await ClientAPI.ListFiles.DoListFiles(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, offset, limit);
                    for (int batchOffset = 0; batchOffset < filesResponse.Files.Count; batchOffset++) {
                        Console.WriteLine($"  {filesResponse.Files[batchOffset]}");
                    }

                    if (offset + filesResponse.Files.Count >= filesResponse.Pagination.Total)
                        break;
                }
            } catch (ClientAPI.ClientAPIException exception) {
                Console.Error.WriteLine($"Error while listing files in store {store}: {exception.Message}");
                return 1;
            }

            return 0;
        }
    }
}