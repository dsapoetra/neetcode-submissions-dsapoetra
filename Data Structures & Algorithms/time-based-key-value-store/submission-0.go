// Idea: for every key keep a list of (timestamp, value) entries.
// The problem promises that Set is called with increasing timestamps, so
// each list is already sorted by time.
// Get must find the LAST entry whose timestamp is <= the asked timestamp,
// which is a binary search on that sorted list.
// Set: O(1). Get: O(log n).

type timeEntry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	entries map[string][]timeEntry
}

func Constructor() TimeMap {
	return TimeMap{entries: map[string][]timeEntry{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.entries[key] = append(this.entries[key], timeEntry{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list := this.entries[key]
	answer := "" // stays empty if nothing is old enough

	left := 0
	right := len(list) - 1
	for left <= right {
		mid := (left + right) / 2
		if list[mid].timestamp <= timestamp {
			answer = list[mid].value // a candidate; look for a later one
			left = mid + 1
		} else {
			right = mid - 1 // too new
		}
	}
	return answer
}
