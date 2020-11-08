package dto

// RequestType is for differetiating between insert and fetch requests
type RequestType int

const (
	// INSERT insert type
	INSERT RequestType = iota
	// FETCH fetch type
	FETCH
)
