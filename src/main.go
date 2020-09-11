package main

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	funcs "github.com/lazmond3/mysql-up-checker/src/utilFunc"
)

func main() {
	args := funcs.GetArg()
	for {
		result := funcs.OpenSql(args)
		if result {
			fmt.Printf("wait is done!: %s:%s\n", args.Host, strconv.Itoa(args.Port))
			return
		}
		time.Sleep(time.Second * 1)
	}
}
