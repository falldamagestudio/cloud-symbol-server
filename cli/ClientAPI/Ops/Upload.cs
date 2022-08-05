﻿using System;
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

        public class UploadException : Exception
        {
            public UploadException(string message) : base(message) { }
        }

        public struct UploadProgress {
            public enum StateEnum { LocalValidation, CreatingUploadEntry, UploadingMissingFiles, UploadingMissingFile, FileAlreadyPresent, Aborting, Done };

            public StateEnum State;
            public string FileName;
        }

        private static HttpClient HttpClient = new HttpClient();

        private static async Task UploadMissingFiles(BackendAPI.Api.DefaultApi api, string store, BackendAPI.Model.CreateStoreUploadResponse createStoreUploadResponse, IEnumerable<HashFiles.FileWithHash> filesWithHashes, IProgress<UploadProgress> progress)
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
                        BackendAPI.Client.ApiResponse<object> markStoreUploadFileUploadedResponse;
                        try {
                            markStoreUploadFileUploadedResponse = api.MarkStoreUploadFileUploadedWithHttpInfo(uploadId, store, fileId);
                            if (markStoreUploadFileUploadedResponse.ErrorText != null)
                                throw new UploadException(markStoreUploadFileUploadedResponse.ErrorText);
                        } catch (BackendAPI.Client.ApiException apiException) {
                            if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                                throw new UploadException($"Error Store {store} / upload {uploadId} / file {fileId} does not exist");
                            else
                                throw;
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

            BackendAPI.Api.DefaultApi api = Helpers.CreateApi(ServiceURL, Email, PAT);

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

            string uploadId = createStoreUploadResponse.Data.Id;

            try {

                if (progress != null)
                    progress.Report(new UploadProgress { State = UploadProgress.StateEnum.UploadingMissingFiles });

                await UploadMissingFiles(api, store, createStoreUploadResponse.Data, filesWithHashes, progress);

                if (progress != null)
                    progress.Report(new UploadProgress { State = UploadProgress.StateEnum.Done });

                BackendAPI.Client.ApiResponse<object> markStoreUploadCompletedResponse;
                try {
                    markStoreUploadCompletedResponse = api.MarkStoreUploadCompletedWithHttpInfo(uploadId, store);
                    if (markStoreUploadCompletedResponse.ErrorText != null)
                        throw new UploadException(markStoreUploadCompletedResponse.ErrorText);
                } catch (BackendAPI.Client.ApiException apiException) {
                    if (apiException.ErrorCode == (int)HttpStatusCode.NotFound)
                        throw new UploadException($"Store {store} / upload {uploadId} does not exist");
                    else
                        throw;
                }

            } catch {

                try {
                    if (progress != null)
                        progress.Report(new UploadProgress { State = UploadProgress.StateEnum.Aborting });
                    api.MarkStoreUploadAbortedWithHttpInfo(uploadId, store);
                } catch {}

                throw;
            }

        }
    }
}
