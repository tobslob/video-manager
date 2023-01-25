package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tobslob/video-manager/utils"
)

func createRandomAnnotation(t *testing.T) Annotation {
	video := createRandomVideo(t)

	start := utils.ParseTimeToStringRepresentation("1h3m")
	end := utils.ParseTimeToStringRepresentation("3h3m")

	annotaId, _ := uuid.NewRandom()
	var annotationArg = CreateAnnotationParams{
		ID:        annotaId,
		VideoID:   video.ID,
		Type:      utils.RandomString(7),
		Note:      utils.RandomString(20),
		Title:     utils.RandomString(8),
		Label:     utils.RandomString(5),
		Pause:     false,
		StartTime: start,
		EndTime:   end,
	}

	annotation, err := testQueries.CreateAnnotation(context.Background(), annotationArg)
	if err != nil {
		log.Fatal("Error while creating annotation:", err)
	}

	require.NotEmpty(t, annotation)

	require.Equal(t, annotationArg.VideoID, annotation.VideoID)
	require.Equal(t, annotationArg.Label, annotation.Label)
	require.Equal(t, annotationArg.Note, annotation.Note)
	require.Equal(t, annotationArg.Pause, annotation.Pause)
	require.Equal(t, annotationArg.StartTime, annotation.StartTime)
	require.Equal(t, annotationArg.EndTime, annotation.EndTime)

	require.NotZero(t, annotation.ID)
	require.NotZero(t, annotation.CreatedAt)

	return annotation
}

func TestCreateAnnotation(t *testing.T) {
	createRandomAnnotation(t)
}

func TestGetAnnotion(t *testing.T) {
	annotation1 := createRandomAnnotation(t)
	annotation2, err := testQueries.GetAnnotation(context.Background(), GetAnnotationParams{
		ID:     annotation1.ID,
		UserID: annotation1.UserID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, annotation2)

	require.Equal(t, annotation1.ID, annotation2.ID)
	require.Equal(t, annotation1.Label, annotation2.Label)
	require.Equal(t, annotation1.VideoID, annotation2.VideoID)
	require.Equal(t, annotation1.Title, annotation2.Title)
	require.Equal(t, annotation1.CreatedAt, annotation2.CreatedAt)
}

func TestUpdateAnnotation(t *testing.T) {
	annotation1 := createRandomAnnotation(t)

	updatedArg := UpdateAnnotationParams{
		ID:      annotation1.ID,
		UserID:  annotation1.UserID,
		VideoID: annotation1.VideoID,
		Type:    utils.RandomString(7),
		Note:    utils.RandomString(20),
		Title:   utils.RandomString(8),
		Label:   utils.RandomString(5),
		Pause:   false,
	}
	annotation2, err := testQueries.UpdateAnnotation(context.Background(), updatedArg)
	require.NoError(t, err)

	require.Equal(t, updatedArg.Label, annotation2.Label)
	require.Equal(t, updatedArg.Note, annotation2.Note)
	require.Equal(t, updatedArg.Pause, annotation2.Pause)
	require.Equal(t, updatedArg.StartTime, annotation2.StartTime)
	require.Equal(t, updatedArg.EndTime, annotation2.EndTime)

	require.NotZero(t, annotation2.ID)
	require.NotZero(t, annotation2.UpdatedAt)

	require.Equal(t, annotation1.VideoID, annotation1.VideoID)
}

func TestDeleteAnnotations(t *testing.T) {
	annotation1 := createRandomAnnotation(t)

	deleteArg := DeleteAnnotationsParams{VideoID: annotation1.VideoID, UserID: annotation1.UserID}
	err := testQueries.DeleteAnnotations(context.Background(), deleteArg)
	require.NoError(t, err)

	annotation2, err := testQueries.GetAnnotation(context.Background(), GetAnnotationParams{
		ID:     annotation1.ID,
		UserID: annotation1.UserID,
	})
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, annotation2)
}

func TestListAccount(t *testing.T) {
	var lastAnnotation Annotation
	for i := 0; i < 10; i++ {
		lastAnnotation = createRandomAnnotation(t)
	}

	arg := ListAnnotationsParams{
		VideoID: lastAnnotation.VideoID,
		UserID:  lastAnnotation.UserID,
		Limit:   5,
		Offset:  0,
	}

	annotations, err := testQueries.ListAnnotations(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, annotations)

	for _, annotation := range annotations {
		require.NotEmpty(t, annotation)
		require.Equal(t, lastAnnotation.VideoID, annotation.VideoID)
	}
}
