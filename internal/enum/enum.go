package enum

// CodeMapRequest 1-Fail,0-Success
var CodeMapRequest = map[string]int{
	"Fail":    1,
	"Success": 0,
}

var CodeMapResponse = map[int]string{
	1: "Fail",
	0: "Success",
}
