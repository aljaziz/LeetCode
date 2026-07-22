import "strings"

func findWords(words []string) []string {
    qwe, asd, zxc := 0, 0, 0
    var res []string
    for _, word := range words {
        for _, r := range strings.ToLower(word) {
            if strings.ContainsRune("qwertyuiop", r) {
                qwe++
            }
            if strings.ContainsRune("asdfghjkl", r) {
                asd++
            }
            if strings.ContainsRune("zxcvbnm", r) {
                zxc++
            }
        }
        if qwe == len(word) || asd == len(word) || zxc == len(word) {
            res = append(res, word)
        }
        qwe, asd, zxc = 0, 0, 0
    }
    return res
}