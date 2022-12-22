package main

import (
	"encoding/json"
	"fmt"
)

type PetStore struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Dogs     []*Pet `json:"dogs,omitempty"`
	Cats     []*Pet `json:"cats,omitempty"`
}
type Pet struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

type PetStoreList = []*PetStore

func main() {
	petstorelist := PetStoreList{}
	petstore := &PetStore{
		Name:     "Kanye's",
		Location: "New York Avenue",
	}
	petstore.Dogs = append(petstore.Dogs, &Pet{
		Name:  "Whiskers",
		Breed: "Pandemoniqs",
	})
	petstore.Dogs = append(petstore.Dogs, &Pet{
		Name: "Anugah",
	})

	petstorelist = append(petstorelist, petstore)

	jsonString, _ := json.MarshalIndent(petstorelist, "", " ")
	fmt.Printf("%s", jsonString)
}
