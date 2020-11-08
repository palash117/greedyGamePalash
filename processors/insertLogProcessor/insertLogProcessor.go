package insertLogProcessor

import (
	"fmt"
	"greedyGamePalash/dto"
)

var insertLogChan chan *dto.FormattedRequest

// InitInsertProcessors initializes insertProcessor workers
func InitInsertLogProcessors() chan *dto.FormattedRequest {
	insertLogChan = make(chan *dto.FormattedRequest, 10)
	go process(insertLogChan)
	return insertLogChan
}

func process(requestChan chan *dto.FormattedRequest) {

	for {
		req := <-requestChan
		fmt.Printf("%+v\n", req)
	}
}
