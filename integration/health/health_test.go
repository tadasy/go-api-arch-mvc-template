package integration

import (
	"go-api-arch-mvc-template/pkg"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	endpoint := pkg.GetEndpoint("/health")
	res, err := http.Get(endpoint)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
