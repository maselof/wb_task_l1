package tasks

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Task20() {
	var str string
	log.Print("Введите строку: ") // Вводим строку, с помощью пакета bufio
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Print(err) // Обработка ошибки при чтении строки
	}
	str = strings.TrimSpace(input) // Убираем пробелы по в начале строки и в конце

	words := strings.Split(str, " ")              // Разбиваем строку на срез строк
	log.Printf("Result: %s", reverseWords(words)) // Запускаем функцию и выводим данные
}

// В данные функции мы передвигаем два индекса вправо и влево с начала и конца строки, меняя их значения
func reverseWords(words []string) string {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words[:], " ") // Объединяем срез в строку
}
