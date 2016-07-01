package dataProvider_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDataProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DataProvider Suite")
}
