using System.IO;

namespace CLI
{
    public static class DownloadBlob
    {
        public static async Task<int> DoDownloadBlob(GlobalOptions globalOptions, string store, string file, string blob)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            try {
                BackendAPI.Model.GetStoreFileBlobDownloadUrlResponse getStoreFileBlobDownloadUrlResponse = await ClientAPI.GetBlobDownloadUrl.DoGetBlobDownloadUrl(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, file, blob);

                if (getStoreFileBlobDownloadUrlResponse.Method != "GET")
                {
                    Console.Error.WriteLine($"Unsupported download method {getStoreFileBlobDownloadUrlResponse.Method}; only GET is supported");
                    return 1;
                }

                HttpClient client = new HttpClient();
                var response = await client.GetAsync(getStoreFileBlobDownloadUrlResponse.Url);

                using (Stream fileStream = File.Create(file))
                {
                    await response.Content.CopyToAsync(fileStream);
                }

                Console.WriteLine($"File-blob for {file} / {blob} downloaded and written to local file: {file}");
                return 0;
            } catch (ClientAPI.ClientAPIException exception) {
                Console.Error.WriteLine($"Error while getting file-blob download URL for store {store} / file {file} / blob {blob}: {exception.Message}");
                return 1;
            }
        }
    }
}