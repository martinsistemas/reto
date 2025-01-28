package main

import (
    "fmt"
)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func perteneceFibonacci(num int) bool {
    i := 0
    for fibonacci(i) <= num {
        if fibonacci(i) == num {
            return true
        }
        i++
    }
    return false
}

func main() {
    var numero int
    fmt.Print("Ingrese un número: ")
    fmt.Scan(&numero)

    if perteneceFibonacci(numero) {
        fmt.Printf("%d pertenece a la secuencia de Fibonacci.\n", numero)
    } else {
        fmt.Printf("%d no pertenece a la secuencia de Fibonacci.\n", numero)
    }
}