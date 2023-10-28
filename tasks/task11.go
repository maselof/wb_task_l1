package tasks

import "log"

func Task11() {
	set1 := []int{1, 3, 4, 2, 4, 9} // Создаем срезы с двумя множествами
	set2 := []int{3, 1, 6, 6, 8, 5}

	valMap := make(map[int]int) // Создаем map

	// Проходим по первому массиву и создаем ключ со значением 0
	for i := 0; i < len(set1); i++ {
		valMap[set1[i]] = 0
	}

	// Проходим по второму срезу и инкремируем значение по ключу, если такой ключ есть в map
	for i := 0; i < len(set2); i++ {
		_, ok := valMap[set2[i]]
		if ok {
			valMap[set2[i]]++
		}
	}

	// Выводим только те ключи, у которых значение 1
	for k, v := range valMap {
		if v == 1 {
			log.Print(k)
		}
	}
}
