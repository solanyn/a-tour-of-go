package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    // Init picture array
    m := make([][]uint8, dy)
    for i := range m {
        m[i] = make([]uint8, dx)
    }

    // Fill picture
    for y := range m {
        for x := range m[y] {
            m[y][x] = uint8(x*y)
        }
    }
    return m
}

func main() {
    pic.Show(Pic)
}
