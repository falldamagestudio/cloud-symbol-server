using System.IO;

namespace CLI
{
    public static class DownloadHash
    {
        public static async Task<int> DoDownloadHash(GlobalOptions globalOptions, string store, string file, string hash)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            try {
                BackendAPI.Model.GetStoreFileHashDownloadUrlResponse getStoreFileHashDownloadUrlResponse = await ClientAPI.GetHashDownloadUrl.DoGetHashDownloadUrl(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, file, hash);

                if (getStoreFileHashDownloadUrlResponse.Method != "GET")
                {
                    Console.Error.WriteLine($"Unsupported download method {getStoreFileHashDownloadUrlResponse.Method}; only GET is supported");
                    return 1;
                }

                HttpClient client = new HttpClient();
                var response = await client.GetAsync(getStoreFileHashDownloadUrlResponse.Url);

                using (Stream fileStream = File.Create(file))
                {
                    await response.Content.CopyToAsync(fileStream);
                }

                Console.WriteLine($"File-hash for {file} / {hash} downloaded and written to local file: {file}");
                return 0;
            } catch (ClientAPI.ClientAPIException exception) {
                Console.Error.WriteLine($"Error while getting file-hash download URL for store {store} / file {file} / hash {hash}: {exception.Message}");
                return 1;
            }
        }
    }
}