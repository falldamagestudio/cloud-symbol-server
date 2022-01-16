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

int exitCode = Parser.Default.ParseArguments<CLI.UploadOptions, CLI.ListStoresOptions, object>(args)
    .MapResult(
        (CLI.UploadOptions o) => CLI.Upload.DoUpload(o),
        (CLI.ListStoresOptions o) => CLI.ListStores.DoListStores(o),
        errs => 1 );

return exitCode;