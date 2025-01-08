package main

import (
    "io"
    "os"
)

type Square struct {
    x int 
    y int
}

func GetVisitedHouses() int {
    return navigate(GetString("./day3/input.txt"))
}


func GetString(fileName string) string {
    file, err := os.Open(fileName)

    if err != nil {
		return ""
	}
	defer file.Close() 

    content, err := io.ReadAll(file)
	if err != nil {
		return ""
	}

    return string(content)
}

func navigate(str string) int{
    santa := Square{0, 0}
    roboSanta := Square{0, 0}
    visits := make(map[Square]int)

    for i, char := range str {
        if i % 2 == 0 {
            changePosition(&roboSanta, char)
            visitHouse(visits, santa)
        }else{
            changePosition(&santa, char)
            visitHouse(visits, roboSanta)
        }
    }
    return len(visits)
}


func visitHouse(visitedHouses map[Square]int,currentHouse Square) {
    if _, exists := visitedHouses[currentHouse]; exists {
        visitedHouses[currentHouse] += 1
    }else {
        visitedHouses[currentHouse] = 1
    }
}

func changePosition(position *Square, direction rune) {
        switch direction{
        case '>':
            position.x += 1
        case '<':
            position.x -= 1
        case '^':
            position.y += 1
        case 'v':
            position.y -= 1
        }
}
