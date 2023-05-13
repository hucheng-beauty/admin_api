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

var StatusMapResponse = map[int]string{
	-1: "STOP",
	1:  "Submitted",
	2:  "Created",
	3:  "Invalid",
	4:  "Effective",
	5:  "Expire",
	6:  "Fail",
}

var BelongToMapResponse = map[int]string{
	1: "Alipay",
	2: "WeChat",
}
