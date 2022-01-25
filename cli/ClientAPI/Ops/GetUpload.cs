using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class GetUpload
    {
        public class GetUploadException : Exception
        {
            public GetUploadException(string message) : base(message) { }
        }

        public static async Task<BackendAPI.Model.GetStoreUploadResponse> DoGetUpload(string ServiceURL, string Email, string PAT, string store, string upload) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadResponse> getStoreUploadResponse;
            try {
                getStoreUploadResponse = await api.GetStoreUploadWithHttpInfoAsync(upload, store);
                if (getStoreUploadResponse.ErrorText != null)
                    throw new GetUploadException(getStoreUploadResponse.ErrorText);
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new GetUploadException($"Upload {upload} does not exist in store {store}");
                else
                    throw;
            }

            return getStoreUploadResponse.Data;
        }
    }
}
