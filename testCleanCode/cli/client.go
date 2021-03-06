package main

import (
	"context"
	"fmt"
	person "github.com/ragilmaulana/Latihan/clean/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)


const (
	address = "localhost:1010"
)

func main() {
printAllPerson()
}

func deletePerson(id int64){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := person.NewPersonServiceClient(conn)

	data := person.PersonRequest{
		IdUser: id,
	}
	response, err := c.DeletePerson(ctx, &data)
	fmt.Println("response from server ",response)

}

func InsertPerson(firstname,lastname string){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := person.NewPersonServiceClient(conn)

	data := person.PersonRequest{
		FirstName: firstname,
		LastName: lastname,
	}
	response, err := c.AddPerson(ctx, &data)
	fmt.Println("response from server ",response)

}

func printAllPerson(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := person.NewPersonServiceClient(conn)

	empty := person.Empty{}

	response, err := c.PrintPerson(ctx, &empty)
	for _,value := range response.PersonResponse {
		fmt.Println(value)
	}

}

//type server struct {
//	input domain.PersonInputPort
//}
//
//func NewPersonService(a domain.PersonInputPort) *server {
//	return &server{
//		input: a,
//	}
//}
//
//func (s server)PrintPerson() (domain.Person, error) {
//	person := domain.Person{
//		Id_User: 123,
//		FirstName: "sdfsdf",
//		LastName: "asdfsdf",
//	}
//	data, err := s.input.PrintPerson(person)
//	if err!= nil{
//		panic(err)
//	}
//	return data,nil
//}
//
//func printPerson(s)
//
//
//func main(){
//
//}