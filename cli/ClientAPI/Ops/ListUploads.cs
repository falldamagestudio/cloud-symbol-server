using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListUploads
    {
        public class StoreUpload
        {
            public string UploadId;
            public BackendAPI.Model.GetStoreUploadResponse Upload;

            public StoreUpload(string uploadId, BackendAPI.Model.GetStoreUploadResponse upload) {
                UploadId = uploadId;
                Upload = upload;
            }
        }

        public static async Task<IEnumerable<StoreUpload>> DoListUploads(string ServiceURL, string Email, string PAT, string store) {

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            BackendAPI.Model.GetStoreUploadIdsResponse getStoreUploadIdsResponse = await backendApiWrapper.GetStoreUploadIdsAsync(store);

            List<StoreUpload> uploads = new List<StoreUpload>();

            foreach (string uploadId in getStoreUploadIdsResponse) {
                BackendAPI.Model.GetStoreUploadResponse getStoreUploadResponse = await backendApiWrapper.GetStoreUploadAsync(store, uploadId);

                uploads.Add(new StoreUpload(uploadId: uploadId, upload: getStoreUploadResponse));
            }

            return uploads;
        }
    }
}
