package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Person struct {
	Id string
	Name string
}

// HTMLテンプレートを読み込む関数
func readerTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: "1", Name: "Keisuke"}

	// docker上にコピーしたtemplateファイルを指定する
	parsedTemplate, err := template.ParseFiles("/go/src/github.com/gouser/templates/template.html")
	if err != nil {
		log.Fatalf("Not found file: %v", err)
	}
	// テンプレートにPersonの値を渡す
	err = parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occured while executing the template or writing its output: %v", err)
		return
	}
}

func main() {
	fmt.Println("API Server Start...")
	// ハンドラー登録
	http.HandleFunc("/", readerTemplate)
	// localhost:7070で動くようにする
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
