using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class ListUploads
    {
        public static async Task<int> DoListUploads(ListUploadsOptions options)
        {
            IEnumerable<string> uploads = await ClientAPI.ListUploads.DoListUploads(options.ServiceURL, options.Email, options.PAT, options.Store);
            Console.WriteLine($"Uploads in store {options.Store}:");
            foreach (string upload in uploads) {
                Console.WriteLine($"  {upload}");
            }

            return 0;
        }
    }
}