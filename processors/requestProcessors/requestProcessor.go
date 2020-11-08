package requestProcessors

import (
	"greedyGamePalash/dto"
	"greedyGamePalash/processors/fetchProcessor"
	"greedyGamePalash/processors/insertLogProcessor"
	"greedyGamePalash/processors/insertProcessor"
	"greedyGamePalash/state"
)

var requestChan chan *dto.FormattedRequest

func processor(requestChan, fetchRequestChan, insertRequestChan, insertLogChan chan *dto.FormattedRequest) {
	for {
		request := <-requestChan
		switch request.RequestType {
		case dto.FETCH:
			fetchRequestChan <- request
			break
		case dto.INSERT:
			insertRequestChan <- request
			insertLogChan <- request

		}
	}
}

func InitializeRequestProcessors(state *state.Node) {
	fetchRequestChan := fetchProcessor.InitFetchProcessors(state)
	insertRequestChan := insertProcessor.InitInsertProcessors(state)
	insertLogChan := insertLogProcessor.InitInsertLogProcessors()
	requestChan = make(chan *dto.FormattedRequest, 10)
	for i := 0; i < 10; i++ {
		go processor(requestChan, fetchRequestChan, insertRequestChan, insertLogChan)
	}
}

// AddRequest used for publishing data
func AddRequest(req *dto.FormattedRequest) {
	requestChan <- req
	return
}
