package request

// DiscoverCellRequest POST api/games/{id}/discover body request
type DiscoverCellRequest struct {
	Row int `json:"rows"`
	Col int `json:"cols"`
}

/*
{
	"row": 0,
	"col": 0
  }
*/
