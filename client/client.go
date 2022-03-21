package main

import (
	"proto/proto/proto"

	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure)
	if err != nil {
		panic(err)
	}

	client := proto.NewTrackActivityClient(conn)
    g:=gin.Default()
    response,err:=
	 

}
