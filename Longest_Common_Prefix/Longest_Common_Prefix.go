package main

func longestCommonPrefix(strs []string) string {
    ch := strs[0]
    i := 0
    for _,content := range(strs[1:]) {
        i = 0
        for ;i < len(ch) && i < len(content) && ch[i] == content[i]; i++ {}
        ch = ch[:i]
    }
    return ch
}