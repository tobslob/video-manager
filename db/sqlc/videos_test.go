package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/tobslob/video-manager/utils"

	"github.com/stretchr/testify/require"
)

func createRandomVideo(t *testing.T) Video {
	user := createRandomUser(t)

	duration := utils.ParseTimeToStringRepresentation("2h")

	videoId, _ := uuid.NewRandom()
	var videoArg = CreateVideoParams{
		ID:       videoId,
		Url:      "https://youtu.be/IDp808foM9I",
		UserID:   user.ID,
		Duration: duration,
		Title:    utils.RandomString(7),
	}

	video, err := testQueries.CreateVideo(context.Background(), videoArg)
	if err != nil {
		log.Fatal("Error while creating video:", err)
	}

	require.NotEmpty(t, video)

	require.Equal(t, videoArg.Url, video.Url)
	require.Equal(t, videoArg.Duration, video.Duration)
	require.Equal(t, videoArg.Title, video.Title)
	require.Equal(t, videoArg.UserID, video.UserID)

	require.NotZero(t, video.ID)
	require.NotZero(t, video.CreatedAt)

	return video
}

func TestCreateVideot(t *testing.T) {
	createRandomVideo(t)
}

func TestGetAvideo(t *testing.T) {
	video1 := createRandomVideo(t)
	video2, err := testQueries.GetVideo(context.Background(), GetVideoParams{ID: video1.ID, UserID: video1.UserID})

	require.NoError(t, err)
	require.NotEmpty(t, video2)

	require.Equal(t, video1.ID, video2.ID)
	require.Equal(t, video1.Duration, video2.Duration)
	require.Equal(t, video1.Title, video2.Title)
	require.Equal(t, video1.Url, video2.Url)
	require.Equal(t, video1.CreatedAt, video2.CreatedAt)
	require.Equal(t, video1.UpdatedAt, video2.UpdatedAt)
}

func TestDeleteVideo(t *testing.T) {
	video1 := createRandomVideo(t)

	deleteArg := DeleteVideoParams{UserID: video1.UserID, ID: video1.ID}
	err := testQueries.DeleteVideo(context.Background(), deleteArg)
	require.NoError(t, err)

	video2, err := testQueries.GetVideo(context.Background(), GetVideoParams{ID: video1.ID, UserID: video1.UserID})
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, video2)
}
