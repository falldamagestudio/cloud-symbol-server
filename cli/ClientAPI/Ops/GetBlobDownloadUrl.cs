using System.Threading.Tasks;

namespace ClientAPI
{
    public class GetBlobDownloadUrl
    {
        public static async Task<BackendAPI.Model.GetStoreFileBlobDownloadUrlResponse> DoGetBlobDownloadUrl(string ServiceURL, string Email, string PAT, string store, string file, string blob) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoreFileBlobDownloadUrlAsync(store, file, blob);
        }
    }
}
