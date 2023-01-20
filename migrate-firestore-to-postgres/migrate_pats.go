package migrate_firestore_to_postgres

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/volatiletech/sqlboiler/v4/boil"

	postgres "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/postgres"
	models "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/generated/sql-db-models"
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

	tx, err := postgres.BeginDBTransaction(ctx)
	if err != nil {
		log.Printf("Error when beginning DB transaction: %v", err)
		return err
	}

	models.PersonalAccessTokens().DeleteAll(ctx, tx)

	for _, userDocRef := range userDocRefs {
		userName := userDocRef.ID

		log.Printf("User: %v", userName)

		patRefIter := firestoreClient.Collection(usersCollectionName).Doc(userName).Collection(patsCollectionName).DocumentRefs(ctx)

		patDocRefs, err := patRefIter.GetAll()
		if err != nil {
			log.Printf("Get all pat doc refs for user failed: %v", err)
			tx.Rollback()
			return err
		}

		for _, patDocRef := range patDocRefs {
			token := patDocRef.ID

			patDoc, err := patDocRef.Get(ctx)
			if err != nil {
				log.Printf("Get pat doc for user / pat %v / %v failed: %v", userName, token, err)
				tx.Rollback()
				return err
			}

			creationTimestamp, _ := patDoc.DataAt("creationTimestamp")
			description, _ := patDoc.DataAt("description")

			log.Printf("  PAT: %v", token)
			log.Printf("    creationTimestamp: %v", creationTimestamp)
			log.Printf("    description: %v", description)

			var pat models.PersonalAccessToken
			pat.Owner = userName
			pat.Token = token
			if description != nil {
				pat.Description = description.(string)
			}
			if creationTimestamp != nil {
				pat.CreationTimestamp = creationTimestamp.(time.Time)
			}

			pat.Insert(ctx, tx, boil.Infer())
		}
	}

	tx.Commit()

	return nil
}
