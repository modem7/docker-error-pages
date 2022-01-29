package metrics_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/modem7/docker-error-pages/internal/metrics"
)

func TestNewRegistry(t *testing.T) {
	registry := metrics.NewRegistry()

	count, err := testutil.GatherAndCount(registry)

	assert.NoError(t, err)
	assert.True(t, count >= 6, "not enough common metrics")
}
