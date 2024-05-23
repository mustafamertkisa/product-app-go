package repository

import (
	"context"
	"product-app-go/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepositoryImpl struct {
	MongoDb *mongo.Client
}

func NewLogRepositoryImpl(MongoDb *mongo.Client) LogRepository {
	return &LogRepositoryImpl{MongoDb: MongoDb}
}

func (r *LogRepositoryImpl) AddLogToMongo(log model.LoginLog) error {
	collection := r.MongoDb.Database("logs").Collection("login_logs")

	_, err := collection.InsertOne(context.TODO(), log)
	if err != nil {
		return err
	}

	return nil
}

func (r *LogRepositoryImpl) GetLogByLogId(logId primitive.ObjectID) (model.LoginLog, error) {
	collection := r.MongoDb.Database("logs").Collection("login_logs")

	var log model.LoginLog
	err := collection.FindOne(context.TODO(), bson.M{"_id": logId}).Decode(&log)
	if err != nil {
		return model.LoginLog{}, err
	}

	return log, nil
}

func (r *LogRepositoryImpl) GetAllLogs() ([]model.LoginLog, error) {
	collection := r.MongoDb.Database("logs").Collection("login_logs")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var logs []model.LoginLog
	if err = cursor.All(context.TODO(), &logs); err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *LogRepositoryImpl) GetLogsByUserId(userId int) ([]model.LoginLog, error) {
	collection := r.MongoDb.Database("logs").Collection("login_logs")

	cursor, err := collection.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var logs []model.LoginLog
	if err = cursor.All(context.TODO(), &logs); err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *LogRepositoryImpl) DeleteLogByLogId(logId primitive.ObjectID) error {
	collection := r.MongoDb.Database("logs").Collection("login_logs")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": logId})
	if err != nil {
		return err
	}

	return nil
}

func (r *LogRepositoryImpl) DeleteAllLogs() error {
	collection := r.MongoDb.Database("logs").Collection("login_logs")

	_, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	return nil
}
