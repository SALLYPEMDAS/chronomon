package main

import (
	"fmt"
	"time"

	"github.com/SALLYPEMDAS/chronomon/pkg/mon"
)

func main() {
	nyc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(nyc)
	pd := mon.ToPlanetDay(now)
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("THE PLANET OF THE DAY IS %v", mon.IntToPlanetString(pd))
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("THE COLOR OF THE DAY IS %v", mon.Color(pd))
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("ADVICE FOR TODAY: %v", mon.DescriptionPlanet(pd))
	fmt.Println("")
	fmt.Println("")

}
