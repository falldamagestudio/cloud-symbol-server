using System;
using System.Collections.Generic;

namespace ClassLib
{
    public class Class1
    {
        public static void Upload(string ServiceURL, string Email, string PAT, IEnumerable<string> Files) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            BackendAPI.Model.UploadTransactionRequest uploadTransactionRequest = new BackendAPI.Model.UploadTransactionRequest();

            uploadTransactionRequest.Files = new List<BackendAPI.Model.UploadFileRequest>();

            foreach (string FileName in Files) {
                uploadTransactionRequest.Files.Add(new BackendAPI.Model.UploadFileRequest{
                    FileName = FileName,
                    Hash = "blah"
                });
            }

            BackendAPI.Model.UploadTransactionResponse uploadTransactionResponse = api.CreateTransaction(uploadTransactionRequest);
            Console.WriteLine(uploadTransactionResponse);

            BackendAPI.Model.GetTransactionResponse getTransactionResponse = api.GetTransaction(uploadTransactionResponse.Id);
            Console.WriteLine(getTransactionResponse);
        }
    }
}
