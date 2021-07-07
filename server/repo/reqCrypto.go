package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type cryptoItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Upvote   int32              `bson:"upvote" json:"upvote"`
	Downvote int32              `bson:"downvote" json:"downvote"`
}

func Create(ctx context.Context, collection *mongo.Collection, crypto cryptoItem) (*cryptoItem, error) {
	res, err := collection.InsertOne(ctx, crypto)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Failed to insert data: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Failed to convert ObjectID"),
		)
	}
	crypto.ID = oid
	return &crypto, nil
}

func Read(ctx context.Context, collection *mongo.Collection, cryptoId string) (*cryptoItem, error) {
	oid, err := primitive.ObjectIDFromHex(cryptoId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("Failed to parse ObjectID"),
		)
	}

	data := &cryptoItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, fmt.Sprintf("Cannot find Crypto ID: %v", err),
		)
	}
	return data, nil
}

func Update(ctx context.Context, collection *mongo.Collection, crypto *cryptoItem) (*cryptoItem, error) {

	filter := bson.M{"_id": crypto.ID}
	_, err := collection.ReplaceOne(ctx, filter, crypto)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Cannot update object %v", err),
		)
	}
	return crypto, nil
}

func Delete(ctx context.Context, collection *mongo.Collection, cryptoId string) (*string, error) {
	oid, err := primitive.ObjectIDFromHex(cryptoId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("Failed to parse ObjectID"),
		)
	}

	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot delete object: %v", err),
		)
	}
	return &cryptoId, nil
}
