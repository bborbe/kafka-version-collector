package main_test

import (
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kafka Version Collector", func() {
	It("Compiles", func() {
		var err error
		_, err = gexec.Build("github.com/bborbe/kafka-k8s-version-collector")
		Expect(err).NotTo(HaveOccurred())
	})
})

func TestKafkaK8sVersionCollector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KafkaK8sVersionCollector Suite")
}