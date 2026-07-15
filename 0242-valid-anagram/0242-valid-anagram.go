func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    hashArr := [26]int{0}
    for i:= 0; i< len(s); i++ {
        hashArr[s[i]-'a']++
        hashArr[t[i]-'a']--
    }
    for _, num := range hashArr{
        if num > 0 {
            return false
        }
    }
    return true
}