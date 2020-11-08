package dto

// Request is request
type Request struct {
	Dim     []KeyVal        `json:"dim"`
	Metrics []KeyWithIntVal `json:"metrics"`
}

// KeyVal struct
type KeyVal struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

// KeyWithIntVal struct
type KeyWithIntVal struct {
	Key string `json:"key"`
	Val int    `json:"val"`
}
