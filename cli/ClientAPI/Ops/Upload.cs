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
        private static BackendAPI.Model.CreateStoreUploadRequest CreateStoreUploadRequest(string description, string buildId, IEnumerable<ComputeFileMetadata.FileWithMetadata> FileWithMetadata)
        {
            BackendAPI.Model.CreateStoreUploadRequest request = new BackendAPI.Model.CreateStoreUploadRequest(
                useProgressApi: true,
                description: description,
                buildId: buildId,
                files: FileWithMetadata.Select(fileWithMetadata => new BackendAPI.Model.CreateStoreUploadFileRequest(
                    fileName: fileWithMetadata.FileWithoutPath,
                    blobIdentifier: fileWithMetadata.BlobIdentifier,
                    type: fileWithMetadata.Type,
                    size: fileWithMetadata.Size,
                    contentHash: fileWithMetadata.ContentHash
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

        private static async Task UploadMissingFiles(BackendApiWrapper backendApiWrapper, string store, BackendAPI.Model.CreateStoreUploadResponse createStoreUploadResponse, IEnumerable<ComputeFileMetadata.FileWithMetadata> filesWithMetadata, IProgress<UploadProgress> progress)
        {
            if (createStoreUploadResponse.Files != null) {

                for (int fileId = 0; fileId < createStoreUploadResponse.Files.Count; fileId++) {
                    BackendAPI.Model.UploadFileResponse uploadFileResponse = createStoreUploadResponse.Files[fileId];

                    ComputeFileMetadata.FileWithMetadata fileWithMetadata = filesWithMetadata.First(fwh => 
                        fwh.FileWithoutPath == uploadFileResponse.FileName && fwh.BlobIdentifier == uploadFileResponse.BlobIdentifier);

                    if (!string.IsNullOrEmpty(uploadFileResponse.Url)) {

                        if (progress != null)
                            progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFile, FileName = fileWithMetadata.FileWithPath });

                        byte[] content = File.ReadAllBytes(fileWithMetadata.FileWithPath);

                        HttpResponseMessage response = await HttpClient.PutAsync(uploadFileResponse.Url, new ByteArrayContent(content));

                        if (!response.IsSuccessStatusCode) {
                            throw new UploadException($"Upload failed with status code {response.StatusCode}; content = {response.Content}");
                        }

                        string uploadId = createStoreUploadResponse.Id;

                        await backendApiWrapper.MarkStoreUploadFileUploadedAsync(store, uploadId, fileId);

                    } else {

                        if (progress != null)
                            progress.Report(new UploadProgress { State = UploadProgress.StateEnum.FileAlreadyPresent, FileName = fileWithMetadata.FileWithPath });
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

            IEnumerable<ComputeFileMetadata.FileWithMetadata> filesWithMetadata = ComputeFileMetadata.DoComputeFileMedatadata(Files);

            if (progress != null)
                progress.Report(new UploadProgress { State = UploadProgress.StateEnum.CreatingUploadEntry });


            BackendAPI.Model.CreateStoreUploadRequest createStoreUploadRequest = CreateStoreUploadRequest(description, buildId, filesWithMetadata);
            BackendAPI.Model.CreateStoreUploadResponse createStoreUploadResponse;
            createStoreUploadResponse = await backendApiWrapper.CreateStoreUploadAsync(store, createStoreUploadRequest);

            string uploadId = createStoreUploadResponse.Id;

            try {

                if (progress != null)
                    progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFiles });

                await UploadMissingFiles(backendApiWrapper, store, createStoreUploadResponse, filesWithMetadata, progress);

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
