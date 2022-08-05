using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class BackendApiWrapper
    {
        BackendAPI.Api.DefaultApi backendApi;

        public BackendApiWrapper(string ServiceURL, string Email, string PAT) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            backendApi = new BackendAPI.Api.DefaultApi(config);
        }

        public class ApiException : ClientAPIException
        {
            public ApiException(string message) : base(message) { }
        }

        public class CreateStoreUploadException : ClientAPIException
        {
            public CreateStoreUploadException(string message) : base(message) { }
        }

        public async Task<BackendAPI.Model.CreateStoreUploadResponse> CreateStoreUploadAsync(string store, BackendAPI.Model.CreateStoreUploadRequest request) {

            try {
                BackendAPI.Client.ApiResponse<BackendAPI.Model.CreateStoreUploadResponse> response = await backendApi.CreateStoreUploadWithHttpInfoAsync(store, request);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return response.Data;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new CreateStoreUploadException($"Store {store} does not exist");
                else
                    throw;
            }
        }
 
        public class MarkStoreUploadCompletedException : ClientAPIException
        {
            public MarkStoreUploadCompletedException(string message) : base(message) { }
        }

        public async Task MarkStoreUploadCompletedAsync(string store, string uploadId) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await backendApi.MarkStoreUploadCompletedWithHttpInfoAsync(uploadId, store);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new MarkStoreUploadCompletedException($"Store {store}  / upload {uploadId} does not exist");
                else
                    throw;
            }
        }

        public class MarkStoreUploadAbortedException : ClientAPIException
        {
            public MarkStoreUploadAbortedException(string message) : base(message) { }
        }

        public async Task MarkStoreUploadAbortedAsync(string store, string uploadId) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await backendApi.MarkStoreUploadAbortedWithHttpInfoAsync(uploadId, store);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new MarkStoreUploadAbortedException($"Store {store} / upload {uploadId} does not exist");
                else
                    throw;
            }
        }

        public class MarkStoreUploadFileUploadedException : ClientAPIException
        {
            public MarkStoreUploadFileUploadedException(string message) : base(message) { }
        }

        public async Task MarkStoreUploadFileUploadedAsync(string store, string uploadId, int fileId) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await backendApi.MarkStoreUploadFileUploadedWithHttpInfoAsync(uploadId, store, fileId);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new MarkStoreUploadFileUploadedException($"Store {store}  / upload {uploadId} / file {fileId} does not exist");
                else
                    throw;
            }
        }

        public class GetStoreUploadIdsException : ClientAPIException
        {
            public GetStoreUploadIdsException(string message) : base(message) { }
        }

        public async Task<BackendAPI.Model.GetStoreUploadIdsResponse> GetStoreUploadIdsAsync(string store) {

            try {
                BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadIdsResponse> response = await backendApi.GetStoreUploadIdsWithHttpInfoAsync(store);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return response.Data;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new GetStoreUploadIdsException($"Store {store} does not exist");
                else
                    throw;
            }
        }

        public class GetStoreUploadException : ClientAPIException
        {
            public GetStoreUploadException(string message) : base(message) { }
        }

        public async Task<BackendAPI.Model.GetStoreUploadResponse> GetStoreUploadAsync(string store, string uploadId) {

            try {
                BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadResponse> response = await backendApi.GetStoreUploadWithHttpInfoAsync(uploadId, store);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return response.Data;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new GetStoreUploadException($"UploadId {uploadId} does not exist in store {store}");
                else
                    throw;
            }
        }

        public async Task<bool> CreateStoreAsync(string store) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await backendApi.CreateStoreWithHttpInfoAsync(store);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return true;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.Conflict)
                    return false;
                else
                    throw;
            }
        }

        public async Task<bool> DeleteStoreAsync(string store) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await backendApi.DeleteStoreWithHttpInfoAsync(store);
                if (response.ErrorText != null)
                    throw new ApiException(response.ErrorText);
                return true;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    return false;
                else
                    throw;
            }
        }

        public async Task<BackendAPI.Model.GetStoresResponse> GetStoresAsync() {

            BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoresResponse> response = await backendApi.GetStoresWithHttpInfoAsync();
            if (response.ErrorText != null)
                throw new ApiException(response.ErrorText);
            return response.Data;
        }
    }
}