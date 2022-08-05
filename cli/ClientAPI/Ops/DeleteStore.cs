using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class DeleteStore
    {
        public class DeleteStoreException : Exception
        {
            public DeleteStoreException(string message) : base(message) { }
        }

        public static async Task<bool> DoDeleteStore(string ServiceURL, string Email, string PAT, string StoreId) {

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

            try {
                BackendAPI.Client.ApiResponse<Object> deleteStoreResponse = await api.DeleteStoreWithHttpInfoAsync(StoreId);
                if (deleteStoreResponse.ErrorText != null)
                    throw new DeleteStoreException(deleteStoreResponse.ErrorText);
                else
                    return true;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    return false;
                else
                    throw;
            }
        }
    }
}
