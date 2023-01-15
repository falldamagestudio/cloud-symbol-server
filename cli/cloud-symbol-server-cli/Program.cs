﻿using System.CommandLine;
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

            // Stores commands

            Command createStoreCommand = new Command("create", "Create a new store within Cloud Symbol Server") {
                new Argument<string>("store", "Name of store to create"),
            };
            createStoreCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store)
                => { return await CLI.CreateStore.DoCreateStore(globalOptions, store); });

            Command deleteStoreCommand = new Command("delete", "Delete a store within Cloud Symbol Server") {
                new Argument<string>("store", "Name of store to delete"),
            };
            deleteStoreCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store)
                => { return await CLI.DeleteStore.DoDeleteStore(globalOptions, store); });

            Command listStoresCommand = new Command("list", "List stores present within Cloud Symbol Server") {
            };
            listStoresCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions)
                => { return await CLI.ListStores.DoListStores(globalOptions); });

            Command storesCommand = new Command("stores", "Manage stores within Cloud Symbol Server") {
                createStoreCommand,
                deleteStoreCommand,
                listStoresCommand,
            };
            storesCommand.Handler = CommandHandler.Create(() => storesCommand.Invoke("--help"));

            // Files commands

            Command listFilesCommand = new Command("list", "List files present within Cloud Symbol Server") {
                new Argument<string>("store", "Name of store containing files"),
            };
            listFilesCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store)
                => { return await CLI.ListFiles.DoListFiles(globalOptions, store); });

            Command filesCommand = new Command("files", "Manage files within Cloud Symbol Server") {
                listFilesCommand,
            };
            filesCommand.Handler = CommandHandler.Create(() => filesCommand.Invoke("--help"));

            // Hashes commands

            Command listHashesCommand = new Command("list", "List hashes of files present within Cloud Symbol Server") {
                new Argument<string>("store", "Name of store containing file"),
                new Argument<string>("file", "Name of file"),
            };
            listHashesCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store, string file)
                => { return await CLI.ListHashes.DoListHashes(globalOptions, store, file); });

            Command hashesCommand = new Command("hashes", "Manage hashes of files within Cloud Symbol Server") {
                listHashesCommand,
            };
            hashesCommand.Handler = CommandHandler.Create(() => hashesCommand.Invoke("--help"));

            // Uploads commands

            Command createUploadCommand = new Command("create", "Upload files to a store") {
                new Option<string>("--description", "Textual description of upload"),
                new Option<string>("--build-id", "Build ID for upload"),
                new Argument<string>("store", "Name of store to upload to"),
                new Argument<string>("patterns", "Globbing patterns of files to upload") { Arity = ArgumentArity.OneOrMore },
            };
            createUploadCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string description, string buildId, string store, string[] patterns)
                => { return await CLI.Upload.DoUpload(globalOptions, description, buildId, store, patterns); });

            Command expireUploadCommand = new Command("expire", "expire an upload and its files") {
                new Argument<string>("store", "Name of store containing upload"),
                new Argument<string>("upload-id", "upload ID to expire"),
            };
            expireUploadCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store, string uploadId)
                => { return await CLI.ExpireStoreUpload.DoExpireStoreUpload(globalOptions, store, uploadId); });

            Command listUploadsCommand = new Command("list", "List existing uploads within a store") {
                new Argument<string>("store", "Name of store to list uploads in"),
            };
            listUploadsCommand.Handler = CommandHandler.Create(async (GlobalOptions globalOptions, string store)
                => { return await CLI.ListUploads.DoListUploads(globalOptions, store); });

            Command uploadsCommand = new Command("uploads", "Upload files, and manage uploaded files within a store") {
                createUploadCommand,
                expireUploadCommand,
                listUploadsCommand,
            };
            uploadsCommand.Handler = CommandHandler.Create(() => uploadsCommand.Invoke("--help"));

            // hash files command

            Command hashFilesCommand = new Command("hash", "Compute hashes for files") {
                new Argument<string>("patterns", "Globbing patterns of files to compute hashes for") { Arity = ArgumentArity.OneOrMore },
            };
            hashFilesCommand.Handler = CommandHandler.Create((string[] patterns)
                => { return CLI.HashFiles.DoHashFiles(patterns); });

            // Root command

            RootCommand rootCommand = new RootCommand("Cloud Symbol Server CLI tool") {
                storesCommand,
                uploadsCommand,
                filesCommand,
                hashesCommand,
                hashFilesCommand,

                // Global options, available to all subcommands
                new Option<string>("--service-url", () => ConfigFile.GetOrDefault("service-url", "")),
                new Option<string>("--email", () => ConfigFile.GetOrDefault("email", "")),
                new Option<string>("--pat", () => ConfigFile.GetOrDefault("pat", "")),
            };

            // When the CLI command is invoked with no arguments at all, print help
            rootCommand.Handler = CommandHandler.Create(() => rootCommand.Invoke("--help"));

            // Parse the incoming args and invoke the handler
            return await rootCommand.InvokeAsync(args);
        }
    }
}