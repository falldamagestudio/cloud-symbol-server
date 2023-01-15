using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListHashes
    {
        public static async Task<BackendAPI.Model.GetStoreFileHashesResponse> DoListHashes(string ServiceURL, string Email, string PAT, string store, string file, int offset, int limit) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoreFileHashesAsync(store, file, offset, limit);
        }
    }
}
