package models_test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-cmp/cmp"
	"github.com/jiaozifs/jiaozifs/models"
	"github.com/jiaozifs/jiaozifs/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCommitRepo(t *testing.T) {
	ctx := context.Background()
	postgres, _, db := testhelper.SetupDatabase(ctx, t)
	defer postgres.Stop() //nolint

	repoID := uuid.New()
	commitRepo := models.NewCommitRepo(db, repoID)
	require.Equal(t, commitRepo.RepositoryID(), repoID)

	commitModel := &models.Commit{}
	require.NoError(t, gofakeit.Struct(commitModel))
	commitModel.RepositoryID = repoID
	newCommitModel, err := commitRepo.Insert(ctx, commitModel)
	require.NoError(t, err)
	commitModel, err = commitRepo.Commit(ctx, commitModel.Hash)
	require.NoError(t, err)

	require.True(t, cmp.Equal(commitModel, newCommitModel, dbTimeCmpOpt))

	t.Run("mis match repo id", func(t *testing.T) {
		mistMatchModel := &models.Commit{}
		require.NoError(t, gofakeit.Struct(mistMatchModel))
		_, err := commitRepo.Insert(ctx, mistMatchModel)
		require.ErrorIs(t, err, models.ErrRepoIDMisMatch)
	})
}
