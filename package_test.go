package osrelease

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOsRelease(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "osrelease package")
}
