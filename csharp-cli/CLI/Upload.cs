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

            string[] Patterns = new string[] { "*.pdb" };

            var Files = FindMatchingFiles(Patterns);

            ClientAPI.Ops.Upload(options.ServiceURL, options.Email, options.PAT, Files);

            return 0;
        }
    }
}