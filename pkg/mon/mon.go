package mon

import (
"fmt"
"time"
)

type Planet int64


const (
	Saturn Planet = iota
	Jupeter
	Mars
	Sun
	Venus
	Mercury
	Moon
)

func Mon () {
fmt.Println("hello world")
}

func ToPlanet(t time.Time) Planet {

return Saturn

}

func Color(planet Planet) string {

switch(planet){
case Saturn:
return "black"
case Jupeter:
return "blue"
case Mars:
return "red"
case Sun:
return "yellow"
case Venus:
return "green"
case Mercury:
return "rainbow"
case Moon:
return "silver"




default:
return "-1";

}

}
