package tasks

import "log"

type SomeData interface{} // Указываем интерфейс

func Type(sm SomeData) { // Создаем функцию для определения и вывода типа данных
	log.Printf("Type: %T", sm)
}

func Task14() {
	var sm SomeData

	sm = 1
	Type(sm)
	sm = "a"
	Type(sm)
	sm = false
	Type(sm)
	sm = make(chan int)
	Type(sm)
}
