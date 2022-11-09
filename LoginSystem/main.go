package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	var config Config

	fmt.Printf("Enter Username: ")
	fmt.Scanln(&config.Username)

	fmt.Printf("Enter Password: ")
	fmt.Scanln(&config.Password)

	var validuname, validpasswd string
	validuname = config.Username
	validpasswd = config.Password
	err = json.Unmarshal([]byte(file), &config)

	if config.Username != validuname || config.Password != validpasswd {
		log.Println("Wrong Login Info!")
		os.Exit(0)
	} else {
		log.Println("Login Success!")
		os.Exit(0)
	}
}
