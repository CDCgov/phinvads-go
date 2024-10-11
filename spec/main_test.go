package spec_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Word struct {
	String string
	Length int
}

var _ = Describe("Spec", func() {
	var helloWorld *Word

	BeforeEach(func() {
		helloWorld = &Word{
			String: "Hello World!",
			Length: 12,
		}
	})

	Describe("Categorizing strings", func() {
		Context("inside the spec", func() {
			It("should have the correct length", func() {
				Expect(helloWorld.Length).To(Equal(12))
			})
		})
	})
})
