package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"user-activity-logger/proto/proto"
	pb "user-activity-logger/proto/proto"
)

func main() {

	c, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	client := pb.NewTrackActivityClient(c)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	g := gin.Default()
	g.GET("/add/:name/:id/:email", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			fmt.Print(err)
		}
		req := &proto.AddUsersrequest{
			People: &pb.Person{
				Name:  name,
				Id:    id,
				Email: email,
			},
		}
	})
	fmt.Println("successful establishment of connection")
	fmt.Print(c)
}
