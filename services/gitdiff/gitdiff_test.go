	"code.gitea.io/gitea/models/db"
	issues_model "code.gitea.io/gitea/models/issues"
	diff := `diff --git a/newfile2 b/newfile2
	diff2 := `diff --git "a/A \\ B" "b/A \\ B"
	diff2a := `diff --git "a/A \\ B" b/A/B
	diff3 := `diff --git a/README.md b/README.md

	issue := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{ID: 2}).(*issues_model.Issue)
	assert.NoError(t, diff.LoadComments(db.DefaultContext, issue, user))
	assert.False(t, (&DiffLine{Type: DiffLineAdd, Comments: []*issues_model.Comment{{Content: "bla"}}}).CanComment())
	assert.Equal(t, "previous", (&DiffLine{Comments: []*issues_model.Comment{{Line: -3}}}).GetCommentSide())
	assert.Equal(t, "proposed", (&DiffLine{Comments: []*issues_model.Comment{{Line: 3}}}).GetCommentSide())
	gitRepo, err := git.OpenRepository(git.DefaultContext, "./testdata/academic-module")