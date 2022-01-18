using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class ListStores
    {
        public class ListStoresException : Exception
        {
            public ListStoresException(string message) : base(message) { }
        }

        public static async Task<IEnumerable<string>> DoListStores(string ServiceURL, string Email, string PAT) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            BackendAPI.Client.ApiResponse<List<string>> getStoresResponse = await api.GetStoresWithHttpInfoAsync();
            if (getStoresResponse.ErrorText != null)
                throw new ListStoresException(getStoresResponse.ErrorText);

            return getStoresResponse.Data;
        }
    }
}
