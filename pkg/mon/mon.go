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
return "black";

default:
return "-1";

}

}
