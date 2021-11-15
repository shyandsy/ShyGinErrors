package ShyGinErrors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestShyGinErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShyGinErrors Suite")
}
