import (
	"slices"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	runes := []rune(s)
	runesT := []rune(t)
	slices.Sort(runes)
	slices.Sort(runesT)

	for i := 0; i < len(runes); i++ {
		if runes[i] != runesT[i] {
			return false
		}
	}

	return true
}
