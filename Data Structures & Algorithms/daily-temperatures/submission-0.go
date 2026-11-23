// Idea: keep a stack of days that are still WAITING for a warmer day.
// The temperatures of the waiting days get lower towards the top.
// For each new day: every waiting day that is colder than today has found
// its answer (today), so pop it and write down the distance.
// Then today starts waiting too.
// Time: O(n), every day is pushed once and popped at most once.
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures)) // 0 = no warmer day
	waiting := []int{}                    // indexes of days waiting for a warmer day

	for today, temp := range temperatures {
		// today is warmer than the day on top of the stack
		for len(waiting) > 0 && temperatures[waiting[len(waiting)-1]] < temp {
			day := waiting[len(waiting)-1]
			waiting = waiting[:len(waiting)-1]
			res[day] = today - day
		}
		waiting = append(waiting, today)
	}
	return res
}
