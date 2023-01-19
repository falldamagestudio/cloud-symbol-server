using System.Threading.Tasks;

namespace ClientAPI
{
    public class ExpireStoreUpload
    {
        public static async Task DoExpireStoreUpload(string ServiceURL, string Email, string PAT, string storeId, int uploadId) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            await backendApiWrapper.ExpireStoreUploadAsync(storeId, uploadId);
        }
    }
}
