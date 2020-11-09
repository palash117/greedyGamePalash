package fetchProcessor

import (
	"greedyGamePalash/dto"
	"greedyGamePalash/state"
	"sync"
	"testing"

	"greedyGamePalash/util"
)

var headState *state.Node

func init() {
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

	child5.Webreq = util.GetUint(15)
	child5.Timespent = util.GetUint(25)
	child4.Webreq = util.GetUint(10)
	child4.Timespent = util.GetUint(20)
	child3.Webreq = util.GetUint(30)
	child3.Timespent = util.GetUint(50)
	child1.Webreq = util.GetUint(45)
	child1.Timespent = util.GetUint(75)
	child2.Webreq = util.GetUint(10)
	child2.Timespent = util.GetUint(20)

	/*
				----> us(req:10,time20) ---->	mobile (req:10,time20)
				||
		global -||
				||							   ||------> mobile (req:15,time25)
				-----> in(req:45,time:75)----->||
											   ||-----> web (req:30,time50)
	*/
}
func Test_countryLevelQueryWithDeviceLevel(t *testing.T) {

	/*
				----> us(req:10,time20) ---->	mobile (req:10,time20)
				||
		global -||
				||							   ||------> mobile (req:15,time25)
				-----> in(req:45,time:75)----->||
											   ||-----> web (req:30,time50)
														⬆️⬆️⬆️⬆️⬆️⬆️⬆️
	*/

	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}

	search(headState, []dto.Level{{LevelName: "country", LevelValue: "in"}, {LevelName: "device", LevelValue: "web"}}, 0, ans)
	if *(ans.Webreq) != (uint64(30)) || *(ans.Timespent) != (uint64(50)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
		t.Fail()
	}

}

func Test_countryLevelQuery_1(t *testing.T) {

	/*
				----> us(req:10,time20) ---->	mobile (req:10,time20)
				||
		global -||
				||							   ||------> mobile (req:15,time25)
				-----> in(req:45,time:75)----->||
						⬆️⬆️⬆️⬆️⬆️⬆️⬆️		||
											   ||-----> web (req:30,time50)
	*/
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}

	search(headState, []dto.Level{{LevelName: "country", LevelValue: "in"}}, 0, ans)
	if *(ans.Webreq) != (uint64(45)) || *(ans.Timespent) != (uint64(75)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
		t.Fail()
	}

}

func Test_countryLevelQuery_2(t *testing.T) {

	/*
				----> us(req:10,time20) ---->	mobile (req:10,time20)
				||	 ⬆️⬆️⬆️⬆️⬆️⬆️⬆️
		global -||
				||							   ||------> mobile (req:15,time25)
				-----> in(req:45,time:75)----->||
												||
											   ||-----> web (req:30,time50)
	*/

	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	search(headState, []dto.Level{{LevelName: "country", LevelValue: "us"}}, 0, ans)
	if *(ans.Webreq) != (uint64(10)) || *ans.Timespent != (uint64(20)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
		t.Fail()
	}
}

func Test_deviceLevelWithoutCountryLevelQuery_1(t *testing.T) {

	/*
				----> us(req:10,time20) ---->	mobile (req:10,time20)
				||
		global -||
				||							   ||------> mobile (req:15,time25)
				-----> in(req:45,time:75)----->||
											   ||-----> web (req:30,time50)
														⬆️⬆️⬆️⬆️⬆️⬆️⬆️
	*/
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	search(headState, []dto.Level{{LevelName: "device", LevelValue: "web"}}, 0, ans)
	if *(ans.Webreq) != (uint64(30)) || *ans.Timespent != (uint64(50)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
		t.Fail()
	}
}

func Test_deviceLevelWithoutCountryLevelQuery_2(t *testing.T) {

	/*
				----> us(req:10,time20) ---->	mobile (req:10,time20)
				||								⬆️⬆️⬆️⬆️⬆️⬆️⬆️
		global -||
				||							   ||------> mobile (req:15,time25)
				-----> in(req:45,time:75)----->||		⬆️⬆️⬆️⬆️⬆️⬆️⬆️
											   ||
											   ||-----> web (req:30,time50)

	*/
	ans := &state.Node{LevelType: "ans", LevelValue: "ans", Webreq: new(uint64), Timespent: new(uint64), Children: make(map[string]*state.Node), Mutex: &sync.Mutex{}}
	search(headState, []dto.Level{{LevelName: "device", LevelValue: "mobile"}}, 0, ans)
	if *(ans.Webreq) != (uint64(25)) || *ans.Timespent != (uint64(45)) {
		t.Fatalf("incorrect %+v\n , %v,%v", ans, int(*(ans.Webreq)), int(*(ans.Timespent)))
		t.Fail()
	}
}
