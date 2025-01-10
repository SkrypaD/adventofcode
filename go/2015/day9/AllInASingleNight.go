// I HATE THIS SOLUTION SO MUCH, IT LOOKS SOO BAD I CAN BRING MYSELF TO CLEAN THIS SHIT UP
// O(n!) solution via recursion


package day9

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type Connection struct {
    firstCity string
    secondCity string
    distance int
}

const SIZE = 8

func GetCitiesConnections() {
    fileName := "./day9/input.txt"
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    var cities []string
    var roadMap []Connection
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        data := strings.Split(scanner.Text(), " ")
        if !Contains(cities, data[0]) {
            cities = append(cities, data[0])
        }
        if !Contains(cities, data[2]) {
            cities = append(cities, data[2])
        }
        val, _ := strconv.Atoi(data[4])
        roadMap = append(roadMap, Connection{data[0], data[2], val})
    }

    var visitedCities [SIZE]string
    maxDistance := 0 


    for _, city := range cities {
        FindWay(roadMap, city, 0, visitedCities, "", &maxDistance)
    }
    fmt.Println(maxDistance)
}

func FindWay(roadMap []Connection, city string, distance int, visitedCities [SIZE]string, lastCity string, maxDistance *int) int {
    if lastCity == "" {
        addValue(&visitedCities, city)
    }else{
        addValue(&visitedCities, lastCity)
    }
    if checkArray(visitedCities) {
        if distance > *maxDistance{
            *maxDistance = distance
        }
        return distance
    }

    for _, connection := range roadMap{
        if connection.firstCity == city && !contains(visitedCities, connection.secondCity){
            FindWay(roadMap, connection.secondCity, distance + connection.distance, visitedCities, connection.secondCity, maxDistance)
        }else if connection.secondCity == city && !contains(visitedCities, connection.firstCity){
            FindWay(roadMap, connection.firstCity, distance + connection.distance, visitedCities, connection.firstCity, maxDistance)
        }
    }
    return *maxDistance
}

func contains(arr [SIZE]string, city string) bool {
    for _, town := range arr {
        if city == town {
            return true
        }
    }
    return false
}

func addValue(arr *[SIZE]string, city string) {
    for i, value := range arr {
        if value == "" {
            arr[i] = city
            return
        }
    }
}

func checkArray(arr [SIZE]string) bool {
    for _, city := range arr {
        if city == "" {
            return false
        }
    }
    return true
}

func Contains(slice []string, city string) bool {
    for _, town := range slice {
        if town == city {
            return true
        }
    }
    return false
}

