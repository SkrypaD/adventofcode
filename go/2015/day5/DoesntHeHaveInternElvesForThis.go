package main

import (
    //"fmt"
    "regexp"
    "os"
    "bufio"
)
func CountNiceStrings()  int64 {
    file, err := os.Open("input.txt")
    if err != nil {
        return -1  
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var result int64
    for scanner.Scan() {
        line := scanner.Text()
        if IsStringNice2(line){
            result++
        }
    }
    return result
    
}

func hasSequence(str string) bool {
    prev := ' '

    for _, char := range str{
        if(char == prev){
            return true
        }else{
            prev = char
        }
    }
    return false
}


func IsStringNice(str string) bool {
    vowelPattern := `.*[aeiou].*[aeiou].*[aeiou].*`
    corruptPattern := `(ab|cd|pq|xy)`

    vowel, _ := regexp.MatchString(vowelPattern, str)
    corrupt, _ := regexp.MatchString(corruptPattern, str)

    if vowel && !corrupt && hasSequence(str) {
        return true
    }else{
        return false
    }
}


//Second part of the 5th day
func hasSeparatedPair(str string) bool {

    for i:= 2; i < len(str); i++ {
        if str[i] == str[i - 2]{
            return true
        }
    }
    return false
}

func hasTwoPairs(str string) bool {
    pairs := make(map[string]int)

    for i := 0; i < len(str) - 1; i++ {
        value, exists := pairs[string(str[i]) + string(str[i + 1])]
        if exists {
            if i - 1 > value {
                return true
            }
        }else{
            pairs[string(str[i]) + string(str[i + 1])] = i;
        }
    }
    return false
}

func IsStringNice2(str string) bool {
    return hasSeparatedPair(str) && hasTwoPairs(str)
}
