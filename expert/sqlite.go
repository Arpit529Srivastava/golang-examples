package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Table struct {
	Id   int
	Text string
	Time int64
}

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.sqlite")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	mode := os.Args[1]
	var qu Table
	if mode == "select" {
		q, err := db.Query("SELECT * FROM \"main\".\"test\" ORDER BY time DESC LIMIT 1;")
		if err != nil {
			fmt.Printf("Error while getting data: %s\n", err)
		} else {
			q.Next()
			err = q.Scan(&qu.Id, &qu.Text, &qu.Time)
			if err != nil {
				fmt.Printf("Error while getting row data: %s\n", err)
			}
			fmt.Println("Last inserted data is: " + qu.Text)
			t := time.Unix(qu.Time, 0)
			fmt.Println("Inserted on " + t.Format(time.RFC1123Z))
		}
	} else {
		if len(os.Args) > 2 {
			data := os.Args[2]
			now := strconv.FormatInt(time.Now().Unix(), 10)
			db.Exec("INSERT INTO \"main\".\"test\" (\"text\", \"time\") VALUES('" + data + "', '" + now + "')")
			fmt.Println(data + " inserted")
		}
	}
}
