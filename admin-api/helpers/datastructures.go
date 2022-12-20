package helpers

const (
	FileDBEntry_Status_Unknown        = "unknown"
	FileDBEntry_Status_AlreadyPresent = "already_present"
	FileDBEntry_Status_Pending        = "pending"
	FileDBEntry_Status_Uploaded       = "uploaded"
	FileDBEntry_Status_Aborted        = "aborted"
	FileDBEntry_Status_Expired        = "expired"
)

type FileDBEntry struct {
	FileName string `firestore:"filename"`
	Hash     string `firestore:"hash"`
	Status   string `firestore:"status"`
}

const (
	StoreUploadEntry_Status_Unknown    = "unknown"
	StoreUploadEntry_Status_InProgress = "in_progress"
	StoreUploadEntry_Status_Completed  = "completed"
	StoreUploadEntry_Status_Aborted    = "aborted"
	StoreUploadEntry_Status_Expired    = "expired"
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

type StoreFileEntry struct {
	RefCount int64
}

type StoreFileHashEntry struct {
	RefCount int64
}

type StoreFileHashUploadEntry struct {
}
