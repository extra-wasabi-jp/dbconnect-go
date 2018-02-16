package main

import (
    "fmt"
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
    "./entity"
)

func main() {
 
    fmt.Printf("Start\n")

    db, err := sql.Open("sqlite3", "./sqlite3.db")
    if err != nil {
        fmt.Printf("DB Open Error\n")
        panic(err.Error())
    }

    fmt.Printf("Db Opened\n")

    dbmap := &gorp.DbMap {Db: db,  Dialect: gorp.SqliteDialect{}}

    list, err := dbmap.Select(entity.Customer{}, "SELECT * FROM customer")

    if err != nil {
        fmt.Printf("Select Error\n")
        panic(err.Error())
    }

    for _, l := range list {
        cust := l.(*entity.Customer)

        fmt.Printf("%s, %s\n", cust.Customer_cd, cust.Customer_nm)
    }


    err = db.Close()

    if err != nil {
        fmt.Printf("DB Open Error\n")
        panic(err.Error())
    }
}

