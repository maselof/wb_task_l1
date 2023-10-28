package tasks

import "log"

func Task23() {
	slice := make([]int, 10) // Создаем и заполняем срез
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	log.Print(slice) // Выводим начальный срез
	ind := 5         // Указываем индекс, который хотим удалить

	log.Print(remove(slice, ind)) // Выводим получившийся срез
}

func remove(slice []int, ind int) []int {
	return append(slice[:ind], slice[ind+1:]...)
}
