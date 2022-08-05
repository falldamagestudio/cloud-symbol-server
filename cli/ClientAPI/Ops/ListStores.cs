using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListStores
    {
        public static async Task<IEnumerable<string>> DoListStores(string ServiceURL, string Email, string PAT) {

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

            return await ApiWrapper.GetStoresAsync(api);
        }
    }
}
