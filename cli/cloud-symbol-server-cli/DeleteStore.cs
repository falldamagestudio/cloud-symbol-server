namespace CLI
{
    public static class DeleteStore
    {
        public static async Task<int> DoDeleteStore(GlobalOptions globalOptions, string store)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            bool deleted = await ClientAPI.DeleteStore.DoDeleteStore(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat, store);
            if (deleted)
            {
                Console.WriteLine($"Deleted existing store: {store}");
                return 0;
            }
            else
            {
                Console.Error.WriteLine($"Store {store} does not exist");
                return 1;
            }
        }
    }
}