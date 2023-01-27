package handlers_test

import (
	"testing"

	"github.com/brunohubner/fc2-hexagonal-architecture/web/handlers"
	"github.com/stretchr/testify/require"
)

func TestHandles_jsonError(t *testing.T) {
	msg := "Error message"
	result := handlers.JsonError(msg)
	require.Equal(t, []byte(`{"message":"Error message"}`), result)
}
