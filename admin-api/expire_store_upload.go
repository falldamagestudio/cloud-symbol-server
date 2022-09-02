package admin_api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ApiService) ExpireStoreUpload(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Expiring upload %v/%v", storeId, uploadId)

	storageClient, err := getStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	type ObjectToDelete struct {
		FileName string
		Hash     string
	}

	objectsToDelete := make([]ObjectToDelete, 0)

	err = runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {
		storeUploadEntry, err := getStoreUploadEntry(client, tx, storeId, uploadId)
		if err != nil {
			return err
		}

		// Fetch all existing store file & file-hash entries and decrease their ref counts accordingly

		type FileHashKey struct {
			FileName string
			Hash     string
		}

		storeFileEntries := make(map[string]StoreFileEntry)
		storeFileHashEntries := make(map[FileHashKey]StoreFileHashEntry)
		filesToDelete := make([]bool, len(storeUploadEntry.Files))

		for fileIndex, file := range storeUploadEntry.Files {

			storeFileEntry, ok := storeFileEntries[file.FileName]
			if !ok {
				storeFileEntryDB, err := getStoreFileEntry(client, tx, storeId, file.FileName)
				if err != nil {
					return err
				} else {
					storeFileEntry = *storeFileEntryDB
				}
			}
			storeFileEntry.RefCount--
			storeFileEntries[file.FileName] = storeFileEntry

			fileHashKey := FileHashKey{FileName: file.FileName, Hash: file.Hash}
			storeFileHashEntry, ok := storeFileHashEntries[fileHashKey]
			if !ok {
				storeFileHashEntryDB, err := getStoreFileHashEntry(client, tx, storeId, file.FileName, file.Hash)
				if err != nil {
					return err
				} else {
					storeFileHashEntry = *storeFileHashEntryDB
				}
			}
			storeFileHashEntry.RefCount--
			storeFileHashEntries[fileHashKey] = storeFileHashEntry

			if storeFileHashEntry.RefCount == 0 {
				filesToDelete[fileIndex] = true
			}
		}

		// Write back results

		for k, v := range storeFileEntries {

			if v.RefCount > 0 {
				if err := updateStoreFileEntry(client, tx, storeId, k, &v); err != nil {
					return err
				}
			} else {
				if err := deleteStoreFileEntry(client, tx, storeId, k); err != nil {
					return err
				}
			}
		}

		for k, v := range storeFileHashEntries {

			if v.RefCount > 0 {
				if err := updateStoreFileHashEntry(client, tx, storeId, k.FileName, k.Hash, &v); err != nil {
					return err
				}
			} else {
				if err := deleteStoreFileHashEntry(client, tx, storeId, k.FileName, k.Hash); err != nil {
					return err
				}
			}
		}

		for fileIndex, file := range storeUploadEntry.Files {

			// Remove reference from upload to file-hash
			err = deleteStoreFileHashUploadEntry(client, tx, storeId, file.FileName, file.Hash, uploadId)
			if err != nil {
				return err
			}

			if filesToDelete[fileIndex] {

				// This was the last reference to the file-hash; delete file-hash

				err = deleteStoreFileHashEntry(client, tx, storeId, file.FileName, file.Hash)
				if err != nil {
					return err
				}

				objectToDelete := ObjectToDelete{
					FileName: file.FileName,
					Hash:     file.Hash,
				}
				objectsToDelete = append(objectsToDelete, objectToDelete)
			}

			storeUploadEntry.Files[fileIndex].Status = FileDBEntry_Status_Expired
		}

		storeUploadEntry.Status = StoreUploadEntry_Status_Expired

		err = updateStoreUploadEntry(client, tx, storeId, uploadId, storeUploadEntry)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("Upload %v/%v not found", storeId, uploadId)
			return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), nil
		} else {
			log.Printf("Error while performing transaction for expiration; err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	log.Printf("Status for %v/%v set to %v", storeId, uploadId, StoreUploadEntry_Status_Expired)

	for _, objectToDelete := range objectsToDelete {
		err = deleteObjectInStore(ctx, storageClient, storeId, objectToDelete.FileName, objectToDelete.Hash)

		if err != nil {
			log.Printf("Error while deleting %s/%s/%s/%s from store; err = %v", storeId, objectToDelete.FileName, objectToDelete.Hash, objectToDelete.FileName, err)
			return openapi.Response(http.StatusInternalServerError, nil), nil
		}

		log.Printf("Deleted file %s/%s/%s/%s from store", storeId, objectToDelete.FileName, objectToDelete.Hash, objectToDelete.FileName)
	}

	log.Printf("Upload %s/%s has been expired", storeId, uploadId)

	return openapi.Response(http.StatusOK, nil), nil
}