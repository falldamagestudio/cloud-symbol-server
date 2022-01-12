using CommandLine;

int exitCode = Parser.Default.ParseArguments<CLI.UploadOptions, CLI.ListStoresOptions, object>(args)
    .MapResult(
        (CLI.UploadOptions o) => CLI.Upload.DoUpload(o),
        (CLI.ListStoresOptions o) => CLI.ListStores.DoListStores(o),
        errs => 1 );
