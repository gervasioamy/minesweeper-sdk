package main

import (
	"fmt"
	"log"

	"github.com/gervasioamy/minesweeper-sdk/sdk"
)

var client *sdk.SDK

const host string = "http://localhost:8080/api/"

func main() {
	log.Print("Starting  SDK test ...")
	sdk.Initialize(host, false)
	client = sdk.GetInstance()

	// Test happy path
	log.Print("Let's test a happy path...")

	gameid, err := client.CreateGame(5, 5, 4, "John Doe")
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	discoverOne(gameid, 0, 0)
	// force and eror
	discoverOne(gameid, 0, 0)

	flagOne(gameid, 4, 4)
	unflagOne(gameid, 4, 4)

	pause(gameid)
	// force error
	pause(gameid)
	resume(gameid)
	//force error
	resume(gameid)
}

func discoverOne(gameid string, row, col int) {
	discoverResponse, err := client.DiscoverCell(gameid, 0, 0)
	if err != nil {
		fmt.Errorf("Error: %v", err)
		return
	}
	if discoverResponse.GameStatus == "GAME OVER" {
		log.Print("GAME OVER")
		return
	}
	log.Printf("Cell (%v, %v) discovered successfully, toghether with other %v cells", 0, 0, len(discoverResponse.DiscoveredCells))
}

func flagOne(gameid string, row, col int) bool {
	response, err := client.FlagCell(gameid, row, col)
	if err != nil {
		log.Println(err)
	}
	return response
}

func unflagOne(gameid string, row, col int) bool {
	response, err := client.UnflagCell(gameid, row, col)
	if err != nil {
		log.Println(err)
	}
	return response
}

func pause(gameid string) bool {
	response, err := client.Pause(gameid)
	if err != nil {
		log.Println(err)
	}
	return response
}

func resume(gameid string) bool {
	response, err := client.Resume(gameid)
	if err != nil {
		log.Println(err)
	}
	return response
}
