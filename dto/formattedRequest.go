package dto

import (
	"greedyGamePalash/constants"
	"greedyGamePalash/state"
)

// FormattedRequest wraps the request data along with reqeust type and resp channel
type FormattedRequest struct {
	RequestData FormattedRequestData

	RequestType RequestType

	RespChannel chan *state.Node
}

// FormattedRequestData contains the request data
type FormattedRequestData struct {
	Dim       []Level
	Webreq    uint64
	Timespent uint64
}

// Level is the dimension values
type Level struct {
	LevelName  string
	LevelValue string
	// LevelHeight i
}

type LevelSlice []Level

func (s LevelSlice) Len() int {
	return len(s)
}

func (s LevelSlice) Less(i, j int) bool {
	c1 := constants.LevelOrderFromString[s[i].LevelName]
	c2 := constants.LevelOrderFromString[s[j].LevelName]
	result := c1 < c2
	return result
}

func (s LevelSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
