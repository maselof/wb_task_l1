package tasks

import "log"

func Task12() {
	set := []string{"cat", "cat", "dog", "cat", "tree"} // Создаем множество
	valMap := make(map[string]int)                      // Создаем map
	for i := 0; i < len(set); i++ {                     // Проходимся по множеству и записываем ключ в map
		_, ok := valMap[set[i]]
		if !ok {
			valMap[set[i]] = 1
		}
		valMap[set[i]]++
	}

	res := make([]string, len(valMap)) // Создаем срез для результата
	i := 0
	// Записываем результаты в res
	for k := range valMap {
		res[i] = k
		i++
	}
	log.Print(res) // Выводим res
}
