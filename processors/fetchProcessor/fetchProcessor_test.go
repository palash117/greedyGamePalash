package fetchProcessor

import (
	"fmt"
	"greedyGamePalash/dto"
	"greedyGamePalash/state"
	"sync"
	"testing"
)

var headState *state.Node

func init() {
	fmt.Println("init called")
	headState = &state.Node{LevelType: "global", LevelValue: "global", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	child1 := &state.Node{LevelType: "country", LevelValue: "in", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	child2 := &state.Node{LevelType: "country", LevelValue: "us", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	child3 := &state.Node{LevelType: "device", LevelValue: "web", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	child4 := &state.Node{LevelType: "device", LevelValue: "mobile", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	child5 := &state.Node{LevelType: "device", LevelValue: "mobile", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}

	headState.Children[state.GetKey("country", "in")] = child1
	headState.Children[state.GetKey("country", "us")] = child2
	child1.Children[state.GetKey(child3.LevelType, child3.LevelValue)] = child3
	child1.Children[state.GetKey(child5.LevelType, child5.LevelValue)] = child5
	child2.Children[state.GetKey(child4.LevelType, child4.LevelValue)] = child4

	child5.Webreq = getUint(15)
	child5.Timespent = getUint(25)
	child4.Webreq = getUint(10)
	child4.Timespent = getUint(20)
	child3.Webreq = getUint(30)
	child3.Timespent = getUint(50)
	child1.Webreq = getUint(45)
	child1.Timespent = getUint(75)
	child2.Webreq = getUint(10)
	child2.Timespent = getUint(20)
}
func Test_search(t *testing.T) {

	// headState.Webreq = 40
	// headState.Timespent =
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}

	search(headState, []dto.Level{{LevelName: "country", LevelValue: "in"}, {LevelName: "device", LevelValue: "web"}}, 0, ans)
	if *(ans.Webreq) != (uint64(30)) || *(ans.Timespent) != (uint64(50)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
	}

}

func Test_search2(t *testing.T) {

	// headState.Webreq = 40
	// headState.Timespent =
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}

	search(headState, []dto.Level{{LevelName: "country", LevelValue: "in"}}, 0, ans)
	if *(ans.Webreq) != (uint64(45)) || *(ans.Timespent) != (uint64(75)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
	}

}

func Test_search3(t *testing.T) {
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	search(headState, []dto.Level{{LevelName: "country", LevelValue: "us"}}, 0, ans)
	if *(ans.Webreq) != (uint64(10)) || *ans.Timespent != (uint64(20)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
	}
}

func Test_search4(t *testing.T) {
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	search(headState, []dto.Level{{LevelName: "device", LevelValue: "web"}}, 0, ans)
	if *(ans.Webreq) != (uint64(30)) || *ans.Timespent != (uint64(50)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
	}
}

func Test_search5(t *testing.T) {
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	search(headState, []dto.Level{{LevelName: "device", LevelValue: "mobile"}}, 0, ans)
	if *(ans.Webreq) != (uint64(25)) || *ans.Timespent != (uint64(45)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
	}
}

func getUint(i int) *uint64 {
	temp := uint64(i)
	return &temp
}
