package sdk

import (
	"errors"
	"sync"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"

	"github.com/gervasioamy/minesweeper-sdk/request"
	"github.com/gervasioamy/minesweeper-sdk/response"
)

// MinesweeperSDK the SDK interface
/*
type MinesweeperSDK interface {
	CreateGame(rows, cols, mines int, player string) (string, error)
	DiscoverCell(gameID string, row, col int) (*response.DiscoverCellResponse, error)
	FlagCell(gameID string, row, col int) (bool, error)
	UnflagCell(gameID string, row, col int) (bool, error)
}
*/

var instance *SDK
var once sync.Once

// GetInstance return sdk instance (if exists)
func GetInstance() *SDK {
	return instance
}

// SDK _
type SDK struct {
	client *resty.Client
}

// Init _
func initSDK(host string, verbose bool) {
	log.Info("Initializing SDK...")
	_sdk := new(SDK)
	_sdk.client = resty.New()
	_sdk.client.SetDebug(verbose)
	_sdk.client.SetHostURL(host)
	instance = _sdk
}

// Initialize _
func Initialize(host string, verbose bool) {
	once.Do(func() {
		initSDK(host, verbose)
	})
}

// CreateGame calls POST api/games/ and returns the new game ID or an error if call failed
func (sdk *SDK) CreateGame(rows, cols, mines int, player string) (string, error) {
	reqBody := request.CreateGameRequest{
		Rows:   rows,
		Cols:   cols,
		Mines:  mines,
		Player: player,
	}
	// call the endpoint
	resp, err := sdk.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetResult(&response.CreateGameResponse{}).
		SetError(&response.Basic4xxResponse{}).
		Post("games")

	if err != nil {
		// handle connection issues
		log.WithField("error", err).Error("Connection issue while createing a game")
	}
	if resp.StatusCode() == 201 {
		// game created ok
		resSucBody := resp.Result().(*response.CreateGameResponse)
		log.Infof("Game created successfully with id %s", resSucBody.ID)
		return resSucBody.ID, nil
	}
	if resp.StatusCode() == 400 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("400 when creating a game: %s", resErrBody.Message)
		return "", &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	return "", nil
}

// DiscoverCell calls POST api/games{id}/discover and returns the cells discovered an an error if call failed
func (sdk *SDK) DiscoverCell(gameID string, row, col int) (*response.DiscoverCellResponse, error) {
	reqBody := request.DiscoverCellRequest{
		Row: row,
		Col: col,
	}
	// call the endpoint
	resp, err := sdk.client.SetDebug(true).R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetResult(&response.DiscoverCellResponse{}).
		SetError(&response.Basic4xxResponse{}).
		SetPathParams(map[string]string{
			"gameId": gameID,
		}).
		Post("games/{gameId}/discover")

	if err != nil {
		// handle connection issues
		log.WithField("error", err).Error("Connection issue while createing a game")
	}
	//logResponse(resp)
	if resp.StatusCode() == 200 {
		// game created ok
		resSucBody := resp.Result().(*response.DiscoverCellResponse)
		log.Infof("Cell (%v, %v) disvoered ok", row, col)
		return resSucBody, nil
	}
	if resp.StatusCode() == 400 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("400 when discovering a cell (%v, %v): %v", row, col, resErrBody.Message)
		return nil, errors.New("Bad Request")
	}
	if resp.StatusCode() == 404 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("404 when discovering a cell: %v", resErrBody.Message)
		return nil, errors.New("Game Not Found")
	}
	return nil, nil
}

