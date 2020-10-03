package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Footballer struct {
    Name        string   `json:"name"`
    Pace        int64    `json:"pace"`
    Dribbling   int64    `json:"dribbling"`
}

func main() {

    player1 := Footballer{
    Name:  "CR7",
    Pace:  90,
    Dribbling: 89,
    }

    player2 := Footballer{
    Name:  "LM10",
    Pace:  85,
    Dribbling: 93,
    }

    players := []Footballer{player1, player2}

    jsonData, err := json.MarshalIndent(players,"", " ")
    if err != nil {
        log.Println(err)
    }
    fmt.Println(string(jsonData))

    var data []Footballer
    err = json.Unmarshal([]byte(jsonData), &data)
    fmt.Printf("%+v\n", data)
}