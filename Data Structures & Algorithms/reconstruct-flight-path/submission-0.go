// Idea: we must use every ticket exactly once, and prefer the smallest
// airport name whenever there is a choice.
//   1) for every airport keep its destinations sorted, biggest name first,
//      so the SMALLEST name is at the end and cheap to take off
//   2) fly (DFS): from an airport, keep taking the smallest unused ticket.
//      Only when an airport has no tickets left do we write it down.
// That writes the route down BACKWARDS (dead ends first), so we reverse
// it at the end.
// Time: O(tickets * log tickets).
func findItinerary(tickets [][]string) []string {
	// 1) destinations per airport, sorted so the smallest name is last
	sort.Slice(tickets, func(a, b int) bool {
		return tickets[a][1] > tickets[b][1]
	})
	destinations := map[string][]string{}
	for _, ticket := range tickets {
		destinations[ticket[0]] = append(destinations[ticket[0]], ticket[1])
	}

	route := []string{}

	// 2) use up all tickets leaving from this airport, then write it down
	var fly func(airport string)
	fly = func(airport string) {
		for len(destinations[airport]) > 0 {
			list := destinations[airport]
			next := list[len(list)-1] // smallest name
			destinations[airport] = list[:len(list)-1]
			fly(next)
		}
		route = append(route, airport)
	}
	fly("JFK")

	// the route was written backwards
	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	return route
}