// FlagCell calls POST api/games{gameID}/flag and returns true if cell was flagged ok or an error if call failed
func (sdk *SDK) FlagCell(gameID string, row, col int) (bool, error) {
	reqBody := request.DiscoverCellRequest{
		Row: row,
		Col: col,
	}
	// call the endpoint
	resp, err := sdk.client.SetDebug(true).R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		//SetResult(&response.DiscoverCellResponse{}).
		SetError(&response.Basic4xxResponse{}).
		SetPathParams(map[string]string{
			"gameId": gameID,
		}).
		Post("games/{gameId}/flag")

	if err != nil {
		// handle connection issues
		log.WithField("error", err).Error("Connection issue while createing a game")
	}
	//logResponse(resp)
	if resp.StatusCode() == 204 {
		// cell flagged created ok
		log.Infof("Cell (%v, %v) flagged ok", row, col)
		return true, nil
	}
	if resp.StatusCode() == 400 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("400 when falgging a cell (%v, %v): %v", row, col, resErrBody.Message)
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	if resp.StatusCode() == 404 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("404 when falgging a cell: %v", resErrBody.Message)
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	return false, nil
}

// UnflagCell calls /api/game/{gameID}/unflag and returns true if cell was unflagged ok or an error if call failed
func (sdk *SDK) UnflagCell(gameID string, row, col int) (bool, error) {
	reqBody := request.DiscoverCellRequest{
		Row: row,
		Col: col,
	}
	// call the endpoint
	resp, err := sdk.client.SetDebug(true).R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		//SetResult(&response.DiscoverCellResponse{}).
		SetError(&response.Basic4xxResponse{}).
		SetPathParams(map[string]string{
			"gameId": gameID,
		}).
		Delete("games/{gameId}/flag")

	if err != nil {
		// handle connection issues
		log.WithField("error", err).Error("Connection issue while createing a game")
	}
	//logResponse(resp)
	if resp.StatusCode() == 204 {
		// cell flagged created ok
		log.Infof("Cell (%v, %v) unflagged ok", row, col)
		return true, nil
	}
	if resp.StatusCode() == 400 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("400 when unfalgging a cell (%v, %v): %v", row, col, resErrBody.Message)
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	if resp.StatusCode() == 404 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("404 when unfalgging a cell: %v", resErrBody.Message)
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	return false, nil
}

// Pause calls POST api/games{gameID}/pause and returns true if game was paused ok or an error if call failed
func (sdk *SDK) Pause(gameID string) (bool, error) {
	// call the endpoint
	resp, err := sdk.client.SetDebug(true).R().
		SetError(&response.Basic4xxResponse{}).
		SetPathParams(map[string]string{
			"gameId": gameID,
		}).
		Post("games/{gameId}/pause")

	if err != nil {
		// handle connection issues
		log.WithField("error", err).Error("Connection issue while createing a game")
	}
	if resp.StatusCode() == 204 {
		// game paused ok
		log.Infof("Game %v paused ok", gameID)
		return true, nil
	}
	if resp.StatusCode() == 400 {
		// bad request, game already paused
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("400 when pause a game: %v", resErrBody.Message)
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	if resp.StatusCode() == 404 {
		// bad request
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("404 when pause a game: %v", resErrBody.Message)
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	return false, nil
}

// Resume calls DELETE api/games{gameID}/pause and returns true if game was resumed ok or an error if call failed
func (sdk *SDK) Resume(gameID string) (bool, error) {
	// call the endpoint
	resp, err := sdk.client.SetDebug(true).R().
		SetError(&response.Basic4xxResponse{}).
		SetPathParams(map[string]string{
			"gameId": gameID,
		}).
		Delete("games/{gameId}/pause")

	if err != nil {
		// handle connection issues
		log.WithField("error", err).Error("Connection issue while createing a game")
	}
	if resp.StatusCode() == 204 {
		// game paused ok
		log.Infof("Game %v resumed ok", gameID)
		return true, nil
	}
	if resp.StatusCode() == 400 {
		// bad request, game not paused
		resErrBody := resp.Error().(response.Basic4xxResponse)
		log.Warnf("400 when resumed a game: %v", resErrBody.Message)
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	if resp.StatusCode() == 404 {
		resErrBody := resp.Error().(*response.Basic4xxResponse)
		log.Warnf("404 when resumed a game: %v", resErrBody.Message)
		//return false, errors.New("Game Not Found")
		return false, &MinesweeperError{resErrBody.ErrorCode, resErrBody.Message}
	}
	return false, nil
}
