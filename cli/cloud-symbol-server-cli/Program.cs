using CommandLine;

try {
    string configLocation = System.IO.Path.Combine(System.IO.Path.GetDirectoryName(System.Environment.ProcessPath ?? throw new ApplicationException("No path available to process; cannot fetch config file"))!, "cloud-symbol-server-cli.config.json");
    if (System.IO.File.Exists(configLocation))
    {
        CLI.ConfigFile.Init(configLocation);
        Console.WriteLine($"Using config file at {configLocation}");
    }
} catch {
    Console.WriteLine("Error while reading config file");
    return 1;
}

int exitCode = await Parser.Default.ParseArguments<CLI.UploadOptions, CLI.ListStoresOptions, object>(args)
    .MapResult(
        async (CLI.UploadOptions o) => await CLI.Upload.DoUpload(o),
        async (CLI.ListStoresOptions o) => await CLI.ListStores.DoListStores(o),
        errs => Task.FromResult(1) );

return exitCode;