using System;
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

        public static async Task<IEnumerable<string>> DoListUploads(string ServiceURL, string Email, string PAT, string store) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoreUploadsResponse> getStoreUploadsResponse = await api.GetStoreUploadsWithHttpInfoAsync(store);
            if (getStoreUploadsResponse.ErrorText != null)
                throw new ListUploadsException(getStoreUploadsResponse.ErrorText);

            return getStoreUploadsResponse.Data;
        }
    }
}
