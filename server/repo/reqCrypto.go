package repository

import (
	"context"
	"fmt"
	model "klever-io-upvote-service/server/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Insert(ctx context.Context, collection *mongo.Collection, crypto model.CryptoItem) (*string, error) {
	res, err := collection.InsertOne(ctx, crypto)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("failed to insert data: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("failed to convert ObjectID"),
		)
	}

	oidString := oid.Hex()
	return &oidString, nil
}

func Read(ctx context.Context, collection *mongo.Collection, cryptoId string) (*model.CryptoItem, error) {
	oid, err := primitive.ObjectIDFromHex(cryptoId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("failed to parse ObjectID"),
		)
	}

	data := &model.CryptoItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, fmt.Sprintf("cannot find Crypto ID: %v", err),
		)
	}
	return data, nil
}

func Update(ctx context.Context, collection *mongo.Collection, crypto *model.CryptoItem) (*model.CryptoItem, error) {

	filter := bson.M{"_id": crypto.ID}
	data, err := collection.ReplaceOne(ctx, filter, crypto)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("cannot update object %v", err),
		)
	}
	fmt.Println(data, "Updated data")
	return crypto, nil
}

func Delete(ctx context.Context, collection *mongo.Collection, cryptoId string) (*string, error) {
	oid, err := primitive.ObjectIDFromHex(cryptoId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("failed to parse ObjectID"),
		)
	}

	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot delete object: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("cannot delete object: %v", err),
		)
	}
	return &cryptoId, nil
}
