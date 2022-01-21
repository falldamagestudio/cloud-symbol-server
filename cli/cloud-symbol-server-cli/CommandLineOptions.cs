using CommandLine;

namespace CLI
{
    public class CommandLineOptions
    {
        // [Value(index: 0, Required = true, HelpText = "Image file Path upload")]
        // public string UploadFilePath { get; set; }

        [Option(longName: "service-url", Required = false, HelpText = "Service URL")]
        public string? ServiceURL { get; set; } = ConfigFile.GetOrDefault("service-url", "");

        [Option(longName: "email", Required = false, HelpText = "Authentication email")]
        public string? Email { get; set; } = ConfigFile.GetOrDefault("email", "");

        [Option(longName: "pat", Required = false, HelpText = "Authentication Personal Access Token")]
        public string? PAT { get; set; } = ConfigFile.GetOrDefault("pat", "");
    }

    [Verb("upload", HelpText = "Upload symbols")]
    public class UploadOptions : CommandLineOptions
    {
        [Option(longName: "description", Required = true, HelpText = "Textual description of upload")]
        public string? Description { get; set; }

        [Option(longName: "build-id", Required = true, HelpText = "Build ID for upload")]
        public string? BuildId { get; set; }

        [Value(0, Required = true, HelpText = "Which store to upload to")]
        public string? Store { get; set; }

        [Value(1, Min = 1, MetaName="<pattern1 pattern2 pattern3 ...>", HelpText = "Wildcard patterns for files to upload, like 'folder1/*.pdb'")]
        public IEnumerable<string>? Patterns { get; set; }
    }

    [Verb("list-stores", HelpText = "List all stores")]
    public class ListStoresOptions : CommandLineOptions
    {
    }

    [Verb("create-store", HelpText = "Create new store")]
    public class CreateStoreOptions : CommandLineOptions
    {
        [Value(0, Required = true, HelpText = "Name of new store")]
        public string? Store { get; set; }
    }

    [Verb("delete-store", HelpText = "Delete existing store")]
    public class DeleteStoreOptions : CommandLineOptions
    {
        [Value(0, Required = true, HelpText = "Name of store to delete")]
        public string? Store { get; set; }
    }
}