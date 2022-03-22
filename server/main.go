package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"user-activity-logger/proto/proto"
	pb "user-activity-logger/proto/proto"
)

type TrackActivityServer struct {
}

func (s *TrackActivityServer) isDone(ctx context.Context, res *proto.UpdateActivityRequest) (*proto.UpdateActivityResponse, error) {
	p := &pb.Activity{
		T:      res.GetAct().GetT(),
		Status: false,
		Valid:  false,
		Lable:  res.GetAct().GetLable(),
		Set:    res.GetAct().GetSet(),
	}
	fmt.Print(p)
	Done := false

	star := time.Now()
	for {

		fmt.Print("want to stop task:Y/N?")
		k := ""
		fmt.Scan(&k)

		if k == "Y" {
			stop := time.Since(star)
			fmt.Print("start time:%v", star)
			fmt.Print("total time:%v", stop)
			Done = true
			break
		} else {
			time.Sleep(1 * time.Hour)

		}
	}

	check := isValid()

	return &proto.UpdateActivityResponse{
		Complete: Done,
		Valid:    check,
	}, nil
}

func isValid() bool {

	valid := true
	return valid
}

func (s *TrackActivityServer) Add(ctx context.Context, res *proto.AddUsersrequest) (*proto.AddUsersresponse, error) {
	p := &pb.Person{
		Name:  res.GetPeople().GetName(),
		Id:    res.GetPeople().GetId(),
		Email: res.GetPeople().GetEmail(),
	}
	return &proto.AddUsersresponse{
		People: p,
	}, nil

}

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("successful establishment of connection")

	fmt.Print(c)

}
