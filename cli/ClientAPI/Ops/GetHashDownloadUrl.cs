using System.Threading.Tasks;

namespace ClientAPI
{
    public class GetHashDownloadUrl
    {
        public static async Task<BackendAPI.Model.GetStoreFileHashDownloadUrlResponse> DoGetHashDownloadUrl(string ServiceURL, string Email, string PAT, string store, string file, string hash) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoreFileHashDownloadUrlAsync(store, file, hash);
        }
    }
}
