using System;

namespace ClassLib
{
    public class Class1
    {
        public static void Upload() {
            Org.OpenAPITools.Api.DefaultApi api = new Org.OpenAPITools.Api.DefaultApi("http://testuser:testpat@localhost:8084");
            Org.OpenAPITools.Model.UploadTransactionRequest uploadTransactionRequest = new Org.OpenAPITools.Model.UploadTransactionRequest();
            Org.OpenAPITools.Model.UploadTransactionResponse uploadTransactionResponse = api.CreateTransaction(uploadTransactionRequest);
            Console.WriteLine(uploadTransactionResponse);
        }
    }
}
