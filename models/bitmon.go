package models

import (
	"context"
	"github.com/bitmon-world/bitmon-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*

	Database Structure:
	mons/ <- General monsters information
		ID/
 		   GeneralMon
	user_mons/ <- Particular monsters information
		ID/ (First 8 bytes of the SHA256(Serialized ADN data)
           ParticularMon
	elements/ <- Elements information
		ID/
           Element


*/

type BitmonDBModel struct {
	db *mongo.Database
}

func (model *BitmonDBModel) GetGenMon(id string) (data types.GeneralMon, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("mons")
	filter := bson.D{{"id", id}}
	err = col.FindOne(ctx, filter).Decode(&data)
	return
}

func (model *BitmonDBModel) GetPartMon(id string) (data types.ParticularMon, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("user_mons")
	filter := bson.D{{"id", id}}
	err = col.FindOne(ctx, filter).Decode(&data)
	return
}

func (model *BitmonDBModel) GetElem(id string) (data types.Element, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("elements")
	filter := bson.D{{"id", id}}
	err = col.FindOne(ctx, filter).Decode(&data)
	return
}

func NewDBModel(dbUri string, dbName string) *BitmonDBModel {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}
	return &BitmonDBModel{db: client.Database(dbName)}
}
