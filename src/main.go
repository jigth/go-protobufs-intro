package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("protobufs action go!")
	fmt.Println()

	daniel := &Person{
		Name: "Daniel",
		Age:  23,
		TrainingItems: []*TrainingItem{
			{
				Name:   "Wooden Rings",
				Weight: 2,
			},
			{
				Name:   "Pullup bar",
				Weight: 1,
			},
			{
				Name:   "Kettlebell 1",
				Weight: 16,
			},
			{
				Name:   "Kettlebell 2",
				Weight: 24,
			},
		},
	}

	fmt.Println("Name", daniel.GetName())
	fmt.Println("Weight", daniel.GetAge())
	fmt.Println()

	for _, el := range daniel.TrainingItems {
		fmt.Println("Name", el.GetName())
		fmt.Println("Weight", el.GetWeight())
		fmt.Println()
	}

	marshalledDaniel, err := proto.Marshal(daniel)
	if err != nil {
		log.Fatal("Couldn't unmarshall person")
	}

	fmt.Println("Marshalled Daniel of length:", marshalledDaniel)

	unmarshalledDaniel := &Person{}

	err = proto.Unmarshal(marshalledDaniel, unmarshalledDaniel)

	if err != nil {
		log.Fatal("Couldn't unmarshall person")
	}

	fmt.Println("Unmarshalled Daniel", unmarshalledDaniel)
}
