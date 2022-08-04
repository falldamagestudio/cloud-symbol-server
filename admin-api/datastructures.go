package admin_api

const (
	FileDBEntry_Status_AlreadyPresent = "already_present"
	FileDBEntry_Status_Pending        = "pending"
	FileDBEntry_Status_Uploaded       = "uploaded"
)

type FileDBEntry struct {
	FileName string `firestore:"filename"`
	Hash     string `firestore:"hash"`
	Status   string `firestore:"status"`
}

const (
	StoreUploadEntry_Status_InProgress = "in_progress"
	StoreUploadEntry_Status_Complete   = "complete"
)

type StoreUploadEntry struct {
	Description string        `firestore:"description"`
	BuildId     string        `firestore:"buildId"`
	Timestamp   string        `firestore:"timestamp"`
	Files       []FileDBEntry `firestore:"files"`
	Status      string        `firestore:"status"`
}

type StoreEntry struct {
	LatestUploadId int64 `firestore:"latestUploadId"`
}
