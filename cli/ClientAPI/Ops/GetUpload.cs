using System;
using System.Net;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class GetUpload
    {
        public static async Task<BackendAPI.Model.GetStoreUploadResponse> DoGetUpload(string ServiceURL, string Email, string PAT, string store, string upload) {

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

            return await ApiWrapper.GetStoreUploadAsync(api, store, upload);
        }
    }
}
