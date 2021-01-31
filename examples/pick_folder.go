package main

import (
	nfd_go "github.com/billyct/nfd-go"
	"log"
)

func main() {
	path, err := nfd_go.PickFolder("")
	if err != nil {
		panic(err)
	}

	log.Println(path)
}
