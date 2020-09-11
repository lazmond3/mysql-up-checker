package funcs

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	util "github.com/lazmond3/mysql-up-checker/utilType"
)

func GetArg() util.ArgType {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 4 { // host, port, user, password, dbname
		panic("error: usage: host, port, user,  password, dbname")
	}

	var portStr = argsWithoutProg[1]
	var portNum, _ = strconv.Atoi(portStr)

	return util.ArgType{
		Host:     argsWithoutProg[0],
		Port:     portNum,
		User:     argsWithoutProg[2],
		Password: argsWithoutProg[3],
		Dbname:   argsWithoutProg[4],
	}
}

func OpenSql(arg util.ArgType) (string, bool) {
	dbLine := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		arg.User,
		arg.Password,
		arg.Host,
		strconv.Itoa(arg.Port),
		arg.Dbname,
	)
	if len(arg.Password) == 0 {
		dbLine = fmt.Sprintf("%s@tcp(%s:%s)/%s",
			arg.User,
			arg.Host,
			strconv.Itoa(arg.Port),
			arg.Dbname,
		)
	}

	db, err := sql.Open("mysql", dbLine)
	defer db.Close()
	if err != nil {
		msg := "db open failed: data source url is wrong."
		return msg, false
	}
	err = db.Ping()
	if err != nil {
		return "Ping failed.", false
	}

	return "succeeded!", true
}
