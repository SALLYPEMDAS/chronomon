package mon

import "fmt"

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
