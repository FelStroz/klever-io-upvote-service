package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type cryptoItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name,omitempty json:"name,omitempty"`
	Upvote   int32              `bson:"upvote" json:"upvote"`
	Downvote int32              `bson:"downvote" json:"downvote"`
}
