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

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

            BackendAPI.Client.ApiResponse<BackendAPI.Model.GetStoresResponse> getStoresResponse = await api.GetStoresWithHttpInfoAsync();
            if (getStoresResponse.ErrorText != null)
                throw new ListStoresException(getStoresResponse.ErrorText);

            return getStoresResponse.Data;
        }
    }
}
