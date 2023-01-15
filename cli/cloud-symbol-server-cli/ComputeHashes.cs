using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class ComputeHashes
    {
        public static IEnumerable<string> FindMatchingFiles(IEnumerable<string> Patterns) {

            Matcher matcher = new Matcher();
            foreach (string Pattern in Patterns) {
                matcher.AddInclude(Pattern);
            }

            return matcher.GetResultsInFullPath(".");
        }

        public static int DoComputeHashes(string[] patterns)
        {
            IReadOnlyCollection<string> files = FindMatchingFiles(patterns).ToList();

            if (!files.Any()) {
                Console.WriteLine($"No files matching patterns: [{String.Join(", ", patterns)}], hash-files skipped");
            } else {
                IEnumerable<ClientAPI.HashFiles.FileWithHash> filesWithHashes = ClientAPI.HashFiles.DoHashFiles(files);
                Console.WriteLine("  Files and hashes:");
                foreach (ClientAPI.HashFiles.FileWithHash fileWithHash in filesWithHashes) {
                    Console.WriteLine($"    {fileWithHash.FileWithPath}: {fileWithHash.Hash}");
                }
            }

            return 0;
        }
    }
}