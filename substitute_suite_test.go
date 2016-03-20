package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSubstitute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Substitute Suite")
}
