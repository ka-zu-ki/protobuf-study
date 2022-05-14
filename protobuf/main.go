package main

import (
	"fmt"
	pb2 "gRPC/protobuf/pb"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"log"
)

func main() {
	employee := &pb2.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb2.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "080-1234-5678"},
		Project:     map[string]*pb2.Company_Project{"ProjectX": &pb2.Company_Project{}},
		Profile: &pb2.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb2.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	binDate, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize", err)
	}

	if err := ioutil.WriteFile("test.bin", binDate, 0666); err != nil {
		log.Fatalln("Can't write", err)
	}

	in, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read file", err)
	}

	readEmployee := &pb2.Employee{}

	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("Can't read deserialize", err)
	}

	fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Can't marshal to json", err)
	}

	fmt.Println(out)

	readEmployee := &pb2.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal from json", err)
	}

	fmt.Println(readEmployee)
}
