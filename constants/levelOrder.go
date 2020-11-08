package constants

type LevelOrder int

const (
	GLOBAL  LevelOrder = 0
	COUNTRY LevelOrder = 1
	DEVICE  LevelOrder = 2
)

// LevelOrderFromString Level Order From String
var LevelOrderFromString = map[string]LevelOrder{
	"global":  GLOBAL,
	"country": COUNTRY,
	"device":  DEVICE,
}

// const LevelOrderToString = map[LevelOrder]string{
// 	GLOBAL:  "global",
// 	COUNTRY: "country",
// 	DEVICE:  "device",
// }
