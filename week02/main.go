package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func fakeQueryFromSql() ([]string, error) {
	return make([]string, 0), sql.ErrNoRows
}


func fakeQuery() ([]string,error) {
	data, err := fakeQueryFromSql()
	if err != nil {
		if err == sql.ErrNoRows{
			err = fmt.Errorf("未查询到相关数据，%w",err)
		} else{
			err = fmt.Errorf("查询失败，%w",err)
		}
	}
	return data, err
}

func main() {
	data, err := fakeQuery()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, item := range data {
		fmt.Println(item)
	}
}
