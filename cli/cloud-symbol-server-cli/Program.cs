using CommandLine;

namespace CLI
{
    public static class Program
    {
        public static async Task<int> Main(string[] args)
        {
            try {
                string configLocation = System.IO.Path.Combine(System.IO.Path.GetDirectoryName(System.Environment.ProcessPath ?? throw new ApplicationException("No path available to process; cannot fetch config file"))!, "cloud-symbol-server-cli.config.json");
                if (System.IO.File.Exists(configLocation)) {
                    CLI.ConfigFile.Init(configLocation);
                    Console.WriteLine($"Using config file at {configLocation}");
                } else {
                    CLI.ConfigFile.Init();
                }
            } catch {
                Console.WriteLine("Error while reading config file");
                return 1;
            }

            int exitCode = await Parser.Default.ParseArguments<
                    CLI.ListUploadsOptions,
                    CLI.UploadOptions,
                    CLI.ListStoresOptions,
                    CLI.CreateStoreOptions,
                    CLI.DeleteStoreOptions,
                    object
                >(args)
                .MapResult(
                    async (CLI.ListUploadsOptions o) => await CLI.ListUploads.DoListUploads(o),
                    async (CLI.UploadOptions o) => await CLI.Upload.DoUpload(o),
                    async (CLI.ListStoresOptions o) => await CLI.ListStores.DoListStores(o),
                    async (CLI.CreateStoreOptions o) => await CLI.CreateStore.DoCreateStore(o),
                    async (CLI.DeleteStoreOptions o) => await CLI.DeleteStore.DoDeleteStore(o),
                    errs => Task.FromResult(1) );

            return exitCode;
        }
    }
}