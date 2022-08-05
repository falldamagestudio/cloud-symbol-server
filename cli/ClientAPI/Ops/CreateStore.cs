using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class CreateStore
    {
        public static async Task<bool> DoCreateStore(string ServiceURL, string Email, string PAT, string StoreId) {

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

            return await ApiWrapper.CreateStoreAsync(api, StoreId);
        }
    }
}
