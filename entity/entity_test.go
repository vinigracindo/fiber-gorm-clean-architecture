package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/entity"
)

func TestEntityEntity_NewID(t *testing.T) {
	id := entity.NewID()
	require.IsType(t, uuid.UUID{}, id)
}
