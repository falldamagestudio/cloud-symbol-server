using System;

namespace ClassLib
{
    public class Class1
    {
        public static void Upload(string ServiceURL, string Email, string PAT) {
            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            BackendAPI.Model.UploadTransactionRequest uploadTransactionRequest = new BackendAPI.Model.UploadTransactionRequest();
            BackendAPI.Model.UploadTransactionResponse uploadTransactionResponse = api.CreateTransaction(uploadTransactionRequest);
            Console.WriteLine(uploadTransactionResponse);

            BackendAPI.Model.GetTransactionResponse getTransactionResponse = api.GetTransaction(uploadTransactionResponse.Id);
            Console.WriteLine(getTransactionResponse);
        }
    }
}
