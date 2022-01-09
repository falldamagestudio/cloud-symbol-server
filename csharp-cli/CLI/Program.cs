using CommandLine;


Parser.Default.ParseArguments<CLI.CommandLineOptions>(args)
        .WithParsed<CLI.CommandLineOptions>(o =>
        {
            CLI.Upload.DoUpload(o.ServiceURL, o.Email, o.PAT, new string[] { "*.pdb" });
        });

