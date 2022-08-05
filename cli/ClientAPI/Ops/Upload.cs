using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;

namespace ClientAPI
{
    public class Ops
    {
        private static BackendAPI.Model.CreateStoreUploadRequest CreateStoreUploadRequest(string description, string buildId, IEnumerable<HashFiles.FileWithHash> FileWithHash)
        {
            BackendAPI.Model.CreateStoreUploadRequest request = new BackendAPI.Model.CreateStoreUploadRequest(
                useProgressApi: true,
                description: description,
                buildId: buildId,
                files: FileWithHash.Select(fileWithHash => new BackendAPI.Model.UploadFileRequest(
                    fileName: fileWithHash.FileWithoutPath,
                    hash: fileWithHash.Hash
                )).ToList()
            );

            return request;
        }

        public class UploadException : ClientAPIException
        {
            public UploadException(string message) : base(message) { }
        }

        public struct UploadProgress {
            public enum StateEnum { LocalValidation, CreatingUploadEntry, UploadingMissingFiles, UploadingMissingFile, FileAlreadyPresent, Aborting, Done };

            public StateEnum State;
            public string FileName;
        }

        private static HttpClient HttpClient = new HttpClient();

        private static async Task UploadMissingFiles(BackendApiWrapper backendApiWrapper, string store, BackendAPI.Model.CreateStoreUploadResponse createStoreUploadResponse, IEnumerable<HashFiles.FileWithHash> filesWithHashes, IProgress<UploadProgress> progress)
        {
            if (createStoreUploadResponse.Files != null) {

                for (int fileId = 0; fileId < createStoreUploadResponse.Files.Count; fileId++) {
                    BackendAPI.Model.UploadFileResponse uploadFileResponse = createStoreUploadResponse.Files[fileId];

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

                        string uploadId = createStoreUploadResponse.Id;

                        await backendApiWrapper.MarkStoreUploadFileUploadedAsync(store, uploadId, fileId);

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

            BackendApiWrapper backendApiWrapper = new BackendApiWrapper(ServiceURL, Email, PAT);

            IEnumerable<HashFiles.FileWithHash> filesWithHashes = HashFiles.GetFilesWithHashes(Files);

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.CreatingUploadEntry });


            BackendAPI.Model.CreateStoreUploadRequest createStoreUploadRequest = CreateStoreUploadRequest(description, buildId, filesWithHashes);
            BackendAPI.Model.CreateStoreUploadResponse createStoreUploadResponse;
            createStoreUploadResponse = await backendApiWrapper.CreateStoreUploadAsync(store, createStoreUploadRequest);

            string uploadId = createStoreUploadResponse.Id;

            try {

                if (progress != null)
                    progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFiles });

                await UploadMissingFiles(backendApiWrapper, store, createStoreUploadResponse, filesWithHashes, progress);

                if (progress != null)
                    progress.Report(new UploadProgress { State = UploadProgress.StateEnum.Done });

                await backendApiWrapper.MarkStoreUploadCompletedAsync(store, uploadId);

            } catch {

                try {
                    if (progress != null)
                        progress.Report(new UploadProgress { State = UploadProgress.StateEnum.Aborting });
                    await backendApiWrapper.MarkStoreUploadAbortedAsync(store, uploadId);
                } catch {}

                throw;
            }

        }
    }
}
