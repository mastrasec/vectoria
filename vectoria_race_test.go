//go:build race

package vectoria

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRaceAdd(t *testing.T) {
	var spaceDim uint32 = 3

	db, err := New("", WithIndexLSH(&LSHConfig{
		SpaceDim: spaceDim,
	}))
	assert.NoError(t, err)

	for i := 0; i <= 1000; i++ {
		go func() {
			itemVec := make([]float64, spaceDim)
			gofakeit.Slice(&itemVec)
			itemID := uuid.NewString()

			err := db.Add(itemID, itemVec)
			assert.NoError(t, err)
		}()
	}
}
