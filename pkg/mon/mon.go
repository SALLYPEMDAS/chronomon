package mon

import (
"fmt"
. "time"
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

func IntToPlanetString(i Planet) string {
switch (i) {
case 0: 
return "Saturn"
case 1:
return "Jupeter"
case 2:
return "Mars"
case 3:
return "Sun"
case 4:
return "Venus"
case 5:
return "Mercury"
case 6:
return "Moon"
default: 
return "-1"
}
}

func Mon () {
fmt.Println("hello world")
}

func ToPlanetDay(t time.Time) Planet {

return DayToPlanet(GetDay(t))

}

func GetDay(t time.Time) time.Weekday {
fmt.Println(time.Hour)
if t.Hour() >= 7 {
fmt.Println("after 7")
return t.Weekday()
}
fmt.Println("before 7")
return t.Add(-1*time.Hour * 24).Weekday()

}

func DayToPlanet(d time.Weekday) Planet {
	switch(d) {
	case Sunday:
		return Sun
	case Monday:
		return Moon
	case Tuesday:
		return Mars
	case Wednesday:
		return Mercury
	case Thursday:
		return Jupeter
	case Friday:
		return Venus
	case Saturday:
		return Saturn

default: return Saturn
}
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
