package main

import "github.com/pmcxs/hexgrid"

func main() {

	hexa := hexgrid.NewHex(1, 2)
	hexb := hexgrid.NewHex(2, 3)

	println(hexgrid.HexDistance(hexa, hexb))

}
