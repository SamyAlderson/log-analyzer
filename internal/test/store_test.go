// File: internal/test/store_test.go
// Purpose: Store unit tests
// Project: log-analyzer
// Description: A log parser analyzer to extract insights from server logs

package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	"log-analyzer/internal/store/store"
)

func TestStore(t *testing.T) {
	t.Run("TestStoreNew", func(t *testing.T) {
		// Arrange
		s, err := store.NewStore()
		require.NoError(t, err)

		// Act
		// Assert
		assert.NotNil(t, s)
	})

	t.Run("TestStorePut", func(t *testing.T) {
		// Arrange
		s, err := store.NewStore()
		require.NoError(t, err)

		// Act
		err = s.Put("key", "value")
		require.NoError(t, err)

		// Assert
		assert.NotNil(t, s.Get("key"))
	})

	t.Run("TestStorePutNotFound", func(t *testing.T) {
		// Arrange
		s, err := store.NewStore()
		require.NoError(t, err)

		// Act
		err = s.Put("key", "value")
		require.NoError(t, err)

		// Act
		err = s.Put("key", "new_value")
		require.NoError(t, err)

		// Assert
		assert.Equal(t, "new_value", s.Get("key"))
	})

	t.Run("TestStoreGet", func(t *testing.T) {
		// Arrange
		s, err := store.NewStore()
		require.NoError(t, err)

		// Act
		err = s.Put("key", "value")
		require.NoError(t, err)

		// Assert
		assert.NotNil(t, s.Get("key"))
	})

	t.Run("TestStoreGetNotFound", func(t *testing.T) {
		// Arrange
		s, err := store.NewStore()
		require.NoError(t, err)

		// Act
		// Assert
		assert.Nil(t, s.Get("key"))
	})

	t.Run("TestStoreDelete", func(t *testing.T) {
		// Arrange
		s, err := store.NewStore()
		require.NoError(t, err)

		// Act
		err = s.Put("key", "value")
		require.NoError(t, err)

		// Act
		err = s.Delete("key")
		require.NoError(t, err)

		// Assert
		assert.Nil(t, s.Get("key"))
	})

	t.Run("TestStoreDeleteNotFound", func(t *testing.T) {
		// Arrange
		s, err := store.NewStore()
		require.NoError(t, err)

		// Act
		err = s.Delete("key")
		assert.Nil(t, s.Get("key"))
	})
}
