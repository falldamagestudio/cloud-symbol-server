using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class ComputeFileMetadata
    {
        public static IEnumerable<string> FindMatchingFiles(IEnumerable<string> Patterns) {

            Matcher matcher = new Matcher();
            foreach (string Pattern in Patterns) {
                matcher.AddInclude(Pattern);
            }

            return matcher.GetResultsInFullPath(".");
        }

        public static int DoComputeFileMetadata(string[] patterns)
        {
            IReadOnlyCollection<string> files = FindMatchingFiles(patterns).ToList();

            if (!files.Any()) {
                Console.WriteLine($"No files matching patterns: [{String.Join(", ", patterns)}], hash-files skipped");
            } else {
                IEnumerable<ClientAPI.ComputeFileMetadata.FileWithMetadata> filesWithMetadata = ClientAPI.ComputeFileMetadata.DoComputeFileMedatadata(files);
                Console.WriteLine("  Files and metadata:");
                foreach (ClientAPI.ComputeFileMetadata.FileWithMetadata fileWithMetadata in filesWithMetadata) {
                    Console.WriteLine($"    File {fileWithMetadata.FileWithPath}:");
                    Console.WriteLine($"      BlobIdentifier: {fileWithMetadata.BlobIdentifier}");
                    Console.WriteLine($"      Type: {fileWithMetadata.Type}");
                    Console.WriteLine($"      Size: {fileWithMetadata.Size}");
                    Console.WriteLine($"      Content SHA256 Hash: {fileWithMetadata.ContentHash}");
                }
            }

            return 0;
        }
    }
}