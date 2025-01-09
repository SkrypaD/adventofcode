package day6

import (
    "strings"
    "fmt"
    "bufio"
    "os"
) 

func Test2() int {
    var lights [1000][1000]bool;
    str := "turn on 0,0 through 999,999"
    TurnOn(&lights, ConvertToInstruction(str))
    couter := 0;
    for i := 0; i < len(lights); i++ {
        for j := 0; j < len(lights); j++ {
            if lights[i][j] {
                couter++
            }
        }
    }
    return couter;
}


func Count2() int {
    file, err := os.Open("./day6/input.txt");
    if ( err != nil) {
        fmt.Println("Cant open file")
        return 0
    }
    defer file.Close();

    var lights [1000][1000]int;

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), "turn on") {
            turnOn(&lights, ConvertToInstruction(scanner.Text()))
        }else if strings.Contains(scanner.Text(), "turn off") {
            turnOff(&lights, ConvertToInstruction(scanner.Text()))
        }else if strings.Contains(scanner.Text(), "toggle") {
            toggle(&lights, ConvertToInstruction(scanner.Text()))
        }
    }

    couter := 0;
    for i := 0; i < len(lights); i++ {
        for j := 0; j < len(lights); j++ {
            couter += lights[i][j];
        }
    }
    return couter;
}


func toggle(lights *[1000][1000]int, instruction Instruction) {
    for i := instruction.startX; i <= instruction.finishX; i++ {
        for j := instruction.startY; j <= instruction.finishY; j++ {
            lights[i][j] += 2;
        }
    }
}

func turnOn(lights *[1000][1000]int, instruction Instruction) {
    for i := instruction.startX; i <= instruction.finishX; i++ {
        for j := instruction.startY; j <= instruction.finishY; j++ {
            lights[i][j] += 1;
        }
    }
}

func turnOff(lights *[1000][1000]int, instruction Instruction) {
    for i := instruction.startX; i <= instruction.finishX; i++ {
        for j := instruction.startY; j <= instruction.finishY; j++ {
            if lights[i][j] > 0 {
                lights[i][j] -= 1;
            }
        }
    }
}
