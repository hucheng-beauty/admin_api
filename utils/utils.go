package utils

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"sync"
	"time"
)

type generator struct {
	lastFlag string
	loc      sync.Mutex
	c        int
}

func NewGenerator() *generator {
	return &generator{}
}

var DefaultGenerator = GenerateNu()

/*
var c int64

	db := db.Table("users").Debug()
	err := db.Where(Filter(db, "name+eq+abc+and+created_at+gt+2023-03-20 07:59:20.061956 +00:00")).Count(&c).Error
*/
func Filter(db *gorm.DB, fs string) *gorm.DB {
	q := fsToQuery(fs)
	//查询数据库活动
	for k, item := range q {
		db = db.Where(k, item)
	}
	return db
}

func (g *generator) gen() string {
	g.loc.Lock()

	g.c++

	n := time.Now()
	s := fmt.Sprintf("%s%s%s",
		string(byte(time.Now().Hour()%24+65)),
		fmt.Sprintf("%02d", n.Minute()),
		fmt.Sprintf("%02d", n.Second()),
	)
	if g.lastFlag == s && g.c >= 1000 {
		time.Sleep(time.Second)
		g.loc.Unlock()
		return g.gen()
	}
	g.lastFlag = s
	if g.c >= 1000 {
		g.c = 1
	}
	s = s + fmt.Sprintf("%03d", g.c)
	g.loc.Unlock()
	return s
}

// 生成活动编号
func GenerateNu() func(prefix string) string {
	g := NewGenerator()
	return func(prefix string) string {
		n := time.Now()
		return fmt.Sprintf("%s%s%02d%02d%s",
			prefix,
			strconv.Itoa(n.Year())[2:],
			n.Month(),
			n.Day(),
			g.gen())
	}
}

var operatorMap = map[string]string{
	"ne":   " != ?",
	"eq":   " = ?",
	"gt":   " > ?",
	"ge":   " >= ?",
	"lt":   " < ?",
	"le":   " <= ?",
	"like": " like ?",
	"in":   " in ?",
}

func fsToQuery(fs string) map[string]interface{} {
	q := map[string]interface{}{}
	//用and/+and+分割
	res := strings.Split(fs, " and ")
	if len(res) == 1 {
		res = strings.Split(fs, "+and+")
	}
	//
	for _, item := range res {
		queryItem := strings.SplitN(item, "+", 3)
		if len(queryItem) != 3 {
			queryItem = strings.SplitN(item, " ", 3)
		}
		if len(queryItem) != 3 {
			continue
		}
		field := queryItem[0]
		var value interface{}
		operator := operatorMap[queryItem[1]]
		if field == `` || operator == `` || value == `` {
			continue
		}

		switch queryItem[1] {
		case "like":
			value = "%" + queryItem[2] + "%"
		case "in":
			value = strings.Split(queryItem[2], ",")
		default:
			value = queryItem[2]
		}
		key := fmt.Sprintf("%s%s", field, operator)
		q[key] = value
	}
	//返回数据库和查询条件
	return q
}
