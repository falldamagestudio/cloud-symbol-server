using System;

namespace ClassLib
{
    public class Class1
    {
        public static void Upload() {
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi("http://testuser:testpat@localhost:8084");
            BackendAPI.Model.UploadTransactionRequest uploadTransactionRequest = new BackendAPI.Model.UploadTransactionRequest();
            BackendAPI.Model.UploadTransactionResponse uploadTransactionResponse = api.CreateTransaction(uploadTransactionRequest);
            Console.WriteLine(uploadTransactionResponse);
        }
    }
}
