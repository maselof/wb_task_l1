package tasks

import (
	"fmt"
	"log"
)

/*var justString string


func someFunc() {
	v := createHugeString(1 << 10) // как я понимаю, переменная, которая передается в функцию это количество символов в строке
	justString = v[:100]           // получает первые 100 символов из строки
	log.Printf("%s", justString)
}

func createHugeString(val int) string {
	str := ""
	for i := 0; i < val; i++ {
		str += fmt.Sprintf("%d", i)
	}
	return str
}

func Task15() {
	someFunc()
}*/

/* Объявление глобальной переменны это уже неприятная ситуация, так как могут возникнуть проблемы с названиями переменны
в данном пакете, изменение переменной, обращение к ней. Кратко говоря, все что связано с чистотой кода будет испорчено*/

/* Еще в коде мы затрачиваем много памяти для хранения строки в таком размере, просто для того, чтобы передать
глобальной переменне первые 100 символом, что тоже не очень хорошо. И сборщик мусора не сможет убрать эту огромную переменную
так как на неё ссылается глобальная переменная*/

func someFunc() string {
	var justString string // Создаем переменную внутри функции
	v := createHugeString(1 << 10)
	justString = v[:100]
	return justString
}

func createHugeString(val int) string {
	str := ""
	for i := 0; i < val; i++ {
		str += fmt.Sprintf("%d", i)
	}
	return str
}

func Task15() {
	str := someFunc()
	log.Printf("%s", str)
}
