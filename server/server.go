package main

import (
	"context"
	"fmt"
	configDB "klever-io-upvote-service/server/config"
	model "klever-io-upvote-service/server/model"
	cryptopb "klever-io-upvote-service/server/proto"
	repository "klever-io-upvote-service/server/repo"

	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
}

func (*server) InsertCrypto(ctx context.Context, req *cryptopb.InsertCryptoRequest) (*cryptopb.CryptoID, error) {
	fmt.Println("Insert crypto request")
	cryptoName := req.Name

	data := model.CryptoItem{
		Name:     cryptoName,
		Upvote:   0,
		Downvote: 0,
	}

	res, _ := repository.Insert(context.Background(), collection, data)

	return &cryptopb.CryptoID{
		Id: *res,
	}, nil
}

func (*server) ReadCrypto(ctx context.Context, req *cryptopb.CryptoID) (*cryptopb.Crypto, error) {
	fmt.Println("Read crypto request")
	cryptoID := req.Id
	data, _ := repository.Read(context.Background(), collection, cryptoID)

	return &cryptopb.Crypto{
		Id:       data.ID.Hex(),
		Name:     data.Name,
		Upvote:   data.Upvote,
		Downvote: data.Downvote,
	}, nil
}

func (*server) UpdateCrypto(ctx context.Context, req *cryptopb.UpdateCryptoRequest) (*cryptopb.UpdateCryptoResponse, error) {
	fmt.Println("Update crypto request")
	crypto := req
	data, _ := repository.Read(context.Background(), collection, crypto.Id)
	data.Upvote = crypto.Upvote
	data.Downvote = crypto.Downvote

	res, _ := repository.Update(context.Background(), collection, data)
	return &cryptopb.UpdateCryptoResponse{
		Crypto: &cryptopb.Crypto{
			Id:       res.ID.Hex(),
			Name:     data.Name,
			Upvote:   res.Upvote,
			Downvote: res.Downvote,
		},
	}, nil
}

func (*server) DeleteCrypto(ctx context.Context, req *cryptopb.CryptoID) (*cryptopb.CryptoID, error) {
	fmt.Println("Delete user request")
	cryptoID := req.Id
	res, _ := repository.Delete(context.Background(), collection, cryptoID)
	return &cryptopb.CryptoID{
		Id: *res,
	}, nil
}

//func (*server) ListCrypto(ctx context.Context, req *cryptopb.ListCryptoRequest) (*cryptopb.ListCryptoResponse, error) {
//fmt.Println("List user request")
//cryptoID := req.GetCryptoId()
//res, _ := repository.Delete(context.Background(), collection, cryptoID)
//var res []*grpc.User
//	for _, u := range *users {
//		res = append(res, u.ToGRPC())
//	}
//	return &grpc.ListUserRes{Users: res}, nil
//return &userpb.DeleteUserResponse{UserId: *res}, nil
//}

// Connect to MongoDB
func connection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(configDB.Db.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10000)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func serverRun() {
	client, _ := connection()
	collection = client.Database(configDB.Db.MongoDB).Collection("crypto")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	cryptopb.RegisterCryptoServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}

	}()

	// --------------------------------------------------------------------------------------------------------------
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := cryptopb.NewCryptoServiceClient(cc)
	// --------------------------------------------------------------------------------------------------------------
	//createCryptoRes, err := c.InsertCrypto(context.Background(), &cryptopb.InsertCryptoRequest{Name: "Bitcoin"})
	ReadedCryptoRes, err := c.ReadCrypto(context.Background(), &cryptopb.CryptoID{Id: "60e68ef626446739d78c566a"})

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	//fmt.Printf("Crypto has been created: %v", createCryptoRes)
	fmt.Printf("Crypto: %v", ReadedCryptoRes)
	// --------------------------------------------------------------------------------------------------------------

	// Wait for exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Signal is received
	<-ch
	s.Stop()
	lis.Close()
	client.Disconnect(context.TODO())
	fmt.Println("Server stopped.")
}

func main() {
	serverRun()
}
