using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class DeleteStore
    {
        public static async Task<bool> DoDeleteStore(string ServiceURL, string Email, string PAT, string StoreId) {

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

            return await ApiWrapper.DeleteStoreAsync(api, StoreId);
        }
    }
}
