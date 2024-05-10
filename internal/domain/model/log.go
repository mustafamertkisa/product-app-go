package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginLog struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Success   bool               `bson:"success" json:"success"`
	Message   string             `bson:"message" json:"message"`
	UserId    int                `bson:"userId" json:"userId"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}
