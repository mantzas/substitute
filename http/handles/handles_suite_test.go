package handles_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHandles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handles Suite")
}
