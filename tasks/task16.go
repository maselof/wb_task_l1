package tasks

import (
	"log"
)

func Task16() {
	slice := []int{5, 6, 32, 7, 35, 3, 4, 34, 2, 8, 7} // Создаем неотсортированный срез
	log.Print(quickSort(slice, 0, len(slice)-1))       // Запускаем функцию и выводим данные
}

// Сортировка является рекурсивной. Для сортировки относительно pivot мы создали еще одну функцию partition
func quickSort(slice []int, l int, r int) []int {
	if r > l {
		slice, m := partition(slice, l, r)
		quickSort(slice, l, m-1)
		quickSort(slice, m+1, r)
		log.Print(".")
	}
	return slice
}

func partition(slice []int, l int, r int) ([]int, int) {
	pivot := slice[r] // Опорным элементом будет самый правый
	m := l            // Вспомогательная индекс, который мы будем перемещать вправо, когда будет перемещать переменные, которые меньше pivot
	for i := l; i < r; i++ {
		if slice[i] < pivot {
			slice[i], slice[m] = slice[m], slice[i] // Меняем местами, если элементом меньше pivot
			m++                                     // Передвигаем индекс
		}
	}
	slice[m], slice[r] = slice[r], slice[m]
	return slice, m
}
