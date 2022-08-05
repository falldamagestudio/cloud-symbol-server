using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListStores
    {
        public static async Task<IEnumerable<string>> DoListStores(string ServiceURL, string Email, string PAT) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoresAsync();
        }
    }
}
