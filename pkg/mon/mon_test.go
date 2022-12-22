package mon_test

import (
"time"
"fmt"

  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"

  "github.com/SALLYPEMDAS/chronomon/pkg/mon"
)

var _ = Describe("Books", func() {
It("test", func() {
        Expect(mon.Color(mon.Saturn)).To(Equal("black"))
        Expect(mon.Color(mon.Jupeter)).To(Equal("blue"))
        Expect(mon.Color(mon.Mars)).To(Equal("red"))
        Expect(mon.Color(mon.Mars)).To(Equal("red"))
        Expect(mon.Color(mon.Sun)).To(Equal("yellow"))
        Expect(mon.Color(mon.Venus)).To(Equal("green"))
        Expect(mon.Color(mon.Mercury)).To(Equal("rainbow"))
        Expect(mon.Color(mon.Moon)).To(Equal("silver"))

//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time

nyc, _ := time.LoadLocation("America/New_York")


// sunday
d := time.Date(1970, 2, 1, 6, 1,1,1, nyc)

fmt.Printf("%v", d)

Expect(mon.ToPlanetDay(d)).To(Equal(mon.Saturn))


d = d.Add(time.Hour)

Expect(mon.ToPlanetDay(d)).To(Equal(mon.Sun))

})

})
