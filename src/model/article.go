package model

import "fmt"

type Article struct {
	number int
	title  string
}

func (a *Article) GetUrl() string {
	return fmt.Sprintf("https://thewaidan.studio.site/%d", a.number)
}

func (a *Article) PrintTitle() {
	fmt.Println("title: ", a.title)
}

func New(number int, title string) *Article {
	return &Article{number: number, title: title}
}
