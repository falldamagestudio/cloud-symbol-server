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

        public static int DoUpload(UploadOptions options)
        {
            IReadOnlyCollection<string> files = FindMatchingFiles(options.Patterns!).ToList();

            if (!files.Any()) {
                Console.WriteLine($"No files matching patterns: [{String.Join(", ", options.Patterns!)}], upload skipped");
            } else {
                Console.WriteLine("Uploading to Cloud Symbol Server...");
                Console.WriteLine($"  Store: {options.Store}");
                Console.WriteLine($"  Description: {options.Description}");
                Console.WriteLine($"  Build ID: {options.BuildId}");
                Console.WriteLine("  Files:");
                foreach (string file in files) {
                    Console.WriteLine($"    {file}");
                }

                try {
                    Progress<ClientAPI.Ops.UploadProgress> uploadProgress = new Progress<ClientAPI.Ops.UploadProgress>();
                    uploadProgress.ProgressChanged += (s, e) => Console.WriteLine($"Progress: {e.State} {e.FileName}");
                    ClientAPI.Ops.Upload(options.ServiceURL, options.Email, options.PAT, options.Store, options.Description, options.BuildId, files, uploadProgress);
                    Console.WriteLine("Upload done.");
                } catch (ClientAPI.Ops.UploadException e) {
                    Console.WriteLine($"Upload failed: {e.Message}");
                }
            }

            return 0;
        }
    }
}