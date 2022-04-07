package mongodb_driver

import (
	"RemoteWorkerService/internal/model"
	"RemoteWorkerService/pkg/database/mongodb"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mDbClientDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbClientDal(Table string) *mDbClientDal {
	return &mDbClientDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbClientDal) Add(data *model.ClientDataModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"ProjectId", data.ProjectId},
		{"IsPaidClient", data.IsPaidClient},
		{"PaidTime", data.PaidTime},
		{"CreatedAt", data.CreatedAt},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *mDbClientDal) UpdateById(Id string, data *model.ClientDataModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.D{{"$set", bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"ProjectId", data.ProjectId},
		{"IsPaidClient", data.IsPaidClient},
		{"PaidTime", data.PaidTime},
		{"CreatedAt", data.CreatedAt},
		{"UpdatedAt", time.Now()},
	}}}

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"ClientId", Id,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}

func (m *mDbClientDal) GetById(Id string) (*model.ClientDataModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", Id,
	}})

	var dataModel = model.ClientDataModel{}
	if result.Err() != nil {
		return &dataModel, result.Err()
	}
	var err = result.Decode(&dataModel)
	if err != nil {
		return &dataModel, err
	}
	return &dataModel, nil
}
