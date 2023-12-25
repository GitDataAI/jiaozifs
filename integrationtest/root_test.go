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

	convey.Convey("user test", t, UserSpec(ctx, urlStr))
	convey.Convey("repo test", t, RepoSpec(ctx, urlStr))
	// convey.Convey("branch test", t, BranchSpec(ctx, urlStr))
	// convey.Convey("wip test", t, WipSpec(ctx, urlStr))
	// convey.Convey("object test", t, ObjectSpec(ctx, urlStr))
	// convey.Convey("wip object test", t, WipObjectSpec(ctx, urlStr))
	// convey.Convey("commit test", t, GetEntriesInRefSpec(ctx, urlStr))
}
