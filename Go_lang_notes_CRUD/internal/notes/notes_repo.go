package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repo -> data access layer
 

type Repo struct {
	coll *mongo.Collection

}


func NewRepo(db *mongo.Database) *Repo{
	 return  &Repo{
		coll:  db.Collection("notes"),
	 }
} 


func (r *Repo) create(ctx context.Context, note Notes) (Notes, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	_, err := r.coll.InsertOne(opCtx,note)

	 if err != nil {
		return  Notes{},fmt.Errorf("Insert note failed")
	 }

	return  note , nil
}