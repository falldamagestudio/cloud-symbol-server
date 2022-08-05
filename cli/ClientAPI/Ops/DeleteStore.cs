using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class DeleteStore
    {
        public static async Task<bool> DoDeleteStore(string ServiceURL, string Email, string PAT, string StoreId) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.DeleteStoreAsync(StoreId);
        }
    }
}
