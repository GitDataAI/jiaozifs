package controller

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/GitDataAI/jiaozifs/auth/rbac"
	"github.com/GitDataAI/jiaozifs/block/params"
	"github.com/GitDataAI/jiaozifs/controller/validator"
	"github.com/GitDataAI/jiaozifs/models/rbacmodel"
	"github.com/GitDataAI/jiaozifs/versionmgr"

	"github.com/GitDataAI/jiaozifs/api"
	"github.com/GitDataAI/jiaozifs/auth"
	"github.com/GitDataAI/jiaozifs/models"
	"github.com/GitDataAI/jiaozifs/utils"
	"go.uber.org/fx"
)

type BranchController struct {
	fx.In
	BaseController

	Repo                models.IRepo
	PublicStorageConfig params.AdapterConfig
}

func (bct BranchController) ListBranches(ctx context.Context, w *api.JiaozifsResponse, _ *http.Request, ownerName string, repositoryName string, params api.ListBranchesParams) {
	owner, err := bct.Repo.UserRepo().Get(ctx, models.NewGetUserParams().SetName(ownerName))
	if err != nil {
		w.Error(err)
		return
	}

	repository, err := bct.Repo.RepositoryRepo().Get(ctx, models.NewGetRepoParams().SetName(repositoryName).SetOwnerID(owner.ID))
	if err != nil {
		w.Error(err)
		return
	}

	if !bct.authorizeMember(ctx, w, repository.ID, rbac.Node{
		Permission: rbac.Permission{
			Action:   rbacmodel.ListBranchesAction,
			Resource: rbacmodel.RepoURArn(owner.ID.String(), repository.ID.String()),
		},
	}) {
		return
	}

	listBranchParams := models.NewListBranchParams()
	if params.Prefix != nil && len(*params.Prefix) > 0 {
		listBranchParams.SetName(*params.Prefix, models.PrefixMatch)
	}
	if params.After != nil && len(*params.After) > 0 {
		listBranchParams.SetAfter(*params.After)
	}
	pageAmount := utils.IntValue(params.Amount)
	if pageAmount > utils.DefaultMaxPerPage || pageAmount <= 0 {
		listBranchParams.SetAmount(utils.DefaultMaxPerPage)
	} else {
		listBranchParams.SetAmount(pageAmount)
	}

	branches, hasMore, err := bct.Repo.BranchRepo().List(ctx, listBranchParams.SetRepositoryID(repository.ID))
	if err != nil {
		w.Error(err)
		return
	}
	results := utils.Silent(utils.ArrMap(branches, branchToDto))
	pagMag := utils.PaginationFor(hasMore, results, "Name")
	pagination := api.Pagination{
		HasMore:    pagMag.HasMore,
		MaxPerPage: pagMag.MaxPerPage,
		NextOffset: pagMag.NextOffset,
		Results:    pagMag.Results,
	}
	w.JSON(api.BranchList{
		Pagination: pagination,
		Results:    results,
	})
}

func (bct BranchController) CreateBranch(ctx context.Context, w *api.JiaozifsResponse, _ *http.Request, body api.CreateBranchJSONRequestBody, ownerName string, repositoryName string) {
	if err := validator.ValidateBranchName(body.Name); err != nil {
		w.BadRequest(err.Error())
		return
	}

	operator, err := auth.GetOperator(ctx)
	if err != nil {
		w.Error(err)
		return
	}

	owner, err := bct.Repo.UserRepo().Get(ctx, models.NewGetUserParams().SetName(ownerName))
	if err != nil {
		w.Error(err)
		return
	}

	// Get repo
	repository, err := bct.Repo.RepositoryRepo().Get(ctx, models.NewGetRepoParams().SetOwnerID(owner.ID).SetName(repositoryName))
	if err != nil {
		w.Error(err)
		return
	}

	if !bct.authorizeMember(ctx, w, repository.ID, rbac.Node{
		Permission: rbac.Permission{
			Action:   rbacmodel.CreateBranchAction,
			Resource: rbacmodel.RepoURArn(owner.ID.String(), repository.ID.String()),
		},
	}) {
		return
	}

	//get source branch
	sourceBranch, err := bct.Repo.BranchRepo().Get(ctx, models.NewGetBranchParams().SetName(body.Source).SetRepositoryID(repository.ID))
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		w.Error(err)
		return
	}

	workRepo, err := versionmgr.NewWorkRepositoryFromConfig(ctx, operator, repository, bct.Repo, bct.PublicStorageConfig)
	if err != nil {
		w.Error(err)
		return
	}

	err = workRepo.CheckOut(ctx, versionmgr.InCommit, sourceBranch.CommitHash.Hex())
	if err != nil {
		w.Error(err)
		return
	}

	newBranch, err := workRepo.CreateBranch(ctx, body.Name)
	if err != nil {
		if strings.Contains(err.Error(), "already exit") {
			w.Code(http.StatusConflict)
			return
		}
		w.Error(err)
		return
	}

	w.JSON(utils.Silent(branchToDto(newBranch)), http.StatusCreated)
}

