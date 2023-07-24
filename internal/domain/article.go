package domain

import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}
