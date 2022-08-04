using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class Ops
    {
        private static BackendAPI.Model.CreateStoreUploadRequest CreateStoreUploadRequest(string description, string buildId, IEnumerable<HashFiles.FileWithHash> FileWithHash)
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
            public enum StateEnum { LocalValidation, CreatingUploadEntry, UploadingMissingFiles, UploadingMissingFile, FileAlreadyPresent, Done };

            public StateEnum State;
            public string FileName;
        }

        private static HttpClient HttpClient = new HttpClient();

        private static async Task UploadMissingFiles(BackendAPI.Model.CreateStoreUploadResponse createStoreUploadResponse, IEnumerable<HashFiles.FileWithHash> filesWithHashes, IProgress<UploadProgress> progress)
        {
            if (createStoreUploadResponse.Files != null) {
                foreach (BackendAPI.Model.UploadFileResponse uploadFileResponse in createStoreUploadResponse.Files) {

                    HashFiles.FileWithHash fileWithHash = filesWithHashes.First(fwh => 
                        fwh.FileWithoutPath == uploadFileResponse.FileName && fwh.Hash == uploadFileResponse.Hash);

                    if (!string.IsNullOrEmpty(uploadFileResponse.Url)) {

                        if (progress != null)
                            progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFile, FileName = fileWithHash.FileWithPath });

                        byte[] content = File.ReadAllBytes(fileWithHash.FileWithPath);

                        HttpResponseMessage response = await HttpClient.PutAsync(uploadFileResponse.Url, new ByteArrayContent(content));

                        if (!response.IsSuccessStatusCode) {
                            throw new UploadException($"Upload failed with status code {response.StatusCode}; content = {response.Content}");
                        }

                    } else {

                        if (progress != null)
                            progress.Report(new UploadProgress { State = UploadProgress.StateEnum.FileAlreadyPresent, FileName = fileWithHash.FileWithPath });
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

            IEnumerable<HashFiles.FileWithHash> filesWithHashes = HashFiles.GetFilesWithHashes(Files);

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.CreatingUploadEntry });

            BackendAPI.Model.CreateStoreUploadRequest createStoreUploadRequest = CreateStoreUploadRequest(description, buildId, filesWithHashes);
            BackendAPI.Client.ApiResponse<BackendAPI.Model.CreateStoreUploadResponse> createStoreUploadResponse;
            try {
                createStoreUploadResponse = api.CreateStoreUploadWithHttpInfo(store, createStoreUploadRequest);
                if (createStoreUploadResponse.ErrorText != null)
                    throw new UploadException(createStoreUploadResponse.ErrorText);
            } catch (BackendAPI.Client.ApiException apiException) {
                if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                    throw new UploadException($"Store {store} does not exist");
                else
                    throw;
            }

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFiles });

            await UploadMissingFiles(createStoreUploadResponse.Data, filesWithHashes, progress);

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.Done });
        }
    }
}