func (bct BranchController) DeleteBranch(ctx context.Context, w *api.JiaozifsResponse, _ *http.Request, ownerName string, repositoryName string, params api.DeleteBranchParams) {
	operator, err := auth.GetOperator(ctx)
	if err != nil {
		w.Error(err)
		return
	}

	owner, err := bct.Repo.UserRepo().Get(ctx, models.NewGetUserParams().SetName(ownerName))
	if err != nil {
		w.Error(err)
		return
	}

	// Get repo
	repository, err := bct.Repo.RepositoryRepo().Get(ctx, models.NewGetRepoParams().SetOwnerID(owner.ID).SetName(repositoryName))
	if err != nil {
		w.Error(err)
		return
	}

	if !bct.authorizeMember(ctx, w, repository.ID, rbac.Node{
		Permission: rbac.Permission{
			Action:   rbacmodel.DeleteBranchAction,
			Resource: rbacmodel.RepoURArn(owner.ID.String(), repository.ID.String()),
		},
	}) {
		return
	}

	if params.RefName == repository.HEAD {
		w.BadRequest("can not delete HEAD branch")
		return
	}

	workRepo, err := versionmgr.NewWorkRepositoryFromConfig(ctx, operator, repository, bct.Repo, bct.PublicStorageConfig)
	if err != nil {
		w.Error(err)
		return
	}

	err = workRepo.CheckOut(ctx, versionmgr.InBranch, params.RefName)
	if err != nil {
		w.Error(err)
		return
	}

	err = workRepo.DeleteBranch(ctx)
	if err != nil {
		w.Error(err)
		return
	}
	w.OK()
}

func (bct BranchController) GetBranch(ctx context.Context, w *api.JiaozifsResponse, _ *http.Request, ownerName string, repositoryName string, params api.GetBranchParams) {
	owner, err := bct.Repo.UserRepo().Get(ctx, models.NewGetUserParams().SetName(ownerName))
	if err != nil {
		w.Error(err)
		return
	}

	// Get repo
	repository, err := bct.Repo.RepositoryRepo().Get(ctx, models.NewGetRepoParams().SetOwnerID(owner.ID).SetName(repositoryName))
	if err != nil {
		w.Error(err)
		return
	}

	if !bct.authorizeMember(ctx, w, repository.ID, rbac.Node{
		Permission: rbac.Permission{
			Action:   rbacmodel.ReadRepositoryAction,
			Resource: rbacmodel.RepoURArn(owner.ID.String(), repository.ID.String()),
		},
	}) {
		return
	}

	// Get branch
	ref, err := bct.Repo.BranchRepo().Get(ctx, models.NewGetBranchParams().SetName(params.RefName).SetRepositoryID(repository.ID))
	if err != nil {
		w.Error(err)
		return
	}
	w.JSON(utils.Silent(branchToDto(ref)))
}

func branchToDto(in *models.Branch) (api.Branch, error) {
	return api.Branch{
		CommitHash:   in.CommitHash.Hex(),
		CreatedAt:    in.CreatedAt.UnixMilli(),
		CreatorId:    in.CreatorID,
		Description:  in.Description,
		Id:           in.ID,
		Name:         in.Name,
		RepositoryId: in.RepositoryID,
		UpdatedAt:    in.UpdatedAt.UnixMilli(),
	}, nil
}
