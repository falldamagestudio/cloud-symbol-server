package admin_api

type FileDBEntry struct {
	FileName string `firestore:"filename"`
	Hash     string `firestore:"hash"`
}

type StoreUploadEntry struct {
	Description string        `firestore:"description"`
	BuildId     string        `firestore:"buildId"`
	Timestamp   string        `firestore:"timestamp"`
	Files       []FileDBEntry `firestore:"files"`
}

type StoreEntry struct {
	LatestUploadId int64 `firestore:"latestUploadId"`
}
