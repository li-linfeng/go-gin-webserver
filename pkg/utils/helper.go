package utils

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gohouse/converter"
)

func Dd(i interface{}) {
	bs, _ := json.Marshal(i)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("%v\n", out.String())
}

func DbToOrm(table string, dsn string, SavePath string) {
	err := converter.NewTable2Struct().
		SavePath(SavePath).
		Dsn(dsn).
		// Dsn("用户名:密码@tcp(IP:端口号)/数据库名?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		Table(table).
		Run()

	fmt.Println(err)
}
