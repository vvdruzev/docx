package main

import (
	"fmt"
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
	docx1.WriteToFile("./test_result.docx")
	fmt.Println("Результат записан в файл test_result.docx ")

}
