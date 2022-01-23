using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class CreateStore
    {
        public static async Task<int> DoCreateStore(CreateStoreOptions options)
        {
            bool created = await ClientAPI.CreateStore.DoCreateStore(options.ServiceURL, options.Email, options.PAT, options.Store);
            if (created)
            {
                Console.WriteLine($"Created new store: {options.Store}");
                return 0;
            }
            else
            {
                Console.Error.WriteLine($"Store {options.Store} already exists");
                return 1;
            }
        }
    }
}