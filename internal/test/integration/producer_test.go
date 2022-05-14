//go:build integration
// +build integration

package integration

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/barcostreams/go-client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	partitionKeyT0Range = "123"
	partitionKeyT1Range = "567"
	partitionKeyT2Range = "234"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration test suite")
}

var _ = Describe("Producer", func ()  {
	It("should work", func ()  {
		host := env("TEST_DISCOVERY_HOST", "barco")
		producer, err := barco.NewProducer(fmt.Sprintf("barco://%s:8083", host))
		Expect(err).NotTo(HaveOccurred())
		err = producer.Send("abc", strings.NewReader(`{"hello": 1}`), partitionKeyT0Range)
		Expect(err).NotTo(HaveOccurred())

		expectedLength, _ := strconv.Atoi(env("TEST_EXPECTED_BROKERS", "1"))
		Expect(producer.BrokersLength()).To(Equal(expectedLength))
	})
})

func env(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
