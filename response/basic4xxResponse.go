package response

// Basic4xxResponse is a basic 4xx response, used for mapping those 400, 404, or others responses which only contains a "message" field
type Basic4xxResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:errorCode`
}

/*
{
	"message":"invalid request",

}
*/
