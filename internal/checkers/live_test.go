package checkers_test

import (
	"testing"

	"github.com/modem7/docker-error-pages/internal/checkers"
	"github.com/stretchr/testify/assert"
)

func TestLiveChecker_Check(t *testing.T) {
	assert.NoError(t, checkers.NewLiveChecker().Check())
}
