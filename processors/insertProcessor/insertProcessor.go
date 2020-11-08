package insertProcessor

import (
	"greedyGamePalash/dto"
	"greedyGamePalash/state"
	"sync"
	"sync/atomic"
)

var insertRequestChan chan *dto.FormattedRequest

// InitInsertProcessors initializes insertProcessor workers
func InitInsertProcessors(state *state.Node) chan *dto.FormattedRequest {
	insertRequestChan = make(chan *dto.FormattedRequest, 10)
	go process(insertRequestChan, state)
	return insertRequestChan
}

func process(requestChan chan *dto.FormattedRequest, rootstate *state.Node) {

	for {
		req := <-requestChan
		currentNode := rootstate
		atomic.AddUint64(currentNode.Webreq, req.RequestData.Webreq)
		atomic.AddUint64(currentNode.Timespent, req.RequestData.Timespent)
		for _, level := range req.RequestData.Dim {

			currentNode.Mutex.Lock()
			child, present := currentNode.Children[state.GetKey(level.LevelName, level.LevelValue)]
			if !present {
				child = &state.Node{LevelType: level.LevelName, LevelValue: level.LevelValue, Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
				currentNode.Children[state.GetKey(level.LevelName, level.LevelValue)] = child
			}
			currentNode.Mutex.Unlock()

			atomic.AddUint64(child.Webreq, req.RequestData.Webreq)
			atomic.AddUint64(child.Timespent, req.RequestData.Timespent)
			currentNode = child

		}
	}
}
