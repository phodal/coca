package main_test

import (
	. "./imp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	Context("Basic", func() {
		It("should return correct result", func() {
			Expect(Add(1, 2)).Should(Equal(uint64(3)))
		})
	})
})