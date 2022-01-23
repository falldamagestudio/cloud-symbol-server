using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;
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

        private static string GetHash(string fileName)
        {
            string pdbHash = PDBParser.GetHash(fileName);
            if (pdbHash != null)
                return pdbHash;

            string peHash = PEParser.GetHash(fileName);
            if (peHash != null)
                return peHash;

            throw new ApplicationException($"File {fileName} is not of a recognized format");
        }

        private static IEnumerable<FileWithHash> GetFilesWithHashes(IEnumerable<string> fileNames)
        {
            IEnumerable<FileWithHash> filesWithHashes = fileNames.Select(fileName => new FileWithHash {
                FileWithPath = fileName,
                FileWithoutPath = Path.GetFileName(fileName),
                Hash = GetHash(fileName)
            }).Where(fileWithHash => fileWithHash.Hash != null);

            return filesWithHashes;
        }

        private static BackendAPI.Model.CreateStoreUploadRequest CreateStoreUploadRequest(string description, string buildId, IEnumerable<FileWithHash> FileWithHash)
        {
            BackendAPI.Model.CreateStoreUploadRequest request = new BackendAPI.Model.CreateStoreUploadRequest(
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

        public struct UploadProgress {
            public enum StateEnum { LocalValidation, CreatingUploadEntry, UploadingMissingFiles, UploadingMissingFile, Done };

            public StateEnum State;
            public string FileName;
        }

        private static HttpClient HttpClient = new HttpClient();

        private static async Task UploadMissingFiles(BackendAPI.Model.CreateStoreUploadResponse createStoreUploadResponse, IEnumerable<FileWithHash> filesWithHashes, IProgress<UploadProgress> progress)
        {
            if (createStoreUploadResponse.Files != null) {
                foreach (BackendAPI.Model.UploadFileResponse uploadFileResponse in createStoreUploadResponse.Files) {

                    FileWithHash fileWithHash = filesWithHashes.First(fwh => 
                        fwh.FileWithoutPath == uploadFileResponse.FileName && fwh.Hash == uploadFileResponse.Hash);

                    if (progress != null)
                        progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFile, FileName = fileWithHash.FileWithPath });

                    byte[] content = File.ReadAllBytes(fileWithHash.FileWithPath);

                    HttpResponseMessage response = await HttpClient.PutAsync(uploadFileResponse.Url, new ByteArrayContent(content));

                    if (!response.IsSuccessStatusCode) {
                        throw new UploadException($"Upload failed with status code {response.StatusCode}; content = {response.Content}");
                    }
                }
            }
        }

        public static async Task Upload(string ServiceURL, string Email, string PAT, string store, string description, string buildId, IReadOnlyCollection<string> Files, IProgress<UploadProgress> progress) {

            if (!Files.Any()) {
                throw new ArgumentException($"Upload requires at least one filename", nameof(Files));
            }

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.LocalValidation });

            BackendAPI.Client.Configuration config = new BackendAPI.Client.Configuration();
            config.BasePath = ServiceURL;
            config.Username = Email;
            config.Password = PAT;
            BackendAPI.Api.DefaultApi api = new BackendAPI.Api.DefaultApi(config);

            IEnumerable<FileWithHash> filesWithHashes = GetFilesWithHashes(Files);

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.CreatingUploadEntry });

            BackendAPI.Model.CreateStoreUploadRequest createStoreUploadRequest = CreateStoreUploadRequest(description, buildId, filesWithHashes);
            BackendAPI.Client.ApiResponse<BackendAPI.Model.CreateStoreUploadResponse> createStoreUploadResponse = api.CreateStoreUploadWithHttpInfo(store, createStoreUploadRequest);
            if (createStoreUploadResponse.ErrorText != null)
                throw new UploadException(createStoreUploadResponse.ErrorText);

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFiles });

            await UploadMissingFiles(createStoreUploadResponse.Data, filesWithHashes, progress);

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.Done });
        }
    }
}
