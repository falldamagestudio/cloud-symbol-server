namespace CLI
{
    public static class ExpireStoreUpload
    {
        public static async Task<int> DoExpireStoreUpload(GlobalOptions globalOptions, string store, int uploadId)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            await ClientAPI.ExpireStoreUpload.DoExpireStoreUpload(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, uploadId);
            Console.WriteLine($"Expired upload: {store} / {uploadId}");
            return 0;
        }
    }
}