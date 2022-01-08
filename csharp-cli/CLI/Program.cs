using CommandLine;

// See https://aka.ms/new-console-template for more information
Console.WriteLine("Hello, World!");

Parser.Default.ParseArguments<CLI.CommandLineOptions>(args)
        .WithParsed<CLI.CommandLineOptions>(o =>
        {
            CLI.Upload.DoUpload(o.ServiceURL, o.Email, o.PAT, new string[] { "*.pdb" });
        });

