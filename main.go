package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	complexpb "protobuf-example-go/src/complex"
	enumpb "protobuf-example-go/src/enum"
	practicepb "protobuf-example-go/src/practice"
	simplepb "protobuf-example-go/src/simple"
)

func main() {
	sm := doSimple()

	//readWriteDome(sm)

	jsonDemo(sm)

	//doEnum()

	//doPractice()
}

func doPractice () {
	ex := practicepb.AddressBook{
		People: []*practicepb.Person{
			&practicepb.Person{
				Id: 1,
				Name: "Rizal",
				Email: "sam.fauzy@gmail.com",
				Phones: []*practicepb.Person_PhoneNumber{
					&practicepb.Person_PhoneNumber{
						Number: "0821-XXXX-XXXX",
						Type: practicepb.Person_MOBILE,
					},
					&practicepb.Person_PhoneNumber{
						Number: "021-XXXX-XXXX",
						Type: practicepb.Person_WORK,
					},
				},
			},
			&practicepb.Person{
				Id: 2,
				Name: "Samsul",
				Email: "sam.fauzy02@gmail.com",
				Phones: []*practicepb.Person_PhoneNumber{
					&practicepb.Person_PhoneNumber{
						Number: "0818-XXXX-XXXX",
						Type: practicepb.Person_MOBILE,
					},
				},
			},
		},
	}

	fmt.Println("declare", ex)

	writeToFile("practice.bin", &ex)

	exFromBin := practicepb.AddressBook{}

	readFromFile("practice.bin", &exFromBin)

	fmt.Println("from bin", exFromBin)
}

func doComplex() {
	cp := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{Id: 1, Name: "first dummy",},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{Id: 2, Name: "second dummy",},
			&complexpb.DummyMessage{Id: 3, Name: "third dummy",},
		},
	}

	fmt.Println(cp)
}

func doEnum() {
	ep := enumpb.EnumMessage{
		Id: 321,
		DayName: enumpb.DayName_MONDAY,
	}

	ep.DayName = enumpb.DayName_SUNDAY

	fmt.Println(ep)
}

func jsonDemo (sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println(smAsString)

	sm3 := &simplepb.SimpleMessage{}

	fromJSON(smAsString, sm3)

	fmt.Println(sm3)
}

func toJSON (pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}

	out, err := marshaler.MarshalToString(pb)

	if err != nil {
		log.Fatalln("Error convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON (in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)

	if err != nil {
		log.Fatalln("Error convert from JSON", err)
	}
}

func readWriteDome(sm proto.Message) {
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}

	readFromFile("simple.bin", sm2)

	fmt.Println(sm2)
}

func readFromFile(filename string, pb proto.Message) error {
	in, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error reading file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)

	if err2 != nil {
		log.Fatalln("Error deserialize", err)
		return err2
	}

	return nil
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Cant serialize to bytes")
		return err
	}

	if err != ioutil.WriteFile(filename, out, 0644) {
		log.Fatalln("Cant write to file")
		return err
	}

	fmt.Println("Data has been written")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id: 123,
		IsSimple: true,
		Name: "name",
		SimpleList: []int32{1, 2, 3, 4},
	}

	fmt.Println(sm)

	sm.Name = "renamed"

	fmt.Println(sm)

	fmt.Println("ID : ", sm.GetId())

	return &sm
}
