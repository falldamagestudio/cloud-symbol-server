using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class DeleteStore
    {
        public static async Task<int> DoDeleteStore(DeleteStoreOptions options)
        {
            bool deleted = await ClientAPI.DeleteStore.DoDeleteStore(options.ServiceURL, options.Email, options.PAT, options.Store);
            if (deleted)
            {
                Console.WriteLine($"Deleted existing store: {options.Store}");
                return 0;
            }
            else
            {
                Console.WriteLine($"Store {options.Store} does not exist");
                return 1;
            }
        }
    }
}