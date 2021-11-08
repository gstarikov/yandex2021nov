package main

func simpleCalsDays(X, K int, tasks []int) int {
	for day := 1; ; day++ {
		//fmt.Printf("day(%3d) X(%2d) K(%2d) tasks%+v\n", day, X, K, tasks)
		var notify bool
		for i := range tasks {
			if tasks[i] == day {
				notify = true
				tasks[i] += X
			}
		}
		if notify {
			K--
			if K <= 0 {
				return day
			}
		}

	}
}
