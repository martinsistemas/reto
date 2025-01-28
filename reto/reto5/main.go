package main

import (
    "fmt"
)

func invertirString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    var texto string
    fmt.Print("Ingrese una cadena de texto: ")
    fmt.Scanln(&texto)

    textoInvertido := invertirString(texto)
    fmt.Printf("Texto invertido: %s\n", textoInvertido)
}