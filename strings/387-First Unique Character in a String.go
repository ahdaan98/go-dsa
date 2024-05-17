package strings

func firstUniqChar(s string) int {
    count := make([]int,26)

    for _,c := range s{
        count[c-'a']++
    }

    for i,char := range s{
        if count[char-'a'] == 1 {
            return i
        }
    }

    return -1
}