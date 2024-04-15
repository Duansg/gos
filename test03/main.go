package main

import (
	"fmt"
	_ "modernc.org/sqlite"
	"path/filepath"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

func main() {
	//driver := string(schemas.SQLITE)
	driver := "sqlite"
	connection := "/Users/duansg/app/sqlite/answer.db"
	engine, err := xorm.NewEngine(driver, connection)
	if err != nil {
		fmt.Println(err)
	}
	if err = engine.Ping(); err != nil {
		fmt.Println(err)
	}
	name := filepath.Join("/Users/duansg/app/go-workspace/gos/test03", fmt.Sprintf("answer_dump_data_%s.sql", time.Now().Format("2006-01-02")))
	engine.DumpAllToFile(name, schemas.SQLITE)
}
