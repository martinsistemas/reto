package main

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "math"
    "path/filepath"
    "strings"
)

type Faturamento struct {
    Dia   int     `json:"dia" xml:"dia"`
    Valor float64 `json:"valor" xml:"valor"`
}

func main() {
    // Ruta del archivo
    filePath := "./dados.json" // o "dados.xml"

    // Leer el archivo
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error al leer el archivo:", err)
        return
    }

    var faturamento []Faturamento

    // Determinar el formato basado en la extensión del archivo
    fileExt := strings.ToLower(filepath.Ext(filePath))
    if fileExt == ".json" {
        err = json.Unmarshal(data, &faturamento)
    } else if fileExt == ".xml" {
        err = xml.Unmarshal(data, &faturamento)
    } else {
        fmt.Println("Formato de archivo no soportado")
        return
    }

    if err != nil {
        fmt.Println("Error al decodificar el archivo:", err)
        return
    }

    menorValor := math.Inf(1)
    maiorValor := math.Inf(-1)
    soma := 0.0
    diasComFaturamento := 0

    for _, dia := range faturamento {
        if dia.Valor > 0 {
            if dia.Valor < menorValor {
                menorValor = dia.Valor
            }
            if dia.Valor > maiorValor {
                maiorValor = dia.Valor
            }
            soma += dia.Valor
            diasComFaturamento++
        }
    }

    mediaMensal := soma / float64(diasComFaturamento)
    diasAcimaMedia := 0

    for _, dia := range faturamento {
        if dia.Valor > mediaMensal {
            diasAcimaMedia++
        }
    }

    fmt.Printf("Menor valor de facturación: %.2f\n", menorValor)
    fmt.Printf("Mayor valor de facturación: %.2f\n", maiorValor)
    fmt.Printf("Número de días por encima de la media mensual: %d\n", diasAcimaMedia)
}
