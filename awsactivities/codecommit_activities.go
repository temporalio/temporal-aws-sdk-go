package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codecommit/codecommitiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CodeCommitActivities struct {
	client codecommitiface.CodeCommitAPI
}

func NewCodeCommitActivities(session *session.Session, config ...*aws.Config) *CodeCommitActivities {
	client := codecommit.New(session, config...)
	return &CodeCommitActivities{client: client}
}

func (a *CodeCommitActivities) AssociateApprovalRuleTemplateWithRepository(ctx context.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) (*codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput, error) {
	return a.client.AssociateApprovalRuleTemplateWithRepositoryWithContext(ctx, input)
}

func (a *CodeCommitActivities) BatchAssociateApprovalRuleTemplateWithRepositories(ctx context.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
	return a.client.BatchAssociateApprovalRuleTemplateWithRepositoriesWithContext(ctx, input)
}

func (a *CodeCommitActivities) BatchDescribeMergeConflicts(ctx context.Context, input *codecommit.BatchDescribeMergeConflictsInput) (*codecommit.BatchDescribeMergeConflictsOutput, error) {
	return a.client.BatchDescribeMergeConflictsWithContext(ctx, input)
}

func (a *CodeCommitActivities) BatchDisassociateApprovalRuleTemplateFromRepositories(ctx context.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
	return a.client.BatchDisassociateApprovalRuleTemplateFromRepositoriesWithContext(ctx, input)
}

func (a *CodeCommitActivities) BatchGetCommits(ctx context.Context, input *codecommit.BatchGetCommitsInput) (*codecommit.BatchGetCommitsOutput, error) {
	return a.client.BatchGetCommitsWithContext(ctx, input)
}

func (a *CodeCommitActivities) BatchGetRepositories(ctx context.Context, input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error) {
	return a.client.BatchGetRepositoriesWithContext(ctx, input)
}

func (a *CodeCommitActivities) CreateApprovalRuleTemplate(ctx context.Context, input *codecommit.CreateApprovalRuleTemplateInput) (*codecommit.CreateApprovalRuleTemplateOutput, error) {
	return a.client.CreateApprovalRuleTemplateWithContext(ctx, input)
}

func (a *CodeCommitActivities) CreateBranch(ctx context.Context, input *codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error) {
	return a.client.CreateBranchWithContext(ctx, input)
}

func (a *CodeCommitActivities) CreateCommit(ctx context.Context, input *codecommit.CreateCommitInput) (*codecommit.CreateCommitOutput, error) {
	return a.client.CreateCommitWithContext(ctx, input)
}

func (a *CodeCommitActivities) CreatePullRequestApprovalRule(ctx context.Context, input *codecommit.CreatePullRequestApprovalRuleInput) (*codecommit.CreatePullRequestApprovalRuleOutput, error) {
	return a.client.CreatePullRequestApprovalRuleWithContext(ctx, input)
}

func (a *CodeCommitActivities) CreateRepository(ctx context.Context, input *codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error) {
	return a.client.CreateRepositoryWithContext(ctx, input)
}

func (a *CodeCommitActivities) CreateUnreferencedMergeCommit(ctx context.Context, input *codecommit.CreateUnreferencedMergeCommitInput) (*codecommit.CreateUnreferencedMergeCommitOutput, error) {
	return a.client.CreateUnreferencedMergeCommitWithContext(ctx, input)
}

func (a *CodeCommitActivities) DeleteApprovalRuleTemplate(ctx context.Context, input *codecommit.DeleteApprovalRuleTemplateInput) (*codecommit.DeleteApprovalRuleTemplateOutput, error) {
	return a.client.DeleteApprovalRuleTemplateWithContext(ctx, input)
}

func (a *CodeCommitActivities) DeleteBranch(ctx context.Context, input *codecommit.DeleteBranchInput) (*codecommit.DeleteBranchOutput, error) {
	return a.client.DeleteBranchWithContext(ctx, input)
}

func (a *CodeCommitActivities) DeleteCommentContent(ctx context.Context, input *codecommit.DeleteCommentContentInput) (*codecommit.DeleteCommentContentOutput, error) {
	return a.client.DeleteCommentContentWithContext(ctx, input)
}

