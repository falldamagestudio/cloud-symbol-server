using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public static class ApiWrapper
    {
        public class ApiException : ClientAPIException
        {
            public ApiException(string message) : base(message) { }
        }

        public class CreateStoreUploadException : ClientAPIException
        {
            public CreateStoreUploadException(string message) : base(message) { }
        }

        public static async Task<BackendAPI.Model.CreateStoreUploadResponse> CreateStoreUploadAsync(BackendAPI.Api.DefaultApi api, string store, BackendAPI.Model.CreateStoreUploadRequest request) {

            try {
                BackendAPI.Client.ApiResponse<BackendAPI.Model.CreateStoreUploadResponse> response = await api.CreateStoreUploadWithHttpInfoAsync(store, request);
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

        public static async Task MarkStoreUploadCompletedAsync(BackendAPI.Api.DefaultApi api, string store, string uploadId) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await api.MarkStoreUploadCompletedWithHttpInfoAsync(uploadId, store);
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

        public static async Task MarkStoreUploadAbortedAsync(BackendAPI.Api.DefaultApi api, string store, string uploadId) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await api.MarkStoreUploadAbortedWithHttpInfoAsync(uploadId, store);
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

        public static async Task MarkStoreUploadFileUploadedAsync(BackendAPI.Api.DefaultApi api, string store, string uploadId, int fileId) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await api.MarkStoreUploadFileUploadedWithHttpInfoAsync(uploadId, store, fileId);
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

        public static async Task<BackendAPI.Model.GetStoreUploadIdsResponse> GetStoreUploadIdsAsync(BackendAPI.Api.DefaultApi api, string store) {

            try {
                BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadIdsResponse> response = await api.GetStoreUploadIdsWithHttpInfoAsync(store);
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

        public static async Task<BackendAPI.Model.GetStoreUploadResponse> GetStoreUploadAsync(BackendAPI.Api.DefaultApi api, string store, string uploadId) {

            try {
                BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadResponse> response = await api.GetStoreUploadWithHttpInfoAsync(uploadId, store);
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

        public static async Task<bool> CreateStoreAsync(BackendAPI.Api.DefaultApi api, string store) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await api.CreateStoreWithHttpInfoAsync(store);
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

        public static async Task<bool> DeleteStoreAsync(BackendAPI.Api.DefaultApi api, string store) {

            try {
                BackendAPI.Client.ApiResponse<object> response = await api.DeleteStoreWithHttpInfoAsync(store);
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

        public static async Task<BackendAPI.Model.GetStoresResponse> GetStoresAsync(BackendAPI.Api.DefaultApi api) {

            BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoresResponse> response = await api.GetStoresWithHttpInfoAsync();
            if (response.ErrorText != null)
                throw new ApiException(response.ErrorText);
            return response.Data;
        }
    }
}