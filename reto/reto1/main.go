package main

import "fmt"

func main() {
    INDICE := 13
    SOMA := 0
    K := 0

    for K < INDICE {
        K = K + 1
        SOMA = SOMA + K
    }

    fmt.Println(SOMA)
}