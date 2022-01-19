package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"pointy/model"

	_ "github.com/lib/pq"
)

const (
	DRIVER_NAME = "postgres"
	IP_ADDRESS  = "127.0.0.1"
	USER        = "postgres"
	PASS        = "postgres"
	DB_NAME     = "postgres"
)

func WriteJson(article_list *model.Article) error {
	f, _ := os.Create("./output.json")
	defer f.Close()
	err := json.NewEncoder(f).Encode(article_list)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	connectionInfo := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", IP_ADDRESS, USER, PASS, DB_NAME)
	db, err := sql.Open(DRIVER_NAME, connectionInfo)

	utils.checkError(err)

	defer db.Close()

	rows, err := db.Query("SELECt * FROM pointy")
	if err != nil {
		fmt.Println(err)
		return
	}
	var articles []model.Article
	for rows.Next() {
		var a model.Article
		rows.Scan(&a.Number, &a.Title)
		articles = append(articles, a)
	}
	fmt.Printf("%v", articles)

}
