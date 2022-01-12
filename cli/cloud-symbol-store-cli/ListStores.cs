using Microsoft.Extensions.FileSystemGlobbing;

namespace CLI
{
    public static class ListStores
    {
        public static int DoListStores(ListStoresOptions options)
        {
            IEnumerable<string> stores = ClientAPI.ListStores.DoListStores(options.ServiceURL, options.Email, options.PAT);
            Console.WriteLine("Stores:");
            foreach (string store in stores) {
                Console.WriteLine($"  {store}");
            }

            return 0;
        }
    }
}