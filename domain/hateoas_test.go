package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDefaultHalResource_AddLink(t *testing.T) {
	h := DefaultHalResource{}

	assert.Empty(t, h.Links)

	h.AddLink(REL_SELF, "selflink")

	assert.Equal(t, 1, len(h.Links))
	assert.Equal(t, "selflink", h.Links[REL_SELF].Href)
}