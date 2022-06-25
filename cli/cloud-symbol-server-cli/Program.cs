using System.CommandLine;
using System.CommandLine.NamingConventionBinder;

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

            Command createStoreCommand = new Command("create") {
                new Argument<string>("store", "Name of store to create"),
            };
            createStoreCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store)
                => { return await CLI.CreateStore.DoCreateStore(globalOptions, store); });

            Command deleteStoreCommand = new Command("delete") {
                new Argument<string>("store", "Name of store to delete"),
            };
            deleteStoreCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store)
                => { return await CLI.DeleteStore.DoDeleteStore(globalOptions, store); });

            Command listStoresCommand = new Command("list") {
            };
            listStoresCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions)
                => { return await CLI.ListStores.DoListStores(globalOptions); });

            Command storesCommand = new Command("stores") {
                createStoreCommand,
                deleteStoreCommand,
                listStoresCommand,
            };

            Command createUploadCommand = new Command("create") {
                new Option<string>("--description", "Textual description of upload"),
                new Option<string>("--build-id", "Build ID for upload"),
                new Argument<string>("store", "Name of store to upload to"),
                new Argument<string>("patterns", "Globbing patterns of files to upload") { Arity = ArgumentArity.OneOrMore },
            };
            createUploadCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string description, string buildId, string store, string[] patterns)
                => { return await CLI.Upload.DoUpload(globalOptions, description, buildId, store, patterns); });

            Command listUploadsCommand = new Command("list") {
                new Argument<string>("store", "Name of store to list uploads in"),
            };
            listUploadsCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store)
                => { return await CLI.ListUploads.DoListUploads(globalOptions, store); });

            Command uploadsCommand = new Command("uploads") {
                createUploadCommand,
                listUploadsCommand,
            };

            Command hashFilesCommand = new Command("hash") {
                new Argument<string>("patterns", "Globbing patterns of files to compute hashes for") { Arity = ArgumentArity.OneOrMore },
            };
            hashFilesCommand.Handler = CommandHandler.Create((string[] patterns)
                => { return CLI.HashFiles.DoHashFiles(patterns); });

            RootCommand rootCommand = new RootCommand {
                storesCommand,
                uploadsCommand,
                hashFilesCommand,

                // Global options, available to all subcommands
                new Option<string>("--service-url", () => ConfigFile.GetOrDefault("service-url", "")),
                new Option<string>("--email", () => ConfigFile.GetOrDefault("email", "")),
                new Option<string>("--pat", () => ConfigFile.GetOrDefault("pat", "")),
            };
            rootCommand.Handler = CommandHandler.Create(() => {});

            // Parse the incoming args and invoke the handler
            return await rootCommand.InvokeAsync(args);
        }
    //     public static async Task<int> Main(string[] args)
    //     {

    //         int exitCode = await Parser.Default.ParseArguments<
    //                 CLI.ListUploadsOptions,
    //                 CLI.UploadOptions,
    //                 CLI.ListStoresOptions,
    //                 CLI.CreateStoreOptions,
    //                 CLI.DeleteStoreOptions,
    //                 object
    //             >(args)
    //             .MapResult(
    //                 async (CLI.ListUploadsOptions o) => await CLI.ListUploads.DoListUploads(o),
    //                 async (CLI.UploadOptions o) => await CLI.Upload.DoUpload(o),
    //                 async (CLI.ListStoresOptions o) => await CLI.ListStores.DoListStores(o),
    //                 async (CLI.CreateStoreOptions o) => await CLI.CreateStore.DoCreateStore(o),
    //                 async (CLI.DeleteStoreOptions o) => await CLI.DeleteStore.DoDeleteStore(o),
    //                 errs => Task.FromResult(1) );

    //         return exitCode;
    //     }
    }
}