package day8

import (
    "fmt"
    "os"
    "bufio"
    //"strconv"
    "strings"
)



func GetStringLiteralsNumber() {
    lines, _ := GetStringsFromFile("./day8/input.txt")

    result := 0
    for _, str := range lines {
        rawString := strings.ReplaceAll(str, `\`, `\\`)
        rawString = strings.ReplaceAll(rawString, `"`, `\"`)
        result += len(rawString) + 2 - len(str)
    }
    fmt.Println(result)
}

func GetStringsFromFile(filepath string) ([]string, error){
    file, err := os.Open(filepath)
    if err != nil {
        return []string{}, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, nil
}

