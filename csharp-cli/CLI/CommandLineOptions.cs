using CommandLine;

namespace CLI
{
    public class CommandLineOptions
    {
        // [Value(index: 0, Required = true, HelpText = "Image file Path upload")]
        // public string UploadFilePath { get; set; }

        [Option(longName: "service-url", Required = false, HelpText = "Service URL, like http://localhos:8084", Default = "http://localhost:8084")]
        public string? ServiceURL { get; set; }

        [Option(longName: "email", Required = false, HelpText = "Authentication email", Default = "")]
        public string? Email { get; set; }

        [Option(longName: "pat", Required = false, HelpText = "Authentication Personal Access Token", Default = "")]
        public string? PAT { get; set; }
    }
}