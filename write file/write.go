package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "log"
)

func main() {
	file, err := os.Create("file.txt") 
      
    if err != nil { 
        log.Fatalf("Failed creating file: %s", err) 
    } 

    defer file.Close() 

    len, err := file.WriteString("Writing Content to the file") 
  
    if err != nil { 
        log.Fatalf("Failed writing to file: %s", err) 
    }

    fileName := "file.txt"

    data, err := ioutil.ReadFile(fileName) 
    if err != nil { 
        log.Panicf("failed reading data from file: %s", err) 
    } 
    fmt.Printf("\nFile Name: %s", fileName) 
    fmt.Printf("\nSize: %d bytes", len)
    fmt.Printf("\nData: %s", data) 
}
