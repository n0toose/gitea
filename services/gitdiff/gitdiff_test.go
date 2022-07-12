	"code.gitea.io/gitea/models"
	var diff = `diff --git a/newfile2 b/newfile2
	var diff2 = `diff --git "a/A \\ B" "b/A \\ B"
	var diff2a = `diff --git "a/A \\ B" b/A/B
	var diff3 = `diff --git a/README.md b/README.md
	issue := unittest.AssertExistsAndLoadBean(t, &models.Issue{ID: 2}).(*models.Issue)
	assert.NoError(t, diff.LoadComments(issue, user))
	assert.False(t, (&DiffLine{Type: DiffLineAdd, Comments: []*models.Comment{{Content: "bla"}}}).CanComment())
	assert.Equal(t, "previous", (&DiffLine{Comments: []*models.Comment{{Line: -3}}}).GetCommentSide())
	assert.Equal(t, "proposed", (&DiffLine{Comments: []*models.Comment{{Line: 3}}}).GetCommentSide())
	gitRepo, err := git.OpenRepository("./testdata/academic-module")