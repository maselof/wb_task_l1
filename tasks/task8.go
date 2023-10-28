package tasks

import (
	"fmt"
	"log"
)

func Task8() {
	var val, res int64 // Вводим переменные
	var i int
	log.Print("Введите значение: ")
	fmt.Scan(&val)
	log.Println("Введите позицию: ")
	fmt.Scan(&i)
	res = val
	res |= 1 << i // Ставим 1 на позицию i и если в дальнейшем res и val будут равны, то ставим 0
	switch {
	case val == res:
		res &^= 1 << i                                 // Ставим 0 на позицию i
		log.Printf("Bytes begginer value: %064b", val) // Вывод результата
		log.Printf("Bytes: %064b, val: %d", res, res)
	case val != res:
		log.Printf("Bytes begginer value: %064b", val) // Вывод результата
		log.Printf("Bytes: %064b, val: %d", res, res)
	}
}
