using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListUploads
    {
        public static async Task<BackendAPI.Model.GetStoreUploadsResponse> DoListUploads(string ServiceURL, string Email, string PAT, string store, int offset, int limit) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoreUploadsAsync(store, offset, limit);
        }
    }
}
