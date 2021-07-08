package main

import (
	"fmt"
	cryptopb "klever-io-upvote-service/server/proto"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// --------------------------------------------------------------------------------------------------------------
	connection, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer connection.Close()

	conn := cryptopb.NewCryptoServiceClient(connection)

	g := gin.Default()

	g.POST("/crypto/:name", func(ctx *gin.Context) {

	})
	g.GET("/crypto", func(ctx *gin.Context) {

	})
	g.GET("/crypto/:id", func(ctx *gin.Context) {
		id := ctx.Query("name")
		fmt.Println(id)
		req := &cryptopb.CryptoID{Id: "60e751887d3b2a8b59812950"}
		if res, err := conn.ReadCrypto(ctx, req); err == nil {
			ctx.JSON(200, gin.H{
				"data": fmt.Sprint(res),
			})
		}
	})
	g.PUT("/crypto/:id", func(ctx *gin.Context) {

	})
	g.DELETE("/crypto/:id", func(ctx *gin.Context) {

	})
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
	if err := g.Run(":8000"); err != nil {
		log.Fatalf("failed to run the server: %v", err)
	}

}
