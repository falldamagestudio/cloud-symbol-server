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

        public static void DoUpload(string ServiceURL, string Email, string PAT, IEnumerable<string> Patterns) {
            var Files = FindMatchingFiles(Patterns);

            ClientAPI.Ops.Upload(ServiceURL, Email, PAT, Files);
        }
    }
}