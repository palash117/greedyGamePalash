package fetchProcessor

import (
	"greedyGamePalash/dto"
	"greedyGamePalash/state"
	"sync/atomic"
)

var fetchRequestChan chan *dto.FormattedRequest

func InitFetchProcessors(state *state.Node) chan *dto.FormattedRequest {
	fetchRequestChan = make(chan *dto.FormattedRequest, 10)
	go process(fetchRequestChan, state)
	return fetchRequestChan
}

func process(requestChan chan *dto.FormattedRequest, currentstate *state.Node) {

	for {
		req := <-requestChan
		ansState := state.GetInitialState()
		search(currentstate, req.RequestData.Dim, 0, ansState)
		req.RespChannel <- ansState
	}
}

func search(node *state.Node, levels []dto.Level, traversalIndex int, ans *state.Node) {
	if traversalIndex == len(levels) || node == nil {
		return
	}
	currentLevel := levels[traversalIndex]

	if currentLevel.LevelName == node.LevelType && currentLevel.LevelValue == node.LevelValue {
		if traversalIndex == len(levels)-1 {

			atomic.AddUint64(ans.Webreq, *node.Webreq)
			atomic.AddUint64(ans.Timespent, *node.Timespent)

		} else {
			nextLevel := levels[traversalIndex+1]
			child, present := node.Children[state.GetKey(nextLevel.LevelName, nextLevel.LevelValue)]
			if present {
				search(child, levels, traversalIndex+1, ans)
			} else {
				for _, child := range node.Children {
					search(child, levels, traversalIndex+1, ans)
				}
			}
		}
	} else {
		child, present := node.Children[state.GetKey(currentLevel.LevelName, currentLevel.LevelValue)]
		if present {
			search(child, levels, traversalIndex, ans)
		} else {
			for _, child := range node.Children {
				search(child, levels, traversalIndex, ans)
			}
		}
	}

}
