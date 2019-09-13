package main

import (
	"github.com/vvdruzev/docx"
	"log"
)

func main() {
	template, err := docx.ReadDocxFile("test.docx")
	if err != nil {
		panic(err)
	}

	docx1 := template.Editable()
	// Что-то делаем с текстовым содержимым
	ctx := map[string]interface{}{
		"title": "Hello World",
		"body":  map[string]string{"content": "Hi"},
		"name": "Peter",
	}
	err = docx1.Render(ctx)

	if err != nil {
		log.Fatal(err)
	}

	//
	docx1.WriteToFile("./new_result_1.docx")

}
