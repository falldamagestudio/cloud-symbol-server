using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListBlobs
    {
        public static async Task<BackendAPI.Model.GetStoreFileBlobsResponse> DoListBlobs(string ServiceURL, string Email, string PAT, string store, string file, int offset, int limit) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoreFileBlobsAsync(store, file, offset, limit);
        }
    }
}
