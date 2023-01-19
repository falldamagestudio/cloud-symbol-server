using System.Threading.Tasks;

namespace ClientAPI
{
    public class GetUpload
    {
        public static async Task<BackendAPI.Model.GetStoreUploadResponse> DoGetUpload(string ServiceURL, string Email, string PAT, string store, int upload) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            return await backendApiWrapper.GetStoreUploadAsync(store, upload);
        }
    }
}
