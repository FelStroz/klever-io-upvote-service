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
		Upvote:   1,
		Downvote: 1,
	}

	res, err := repository.Insert(context.Background(), collection, data)

	if err != nil {
		return nil, err
	}

	return &cryptopb.CryptoID{
		Id: *res,
	}, nil
}

func (*server) ReadCrypto(ctx context.Context, req *cryptopb.CryptoID) (*cryptopb.Crypto, error) {
	fmt.Println("Read crypto request")
	cryptoID := req.Id
	data, err := repository.Read(context.Background(), collection, cryptoID)

	if err != nil {
		return nil, err
	}

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
	data, err := repository.Read(context.Background(), collection, crypto.Id)

	if err != nil {
		return nil, err
	}

	data.Upvote += crypto.Upvote
	data.Downvote += crypto.Downvote

	res, err := repository.Update(context.Background(), collection, data)

	if err != nil {
		return nil, err
	}

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
	fmt.Println("Delete crypto request")
	cryptoID := req.Id
	res, err := repository.Delete(context.Background(), collection, cryptoID)

	if err != nil {
		return nil, err
	}

	return &cryptopb.CryptoID{
		Id: *res,
	}, nil
}

func (*server) ListCrypto(req *cryptopb.ListCryptoRequest, stream cryptopb.CryptoService_ListCryptoServer) error {
	fmt.Println("List crypto request")

	data := &model.CryptoItem{}
	res, err := repository.List(context.Background(), collection)

	if err != nil {
		return err
	}

	defer res.Close(context.Background())
	for res.Next(context.Background()) {
		err := res.Decode(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
		stream.Send(&cryptopb.ListCryptoResponse{
			Crypto: &cryptopb.Crypto{
				Id:       data.ID.Hex(),
				Name:     data.Name,
				Upvote:   data.Upvote,
				Downvote: data.Downvote,
			},
		})
	}
	return nil
}

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
	lis, err := net.Listen("tcp", "0.0.0.0:4040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	cryptopb.RegisterCryptoServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server on port 4040...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}

	}()

	// --------------------------------------------------------------------------------------------------------------
	//cc, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("could not connect: %v", err)
	//}
	//defer cc.Close()

	//c := cryptopb.NewCryptoServiceClient(cc)
	// --------------------------------------------------------------------------------------------------------------
	//createCryptoRes, err := c.InsertCrypto(context.Background(), &cryptopb.InsertCryptoRequest{Name: "Bitcoin"})
	//ReadedCryptoRes, err := c.ReadCrypto(context.Background(), &cryptopb.CryptoID{Id: "60e751887d3b2a8b59812950"})
	//UpdatedCryptoRes, err := c.UpdateCrypto(context.Background(), &cryptopb.UpdateCryptoRequest{Id: "60e751887d3b2a8b59812950", Upvote: 1, Downvote: 0})
	//DeletedCryptoRes, err := c.DeleteCrypto(context.Background(), &cryptopb.CryptoID{Id: "60e751887d3b2a8b59812950"})
	// Call ListBlogs that returns a stream
	//stream, err := c.ListCrypto(context.Background(), &cryptopb.ListCryptoRequest{})
	// Check for errors
	// Start iterating
	//for {
	//res, err := stream.Recv()
	//if err == io.EOF {
	//	break
	//}
	//if err != nil {
	//	return log.Println(err)
	//}
	//	fmt.Println(res.GetCrypto())
	//}

	//if err != nil {
	//	log.Println(err)
	//} else {
	//fmt.Printf("Crypto has been created: %v", createCryptoRes)
	//fmt.Printf("Crypto: %v", ReadedCryptoRes)
	//fmt.Printf("Crypto has been updated: %v", UpdatedCryptoRes)
	//fmt.Printf("Crypto has been deleted: %v", DeletedCryptoRes)
	//}
	// --------------------------------------------------------------------------------------------------------------

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	s.Stop()
	lis.Close()
	client.Disconnect(context.TODO())
	fmt.Println("Server stopped.")
}

func main() {
	serverRun()
}
