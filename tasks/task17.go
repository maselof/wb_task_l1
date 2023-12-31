package tasks

import (
	"fmt"
	"log"
)

func Task17() {
	var target int                           // Создаем переменную, которую будем искать
	slice := []int{1, 2, 3, 4, 7, 9, 23, 53} // Создаем отсортированный слайс, так как этот алгоритм работает только так
	log.Print("Введите искомое значение: ")
	fmt.Scan(&target)                      // Вводим значение искомой переменной
	log.Print(binarySearch(slice, target)) // Выводим и вызываем функцию
}

func binarySearch(slice []int, target int) (int, bool) {
	l, r := 0, len(slice)-1 // Указываем левый и правый индекс слайса, то есть 0 и длина слайса - 1

	for l <= r { // Запускаем цикл бинарного поиска
		m := (l + r) / 2       // Вычисляем индекс среднего элемента
		if slice[m] < target { // Если искомая больше среднего элемента, то мы сдвигаем левый индекс, а иначе правый
			l = m + 1
		} else if slice[m] > target {
			r = m - 1
		} else { // При нахождении возвращаем индекс и флаг
			return m, true
		}
	}
	return -1, false // Вывод если не найдем
}
