package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginLog struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Status    int                `bson:"status" json:"status"`
	Message   string             `bson:"message" json:"message"`
	UserId    int                `bson:"userId" json:"userId"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}
