namespace CLI
{
    public static class ListStores
    {
        public static async Task<int> DoListStores(GlobalOptions globalOptions)
        {
            if (!globalOptions.Validate()) {
                Console.Error.WriteLine("Please set service-url, email and pat via config.json or commandline options");
                return 1;
            }

            IEnumerable<string> stores = await ClientAPI.ListStores.DoListStores(globalOptions.ServiceUrl, globalOptions.Email, globalOptions.Pat);
            Console.WriteLine("Stores:");
            foreach (string store in stores) {
                Console.WriteLine($"  {store}");
            }

            return 0;
        }
    }
}