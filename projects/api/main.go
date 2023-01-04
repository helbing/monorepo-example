package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	greetv1 "github.com/helbing/monorepo-example/gen/go/greet/v1"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(":8881", opts...)
	if err != nil {
		log.Fatalf("failed to create grpc connection: %v", err)
	}
	client := greetv1.NewGreetServiceClient(conn)

	r := gin.Default()
	r.GET("/greet", func(ctx *gin.Context) {
		resp, err := client.Greet(ctx, &greetv1.GreetRequest{
			Name: "world",
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"greeting": resp.Greeting,
		})
	})
	err = r.Run(":8882")
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
