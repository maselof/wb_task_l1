package tasks

import (
	"fmt"
	"log"
	"strings"
)

func Task26() {
	var str string // Создаем строку
	res := true    // Создаем переменную для вывода результата
	log.Print("Введите строку: ")
	fmt.Scan(&str)
	str = strings.ToLower(str)                    // Переводим все символы в нижний регистр
	strRune := []rune(str)                        // Конвертируем строку в срез из рун
	helperMap := make(map[rune]int, len(strRune)) // Создаем словарь для того, чтобы считать количество вхождение символа в строку
	for _, ch := range strRune {                  // Проходимся по срезу
		_, ok := helperMap[ch]
		if !ok {
			helperMap[ch] = 0
		}
		helperMap[ch]++
		if helperMap[ch] > 1 {
			res = false // Если символ встречается уже не в первый раз
			break
		}
	}
	log.Print(res) // Выводим результат

}