func (a *CodeCommitActivities) DeleteFile(ctx context.Context, input *codecommit.DeleteFileInput) (*codecommit.DeleteFileOutput, error) {
	return a.client.DeleteFileWithContext(ctx, input)
}

func (a *CodeCommitActivities) DeletePullRequestApprovalRule(ctx context.Context, input *codecommit.DeletePullRequestApprovalRuleInput) (*codecommit.DeletePullRequestApprovalRuleOutput, error) {
	return a.client.DeletePullRequestApprovalRuleWithContext(ctx, input)
}

func (a *CodeCommitActivities) DeleteRepository(ctx context.Context, input *codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error) {
	return a.client.DeleteRepositoryWithContext(ctx, input)
}

func (a *CodeCommitActivities) DescribeMergeConflicts(ctx context.Context, input *codecommit.DescribeMergeConflictsInput) (*codecommit.DescribeMergeConflictsOutput, error) {
	return a.client.DescribeMergeConflictsWithContext(ctx, input)
}

func (a *CodeCommitActivities) DescribePullRequestEvents(ctx context.Context, input *codecommit.DescribePullRequestEventsInput) (*codecommit.DescribePullRequestEventsOutput, error) {
	return a.client.DescribePullRequestEventsWithContext(ctx, input)
}

func (a *CodeCommitActivities) DisassociateApprovalRuleTemplateFromRepository(ctx context.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) (*codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput, error) {
	return a.client.DisassociateApprovalRuleTemplateFromRepositoryWithContext(ctx, input)
}

