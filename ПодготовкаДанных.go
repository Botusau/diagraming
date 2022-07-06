package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ПодготовитьДанныеИзФайла(Путь string) (СтруктураФайла, error) {

	var Данные СтруктураФайла

	f, err := os.Open(Путь)
	if err != nil {
		return Данные, err
	}

	Файл, err := ioutil.ReadAll(f)
	if err != nil {
		return Данные, err
	}

	err = json.Unmarshal(Файл, &Данные)
	if err != nil {
		return Данные, err
	}

	return Данные, err
}
