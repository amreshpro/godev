package main

import (
	"errors"
	"fmt"
	"log"
)

type Music struct {
	id     string
	status string
}

func processMusic(music Music) error {
	fmt.Printf("id:%s \nEnjoying Music : %s ", music.id, music.status)
	return errors.New("\n not implemented")
}

func main() {
	musics := []Music{
		{id: "1",
			status: "enjoying"},
		{id: "2",
			status: "not really"}, {id: "3",
			status: "neutral"},
	}

	for _, music := range musics {
		fmt.Printf("\nMusic file %s is recevied.\n", music.status)
	
	if	err := processMusic(music); err!=nil{
		log.Fatalf("Error in processing : %s",err)
	}
	}

}
