
func hasDuplicate(nums []int) bool {
    cont := make(map[int]bool)
    l := len(nums)

    for i := 0; i<l; i++ {
        cont[nums[i]] = true
    }

    fmt.Println(nums)
    fmt.Println(cont)

    if l != len(cont) {
        return true
    }

    return false
}
