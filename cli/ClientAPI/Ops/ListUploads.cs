using System;
using System.Net;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListUploads
    {
        public class ListUploadsException : Exception
        {
            public ListUploadsException(string message) : base(message) { }
        }

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

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadIdsResponse> getStoreUploadIdsResponse;
            try {
                getStoreUploadIdsResponse = await api.GetStoreUploadIdsWithHttpInfoAsync(store);
                if (getStoreUploadIdsResponse.ErrorText != null)
                    throw new ListUploadsException(getStoreUploadIdsResponse.ErrorText);
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new ListUploadsException($"Store {store} does not exist");
                else
                    throw;
            }

            List<StoreUpload> uploads = new List<StoreUpload>();

            foreach (string uploadId in getStoreUploadIdsResponse.Data) {
                BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadResponse> getStoreUploadResponse;
                try {
                    getStoreUploadResponse = await api.GetStoreUploadWithHttpInfoAsync(uploadId, store);
                    if (getStoreUploadResponse.ErrorText != null)
                        throw new ListUploadsException(getStoreUploadResponse.ErrorText);
                } catch (BackendAPI.Client.ApiException apiException) {
                    if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                        throw new ListUploadsException($"UploadId {uploadId} does not exist in store {store}");
                    else
                        throw;
                }

                uploads.Add(new StoreUpload(uploadId: uploadId, upload: getStoreUploadResponse.Data));
            }

            return uploads;
        }
    }
}
