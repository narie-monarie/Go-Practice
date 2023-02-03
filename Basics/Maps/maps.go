package main

import "fmt"

func main() {
	heroes := make(map[string]string)

	heroes["Batman"] = "Bruce Kid"
	heroes["Superman"] = "Klerk Kent"

	for i, v := range heroes {
		fmt.Println(i, v)
	}

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}
	fmt.Println(superPets[1])

	//check if exists
	_, ok := superPets[3]
	fmt.Println(ok)

	delete(heroes, "Superman")

	for i, v := range heroes {
		fmt.Println(i, v)
	}

}
