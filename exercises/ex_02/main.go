package main

import (
	"fmt"
	"os"
	"strconv"
)

type Miles float64
type Kilometers float64

const ConvertionValue float64 = 1.60934

func main() {
	for _, arg := range os.Args[1:] {
		distance, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		km := Kilometers(distance)
		mi := Miles(distance)

		fmt.Printf("%s = %s, %s = %s\n",
			km, KMtoMi(km), mi, MitoKM(mi))
	}
}

func (mi Miles) String() string      { return fmt.Sprintf("%g mi", mi) }
func (km Kilometers) String() string { return fmt.Sprintf("%g km", km) }

func KMtoMi(km Kilometers) Miles { return Miles(km / Kilometers(ConvertionValue)) }
func MitoKM(mi Miles) Kilometers { return Kilometers(mi * Miles(ConvertionValue)) }
