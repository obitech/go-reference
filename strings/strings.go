package main

import (
    "fmt"
    "unicode/utf8"
)

func printBytes(s string) {
    for i := 0; i < len(s); i++ {
        // %x = Base16, Hex formatter
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {
    /*
       Below will struggle with other unicode charactes, because some are represented
       with 2 - 4 Bytes...
       for i := range s {
           fmt.Printf("%c ", s[i])
       }
    */

    // ... runes to the rescue!
    // Runes = int32, represents unicode codepoints
    runes := []rune(s)
    for i, rune := range runes {
        fmt.Printf("Index: %d, Rune: %c\n", i, rune)
    }
}

// Changes a string
func mutate(s []rune) string {
    s[0] = 'a'
    return string(s)
}

func main() {
    // Strings are slices of bytes, UTF-8 encoded
    name := "Hello Señor"
    fmt.Println(name)
    printBytes(name)
    fmt.Println()
    printChars(name)

    // We can print strings with byte slices
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)

    // Same with decimal representations
    byteSlice2 := []byte{67, 97, 102, 195, 169}
    str = string(byteSlice2)
    fmt.Println(str)

    // Constructing string from rune slice
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str = string(runeSlice)
    fmt.Println(str)

    // This is getting the number of bytes!
    fmt.Println(len(str))

    // utf8 can get string length
    fmt.Printf("Length of %s: %d\n", str, utf8.RuneCountInString(str))

    // Strings are immutable
    // str[0] = 'a' won't work!

    // To mutate strings we have to cast them as rune slices
    str = mutate([]rune(str))
    fmt.Println(str)
}
