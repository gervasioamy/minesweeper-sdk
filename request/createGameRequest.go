package request

// Create Game request
type CreateGameRequest struct {
	Rows   int    `json:"rows"`
	Cols   int    `json:"cols"`
	Mines  int    `json:"mines"`
	Player string `json:"player"`
}

/*
{
	"rows": 0,
	"cols": 0,
	"mines": 0,
	"player": "string"
  }
*/
