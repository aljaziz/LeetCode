import "strings"

type Key struct {
	name  byte
	value string
}

func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	if len(pattern) != len(split) {
		return false
	}

	hashMap := map[Key]string{}
	for i := 0; i < len(pattern); i++ {
		key_p := Key{name: 'p', value: string(pattern[i])}
		key_w := Key{name: 'w', value: string(split[i])}

		if val, ok := hashMap[key_p]; ok && val != split[i] {
			return false
		}

		if val, ok := hashMap[key_w]; ok && val != string(pattern[i]) {
			return false
		}

		hashMap[key_p] = split[i]
		hashMap[key_w] = string(pattern[i])
	}

	return true
}
