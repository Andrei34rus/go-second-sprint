package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга данных (строка %d): %v\n", i+1, err)
			continue
		}
		infoStr, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка фщрмирования информации(строка %d): %v\n", i+1, err)
			continue
		}
		fmt.Println(infoStr)
	}
}
