package util

import (
	"greedyGamePalash/dto"
	"greedyGamePalash/state"
	"sort"
	"strings"
)

func FormatRequest(request dto.Request, requestType dto.RequestType) *dto.FormattedRequest {

	//todo

	var fRequest = dto.FormattedRequest{}
	levels := getFormattedLevels(request.Dim)
	fRequest.RequestData = dto.FormattedRequestData{}
	fRequest.RequestData.Dim = levels
	webreq, timespent := getMetrics(request.Metrics)
	fRequest.RequestData.Timespent = uint64(timespent)
	fRequest.RequestData.Webreq = uint64(webreq)

	fRequest.RequestType = requestType
	if requestType == dto.FETCH {
		fRequest.RespChannel = make(chan *state.Node)
	}
	return &fRequest
}

func getFormattedLevels(dimensions []dto.KeyVal) []dto.Level {
	// levels := []dto.Level{{LevelName: "global", LevelValue: "global"}}
	levels := []dto.Level{}
	for _, keyVal := range dimensions {
		levels = append(levels, dto.Level{LevelName: strings.ToLower(keyVal.Key), LevelValue: strings.ToLower(keyVal.Val)})
	}
	levels = sortLevels(levels)
	return levels
}

func getMetrics(metrics []dto.KeyWithIntVal) (int, int) {
	webreq := 0
	timespent := 0
	for _, keyVal := range metrics {
		// fmt.Println("key val is ", keyVal.Key, keyVal.Val)
		switch keyVal.Key {
		case "webreq":
			webreq = keyVal.Val
			break
		case "timespent":
			timespent = keyVal.Val
			break
		}
	}
	// fmt.Println("webreq, timespent", webreq, timespent)
	return webreq, timespent
}

func sortLevels(levels []dto.Level) dto.LevelSlice {
	levelSlice := (dto.LevelSlice)(levels)
	sort.Sort(&levelSlice)
	return levelSlice
}
