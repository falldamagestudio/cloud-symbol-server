﻿using CommandLine;

int exitCode = Parser.Default.ParseArguments<CLI.UploadOptions, object>(args)
    .MapResult(
        (CLI.UploadOptions o) => CLI.Upload.DoUpload(o),
        errs => 1 );
