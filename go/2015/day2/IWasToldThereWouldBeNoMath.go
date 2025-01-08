package main

import (
	"bufio"
    "os"
	"strconv"
	"strings"
    "fmt"
    "sort"
)

func TotalOrder() int64 {
    file, err := os.Open("./day2/input.txt")
    if err != nil {
        return -1 
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var result int64
    for scanner.Scan() {
        line := scanner.Text()
        result += getRibbon(convertToArr(line))
    }
    return result
    
}

func getRibbon(arr []int64) int64 {

    slice := []int{}

    for _, val := range arr {
        slice = append(slice, int(val))
    }
    
    sort.Ints(slice)
    fmt.Println(slice)
    
    result := int64(slice[0] * 2 + slice[1] * 2 + slice[0] * slice[1] * slice[2])
    fmt.Println(result)

    return  result
}

func convertToArr(str string) []int64 {
    strings := strings.Split(str, "x")
    var arr []int64
    for _, char := range strings{
        val, _ := strconv.ParseInt(char, 10, 64)
        arr = append(arr, val)
    }
    fmt.Println(arr)
    return arr
}

func getBoxSize(arr []int64) int64  {
    one := 2 * arr[0] * arr[1] 
    two := 2 * arr[1] * arr[2]
    three := 2 * arr[0] * arr[2]
    
    arr[0] = one
    arr[1] = two
    arr[2] = three
    smallest := one 
    for _, v := range arr {
        if v  < smallest {
            smallest = v
        }
    }

    return one + two + three + smallest / 2
}
