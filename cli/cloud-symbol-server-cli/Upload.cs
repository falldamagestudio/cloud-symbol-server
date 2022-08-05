using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class Upload
    {
        public static IEnumerable<string> FindMatchingFiles(IEnumerable<string> Patterns) {

            Matcher matcher = new Matcher();
            foreach (string Pattern in Patterns) {
                matcher.AddInclude(Pattern);
            }

            return matcher.GetResultsInFullPath(".");
        }

        public static async Task<int> DoUpload(GlobalOptions globalOptions, string description, string buildId, string store, string[] patterns)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            IReadOnlyCollection<string> files = FindMatchingFiles(patterns).ToList();

            if (!files.Any()) {
                Console.WriteLine($"No files matching patterns: [{String.Join(", ", patterns)}], upload skipped");
            } else {
                Console.WriteLine("Uploading to Cloud Symbol Server...");
                Console.WriteLine($"  Store: {store}");
                Console.WriteLine($"  Description: {description}");
                Console.WriteLine($"  Build ID: {buildId}");
                Console.WriteLine("  Files:");
                foreach (string file in files) {
                    Console.WriteLine($"    {file}");
                }

                try {
                    Progress<ClientAPI.Ops.UploadProgress> uploadProgress = new Progress<ClientAPI.Ops.UploadProgress>();
                    uploadProgress.ProgressChanged += (s, e) => Console.WriteLine($"Progress: {e.State} {e.FileName}");
                    await ClientAPI.Ops.Upload(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store, description, buildId, files, uploadProgress);
                    Console.WriteLine("Upload done.");
                } catch (ClientAPI.ClientAPIException e) {
                    Console.Error.WriteLine($"Upload failed: {e.Message}");
                    return 1;
                }
            }

            return 0;
        }
    }
}