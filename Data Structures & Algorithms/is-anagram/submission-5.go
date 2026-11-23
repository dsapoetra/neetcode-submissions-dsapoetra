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

	return slices.Equal(runes, runesT)
}
