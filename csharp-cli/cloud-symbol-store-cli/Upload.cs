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
            var Files = FindMatchingFiles(options.Patterns!);

            ClientAPI.Ops.Upload(options.ServiceURL, options.Email, options.PAT, options.Description, options.BuildId, Files);

            return 0;
        }
    }
}