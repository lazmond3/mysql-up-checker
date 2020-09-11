package main

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	funcs "github.com/lazmond3/mysql-up-checker/utilFunc"
)

func main() {
	args := funcs.GetArg()
	for {
		msg, result := funcs.OpenSql(args)
		fmt.Printf("- - - - - status - - :  %s", msg)
		if result {
			fmt.Printf("wait is done!: %s:%s\n", args.Host, strconv.Itoa(args.Port))
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Printf(" failed.\n")
	}
}
