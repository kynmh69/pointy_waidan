package utils

import (
	"fmt"
	"pointy/model"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReplaceTitle(a *model.Article) *model.Article {
	a.Title = strings.ReplaceAll(a.Title, "<title>", "")
	a.Title = strings.ReplaceAll(a.Title, "</title>", "")
	fmt.Println("Replace str: ", a)
	return a
}
