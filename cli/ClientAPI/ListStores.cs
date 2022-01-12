using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using BackendAPI.Model;
using RestSharp;

namespace ClientAPI
{
    public class ListStores
    {
        public class ListStoresException : Exception
        {
            public ListStoresException(string message) : base(message) { }
        }

        public static IEnumerable<string> DoListStores(string ServiceURL, string Email, string PAT) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            BackendAPI.Client.ApiResponse<List<string>> getStoresResponse = api.GetStoresWithHttpInfo();
            if (getStoresResponse.ErrorText != null)
                throw new ListStoresException(getStoresResponse.ErrorText);

            return getStoresResponse.Data;
        }
    }
}
