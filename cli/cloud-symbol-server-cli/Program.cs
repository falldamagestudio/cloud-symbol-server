using CommandLine;

string configLocation = System.IO.Path.ChangeExtension(System.Environment.ProcessPath ?? throw new ApplicationException("No path available to process; cannot fetch config file"), "config.json");
CLI.ConfigFile.Init(configLocation);

int exitCode = Parser.Default.ParseArguments<CLI.UploadOptions, CLI.ListStoresOptions, object>(args)
    .MapResult(
        (CLI.UploadOptions o) => CLI.Upload.DoUpload(o),
        (CLI.ListStoresOptions o) => CLI.ListStores.DoListStores(o),
        errs => 1 );
