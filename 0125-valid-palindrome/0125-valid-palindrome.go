package main

import "unicode"

func isPalindrome(s string) bool {
    
    i, j := 0, len(s)-1

    for i<j {
        for i < j && !isAlnum(s[i]) {i++}
        for i < j && !isAlnum(s[j]) {j--}

        if toLower(s[i]) != toLower(s[j]) {
            return false
        }
        i++
        j--
    }
    return true
}

func isAlnum(b byte) bool {
        r := rune(b)

        return unicode.IsLetter(r) || unicode.IsDigit(r)
    } 

    func toLower(b byte)byte{
        if 'A' <= b && b <= 'Z' {

            return b + 32
        }
        return b
    }