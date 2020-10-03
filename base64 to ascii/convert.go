package main

import (
    "encoding/base64"
    "fmt"
)

func main() {

    data := "VGhpcyBpcyBkYXRhIGluIGFzY2lpIGZvcm1hdA=="

   // Base64 Standard Decoding
    decoded_data, err := base64.StdEncoding.DecodeString(data)
    if err != nil {
        fmt.Printf("Error decoding string: %s ", err.Error())
        return
    }

    fmt.Println(string(decoded_data))
}