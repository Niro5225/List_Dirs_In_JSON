package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Add_file_JSON(name string, value string, JSON map[string]string) {
	//Проверка существует ли дагный ключ
	if len(JSON[name]) == 0 {
		JSON[name] = value
	} else {
		JSON[name] += "," + value
	}
}

func main() {
	Res := map[string]string{}
	fmt.Print("Enter path:")
	var path string
	fmt.Scanln(&path)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}
		//сортировка по ключам папки и файлы
		if info.IsDir() {
			Add_file_JSON("DIRS:", path, Res)
		} else {
			Add_file_JSON("FILES:", path, Res)
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	} else {
		// Res_json, _ := json.Marshal(Res)
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "		")
		enc.Encode(Res)
	}
}
