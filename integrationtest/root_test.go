package integrationtest

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	ctx := context.Background()
	urlStr, cancel := SetupDaemon(t, ctx)
	defer cancel()

	convey.Convey("status test", t, StatusSpec(ctx, urlStr))

	convey.Convey("user test", t, UserSpec(ctx, urlStr))
	convey.Convey("aksk test", t, AkSkSpec(ctx, urlStr))
	convey.Convey("repo test", t, RepoSpec(ctx, urlStr))
	convey.Convey("branch test", t, BranchSpec(ctx, urlStr))
	convey.Convey("object test", t, ObjectSpec(ctx, urlStr))
	convey.Convey("wip test", t, WipSpec(ctx, urlStr))
	convey.Convey("wip object test", t, WipObjectSpec(ctx, urlStr))
	convey.Convey("update wip test", t, UpdateWipSpec(ctx, urlStr))
	convey.Convey("get entries test", t, GetEntriesInRefSpec(ctx, urlStr))
	convey.Convey("commit changes test", t, GetCommitChangesSpec(ctx, urlStr))
	convey.Convey("merge request test", t, MergeRequestSpec(ctx, urlStr))
	convey.Convey("group test", t, GroupSpec(ctx, urlStr))
	convey.Convey("member test", t, MemberSpec(ctx, urlStr))
}
