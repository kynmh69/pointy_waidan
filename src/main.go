package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pointy/control"
	"pointy/handler"
	"pointy/utils"
)

var templates = make(map[string]*template.Template)

func updateDB() {
	control.CreateTable("pointy")

	max_id := 0
	article := control.GetLastRow()
	if article != nil {
		max_id = article.Number
	}
	fmt.Println("max id: ", max_id)

	// htmlから情報吸い出し
	article_list := *utils.GetArticle(max_id)
	if len(article_list) <= 0 {
		fmt.Println("空リスト")
	}

	// テーブルに挿入
	control.InsertData(&article_list)
	// 全てのデータ取得
	article_list = *control.GetAllData()
	for _, v := range article_list {
		fmt.Println(v.GetUrl())
	}

}
func main() {
	updateDB()

	templates["index"] = handler.LoadTemplate("index")

	port := "8080"
	http.HandleFunc("/", handler.HandleIndex)
	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}
