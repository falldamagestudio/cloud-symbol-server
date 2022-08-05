using System;
using System.Net;
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

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

            BackendAPI.Model.GetStoreUploadIdsResponse getStoreUploadIdsResponse = await ApiWrapper.GetStoreUploadIdsAsync(api, store);

            List<StoreUpload> uploads = new List<StoreUpload>();

            foreach (string uploadId in getStoreUploadIdsResponse) {
                BackendAPI.Model.GetStoreUploadResponse getStoreUploadResponse = await ApiWrapper.GetStoreUploadAsync(api, store, uploadId);

                uploads.Add(new StoreUpload(uploadId: uploadId, upload: getStoreUploadResponse));
            }

            return uploads;
        }
    }
}
