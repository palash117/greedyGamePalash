package insertProcessor

import (
	"greedyGamePalash/dto"
	"greedyGamePalash/state"
	"sync"
	"testing"
)

func Test_process(t *testing.T) {
	headState := &state.Node{LevelType: "global", LevelValue: "global", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}

	request := dto.FormattedRequest{
		RequestData: dto.FormattedRequestData{
			Dim: []dto.Level{
				{
					LevelName:  "country",
					LevelValue: "in",
				},
				{
					LevelName:  "device",
					LevelValue: "web",
				},
			},
			Webreq:    10,
			Timespent: 20,
		},
	}
	performInsertion(headState, &request)

	child, present := headState.Children[state.GetKey("country", "in")]
	if !present {
		t.Fail()
	}

	if !(*(child.Webreq) == (uint64(10)) && *(child.Timespent) == (uint64(20))) {

		t.Fatalf("incorrect %+v\n , %v,%v , expected: %v, %v", child, int(*(child.Webreq)), int(*(child.Timespent)), 10, 20)
		t.Fail()
	}
	child2, present := child.Children[state.GetKey("device", "web")]
	if !present {
		t.Fail()
	}

	if !(*(child2.Webreq) == (uint64(10)) && *(child2.Timespent) == (uint64(20))) {
		t.Fatalf("incorrect %+v\n , %v,%v , expected: %v, %v", child2, int(*(child2.Webreq)), int(*(child2.Timespent)), 10, 20)
		t.Fail()
	}
}

func Test_process2(t *testing.T) {
	headState := &state.Node{LevelType: "global", LevelValue: "global", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}

	request := dto.FormattedRequest{
		RequestData: dto.FormattedRequestData{
			Dim: []dto.Level{
				{
					LevelName:  "country",
					LevelValue: "in",
				},
				{
					LevelName:  "device",
					LevelValue: "web",
				},
			},
			Webreq:    10,
			Timespent: 20,
		},
	}
	performInsertion(headState, &request)
	request = dto.FormattedRequest{
		RequestData: dto.FormattedRequestData{
			Dim: []dto.Level{
				{
					LevelName:  "country",
					LevelValue: "in",
				},
				{
					LevelName:  "device",
					LevelValue: "mobile",
				},
			},
			Webreq:    15,
			Timespent: 25,
		},
	}

	performInsertion(headState, &request)

	child, present := headState.Children[state.GetKey("country", "in")]
	if !present {
		t.Fail()
	}

	if !(*(child.Webreq) == (uint64(25)) && *(child.Timespent) == (uint64(45))) {

		t.Fatalf("incorrect %+v\n , %v,%v , expected: %v, %v", child, int(*(child.Webreq)), int(*(child.Timespent)), 25, 45)
		t.Fail()
	}
	_, present = child.Children[state.GetKey("device", "web")]
	if !present {
		t.Fail()
	}
	_, present = child.Children[state.GetKey("device", "mobile")]
	if !present {
		t.Fail()
	}

}
