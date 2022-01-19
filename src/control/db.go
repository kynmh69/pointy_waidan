package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"pointy/model"

	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=pointy sslmode=disable")

	if err != nil {
		fmt.Println(err)
		return
	}
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
