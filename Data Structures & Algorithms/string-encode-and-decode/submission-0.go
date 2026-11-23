// Idea: write every string as  <length>#<string>.
// Example: ["neet", "code"] -> "4#neet4#code"
// When decoding we read the number up to '#', so we know exactly how
// many characters belong to the string, even if it contains '#' itself.

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, str := range strs {
		encoded += strconv.Itoa(len(str)) + "#" + str
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0
	for i < len(encoded) {
		// find the '#': the digits before it are the length
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])

		// the string starts right after '#' and is "length" characters long
		start := j + 1
		res = append(res, encoded[start:start+length])

		// jump to the beginning of the next piece
		i = start + length
	}
	return res
}
