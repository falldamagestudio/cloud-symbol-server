using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class CreateStore
    {
        public class CreateStoreException : Exception
        {
            public CreateStoreException(string message) : base(message) { }
        }

        public static async Task<bool> DoCreateStore(string ServiceURL, string Email, string PAT, string StoreId) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            try {
                BackendAPI.Client.ApiResponse<Object> createStoreResponse = await api.CreateStoreWithHttpInfoAsync(StoreId);
                if (createStoreResponse.ErrorText != null)
                    throw new CreateStoreException(createStoreResponse.ErrorText);
                else
                    return true;
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.Conflict)
                    return false;
                else
                    throw;
            }
        }
    }
}
