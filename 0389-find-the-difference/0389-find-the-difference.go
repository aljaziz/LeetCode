func findTheDifference(s string, t string) byte {
    hashMap := [26]int{}
    for i := 0; i < len(s); i++ {
        hashMap[s[i]-'a']++
    }
    var ans byte
    for i := 0; i< len(t); i++ {
        if hashMap[t[i]-'a'] == 0 {
            ans = t[i]
        }
        hashMap[t[i]-'a']--
    }
    return ans
}