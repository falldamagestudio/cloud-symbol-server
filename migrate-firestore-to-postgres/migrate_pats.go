package migrate_firestore_to_postgres

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

const (
	usersCollectionName = "users"
	patsCollectionName  = "pats"
)

func MigratePATs(ctx context.Context, firestoreClient *firestore.Client) error {

	documentRefIter := firestoreClient.Collection(usersCollectionName).DocumentRefs(ctx)

	userDocRefs, err := documentRefIter.GetAll()
	if err != nil {
		log.Printf("Get all user doc refs failed: %v", err)
		return err
	}

	for _, userDocRef := range userDocRefs {
		userName := userDocRef.ID

		log.Printf("User: %v", userName)

		patRefIter := firestoreClient.Collection(usersCollectionName).Doc(userName).Collection(patsCollectionName).DocumentRefs(ctx)

		patDocRefs, err := patRefIter.GetAll()
		if err != nil {
			log.Printf("Get all pat doc refs for user failed: %v", err)
			return err
		}

		for _, patDocRef := range patDocRefs {
			pat := patDocRef.ID

			patDoc, err := patDocRef.Get(ctx)
			if err != nil {
				log.Printf("Get pat doc for user / pat %v / %v failed: %v", userName, pat, err)
				return err
			}

			creationTimestamp, _ := patDoc.DataAt("creationTimestamp")
			description, _ := patDoc.DataAt("description")

			log.Printf("  PAT: %v", pat)
			log.Printf("    creationTimestamp: %v", creationTimestamp)
			log.Printf("    description: %v", description)
		}
	}

	return nil
}
