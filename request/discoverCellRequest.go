package request

// DiscoverCellRequest POST api/games/{id}/discover body request
type DiscoverCellRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

/*
{
	"row": 0,
	"col": 0
  }
*/
