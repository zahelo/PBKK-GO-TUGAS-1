package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func reverseString(word string) string {
    runes := []rune(word)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Input (minimal 3 kata): ")
    input, _ := reader.ReadString('\n') 
    input = strings.TrimSpace(input)

    words := strings.Fields(input)
    if len(words) < 3 {
        fmt.Println("Inputan minimla harus terdiri dari 3 kata.")
        return
    }

    for _, word := range words {
        fmt.Println(reverseString(word))
    }
}
