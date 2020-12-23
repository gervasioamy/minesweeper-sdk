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
	sdk.Initialize(host)
	client = sdk.GetInstance()

	// Test happy path
	log.Print("Let's test a happy path...")

	gameid, err := client.CreateGame(5, 5, 4, "John Doe")
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	discoverAll(gameid)

	flagOne(gameid, 4, 4)
	unflagOne(gameid, 4, 4)

	_, err = client.Pause(gameid)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	_, err = client.Resume(gameid)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	log.Print("=====\n Now let's test error cases...")
	gameid, err = client.CreateGame(5, 5, 400, "John Doe")
	if err == nil {
		fmt.Errorf("Bad reauest was expected")
	}
	log.Printf("Error received as expected: %v", err)

	_, err = client.DiscoverCell("WRONG-ID", 1, 1)
	if err == nil {
		log.Printf("game with WRONG-ID was not found, AS EXPECTED")
	}

}

// start discovering as mach as possible starting in (0,0) without being so smart but using the brute force
func discoverAll(gameid string) bool {
	return discoverOne(gameid, 0, 0)
}

func discoverOne(gameid string, row, col int) bool {
	response, err := client.DiscoverCell(gameid, row, col)
	if err != nil {
		fmt.Errorf("Error: %v", err)
		return false
	}
	return response.GameStatus == "GAME OVER"
}

func flagOne(gameid string, row, col int) bool {
	response, err := client.FlagCell(gameid, row, col)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func unflagOne(gameid string, row, col int) bool {
	response, err := client.UnflagCell(gameid, row, col)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
