package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tobslob/video-manager/utils"
)

func TestCreatVideoWithMetadataTx(t *testing.T) {
	store := NewStore(testDb)

	user := createRandomUser(t)

	duration := utils.ParseTimeToStringRepresentation("2h30m")
	result, err := store.CreatVideoWithMetadataTx(context.Background(), CreateVideoWithMetadata{
		CreateVideoParams: CreateVideoParams{
			Url:      "https://youtube.com/video",
			UserID:   user.ID,
			Duration: duration,
			Title:    utils.RandomString(9),
		}, CreateMetadataParams: CreateMetadataParams{
			Width:        100,
			Height:       300,
			FileType:     "mp4",
			FileSize:     sql.NullString{"3000", true},
			AccessedDate: time.Now(),
			Resolutions:  1000,
			LastModify:   time.Now(),
			Keywords:     "key words",
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

}
