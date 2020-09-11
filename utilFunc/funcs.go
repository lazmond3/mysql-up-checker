package funcs

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

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

func OpenSql(arg util.ArgType) bool {
	dbLine := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=skip-verify&autocommit=true",
		arg.User,
		arg.Password,
		arg.Host,
		strconv.Itoa(arg.Port),
		arg.Dbname,
	)

	db, err := sql.Open("mysql", dbLine)
	if err != nil {
		println(err)
		return false
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return true
}
