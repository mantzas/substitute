package mux_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMux(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mux Suite")
}
