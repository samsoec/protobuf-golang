package main

import (
	"fmt"
	practicepb "protobuf-example-go/src/practice"
)

func main() {

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