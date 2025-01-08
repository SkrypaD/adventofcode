package day7

import (
    "os"
    "bufio"
    "strings"
    "strconv"
    "fmt"
)

const (
    ASSIGN = iota
    NOT
    OR
    AND
    LSHIFT
    RSHIFT
)

type InputSignal struct {
    assigned bool
    value uint16
    inputType int
    inputWire []string
}

func Entry() {
    gates := make(map[string]InputSignal)
    for _, str := range GetStrings("./day7/input.txt") {
        ConvertStringToGate(gates, str)
    }

    gates["b"] = InputSignal{true, 16076, gates["b"].inputType, gates["b"].inputWire}
    fmt.Println(FindWireSignal(gates, "a"))

}

func GetStrings(fileName string) []string {
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    lines := make([]string, 0, 340)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func ConvertStringToGate(gates map[string]InputSignal, str string) {
    arr := strings.Split(str, " ")

    newSignal := InputSignal{}
    if len(arr) == 3 {
        newSignal.inputType = ASSIGN
        value, err := strconv.Atoi(arr[0]) 
        val := uint16(value)
        if err == nil {
            newSignal.assigned = true
            newSignal.value = val
        }else{
            newSignal.inputWire = append(newSignal.inputWire, arr[0])
        }
    } else if len(arr) == 4 {
        newSignal.inputType = NOT
        value, err := strconv.Atoi(arr[1]) 
        val := uint16(value)
        if err == nil {
            newSignal.assigned = true
            newSignal.value = val
        }else{
            newSignal.inputWire = append(newSignal.inputWire, arr[1])
        }
    } else {
        if strings.Contains(str, "AND"){ newSignal.inputType = AND
        }else if strings.Contains(str, "OR") { newSignal.inputType = OR
        }else if strings.Contains(str, "LSHIFT") { newSignal.inputType = LSHIFT 
        }else if strings.Contains(str, "RSHIFT") { newSignal.inputType = RSHIFT }
        
        newSignal.inputWire = append(newSignal.inputWire, arr[0], arr[2])
    }
    gates[arr[len(arr)-1]] = newSignal
}

func FindWireSignal(gates map[string]InputSignal, key string) uint16 {
    if gates[key].assigned {
        return gates[key].value
    }else {
        switch gates[key].inputType {
            case ASSIGN:
                value, err := strconv.Atoi(gates[key].inputWire[0]) 
                val := uint16(value)
                if err == nil {
                    gates[key] = InputSignal{true, val, ASSIGN, gates[key].inputWire}
                }else{
                    value := FindWireSignal(gates, gates[key].inputWire[0])
                    gates[key] = InputSignal{true, value, ASSIGN, gates[key].inputWire}
                }
                break
            case NOT:
                value, err := strconv.Atoi(gates[key].inputWire[0]) 
                val := uint16(value)
                if err == nil {
                    gates[key] = InputSignal{true, ^val, NOT, gates[key].inputWire}
                }else {
                    value := uint16(FindWireSignal(gates, gates[key].inputWire[0]))
                    gates[key] = InputSignal{true, ^value, NOT, gates[key].inputWire}
                }
                break
            default:
                val, err := strconv.Atoi(gates[key].inputWire[0]) 
                value := uint16(val)
                val1, err1 := strconv.Atoi(gates[key].inputWire[1]) 
                value1 := uint16(val1)
                if err != nil {
                    value = FindWireSignal(gates, gates[key].inputWire[0])
                }
                if err1 != nil {
                    value1 = FindWireSignal(gates, gates[key].inputWire[1])
                }

                var finalValue uint16
                switch gates[key].inputType {
                case AND:
                    finalValue = AndGate(value, value1)
                    break
                case OR:
                    finalValue = OrGate(value, value1)
                    break
                case LSHIFT:
                    finalValue = LSGate(value, int(value1))
                    break
                case RSHIFT:
                    finalValue = RSGate(value, int(value1))
                    break
                }
                gates[key] = InputSignal{true, finalValue, AND, gates[key].inputWire}
                break
        }
        return gates[key].value
    }
}

func convertToInt16(str string) int16 {
    val, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
        return int16(0)
    }
    return int16(val)
}

func OrGate(first uint16, second uint16) uint16 {
    return first | second
}

func AndGate(first uint16, second uint16) uint16 {
    return first & second
}

func LSGate(input uint16, shift int) uint16 {
    return input << shift
}

func RSGate(input uint16, shift int) uint16 {
    return input >> shift
}

func Not(input uint16) uint16 {
    return ^input
}
