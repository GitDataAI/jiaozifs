package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/jiaozifs/jiaozifs/utils"

	"github.com/jiaozifs/jiaozifs/testhelper"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/google/go-cmp/cmp"

	"github.com/google/uuid"

	"github.com/jiaozifs/jiaozifs/models"

	"github.com/stretchr/testify/require"
)

var dbTimeCmpOpt = cmp.Comparer(func(x, y time.Time) bool {
	return x.Unix() == y.Unix()
})

func TestNewUserRepo(t *testing.T) {
	ctx := context.Background()
	postgres, _, db := testhelper.SetupDatabase(ctx, t)
	defer postgres.Stop() //nolint

	repo := models.NewUserRepo(db)

	userModel := &models.User{}
	require.NoError(t, gofakeit.Struct(userModel))
	newUser, err := repo.Insert(ctx, userModel)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, newUser.ID)

	user, err := repo.Get(ctx, &models.GetUserParam{ID: newUser.ID})
	require.NoError(t, err)

	require.True(t, cmp.Equal(userModel, user, dbTimeCmpOpt))

	ep, err := repo.GetEPByName(ctx, newUser.Name)
	require.NoError(t, err)
	require.True(t, cmp.Equal(userModel.EncryptedPassword, ep))

	userByEmail, err := repo.Get(ctx, &models.GetUserParam{Email: utils.String(newUser.Email)})
	require.NoError(t, err)
	require.True(t, cmp.Equal(userModel, userByEmail, dbTimeCmpOpt))

	userByName, err := repo.Get(ctx, &models.GetUserParam{Name: utils.String(newUser.Name)})
	require.NoError(t, err)
	require.True(t, cmp.Equal(userModel, userByName, dbTimeCmpOpt))
}
