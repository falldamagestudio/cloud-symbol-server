using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListFiles
    {
        public static async Task<BackendAPI.Model.GetStoreFilesResponse> DoListFiles(string ServiceURL, string Email, string PAT, string store, int offset, int limit) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoreFilesAsync(store, offset, limit);
        }
    }
}
