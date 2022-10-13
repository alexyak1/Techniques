package migrations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllTechniques(t *testing.T) {
	techniques, err := getAllTechniques()
	require.NoError(t, err)
	require.Equal(t, 81, len(techniques))
}

func TestGetImageIds(t *testing.T) {
	techniques, err := getAllTechniques()
	require.NoError(t, err)
	imageIds := getImageIds(techniques)

	require.Equal(t, 81, len(imageIds))
}

func TestInsertImageIdToTechniques(t *testing.T) {
	techniques, err := getAllTechniques()
	require.NoError(t, err)

	imageIds := getImageIds(techniques)

	err = insertImageIdToTechniques(imageIds)

	require.NoError(t, err)

}
