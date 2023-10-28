package tasks

import (
	"fmt"
	"log"
)

func Task19() {
	var str string
	log.Print("Введите строку:")
	fmt.Scan(&str)
	log.Printf("Первый способ: %s", reverseString1(str))
	log.Printf("Второй способ: %s", reverseString2(str))
	log.Printf("Третий способ: %s", reverseString3(str))
}

// Первый способ. Создаем строку и в начала строки добавляем элемент из входящей строки
func reverseString1(str string) string {
	s := ""
	for _, v := range str {
		s = string(v) + s
	}
	return s
}

// Второй способ. Обход в обратную сторону срез из рун и добавление в другой срез, после чего конвертируем срез в строку
func reverseString2(str string) string {
	strRunes := []rune(str)
	var reverseStr []rune
	for i := len(strRunes) - 1; i >= 0; i-- {
		reverseStr = append(reverseStr, strRunes[i])
	}
	return string(reverseStr)
}

// Трейти способ. Двигаем два индекса влево и вправо с конца и начала среза рун и просто меняем местами символы
func reverseString3(str string) string {
	strRunes := []rune(str)
	for i, j := 0, len(strRunes)-1; i < j; i, j = i+1, j-1 {
		strRunes[i], strRunes[j] = strRunes[j], strRunes[i]
	}
	return string(strRunes)
}