func (a *CodeCommitActivities) EvaluatePullRequestApprovalRules(ctx context.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) (*codecommit.EvaluatePullRequestApprovalRulesOutput, error) {
	return a.client.EvaluatePullRequestApprovalRulesWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetApprovalRuleTemplate(ctx context.Context, input *codecommit.GetApprovalRuleTemplateInput) (*codecommit.GetApprovalRuleTemplateOutput, error) {
	return a.client.GetApprovalRuleTemplateWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetBlob(ctx context.Context, input *codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error) {
	return a.client.GetBlobWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetBranch(ctx context.Context, input *codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error) {
	return a.client.GetBranchWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetComment(ctx context.Context, input *codecommit.GetCommentInput) (*codecommit.GetCommentOutput, error) {
	return a.client.GetCommentWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetCommentReactions(ctx context.Context, input *codecommit.GetCommentReactionsInput) (*codecommit.GetCommentReactionsOutput, error) {
	return a.client.GetCommentReactionsWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetCommentsForComparedCommit(ctx context.Context, input *codecommit.GetCommentsForComparedCommitInput) (*codecommit.GetCommentsForComparedCommitOutput, error) {
	return a.client.GetCommentsForComparedCommitWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetCommit(ctx context.Context, input *codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error) {
	return a.client.GetCommitWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetDifferences(ctx context.Context, input *codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error) {
	return a.client.GetDifferencesWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetFile(ctx context.Context, input *codecommit.GetFileInput) (*codecommit.GetFileOutput, error) {
	return a.client.GetFileWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetFolder(ctx context.Context, input *codecommit.GetFolderInput) (*codecommit.GetFolderOutput, error) {
	return a.client.GetFolderWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetMergeCommit(ctx context.Context, input *codecommit.GetMergeCommitInput) (*codecommit.GetMergeCommitOutput, error) {
	return a.client.GetMergeCommitWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetMergeConflicts(ctx context.Context, input *codecommit.GetMergeConflictsInput) (*codecommit.GetMergeConflictsOutput, error) {
	return a.client.GetMergeConflictsWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetMergeOptions(ctx context.Context, input *codecommit.GetMergeOptionsInput) (*codecommit.GetMergeOptionsOutput, error) {
	return a.client.GetMergeOptionsWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetPullRequestApprovalStates(ctx context.Context, input *codecommit.GetPullRequestApprovalStatesInput) (*codecommit.GetPullRequestApprovalStatesOutput, error) {
	return a.client.GetPullRequestApprovalStatesWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetPullRequestOverrideState(ctx context.Context, input *codecommit.GetPullRequestOverrideStateInput) (*codecommit.GetPullRequestOverrideStateOutput, error) {
	return a.client.GetPullRequestOverrideStateWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetRepository(ctx context.Context, input *codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error) {
	return a.client.GetRepositoryWithContext(ctx, input)
}

func (a *CodeCommitActivities) GetRepositoryTriggers(ctx context.Context, input *codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error) {
	return a.client.GetRepositoryTriggersWithContext(ctx, input)
}

func (a *CodeCommitActivities) ListApprovalRuleTemplates(ctx context.Context, input *codecommit.ListApprovalRuleTemplatesInput) (*codecommit.ListApprovalRuleTemplatesOutput, error) {
	return a.client.ListApprovalRuleTemplatesWithContext(ctx, input)
}

func (a *CodeCommitActivities) ListAssociatedApprovalRuleTemplatesForRepository(ctx context.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
	return a.client.ListAssociatedApprovalRuleTemplatesForRepositoryWithContext(ctx, input)
}

func (a *CodeCommitActivities) ListBranches(ctx context.Context, input *codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error) {
	return a.client.ListBranchesWithContext(ctx, input)
}

func (a *CodeCommitActivities) ListPullRequests(ctx context.Context, input *codecommit.ListPullRequestsInput) (*codecommit.ListPullRequestsOutput, error) {
	return a.client.ListPullRequestsWithContext(ctx, input)
}

func (a *CodeCommitActivities) ListRepositories(ctx context.Context, input *codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error) {
	return a.client.ListRepositoriesWithContext(ctx, input)
}

func (a *CodeCommitActivities) ListRepositoriesForApprovalRuleTemplate(ctx context.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error) {
	return a.client.ListRepositoriesForApprovalRuleTemplateWithContext(ctx, input)
}

func (a *CodeCommitActivities) ListTagsForResource(ctx context.Context, input *codecommit.ListTagsForResourceInput) (*codecommit.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CodeCommitActivities) MergeBranchesByFastForward(ctx context.Context, input *codecommit.MergeBranchesByFastForwardInput) (*codecommit.MergeBranchesByFastForwardOutput, error) {
	return a.client.MergeBranchesByFastForwardWithContext(ctx, input)
}

func (a *CodeCommitActivities) MergeBranchesBySquash(ctx context.Context, input *codecommit.MergeBranchesBySquashInput) (*codecommit.MergeBranchesBySquashOutput, error) {
	return a.client.MergeBranchesBySquashWithContext(ctx, input)
}

func (a *CodeCommitActivities) MergeBranchesByThreeWay(ctx context.Context, input *codecommit.MergeBranchesByThreeWayInput) (*codecommit.MergeBranchesByThreeWayOutput, error) {
	return a.client.MergeBranchesByThreeWayWithContext(ctx, input)
}

func (a *CodeCommitActivities) MergePullRequestByFastForward(ctx context.Context, input *codecommit.MergePullRequestByFastForwardInput) (*codecommit.MergePullRequestByFastForwardOutput, error) {
	return a.client.MergePullRequestByFastForwardWithContext(ctx, input)
}

func (a *CodeCommitActivities) MergePullRequestBySquash(ctx context.Context, input *codecommit.MergePullRequestBySquashInput) (*codecommit.MergePullRequestBySquashOutput, error) {
	return a.client.MergePullRequestBySquashWithContext(ctx, input)
}

func (a *CodeCommitActivities) MergePullRequestByThreeWay(ctx context.Context, input *codecommit.MergePullRequestByThreeWayInput) (*codecommit.MergePullRequestByThreeWayOutput, error) {
	return a.client.MergePullRequestByThreeWayWithContext(ctx, input)
}

func (a *CodeCommitActivities) OverridePullRequestApprovalRules(ctx context.Context, input *codecommit.OverridePullRequestApprovalRulesInput) (*codecommit.OverridePullRequestApprovalRulesOutput, error) {
	return a.client.OverridePullRequestApprovalRulesWithContext(ctx, input)
}

func (a *CodeCommitActivities) PostCommentForComparedCommit(ctx context.Context, input *codecommit.PostCommentForComparedCommitInput) (*codecommit.PostCommentForComparedCommitOutput, error) {
	return a.client.PostCommentForComparedCommitWithContext(ctx, input)
}

func (a *CodeCommitActivities) PostCommentReply(ctx context.Context, input *codecommit.PostCommentReplyInput) (*codecommit.PostCommentReplyOutput, error) {
	return a.client.PostCommentReplyWithContext(ctx, input)
}

func (a *CodeCommitActivities) PutCommentReaction(ctx context.Context, input *codecommit.PutCommentReactionInput) (*codecommit.PutCommentReactionOutput, error) {
	return a.client.PutCommentReactionWithContext(ctx, input)
}

func (a *CodeCommitActivities) PutFile(ctx context.Context, input *codecommit.PutFileInput) (*codecommit.PutFileOutput, error) {
	return a.client.PutFileWithContext(ctx, input)
}

func (a *CodeCommitActivities) PutRepositoryTriggers(ctx context.Context, input *codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error) {
	return a.client.PutRepositoryTriggersWithContext(ctx, input)
}

func (a *CodeCommitActivities) TagResource(ctx context.Context, input *codecommit.TagResourceInput) (*codecommit.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CodeCommitActivities) TestRepositoryTriggers(ctx context.Context, input *codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error) {
	return a.client.TestRepositoryTriggersWithContext(ctx, input)
}

func (a *CodeCommitActivities) UntagResource(ctx context.Context, input *codecommit.UntagResourceInput) (*codecommit.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdateApprovalRuleTemplateContent(ctx context.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) (*codecommit.UpdateApprovalRuleTemplateContentOutput, error) {
	return a.client.UpdateApprovalRuleTemplateContentWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdateApprovalRuleTemplateDescription(ctx context.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) (*codecommit.UpdateApprovalRuleTemplateDescriptionOutput, error) {
	return a.client.UpdateApprovalRuleTemplateDescriptionWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdateApprovalRuleTemplateName(ctx context.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) (*codecommit.UpdateApprovalRuleTemplateNameOutput, error) {
	return a.client.UpdateApprovalRuleTemplateNameWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdateComment(ctx context.Context, input *codecommit.UpdateCommentInput) (*codecommit.UpdateCommentOutput, error) {
	return a.client.UpdateCommentWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdateDefaultBranch(ctx context.Context, input *codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error) {
	return a.client.UpdateDefaultBranchWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdatePullRequestApprovalRuleContent(ctx context.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) (*codecommit.UpdatePullRequestApprovalRuleContentOutput, error) {
	return a.client.UpdatePullRequestApprovalRuleContentWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdatePullRequestApprovalState(ctx context.Context, input *codecommit.UpdatePullRequestApprovalStateInput) (*codecommit.UpdatePullRequestApprovalStateOutput, error) {
	return a.client.UpdatePullRequestApprovalStateWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdatePullRequestDescription(ctx context.Context, input *codecommit.UpdatePullRequestDescriptionInput) (*codecommit.UpdatePullRequestDescriptionOutput, error) {
	return a.client.UpdatePullRequestDescriptionWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdatePullRequestStatus(ctx context.Context, input *codecommit.UpdatePullRequestStatusInput) (*codecommit.UpdatePullRequestStatusOutput, error) {
	return a.client.UpdatePullRequestStatusWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdatePullRequestTitle(ctx context.Context, input *codecommit.UpdatePullRequestTitleInput) (*codecommit.UpdatePullRequestTitleOutput, error) {
	return a.client.UpdatePullRequestTitleWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdateRepositoryDescription(ctx context.Context, input *codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error) {
	return a.client.UpdateRepositoryDescriptionWithContext(ctx, input)
}

func (a *CodeCommitActivities) UpdateRepositoryName(ctx context.Context, input *codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error) {
	return a.client.UpdateRepositoryNameWithContext(ctx, input)
}
