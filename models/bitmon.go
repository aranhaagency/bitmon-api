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
	modifiers/
		adv/ <- adventure algorithm modifiers

*/

type BitmonDBModel struct {
	db *mongo.Database
}

func (model *BitmonDBModel) GetMonsList() (data []types.GeneralMon, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("mons")
	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer func() {
		_ = cur.Close(ctx)
	}()
	for cur.Next(ctx) {
		var mon types.GeneralMon
		err = cur.Decode(&data)
		if err != nil {
			return
		}
		data = append(data, mon)
	}
	return
}

func (model *BitmonDBModel) GetGenMon(id string) (data types.GeneralMon, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("mons")
	filter := bson.D{{Key: "id", Value: id}}
	err = col.FindOne(ctx, filter).Decode(&data)
	return
}

func (model *BitmonDBModel) AddMon(mon types.GeneralMon) (data types.GeneralMon, err error) {
	return
}

func (model *BitmonDBModel) GetAdventureModifiers() (data types.AdventureModifiers, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	col := model.db.Collection("modifiers")
	filter := bson.D{{Key: "id", Value: "adv"}}
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
