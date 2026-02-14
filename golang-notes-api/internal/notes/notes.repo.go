package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository -> data access layer

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context , note Note) (Note, error){

	opCtx , cancel := context.WithTimeout(ctx , 5 *time.Second)

	defer cancel()

	_, err := r.coll.InsertOne(opCtx , note)
	if err != nil {
		return Note{} , fmt.Errorf("Insert note failed")
	}

	return note , nil
}

func (r *Repo) List(ctx context.Context) ([]Note , error) {

	opCtx , cancel := context.WithTimeout(ctx , 5 * time.Second)
	defer cancel()

	filter := bson.M{} // match all docs

	// Find returns a cursor (like an iterator)
	// over matching documents
	cursor , err := r.coll.Find(opCtx , filter)
	if err != nil {
		return nil , fmt.Errorf("Find noters failed: %w" , err)
	}

	// cursor must be closed after use
	defer cursor.Close(opCtx)

	var notes []Note

	if err := cursor.All(opCtx , &notes); err != nil {
		return nil , fmt.Errorf("Decode notes failed: %w" , err)
	}

	return notes , nil
}