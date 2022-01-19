package model

import (
	"fmt"
)

type Article struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func (a *Article) GetUrl() string {
	return fmt.Sprintf("https://thewaidan.studio.site/%d", a.Number)
}

func (a *Article) PrintTitle() {
	fmt.Println("title: ", a.Title)
}

func New(number int, title string) *Article {
	return &Article{Number: number, Title: title}
}
