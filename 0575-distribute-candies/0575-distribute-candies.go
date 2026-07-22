func distributeCandies(candyType []int) int {
    hashMap := map[int]bool{}
    for _, n := range candyType {
        hashMap[n] = true
    }

    return min(len(candyType)/2, len(hashMap))
}