package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	/*struct to json*/
	var dog = cachorro{"Capa", "Pasto", 1}
	fmt.Println(dog)
	cachorroJson, erro := json.Marshal(dog)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroJson)
	var buffer = bytes.NewBuffer(cachorroJson)
	fmt.Println(buffer)

	/*json to struct*/
	var cachorroJsonString = `{"nome": "Fera","raca": "Vira Lixo","idade": 10}`
	var cachorroDoJson cachorro
	erro = json.Unmarshal([]byte(cachorroJsonString), &cachorroDoJson)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroDoJson)
}
