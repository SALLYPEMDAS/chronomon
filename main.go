package main

import (
"github.com/SALLYPEMDAS/chronomon/pkg/mon"
"fmt"
"time"
)

func main () {
nyc, _ := time.LoadLocation("America/New_York")
now := time.Now().In(nyc)

fmt.Printf("THE PLANET OF THE DAY IS %v", mon.IntToPlanetString(mon.ToPlanetDay(now)))
fmt.Printf("THE COLOR OF THE DAY IS %v", mon.Color(mon.ToPlanetDay(now)))

}
