package response

// DiscoverCellResponse is a the json response with the discovered cells when calling POST api/games/{id}/discover
type DiscoverCellResponse struct {
	GameStatus      string         `json:"gameStatus"`
	DiscoveredCells []CellResponse `json:discoveredCells`
}

// CellResponse each of the cells discovered
type CellResponse struct {
	Row   int `json:"row"`
	Col   int `json:"col"`
	Value int `json:value`
}

/*
{
    "gameStatus": "STARTED",
    "discoveredCells": [
        {
            "row": 5,
            "col": 0,
            "discovered": true,
            "value": 0,
            "flagged": false
        },
        {
            "row": 5,
            "col": 1,
            "discovered": true,
            "value": 1,
            "flagged": false
        },
        {
            "row": 4,
            "col": 1,
            "discovered": true,
            "value": 2,
            "flagged": false
        },
        {
            "row": 4,
            "col": 0,
            "discovered": true,
            "value": 1,
            "flagged": false
        }
    ]
}
*/
