package main

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	cryptopb "klever-io-upvote-service/server/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	cryptopb.RegisterCryptoServiceServer(s, &server{})

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func resetMongo() {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://eu:1234qwer@cluster0.ogm0x.mongodb.net/test?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	err = dbClient.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = dbClient.Database("test").Collection("cryptos")
	collection.Drop(context.Background())
}

func TestCreateCrypto(t *testing.T) {
	resetMongo()
	grpcServer := server{}

	emptyReq := &cryptopb.InsertCryptoRequest{Name: ""}
	_, err := grpcServer.InsertCrypto(context.Background(), emptyReq)

	require.NotNil(t, err)

	assert.Equal(t, "rpc error: code = InvalidArgument desc = crypto name must not be empty", err.Error())

	goodReq := &cryptopb.InsertCryptoRequest{Name: "Klever"}
	res, err := grpcServer.InsertCrypto(context.Background(), goodReq)

	require.Nil(t, err)
	assert.NotNil(t, res.GetId())

	//crypto already exists
	_, err = grpcServer.InsertCrypto(context.Background(), goodReq)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = AlreadyExists desc = crypto already exists: <nil>", err.Error())

}

func TestReadCrypto(t *testing.T) {
	resetMongo()
	grpcServer := server{}

	emptyReq := &cryptopb.CryptoID{Id: "123"}
	_, err := grpcServer.ReadCrypto(context.Background(), emptyReq)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = InvalidArgument desc = failed to parse ObjectID", err.Error())

	NotFoundReq := &cryptopb.CryptoID{Id: primitive.NewObjectID().Hex()}
	_, err = grpcServer.ReadCrypto(context.Background(), NotFoundReq)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = NotFound desc = cannot find Crypto ID: mongo: no documents in result", err.Error())

	goodReq := &cryptopb.InsertCryptoRequest{Name: "Klever"}
	res, err := grpcServer.InsertCrypto(context.Background(), goodReq)

	require.Nil(t, err)

	validReq := &cryptopb.CryptoID{Id: res.GetId()}

	response, err := grpcServer.ReadCrypto(context.Background(), validReq)

	require.Nil(t, err)
	require.NotNil(t, response.Id)
	require.NotNil(t, response.Name)
	require.NotNil(t, response.Downvote)
	require.NotNil(t, response.Upvote)

}

func TestListCrypto(t *testing.T) {
	resetMongo()
	grpcServer := server{}

	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(connect()))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := cryptopb.NewCryptoServiceClient(connection)
	request := &cryptopb.ListCryptoRequest{}

	stream, err := client.ListCrypto(ctx, request)

	var result []*cryptopb.ListCryptoResponse
	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			require.Nil(t, err)
			result = append(result, resp)

		}
	}()

	<-done

	require.Nil(t, err)
	require.Nil(t, result)

	cryptos := []*cryptopb.Crypto{
		{
			Name: "Klever",
		},
		{
			Name: "Ethereum",
		},
		{
			Name: "Theter",
		},
	}

	for _, crypto := range cryptos {
		createRequest := &cryptopb.InsertCryptoRequest{Name: crypto.Name}
		_, err := grpcServer.InsertCrypto(context.Background(), createRequest)
		require.Nil(t, err)
	}

	stream, err = client.ListCrypto(ctx, request)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			require.Nil(t, err)

			result = append(result, resp)

		}
	}()

	<-done

	require.Nil(t, err)
	for i, crypto := range cryptos {
		assert.Equal(t, crypto.GetId(), result[i].GetCrypto().GetId())
		assert.Equal(t, crypto.GetName(), result[i].GetCrypto().GetName())
		assert.Equal(t, crypto.GetUpvote(), result[i].GetCrypto().GetUpvote())
		assert.Equal(t, crypto.GetDownvote(), result[i].GetCrypto().GetDownvote())
	}

}

func TestDeleteCrypto(t *testing.T) {
	resetMongo()
	grpcServer := server{}

	emptyReq := &cryptopb.CryptoID{Id: ""}
	_, err := grpcServer.DeleteCrypto(context.Background(), emptyReq)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = InvalidArgument desc = crypto id must not be empty", err.Error())

	NotFoundReq := &cryptopb.CryptoID{Id: primitive.NewObjectID().Hex()}

	_, err = grpcServer.DeleteCrypto(context.Background(), NotFoundReq)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = NotFound desc = cannot delete object: <nil>", err.Error())

	goodReq := &cryptopb.InsertCryptoRequest{Name: "Klever"}
	res, err := grpcServer.InsertCrypto(context.Background(), goodReq)

	require.Nil(t, err)

	validReq := &cryptopb.CryptoID{Id: res.GetId()}
	response, err := grpcServer.DeleteCrypto(context.Background(), validReq)

	require.Nil(t, err)
	assert.NotNil(t, response.GetId())

}

func TestUpvoteCrypto(t *testing.T) {
	resetMongo()
	grpcServer := server{}

	emptyIDRequest := &cryptopb.UpdateCryptoRequest{Id: "", Upvote: 1, Downvote: 0}
	_, err := grpcServer.UpdateCrypto(context.Background(), emptyIDRequest)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = InvalidArgument desc = crypto id must not be empty", err.Error())

	NotFountReq := &cryptopb.UpdateCryptoRequest{Id: primitive.NewObjectID().Hex(), Upvote: 1, Downvote: 0}
	_, err = grpcServer.UpdateCrypto(context.Background(), NotFountReq)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = NotFound desc = cannot find Crypto ID: mongo: no documents in result", err.Error())

	goodReq := &cryptopb.InsertCryptoRequest{Name: "Klever"}
	res, err := grpcServer.InsertCrypto(context.Background(), goodReq)

	require.Nil(t, err)

	validReq := &cryptopb.UpdateCryptoRequest{Id: res.GetId(), Upvote: 1, Downvote: 0}
	response, err := grpcServer.UpdateCrypto(context.Background(), validReq)

	require.Nil(t, err)

	require.NotNil(t, response.GetCrypto().GetId())
	require.Equal(t, "Klever", response.GetCrypto().GetName())
}

func TestDownvoteCrypto(t *testing.T) {
	resetMongo()
	grpcServer := server{}

	emptyIDRequest := &cryptopb.UpdateCryptoRequest{Id: "", Upvote: 0, Downvote: 1}
	_, err := grpcServer.UpdateCrypto(context.Background(), emptyIDRequest)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = InvalidArgument desc = crypto id must not be empty", err.Error())

	NotFountReq := &cryptopb.UpdateCryptoRequest{Id: primitive.NewObjectID().Hex(), Upvote: 0, Downvote: 1}
	_, err = grpcServer.UpdateCrypto(context.Background(), NotFountReq)

	require.NotNil(t, err)
	assert.Equal(t, "rpc error: code = NotFound desc = cannot find Crypto ID: mongo: no documents in result", err.Error())

	goodReq := &cryptopb.InsertCryptoRequest{Name: "Klever"}
	res, err := grpcServer.InsertCrypto(context.Background(), goodReq)

	require.Nil(t, err)

	validReq := &cryptopb.UpdateCryptoRequest{Id: res.GetId(), Upvote: 0, Downvote: 1}
	response, err := grpcServer.UpdateCrypto(context.Background(), validReq)

	require.Nil(t, err)

	require.NotNil(t, response.GetCrypto().GetId())
	require.Equal(t, "Klever", response.GetCrypto().GetName())

}
