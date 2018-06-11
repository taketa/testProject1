package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"io"
	//"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql","root:root@tcp(127.0.0.1:33061)/world")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)
	top5()
	engEuro()

	//columns,err:=rows.Columns()
	//check(err)
	//fmt.Println(columns)
	}



func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func top5(){
	rows, err := db.Query("SELECT NAME FROM city ORDER BY Population DESC LIMIT 5")
	check(err)
	defer rows.Close()
	fmt.Print("Топ-5 стран по количеству населения в столице: ")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s/", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func engEuro(){
	rows, err := db.Query("SELECT ROUND(SUM(con.Population * cl.Percentage)) AS Total 	FROM `countrylanguage` as cl JOIN country as con ON cl.CountryCode = con.Code WHERE cl.LANGUAGE='English'")
	check(err)
	defer rows.Close()
	fmt.Print("Суммарное кол-во людей, говорящих на английском языке в Европе: ")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}