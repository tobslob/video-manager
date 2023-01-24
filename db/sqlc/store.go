package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Store provides all  functions to excute db queries and transactions
type Store interface {
	Querier
	CreatVideoWithMetadataTx(ctx context.Context, arg CreateVideoWithMetadata) (CreatVideoWithMetadataTxResult, error)
}

// SQLStore provides all  functions to excute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreateVideoWithMetadata struct {
	CreateVideoParams
	CreateMetadataParams
}

type CreatVideoWithMetadataTxResult struct {
	Video    Video
	Metadata Metadata
}

func (store *SQLStore) CreatVideoWithMetadataTx(ctx context.Context, arg CreateVideoWithMetadata) (CreatVideoWithMetadataTxResult, error) {
	var result CreatVideoWithMetadataTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		videoId, _ := uuid.NewRandom()
		metaId, _ := uuid.NewRandom()

		result.Video, err = q.CreateVideo(ctx, CreateVideoParams{
			ID:       videoId,
			Url:      arg.Url,
			UserID:   arg.UserID,
			Duration: arg.Duration,
			Title:    arg.Title,
		})
		if err != nil {
			return err
		}

		result.Metadata, err = q.CreateMetadata(ctx, CreateMetadataParams{
			ID:           metaId,
			VideoID:      result.Video.ID,
			Width:        arg.Width,
			Height:       arg.Height,
			FileType:     arg.FileType,
			FileSize:     arg.FileSize,
			AccessedDate: time.Now(),
			Resolutions:  arg.Resolutions,
			LastModify:   time.Now(),
			Keywords:     arg.Keywords,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
