package main

import "fmt"

func main() {
    faturamento := map[string]float64{
        "SP":     67836.43,
        "RJ":     36678.66,
        "MG":     29229.88,
        "ES":     27165.48,
        "Outros": 19849.53,
    }

    total := 0.0
    for _, valor := range faturamento {
        total += valor
    }

    fmt.Println("Porcentaje de representación de cada estado:")
    for estado, valor := range faturamento {
        percentual := (valor / total) * 100
        fmt.Printf("%s: %.2f%%\n", estado, percentual)
    }
}