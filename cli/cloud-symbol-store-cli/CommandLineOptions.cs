using CommandLine;

namespace CLI
{
    public class CommandLineOptions
    {
        // [Value(index: 0, Required = true, HelpText = "Image file Path upload")]
        // public string UploadFilePath { get; set; }

        [Option(longName: "service-url", Required = false, HelpText = "Service URL", Default = "http://localhost:8084")]
        public string? ServiceURL { get; set; }

        [Option(longName: "email", Required = false, HelpText = "Authentication email", Default = "")]
        public string? Email { get; set; }

        [Option(longName: "pat", Required = false, HelpText = "Authentication Personal Access Token", Default = "")]
        public string? PAT { get; set; }
    }

    [Verb("upload", HelpText = "Upload symbols")]
    public class UploadOptions : CommandLineOptions
    {
        [Option(longName: "store", Required = true, HelpText = "Which store to upload to")]
        public string? Store { get; set; }

        [Option(longName: "description", Required = true, HelpText = "Textual description of upload")]
        public string? Description { get; set; }

        [Option(longName: "build-id", Required = true, HelpText = "Build ID for upload")]
        public string? BuildId { get; set; }

        [Value(0, MetaName="<pattern1 pattern2 pattern3 ...>", Min = 1, HelpText = "Wildcard patterns for files to upload, like 'folder1/*.pdb'")]
        public IEnumerable<string>? Patterns { get; set; }
    }

    [Verb("list-stores", HelpText = "List all stores")]
    public class ListStoresOptions : CommandLineOptions
    {
    }
}