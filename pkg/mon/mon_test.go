package mon_test

import (
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



})

})
