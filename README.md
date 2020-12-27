# minesweeper-sdk
Author: Gervasio Amy

### Run a demo of the SDK usage
```
go run sdk-demo.go
```
_It will use the deployed minesweeper-api at `https://gamy-minesweeper-api.herokuapp.com/api/`_

### Implementation details
A simple SDK to call [minesweeper-api](https://github.com/gervasioamy/minesweeper-api) endpoints was developed besed on [Resty](https://github.com/go-resty/resty), a simple http client library.

A package called `sdk` was implemented to wrap all the posible endpoint calls. It could be plugged into a service or other applicaton by just importing the package
```
import "github.com/gervasioamy/minesweeper-sdk/sdk"
```
and then call the methods to interact with the API:
```
var client *sdk.SDK
sdk.Initialize("http://localhost:8080/api/")
client = sdk.GetInstance()

//create a game
gameid, err := client.CreateGame(5, 5, 4, "John Doe")
if err != nil {
	fmt.Errorf("Game was not created. Error: %v", err)
}

//discover a cell
discoverResponse, err := client.DiscoverCell(gameid, row, col)
if err != nil {
    fmt.Errorf("Error: %v", err)
}

// flag a cell
flagResponse, err := client.FlagCell(gameid, row, col)
if err != nil {
    fmt.Errorf(err)
}

// unflag a cell
unflagResponse, err := client.UnlagCell(gameid, row, col)
if err != nil {
    fmt.Errorf(err)
}

// pause a game
_, err = client.Pause(gameid)
if err != nil {
    fmt.Errorf("Error: %v", err)
}

// resume a game
_, err = client.Resume(gameid)
if err != nil {
    fmt.Errorf("Error: %v", err)
}
```
