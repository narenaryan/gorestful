package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/narenaryan/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "John D",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	body, _ := proto.Marshal(p)
	fmt.Println(body)
}