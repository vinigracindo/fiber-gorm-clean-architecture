package id_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/pkg/id"
)

func TestPkgID_NewID(t *testing.T) {
	id := id.NewID()
	require.IsType(t, uuid.UUID{}, id)
}
