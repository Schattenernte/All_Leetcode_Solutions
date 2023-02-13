package main

func reverseString(s []byte)  {
    st := len(s) -1
    en := 0
    for st > en {
        s[st] ^= s[en]
        s[en] ^= s[st]
        s[st] ^= s[en]
    st -= 1
    en += 1
    }
}