
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    s1 := strings.Split(s, "")
    sort.Strings(s1)

    t1 := strings.Split(t, "")
    sort.Strings(t1)

    for i := 0; i < len(s1); i++ {
        if s1[i] != t1[i] {
            return false
        }
    }

    return true
}
