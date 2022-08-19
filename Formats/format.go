package formats

import (
	"LearnGin/messages"
	"encoding/json"
	"encoding/xml"
	"log"

	"gopkg.in/yaml.v2"
)

func FormatFile() {
	person := messages.Person{
		ID:      "36251",
		Name:    "John",
		Email:   "john@example.com",
		Phone:   "123-456-1234",
		Address: "Shen Zheng Street",
	}

	b, _ := xml.MarshalIndent(person, " ", " ")
	log.Printf("[DEBUG] FormatFile: person: \n%s\n\n", b)

	b2, _ := json.MarshalIndent(person, " ", " ")
	log.Printf("[DEBUG] FormatFile: person: \n%s\n\n", b2)

	b3, _ := yaml.Marshal(person)
	log.Printf("[DEBUG] FormatFile: person: \n%s\n\n", b3)
}
