package main

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x < 9 {
        return true
    }
    x_copy := x
    tmp := 0
    for x > 0 {
        tmp = (tmp * 10) + (x % 10)
        x = x / 10
    }
    return x_copy == tmp
}