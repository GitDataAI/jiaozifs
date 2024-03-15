package integrationtest

import (
	"context"
	"net/http"

	"github.com/GitDataAI/jiaozifs/utils/hash"

	"github.com/GitDataAI/jiaozifs/api"
	apiimpl "github.com/GitDataAI/jiaozifs/api/api_impl"
	"github.com/GitDataAI/jiaozifs/utils"
	"github.com/smartystreets/goconvey/convey"
)

func WipObjectSpec(ctx context.Context, urlStr string) func(c convey.C) {
	client, _ := api.NewClient(urlStr + apiimpl.APIV1Prefix)
	return func(c convey.C) {
		userName := "jude"
		repoName := "hash"
		branchName := "feat/wip_obj_test"

		c.Convey("init", func(_ convey.C) {
			_ = createUser(ctx, client, userName)
			loginAndSwitch(ctx, client, userName, false)
			_ = createRepo(ctx, client, repoName, false)
			_ = createBranch(ctx, client, userName, repoName, "main", branchName)
			_ = createWip(ctx, client, userName, repoName, branchName)
			_ = uploadObject(ctx, client, userName, repoName, branchName, "m.dat")
			_ = uploadObject(ctx, client, userName, repoName, branchName, "g/m.dat")
		})

		c.Convey("head object", func(c convey.C) {
			c.Convey("no auth", func() {
				re := client.RequestEditors
				client.RequestEditors = nil
				resp, err := client.HeadObject(ctx, userName, repoName, &api.HeadObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				client.RequestEditors = re
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("fail to head object in non exit user", func() {
				resp, err := client.HeadObject(ctx, "mock user", repoName, &api.HeadObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to head object in non exit repo", func() {
				resp, err := client.HeadObject(ctx, userName, "fakerepo", &api.HeadObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to head object in non exit branch", func() {
				resp, err := client.HeadObject(ctx, userName, repoName, &api.HeadObjectParams{
					RefName: "mockref",
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("forbidden head object in others", func() {
				resp, err := client.HeadObject(ctx, "jimmy", "happygo", &api.HeadObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("empty path", func() {
				resp, err := client.HeadObject(ctx, userName, repoName, &api.HeadObjectParams{
					Path: "",
					Type: api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
			})

			c.Convey("not exit path", func() {
				resp, err := client.HeadObject(ctx, userName, repoName, &api.HeadObjectParams{
					RefName: branchName,
					Path:    "c/d.txt",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
			})

			c.Convey("success to head object", func() {
				resp, err := client.HeadObject(ctx, userName, repoName, &api.HeadObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
			})
		})

		c.Convey("get object", func(c convey.C) {
			c.Convey("no auth", func() {
				re := client.RequestEditors
				client.RequestEditors = nil
				resp, err := client.GetObject(ctx, userName, repoName, &api.GetObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				client.RequestEditors = re
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("fail to get object in non exit user", func() {
				resp, err := client.GetObject(ctx, "mock user", repoName, &api.GetObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to get object in non exit repo", func() {
				resp, err := client.GetObject(ctx, userName, "fakerepo", &api.GetObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to get object in non exit branch", func() {
				resp, err := client.GetObject(ctx, userName, repoName, &api.GetObjectParams{
					RefName: "mockref",
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("forbidden get object in others", func() {
				resp, err := client.GetObject(ctx, "jimmy", "happygo", &api.GetObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("empty path", func() {
				resp, err := client.GetObject(ctx, userName, repoName, &api.GetObjectParams{
					RefName: branchName,
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
			})

			c.Convey("not exit path", func() {
				resp, err := client.GetObject(ctx, userName, repoName, &api.GetObjectParams{
					RefName: branchName,
					Path:    "c/d.txt",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
			})

			c.Convey("success to get object", func() {
				resp, err := client.GetObject(ctx, userName, repoName, &api.GetObjectParams{
					RefName: branchName,
					Path:    "m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
			})
		})

		c.Convey("delete object", func(c convey.C) {
			c.Convey("no auth", func() {
				re := client.RequestEditors
				client.RequestEditors = nil
				resp, err := client.DeleteObject(ctx, userName, repoName, &api.DeleteObjectParams{
					RefName: branchName,
					Path:    "g/m.dat",
				})
				client.RequestEditors = re
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("fail to delete object in non exit user", func() {
				resp, err := client.DeleteObject(ctx, "mockUser", repoName, &api.DeleteObjectParams{
					RefName: branchName,
					Path:    "g/m.dat",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to delete object in non exit repo", func() {
				resp, err := client.DeleteObject(ctx, userName, "fakerepo", &api.DeleteObjectParams{
					RefName: branchName,
					Path:    "g/m.dat",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to delete object in non exit branch", func() {
				resp, err := client.DeleteObject(ctx, userName, repoName, &api.DeleteObjectParams{
					RefName: "mockref",
					Path:    "g/m.dat",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("forbidden delete object in others", func() {
				resp, err := client.DeleteObject(ctx, "jimmy", "happygo", &api.DeleteObjectParams{
					RefName: "main",
					Path:    "g/m.dat",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("empty path", func() {
				resp, err := client.DeleteObject(ctx, userName, repoName, &api.DeleteObjectParams{
					RefName: branchName,
					Path:    "",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
			})

			c.Convey("not exit path", func() {
				resp, err := client.DeleteObject(ctx, userName, repoName, &api.DeleteObjectParams{
					RefName: branchName,
					Path:    "mm/t.dat",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
			})

			c.Convey("success to delete object", func(c convey.C) {
				//ensure exit
				resp, err := client.HeadObject(ctx, userName, repoName, &api.HeadObjectParams{
					RefName: branchName,
					Path:    "g/m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				resp, err = client.DeleteObject(ctx, userName, repoName, &api.DeleteObjectParams{
					RefName: branchName,
					Path:    "g/m.dat",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				c.Convey("commit changes", func() {
					_ = commitWip(ctx, client, userName, repoName, branchName, "test")
				})

				//ensure not exit
				resp, err = client.HeadObject(ctx, userName, repoName, &api.HeadObjectParams{
					RefName: branchName,
					Path:    "g/m.dat",
					Type:    api.RefTypeWip,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
			})
		})
		testBranchName := "test/empty_branch"

		c.Convey("update objects and create empty branch", func(_ convey.C) {
			uploadObject(ctx, client, userName, repoName, branchName, "a/m.dat")
			uploadObject(ctx, client, userName, repoName, branchName, "a/b.dat")
			uploadObject(ctx, client, userName, repoName, branchName, "b.dat")
			uploadObject(ctx, client, userName, repoName, branchName, "c.dat")

			createBranch(ctx, client, userName, repoName, "main", testBranchName)
			createWip(ctx, client, userName, repoName, testBranchName)
		})

		c.Convey("get wip success on init", func(_ convey.C) {
			resp, err := client.GetWipChanges(ctx, userName, repoName, &api.GetWipChangesParams{
				RefName: testBranchName,
			})
			convey.So(err, convey.ShouldBeNil)
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

			result, err := api.ParseGetWipChangesResponse(resp)
			convey.So(err, convey.ShouldBeNil)
			convey.So(*result.JSON200, convey.ShouldHaveLength, 0)
		})

		c.Convey("get wip changes", func(c convey.C) {
			c.Convey("no auth", func() {
				re := client.RequestEditors
				client.RequestEditors = nil
				resp, err := client.GetWipChanges(ctx, userName, repoName, &api.GetWipChangesParams{
					RefName: branchName,
				})
				client.RequestEditors = re
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("fail to get object in non exit user", func() {
				resp, err := client.GetWipChanges(ctx, "mock user", repoName, &api.GetWipChangesParams{
					RefName: branchName,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to get object in non exit repo", func() {
				resp, err := client.GetWipChanges(ctx, userName, "fakerepo", &api.GetWipChangesParams{
					RefName: branchName,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to get object in non exit branch", func() {
				resp, err := client.GetWipChanges(ctx, userName, repoName, &api.GetWipChangesParams{
					RefName: "mockref",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("forbidden get object in others", func() {
				resp, err := client.GetWipChanges(ctx, "jimmy", "happygo", &api.GetWipChangesParams{
					RefName: "main",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("not exit path", func() {
				resp, err := client.GetWipChanges(ctx, userName, repoName, &api.GetWipChangesParams{
					RefName: branchName,
					Path:    utils.String("a/b/c/d"),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				result, err := api.ParseGetWipChangesResponse(resp)
				convey.So(err, convey.ShouldBeNil)
				convey.So(*result.JSON200, convey.ShouldHaveLength, 0)
			})

			c.Convey("success to get object", func() {
				resp, err := client.GetWipChanges(ctx, userName, repoName, &api.GetWipChangesParams{
					RefName: branchName,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				result, err := api.ParseGetWipChangesResponse(resp)
				convey.So(err, convey.ShouldBeNil)
				convey.So(*result.JSON200, convey.ShouldHaveLength, 4)
			})
		})

		c.Convey("revert wip changes", func(c convey.C) {
			c.Convey("no auth", func() {
				re := client.RequestEditors
				client.RequestEditors = nil
				resp, err := client.RevertWipChanges(ctx, userName, repoName, &api.RevertWipChangesParams{
					RefName: branchName,
				})
				client.RequestEditors = re
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("fail to revert changes in non exit user", func() {
				resp, err := client.RevertWipChanges(ctx, "mockUser", repoName, &api.RevertWipChangesParams{
					RefName: branchName,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to revert changes in non exit repo", func() {
				resp, err := client.RevertWipChanges(ctx, userName, "fakeRepo", &api.RevertWipChangesParams{
					RefName: branchName,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to revert changes in non exit branch", func() {
				resp, err := client.RevertWipChanges(ctx, userName, repoName, &api.RevertWipChangesParams{
					RefName: "mockref",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("forbidden revert changes in others", func() {
				resp, err := client.RevertWipChanges(ctx, "jimmy", "happygo", &api.RevertWipChangesParams{
					RefName: branchName,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusForbidden)
			})

			c.Convey("not exit path", func() {
				resp, err := client.RevertWipChanges(ctx, userName, repoName, &api.RevertWipChangesParams{
					RefName:    branchName,
					PathPrefix: utils.String("a/b/c/d"),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("success to revert changes", func() {
				resp, err := client.RevertWipChanges(ctx, userName, repoName, &api.RevertWipChangesParams{
					RefName: branchName,
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				{
					resp, err = client.GetWipChanges(ctx, userName, repoName, &api.GetWipChangesParams{
						RefName: branchName,
					})
					convey.So(err, convey.ShouldBeNil)
					convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

					result, err := api.ParseGetWipChangesResponse(resp)
					convey.So(err, convey.ShouldBeNil)
					convey.So(*result.JSON200, convey.ShouldHaveLength, 0)
				}
			})
		})

	}
}

func UpdateWipSpec(ctx context.Context, urlStr string) func(c convey.C) {
	client, _ := api.NewClient(urlStr + apiimpl.APIV1Prefix)
	var wip *api.Wip
	return func(c convey.C) {
		userName := "milly"
		repoName := "update_wip_test"
		branchName := "main"

		c.Convey("create wip", func(_ convey.C) {
			_ = createUser(ctx, client, userName)
			loginAndSwitch(ctx, client, userName, false)
			_ = createRepo(ctx, client, repoName, false)
			_ = createWip(ctx, client, userName, repoName, branchName)

			//make wip base commit has value
			_ = uploadObject(ctx, client, userName, repoName, branchName, "a.txt")
			_ = commitWip(ctx, client, userName, repoName, branchName, "test")

			_ = uploadObject(ctx, client, userName, repoName, branchName, "m.dat")
			_ = uploadObject(ctx, client, userName, repoName, branchName, "g/m.dat")
		})

		c.Convey("get wip", func(_ convey.C) {
			resp, err := client.GetWip(ctx, userName, repoName, &api.GetWipParams{
				RefName: branchName,
			})
			convey.So(err, convey.ShouldBeNil)
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

			result, err := api.ParseGetWipResponse(resp)
			convey.So(err, convey.ShouldBeNil)
			wip = result.JSON200
		})

		c.Convey("update wip", func(c convey.C) {
			c.Convey("no auth", func() {
				re := client.RequestEditors
				client.RequestEditors = nil
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(hash.Empty.Hex()),
					BaseCommit:  utils.String(hash.Empty.Hex()),
				})
				client.RequestEditors = re
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("update wip in not exit repo", func() {
				resp, err := client.UpdateWip(ctx, userName, "mock_repo", &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(hash.Empty.Hex()),
					BaseCommit:  utils.String(hash.Empty.Hex()),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("update wip in non exit user", func() {
				resp, err := client.UpdateWip(ctx, "telo", repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(hash.Empty.Hex()),
					BaseCommit:  utils.String(hash.Empty.Hex()),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("update wip in other's repo", func() {
				resp, err := client.UpdateWip(ctx, "jimmy", "happygo", &api.UpdateWipParams{RefName: "main"}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(hash.Empty.Hex()),
					BaseCommit:  utils.String(hash.Empty.Hex()),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusUnauthorized)
			})

			c.Convey("update wip in non exit branch", func() {
				//delete
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: "feat/mock_ref"}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(hash.Empty.Hex()),
					BaseCommit:  utils.String(hash.Empty.Hex()),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("update wip with fail base commit", func() {
				//delete
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(hash.Empty.Hex()),
					BaseCommit:  utils.String("ddd"),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusInternalServerError)
			})

			c.Convey("update wip with fail tree hash", func() {
				//delete
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String("ddd"),
					BaseCommit:  utils.String(hash.Empty.Hex()),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusInternalServerError)
			})

			c.Convey("update wip successful", func() {
				//delete
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(hash.Empty.Hex()),
					BaseCommit:  utils.String(hash.Empty.Hex()),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				//ensure delete work
				getResp, err := client.GetWip(ctx, userName, repoName, &api.GetWipParams{RefName: branchName})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				updatedWip, err := api.ParseGetWipResponse(getResp)
				convey.So(err, convey.ShouldBeNil)
				convey.So((*updatedWip.JSON200).BaseCommit, convey.ShouldEqual, "")
				convey.So((*updatedWip.JSON200).CurrentTree, convey.ShouldEqual, "")
			})

			c.Convey("fail to update non exit tree hash", func() {
				//delete
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String("6161616161"),
					BaseCommit:  utils.String(wip.BaseCommit),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("fail to update non exit base commit", func() {
				//delete
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(wip.CurrentTree),
					BaseCommit:  utils.String("6161616161"),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
			})

			c.Convey("update wip to non empty successful", func() {
				//delete
				resp, err := client.UpdateWip(ctx, userName, repoName, &api.UpdateWipParams{RefName: branchName}, api.UpdateWipJSONRequestBody{
					CurrentTree: utils.String(wip.CurrentTree),
					BaseCommit:  utils.String(wip.BaseCommit),
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				//ensure delete work
				getResp, err := client.GetWip(ctx, userName, repoName, &api.GetWipParams{RefName: branchName})
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				updatedWip, err := api.ParseGetWipResponse(getResp)
				convey.So(err, convey.ShouldBeNil)
				convey.So((*updatedWip.JSON200).BaseCommit, convey.ShouldEqual, wip.BaseCommit)
				convey.So((*updatedWip.JSON200).CurrentTree, convey.ShouldEqual, wip.CurrentTree)
			})
		})
	}
}
