package time

import "time"

/*
	公用常量:
		CustomizeFormat: 自定义时间格式化格式
		StandardFormat: 标准库时间格式化格式
*/

var StandardFormat = []string{
	"2006-01-02",
	"2006-01-02 15:04:05.000",
	"Mon Jan _2 15:04:05 MST 2006",
	"Time is: 03:04:05 PM",
	"2006-01-02T15:04:05.000000000Z07:00 MST",
	"2006-01",
	"2006",
}

// ToStandardFormat time.Time{}.Format("layout")
func ToStandardFormat(t *time.Time, format string) string {
	return t.Format(format)
}
