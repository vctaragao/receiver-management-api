package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test suite")
	}
	suite.Run(t, new(IntegrationSuite))
}
