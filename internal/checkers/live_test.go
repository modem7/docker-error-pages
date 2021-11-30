package checkers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/modem7/docker-error-pages/internal/checkers"
)

func TestLiveChecker_Check(t *testing.T) {
	assert.NoError(t, checkers.NewLiveChecker().Check())
}