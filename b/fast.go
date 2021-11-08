package main

import (
	"sort"
)

func fastCalsDays(X, K int, tasks []int) int {

	//дедупликация
	tasksMap := make(map[int]struct{})
	for _, v := range tasks {
		tasksMap[v] = struct{}{}
	}
	tasks = tasks[:0]
	for k := range tasksMap {
		tasks = append(tasks, k)
	}

	// упорядочивание
	sort.Ints(tasks)

	for day := tasks[0]; ; day = tasks[0] {
		//fmt.Printf("day[%2d] X[%2d] K[%2d] tasks%+v\n", day, X, K, tasks)

		K--
		if K <= 0 {
			return day
		}

		target := day + X

		if len(tasks) == 1 { // вырожденный случай
			tasks[0] = target
			continue
		}

		index := sort.SearchInts(tasks, target)

		switch {
		case index >= len(tasks): // надо добавить после последнего конец
			copy(tasks, tasks[1:])
			tasks[index-1] = target
		case target == tasks[index]: // попали в уже существующий элемент, ничего вставлять не надо
			copy(tasks, tasks[1:])       // сдвинули очередь
			tasks = tasks[:len(tasks)-1] // удалили последний элемент (он и так уже скопирован в предпоследний)
		default: // новый элемент
			copy(tasks, tasks[1:index+1]) // сдвинули очередь до места вставки
			tasks[index-1] = target       // вставили
			// хвост не трогаем
		}
	}
}
