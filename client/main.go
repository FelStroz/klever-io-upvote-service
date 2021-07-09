package main

import (
	"context"
	"fmt"
	"io"
	cryptopb "klever-io-upvote-service/server/proto"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer connection.Close()

	conn := cryptopb.NewCryptoServiceClient(connection)

	g := gin.Default()

	g.POST("/crypto", func(ctx *gin.Context) {
		name := ctx.Query("name")
		req := &cryptopb.InsertCryptoRequest{Name: name}
		if res, err := conn.InsertCrypto(ctx, req); err == nil {
			ctx.JSON(200, gin.H{
				"data": gin.H{"id": res.Id},
				"erro": nil,
			})
		} else {
			ctx.JSON(500, gin.H{
				"data": nil,
				"erro": fmt.Sprint(err),
			})
		}
	})
	g.GET("/crypto", func(ctx *gin.Context) {
		stream, err := conn.ListCrypto(context.Background(), &cryptopb.ListCryptoRequest{})

		if err != nil {
			ctx.JSON(404, gin.H{
				"data": nil,
				"erro": fmt.Sprint(err),
			})
		}

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				ctx.JSON(500, gin.H{
					"data": nil,
					"erro": fmt.Sprint(err),
				})
				break
			}
			ctx.JSON(200, gin.H{
				"data": gin.H{"id": res.Crypto.Id, "name": res.Crypto.Name, "upvote": res.Crypto.Upvote, "downvote": res.Crypto.Downvote},
				"erro": nil,
			})
		}

	})
	g.GET("/crypto/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		req := &cryptopb.CryptoID{Id: id}
		if res, err := conn.ReadCrypto(ctx, req); err == nil {
			ctx.JSON(200, gin.H{
				"data": gin.H{"id": res.Id, "name": res.Name, "upvote": res.Upvote, "downvote": res.Downvote},
				"erro": nil,
			})
		} else {
			ctx.JSON(404, gin.H{
				"data": nil,
				"erro": fmt.Sprint(err),
			})
		}
	})
	g.PUT("/crypto/:id", func(ctx *gin.Context) {
		upvote, err := strconv.ParseInt(ctx.Query("upvote"), 10, 32)
		if err != nil {
			ctx.JSON(500, gin.H{"data": nil, "erro": fmt.Sprint(err)})
		}
		downvote, err := strconv.ParseInt(ctx.Query("downvote"), 10, 32)
		if err != nil {
			ctx.JSON(500, gin.H{"data": nil, "erro": fmt.Sprint(err)})
		}

		id := ctx.Param("id")
		req := &cryptopb.UpdateCryptoRequest{Id: id, Upvote: int32(upvote), Downvote: int32(downvote)}

		if res, err := conn.UpdateCrypto(ctx, req); err == nil {
			ctx.JSON(200, gin.H{
				"data": gin.H{"id": res.Crypto.Id, "name": res.Crypto.Name},
				"erro": nil,
			})
		} else {
			ctx.JSON(404, gin.H{
				"data": nil,
				"erro": fmt.Sprint(err),
			})
		}

	})
	g.DELETE("/crypto/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		req := &cryptopb.CryptoID{Id: id}

		if res, err := conn.DeleteCrypto(ctx, req); err == nil {
			ctx.JSON(200, gin.H{
				"data": gin.H{"id": res.Id},
				"erro": nil,
			})
		} else {
			ctx.JSON(404, gin.H{
				"data": nil,
				"erro": fmt.Sprint(err),
			})
		}
	})

	if err := g.Run(":8000"); err != nil {
		log.Fatalf("failed to run the server: %v", err)
	}
}
