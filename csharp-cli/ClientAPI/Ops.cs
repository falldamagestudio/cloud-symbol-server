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
                        Console.WriteLine($"Upload failed with status code {rrr.StatusCode}; content = {rrr.Content}");
                        throw new ApplicationException($"Upload failed with status code {rrr.StatusCode}; content = {rrr.Content}");
                    }
                }
            }
        }

        public static void Upload(string ServiceURL, string Email, string PAT, IEnumerable<string> Files) {

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            IEnumerable<FileWithHash> filesWithHashes = GetFilesWithHashes(Files);
            BackendAPI.Model.UploadTransactionRequest uploadTransactionRequest = CreateUploadTransactionRequest("", "", filesWithHashes);
            BackendAPI.Model.UploadTransactionResponse uploadTransactionResponse = api.CreateTransaction(uploadTransactionRequest);
            if (uploadTransactionResponse == null)
                throw new ApplicationException("Upload transaction failed");

            UploadMissingFiles(uploadTransactionResponse, filesWithHashes);
        }
    }
}
