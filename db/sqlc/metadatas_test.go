package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type NullString struct {
	String string
	Valid  bool
}

func createRandomVideoMetadata(t *testing.T) Metadata {
	video := createRandomVideo(t)

	metaId, _ := uuid.NewRandom()
	var metadataArg = CreateMetadataParams{
		ID:       metaId,
		VideoID:  video.ID,
		Width:    500,
		Height:   700,
		FileType: "mp4",
		Keywords: sql.NullString{String: "nice funny", Valid: true},
	}

	metadata, err := testQueries.CreateMetadata(context.Background(), metadataArg)
	if err != nil {
		log.Fatal("Error while creating metadata:", err)
	}

	require.NotEmpty(t, metadata)

	require.Equal(t, metadataArg.VideoID, metadata.VideoID)
	require.Equal(t, metadataArg.Width, metadata.Width)
	require.Equal(t, metadataArg.Height, metadata.Height)
	require.Equal(t, metadataArg.FileType, metadata.FileType)

	require.NotZero(t, metadata.ID)
	require.NotZero(t, metadata.CreatedAt)

	return metadata
}

func TestCreateVideoMetadata(t *testing.T) {
	createRandomVideoMetadata(t)
}

func TestGetVideoMetadata(t *testing.T) {
	metadata1 := createRandomVideoMetadata(t)
	metadata2, err := testQueries.GetMetadata(context.Background(), metadata1.VideoID)

	require.NoError(t, err)
	require.NotEmpty(t, metadata2)

	require.Equal(t, metadata1.Keywords, metadata2.Keywords)
	require.Equal(t, metadata1.Height, metadata2.Height)
	require.Equal(t, metadata1.VideoID, metadata2.VideoID)
	require.Equal(t, metadata1.Width, metadata2.Width)
	require.Equal(t, metadata1.CreatedAt, metadata2.CreatedAt)
}

func TestUpdateVideoMetadatat(t *testing.T) {
	metadata1 := createRandomVideoMetadata(t)

	updatedArg := UpdateMetadataParams{
		ID:       metadata1.ID,
		Width:    40,
		Height:   50,
		Keywords: sql.NullString{String: "new keyword", Valid: true},
	}
	metadata2, err := testQueries.UpdateMetadata(context.Background(), updatedArg)
	require.NoError(t, err)

	require.NotSame(t, updatedArg.Height, metadata2.Height)
	require.NotSame(t, metadata1.Width, metadata2.Width)
	require.NotSame(t, metadata1.Keywords, metadata2.Keywords)

	require.Equal(t, metadata1.VideoID, metadata2.VideoID)
}
