package response

// CreateGameResponse body to be sent to POST /api/games
type CreateGameResponse struct {
	ID string `json:"id"`
}

/*
{
	"id":"5a7de3cc-e758-4041-9ff7-0a40c36e7c34",
	"cells":[
		[],
		[],
		[],
		[],
		[]
	],
	"player":"John Doe",
	"startedTimestamp":null,
	"endedTimestamp":null,
	"mines":9,
	"rows":5,
	"cols":6,
	"status":"CREATED",
	"millisecondsElapsed":0}
*/
