package main

import (
    "fmt"
)

func main() {

    var (
        di, df int
        hi, mi, si int
        hf, mf, sf int
    )

    fmt.Scanf("Dia %d", &di)
    fmt.Scanf("%d : %d : %d", &hi, &mi, &si)
    fmt.Scanf("Dia %d", &df)
    fmt.Scanf("%d : %d : %d", &hf, &mf, &sf)

    var minutoSeg int = 60
    var horaSeg int = minutoSeg * 60
    var diaSeg int = horaSeg * 24

    i := si + mi*minutoSeg + hi*horaSeg + di*diaSeg
    f := sf + mf*minutoSeg + hf*horaSeg + df*diaSeg

    // Calcula a diferença de tempo e converte para dias, horas, minutos e segundos
    if i < f {
        tempo := f - i
        dias := tempo / diaSeg
        tempo %= diaSeg
        horas := tempo / horaSeg
        tempo %= horaSeg
        minutos := tempo / minutoSeg
        segundos := tempo % minutoSeg

        fmt.Printf("%d dia(s)\n%d hora(s)\n%d minuto(s)\n%d segundo(s)\n", dias, horas, minutos, segundos)
    }
}
