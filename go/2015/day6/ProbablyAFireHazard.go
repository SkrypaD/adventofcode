package day6

import (
    "strings"
    "strconv"
    "fmt"
    "bufio"
    "os"
) 

func Test() int {
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


func Count() int {
    file, err := os.Open("./day6/input.txt");
    if ( err != nil) {
        fmt.Println("Cant open file")
        return 0
    }
    defer file.Close();

    var lights [1000][1000]bool;

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), "turn on") {
            TurnOn(&lights, ConvertToInstruction(scanner.Text()))
        }else if strings.Contains(scanner.Text(), "turn off") {
            TurnOff(&lights, ConvertToInstruction(scanner.Text()))
        }else if strings.Contains(scanner.Text(), "toggle") {
            Toggle(&lights, ConvertToInstruction(scanner.Text()))
        }
    }

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

func ConvertToInstruction(str string) Instruction {
    var temp []string = strings.Split(str, " ");
    var instruction Instruction;
    var start []string 
    var finish []string
    if strings.Contains(str, "toggle"){
        start = strings.Split(temp[1], ",");
        finish = strings.Split(temp[3], ",");
    }else{
        start = strings.Split(temp[2], ",");
        finish = strings.Split(temp[4], ",");
    }

    instruction.startX, _ = strconv.Atoi(start[0])
    instruction.startY, _ = strconv.Atoi(start[1])
    instruction.finishX, _ = strconv.Atoi(finish[0])
    instruction.finishY, _ = strconv.Atoi(finish[1])
    return instruction
}

type Instruction struct{
    startX int
    startY int
    finishX int
    finishY int
}

func Toggle(lights *[1000][1000]bool, instruction Instruction) {
    for i := instruction.startX; i <= instruction.finishX; i++ {
        for j := instruction.startY; j <= instruction.finishY; j++ {
            lights[i][j] = !lights[i][j];
        }
    }
}

func TurnOn(lights *[1000][1000]bool, instruction Instruction) {
    for i := instruction.startX; i <= instruction.finishX; i++ {
        for j := instruction.startY; j <= instruction.finishY; j++ {
            lights[i][j] = true;
        }
    }
}

func TurnOff(lights *[1000][1000]bool, instruction Instruction) {
    for i := instruction.startX; i <= instruction.finishX; i++ {
        for j := instruction.startY; j <= instruction.finishY; j++ {
            lights[i][j] = false;
        }
    }
}
