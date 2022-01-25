using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class CreateStore
    {
        public static async Task<int> DoCreateStore(GlobalOptions globalOptions, string store)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            bool created = await ClientAPI.CreateStore.DoCreateStore(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store);
            if (created)
            {
                Console.WriteLine($"Created new store: {store}");
                return 0;
            }
            else
            {
                Console.Error.WriteLine($"Store {store} already exists");
                return 1;
            }
        }
    }
}