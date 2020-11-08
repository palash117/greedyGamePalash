package server

import (
	"encoding/json"
	"fmt"
	"greedyGamePalash/dto"
	"greedyGamePalash/processors/requestProcessors"
	"net/http"

	"greedyGamePalash/util"
)

var (
	// PORT = os.Getenv("DC_PORT")
	// IP   = os.Getenv("DC_IP")
	PORT = ":8080"
	IP   = "localhost"
)

func insertHandler(w http.ResponseWriter, req *http.Request) {
	var request dto.Request
	unmarshalErr := json.NewDecoder(req.Body).Decode(&request)
	if unmarshalErr != nil {
		fmt.Errorf("error while unmarshalling %v", unmarshalErr)
	}
	formattedReqeust := util.FormatRequest(request, dto.INSERT)
	requestProcessors.AddRequest(formattedReqeust)
	// request_processors.
	w.Write([]byte("ok"))
}

var fetchHandler = func(w http.ResponseWriter, req *http.Request) {
	var request dto.Request
	unmarshalErr := json.NewDecoder(req.Body).Decode(&request)
	if unmarshalErr != nil {
		fmt.Errorf("error while unmarshalling %v", unmarshalErr)
	}
	formattedReqeust := util.FormatRequest(request, dto.FETCH)
	requestProcessors.AddRequest(formattedReqeust)
	// request_processors.
	data := <-formattedReqeust.RespChannel
	response := dto.Request{}
	response.Dim = request.Dim
	response.Metrics = []dto.KeyWithIntVal{{Key: "webreq", Val: int(*(data.Webreq))}, {Key: "timespent", Val: int(*(data.Timespent))}}
	byteResp, _ := json.Marshal(response)
	w.Write(byteResp)
}

func Start() {
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/query", fetchHandler)

	err := http.ListenAndServe(PORT, nil)
	fmt.Errorf("error", err)
}
