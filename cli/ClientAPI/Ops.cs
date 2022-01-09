using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using BackendAPI.Model;
using RestSharp;

namespace ClientAPI
{
    public class Ops
    {
        private struct FileWithHash
        {
        	public string FileWithPath;
        	public string FileWithoutPath;
        	public string Hash;
        }

        private static IEnumerable<FileWithHash> GetFilesWithHashes(IEnumerable<string> fileNames)
        {
            IEnumerable<FileWithHash> filesWithHashes = fileNames.Select(fileName => new FileWithHash {
                FileWithPath = fileName,
                FileWithoutPath = Path.GetFileName(fileName),
                Hash = PDBParser.GetHash(fileName)
            });

            return filesWithHashes;
        }

        private static BackendAPI.Model.UploadTransactionRequest CreateUploadTransactionRequest(string description, string buildId, IEnumerable<FileWithHash> FileWithHash)
        {
            BackendAPI.Model.UploadTransactionRequest request = new BackendAPI.Model.UploadTransactionRequest(
                description: description,
                buildId: buildId,
                files: FileWithHash.Select(fileWithHash => new BackendAPI.Model.UploadFileRequest(
                    fileName: fileWithHash.FileWithoutPath,
                    hash: fileWithHash.Hash
                )).ToList()
            );

            return request;
        }

        public class UploadException : Exception
        {
            public UploadException(string message) : base(message) { }
        }

        private static void UploadMissingFiles(BackendAPI.Model.UploadTransactionResponse uploadTransactionResponse, IEnumerable<FileWithHash> filesWithHashes)
        {
            if (uploadTransactionResponse.Files != null) {
                foreach (BackendAPI.Model.UploadFileResponse uploadFileResponse in uploadTransactionResponse.Files) {

                    FileWithHash fileWithHash = filesWithHashes.First(fwh => 
                        fwh.FileWithoutPath == uploadFileResponse.FileName && fwh.Hash == uploadFileResponse.Hash);

                    Console.WriteLine($"Uploading file {fileWithHash.FileWithPath}...");

                    RestClient restClient = new RestClient();
                    RestRequest request = new RestRequest(uploadFileResponse.Url, Method.PUT);
                    IRestResponse rrr = restClient.Execute(request);

                    if (!rrr.IsSuccessful) {
                        throw new UploadException($"Upload failed with status code {rrr.StatusCode}; content = {rrr.Content}");
                    }
                }
            }
        }

        public static void Upload(string ServiceURL, string Email, string PAT, string description, string buildId, IEnumerable<string> Files) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            IEnumerable<FileWithHash> filesWithHashes = GetFilesWithHashes(Files);
            BackendAPI.Model.UploadTransactionRequest uploadTransactionRequest = CreateUploadTransactionRequest(description, buildId, filesWithHashes);
            BackendAPI.Client.ApiResponse<BackendAPI.Model.UploadTransactionResponse> uploadTransactionResponse = api.CreateTransactionWithHttpInfo(uploadTransactionRequest);
            if (uploadTransactionResponse.ErrorText != null)
                throw new UploadException(uploadTransactionResponse.ErrorText);

            UploadMissingFiles(uploadTransactionResponse.Data, filesWithHashes);
        }
    }
}
