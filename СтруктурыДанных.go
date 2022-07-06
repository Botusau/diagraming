package main

import "unsafe"

type СтруктураФайла struct {
	РабочиеЦентры []typeРабочиеЦентры
	ГрафыОпераций []typeГрафыОпераций
}

type typeРабочиеЦентры struct {
	ВидРабочегоЦентра string
	КлючСтроки        int
	РабочийЦентр      string
}

type typeГрафыОпераций struct {
	ВидРабочегоЦентра       string
	КлючСтроки              int
	НомерОперации           int
	Операция                string
	ПлановоеВремяВыполнения int
	Подразделение           string
	Приоритет               int
	Распоряжение            string
	Уровень                 int
	УровеньОперации         int
	ФинальныйЭтап           string
	Этап                    string
	ЭтапПотребитель         string
}

type typeКлючЭтап struct {
	ФинальныйЭтап string
	Этап          *typeВершина
}

type typeВершина struct {
	Этап          string
	Последователи []unsafe.Pointer
	Удалить       bool
}

type typeЯчейка struct {
	Этап, Операция string
}
