package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCoca(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coca Suite")
}
