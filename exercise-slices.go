package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	 pic := make([][]uint8, dy) /* type declaration */
    for i := range pic {
        pic[i] = make([]uint8, dx) /* again the type? */
        for j := range pic[i] {
            pic[i][j] = uint8((i^j))
        }
    }
    return pic
}

func main() {
	pic.Show(Pic)
}
