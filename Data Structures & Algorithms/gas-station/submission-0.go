// Idea (greedy):
//   1) if the total gas is less than the total cost, no start can work
//   2) otherwise an answer exists. Drive from station 0 and track the
//      tank. If the tank goes below 0 at station i, then NO station from
//      our start up to i can be the start (they would all arrive here
//      with even less gas). So try again with start = i + 1.
// Time: O(n).
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	if totalGas < totalCost {
		return -1
	}

	start := 0
	tank := 0
	for i := range gas {
		tank += gas[i] - cost[i]
		if tank < 0 {
			// cannot get past station i: start fresh after it
			start = i + 1
			tank = 0
		}
	}
	return start
}
