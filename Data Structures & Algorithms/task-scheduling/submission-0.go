// Idea: the task that appears most often decides the schedule.
// Say A appears maxCount times and the cooldown is n:
//
//	A _ _ A _ _ A        (n = 2, maxCount = 3)
//
// That is (maxCount - 1) blocks of length (n + 1), plus the last A.
// If several tasks appear maxCount times, they all sit in the last block:
//
//	A B _ A B _ A B
//
// The other tasks fill the gaps. If there are more tasks than gaps, no
// idle time is needed at all and the answer is just the number of tasks.
// Time: O(n).
func leastInterval(tasks []byte, n int) int {
	var count [26]int
	for _, task := range tasks {
		count[task-'A']++
	}

	// the highest count, and how many tasks have that count
	maxCount := 0
	tasksWithMax := 0
	for _, c := range count {
		if c > maxCount {
			maxCount = c
			tasksWithMax = 1
		} else if c == maxCount {
			tasksWithMax++
		}
	}

	slots := (maxCount-1)*(n+1) + tasksWithMax
	if len(tasks) > slots {
		return len(tasks)
	}
	return slots
}
