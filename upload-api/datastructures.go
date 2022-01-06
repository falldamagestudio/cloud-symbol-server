package upload_api

type FileDBEntry struct {
	FileName string `firestore:"filename"`
	Hash     string `firestore:"hash"`
}

type TransactionDBEntry struct {
	Description string        `firestore:"description"`
	BuildId     string        `firestore:"buildId"`
	Timestamp   string        `firestore:"timestamp"`
	Files       []FileDBEntry `firestore:"files"`
}
