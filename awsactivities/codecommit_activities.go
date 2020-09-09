
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codecommit/codecommitiface"
)

type CodeCommitActivities struct {
	client codecommitiface.CodeCommitAPI
}

func NewCodeCommitActivities(session *session.Session, config... *aws.Config) *CodeCommitActivities {
    client := codecommit.New(session, config...)
    return &CodeCommitActivities{client: client}
}

func (a *CodeCommitActivities) AssociateApprovalRuleTemplateWithRepository(input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) (*codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput, error) {
    return a.client.AssociateApprovalRuleTemplateWithRepository(input)
}

func (a *CodeCommitActivities) BatchAssociateApprovalRuleTemplateWithRepositories(input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
    return a.client.BatchAssociateApprovalRuleTemplateWithRepositories(input)
}

func (a *CodeCommitActivities) BatchDescribeMergeConflicts(input *codecommit.BatchDescribeMergeConflictsInput) (*codecommit.BatchDescribeMergeConflictsOutput, error) {
    return a.client.BatchDescribeMergeConflicts(input)
}

func (a *CodeCommitActivities) BatchDisassociateApprovalRuleTemplateFromRepositories(input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
    return a.client.BatchDisassociateApprovalRuleTemplateFromRepositories(input)
}

func (a *CodeCommitActivities) BatchGetCommits(input *codecommit.BatchGetCommitsInput) (*codecommit.BatchGetCommitsOutput, error) {
    return a.client.BatchGetCommits(input)
}

func (a *CodeCommitActivities) BatchGetRepositories(input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error) {
    return a.client.BatchGetRepositories(input)
}

func (a *CodeCommitActivities) CreateApprovalRuleTemplate(input *codecommit.CreateApprovalRuleTemplateInput) (*codecommit.CreateApprovalRuleTemplateOutput, error) {
    return a.client.CreateApprovalRuleTemplate(input)
}

func (a *CodeCommitActivities) CreateBranch(input *codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error) {
    return a.client.CreateBranch(input)
}

func (a *CodeCommitActivities) CreateCommit(input *codecommit.CreateCommitInput) (*codecommit.CreateCommitOutput, error) {
    return a.client.CreateCommit(input)
}

func (a *CodeCommitActivities) CreatePullRequestApprovalRule(input *codecommit.CreatePullRequestApprovalRuleInput) (*codecommit.CreatePullRequestApprovalRuleOutput, error) {
    return a.client.CreatePullRequestApprovalRule(input)
}

func (a *CodeCommitActivities) CreateRepository(input *codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error) {
    return a.client.CreateRepository(input)
}

func (a *CodeCommitActivities) CreateUnreferencedMergeCommit(input *codecommit.CreateUnreferencedMergeCommitInput) (*codecommit.CreateUnreferencedMergeCommitOutput, error) {
    return a.client.CreateUnreferencedMergeCommit(input)
}

func (a *CodeCommitActivities) DeleteApprovalRuleTemplate(input *codecommit.DeleteApprovalRuleTemplateInput) (*codecommit.DeleteApprovalRuleTemplateOutput, error) {
    return a.client.DeleteApprovalRuleTemplate(input)
}

func (a *CodeCommitActivities) DeleteBranch(input *codecommit.DeleteBranchInput) (*codecommit.DeleteBranchOutput, error) {
    return a.client.DeleteBranch(input)
}

func (a *CodeCommitActivities) DeleteCommentContent(input *codecommit.DeleteCommentContentInput) (*codecommit.DeleteCommentContentOutput, error) {
    return a.client.DeleteCommentContent(input)
}

func (a *CodeCommitActivities) DeleteFile(input *codecommit.DeleteFileInput) (*codecommit.DeleteFileOutput, error) {
    return a.client.DeleteFile(input)
}

func (a *CodeCommitActivities) DeletePullRequestApprovalRule(input *codecommit.DeletePullRequestApprovalRuleInput) (*codecommit.DeletePullRequestApprovalRuleOutput, error) {
    return a.client.DeletePullRequestApprovalRule(input)
}

func (a *CodeCommitActivities) DeleteRepository(input *codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error) {
    return a.client.DeleteRepository(input)
}

func (a *CodeCommitActivities) DescribeMergeConflicts(input *codecommit.DescribeMergeConflictsInput) (*codecommit.DescribeMergeConflictsOutput, error) {
    return a.client.DescribeMergeConflicts(input)
}

func (a *CodeCommitActivities) DescribePullRequestEvents(input *codecommit.DescribePullRequestEventsInput) (*codecommit.DescribePullRequestEventsOutput, error) {
    return a.client.DescribePullRequestEvents(input)
}

func (a *CodeCommitActivities) DisassociateApprovalRuleTemplateFromRepository(input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) (*codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput, error) {
    return a.client.DisassociateApprovalRuleTemplateFromRepository(input)
}

func (a *CodeCommitActivities) EvaluatePullRequestApprovalRules(input *codecommit.EvaluatePullRequestApprovalRulesInput) (*codecommit.EvaluatePullRequestApprovalRulesOutput, error) {
    return a.client.EvaluatePullRequestApprovalRules(input)
}

func (a *CodeCommitActivities) GetApprovalRuleTemplate(input *codecommit.GetApprovalRuleTemplateInput) (*codecommit.GetApprovalRuleTemplateOutput, error) {
    return a.client.GetApprovalRuleTemplate(input)
}

func (a *CodeCommitActivities) GetBlob(input *codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error) {
    return a.client.GetBlob(input)
}

func (a *CodeCommitActivities) GetBranch(input *codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error) {
    return a.client.GetBranch(input)
}

func (a *CodeCommitActivities) GetComment(input *codecommit.GetCommentInput) (*codecommit.GetCommentOutput, error) {
    return a.client.GetComment(input)
}

func (a *CodeCommitActivities) GetCommentReactions(input *codecommit.GetCommentReactionsInput) (*codecommit.GetCommentReactionsOutput, error) {
    return a.client.GetCommentReactions(input)
}

func (a *CodeCommitActivities) GetCommentsForComparedCommit(input *codecommit.GetCommentsForComparedCommitInput) (*codecommit.GetCommentsForComparedCommitOutput, error) {
    return a.client.GetCommentsForComparedCommit(input)
}

func (a *CodeCommitActivities) GetCommit(input *codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error) {
    return a.client.GetCommit(input)
}

func (a *CodeCommitActivities) GetDifferences(input *codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error) {
    return a.client.GetDifferences(input)
}

func (a *CodeCommitActivities) GetFile(input *codecommit.GetFileInput) (*codecommit.GetFileOutput, error) {
    return a.client.GetFile(input)
}

func (a *CodeCommitActivities) GetFolder(input *codecommit.GetFolderInput) (*codecommit.GetFolderOutput, error) {
    return a.client.GetFolder(input)
}

func (a *CodeCommitActivities) GetMergeCommit(input *codecommit.GetMergeCommitInput) (*codecommit.GetMergeCommitOutput, error) {
    return a.client.GetMergeCommit(input)
}

func (a *CodeCommitActivities) GetMergeConflicts(input *codecommit.GetMergeConflictsInput) (*codecommit.GetMergeConflictsOutput, error) {
    return a.client.GetMergeConflicts(input)
}

func (a *CodeCommitActivities) GetMergeOptions(input *codecommit.GetMergeOptionsInput) (*codecommit.GetMergeOptionsOutput, error) {
    return a.client.GetMergeOptions(input)
}

func (a *CodeCommitActivities) GetPullRequestApprovalStates(input *codecommit.GetPullRequestApprovalStatesInput) (*codecommit.GetPullRequestApprovalStatesOutput, error) {
    return a.client.GetPullRequestApprovalStates(input)
}

func (a *CodeCommitActivities) GetPullRequestOverrideState(input *codecommit.GetPullRequestOverrideStateInput) (*codecommit.GetPullRequestOverrideStateOutput, error) {
    return a.client.GetPullRequestOverrideState(input)
}

func (a *CodeCommitActivities) GetRepository(input *codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error) {
    return a.client.GetRepository(input)
}

func (a *CodeCommitActivities) GetRepositoryTriggers(input *codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error) {
    return a.client.GetRepositoryTriggers(input)
}

func (a *CodeCommitActivities) ListApprovalRuleTemplates(input *codecommit.ListApprovalRuleTemplatesInput) (*codecommit.ListApprovalRuleTemplatesOutput, error) {
    return a.client.ListApprovalRuleTemplates(input)
}

func (a *CodeCommitActivities) ListAssociatedApprovalRuleTemplatesForRepository(input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
    return a.client.ListAssociatedApprovalRuleTemplatesForRepository(input)
}

func (a *CodeCommitActivities) ListBranches(input *codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error) {
    return a.client.ListBranches(input)
}

func (a *CodeCommitActivities) ListPullRequests(input *codecommit.ListPullRequestsInput) (*codecommit.ListPullRequestsOutput, error) {
    return a.client.ListPullRequests(input)
}

func (a *CodeCommitActivities) ListRepositories(input *codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error) {
    return a.client.ListRepositories(input)
}

func (a *CodeCommitActivities) ListRepositoriesForApprovalRuleTemplate(input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error) {
    return a.client.ListRepositoriesForApprovalRuleTemplate(input)
}

func (a *CodeCommitActivities) ListTagsForResource(input *codecommit.ListTagsForResourceInput) (*codecommit.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CodeCommitActivities) MergeBranchesByFastForward(input *codecommit.MergeBranchesByFastForwardInput) (*codecommit.MergeBranchesByFastForwardOutput, error) {
    return a.client.MergeBranchesByFastForward(input)
}

func (a *CodeCommitActivities) MergeBranchesBySquash(input *codecommit.MergeBranchesBySquashInput) (*codecommit.MergeBranchesBySquashOutput, error) {
    return a.client.MergeBranchesBySquash(input)
}

func (a *CodeCommitActivities) MergeBranchesByThreeWay(input *codecommit.MergeBranchesByThreeWayInput) (*codecommit.MergeBranchesByThreeWayOutput, error) {
    return a.client.MergeBranchesByThreeWay(input)
}

func (a *CodeCommitActivities) MergePullRequestByFastForward(input *codecommit.MergePullRequestByFastForwardInput) (*codecommit.MergePullRequestByFastForwardOutput, error) {
    return a.client.MergePullRequestByFastForward(input)
}

func (a *CodeCommitActivities) MergePullRequestBySquash(input *codecommit.MergePullRequestBySquashInput) (*codecommit.MergePullRequestBySquashOutput, error) {
    return a.client.MergePullRequestBySquash(input)
}

func (a *CodeCommitActivities) MergePullRequestByThreeWay(input *codecommit.MergePullRequestByThreeWayInput) (*codecommit.MergePullRequestByThreeWayOutput, error) {
    return a.client.MergePullRequestByThreeWay(input)
}

func (a *CodeCommitActivities) OverridePullRequestApprovalRules(input *codecommit.OverridePullRequestApprovalRulesInput) (*codecommit.OverridePullRequestApprovalRulesOutput, error) {
    return a.client.OverridePullRequestApprovalRules(input)
}

func (a *CodeCommitActivities) PostCommentForComparedCommit(input *codecommit.PostCommentForComparedCommitInput) (*codecommit.PostCommentForComparedCommitOutput, error) {
    return a.client.PostCommentForComparedCommit(input)
}

func (a *CodeCommitActivities) PostCommentReply(input *codecommit.PostCommentReplyInput) (*codecommit.PostCommentReplyOutput, error) {
    return a.client.PostCommentReply(input)
}

func (a *CodeCommitActivities) PutCommentReaction(input *codecommit.PutCommentReactionInput) (*codecommit.PutCommentReactionOutput, error) {
    return a.client.PutCommentReaction(input)
}

func (a *CodeCommitActivities) PutFile(input *codecommit.PutFileInput) (*codecommit.PutFileOutput, error) {
    return a.client.PutFile(input)
}

func (a *CodeCommitActivities) PutRepositoryTriggers(input *codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error) {
    return a.client.PutRepositoryTriggers(input)
}

func (a *CodeCommitActivities) TagResource(input *codecommit.TagResourceInput) (*codecommit.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CodeCommitActivities) TestRepositoryTriggers(input *codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error) {
    return a.client.TestRepositoryTriggers(input)
}

func (a *CodeCommitActivities) UntagResource(input *codecommit.UntagResourceInput) (*codecommit.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CodeCommitActivities) UpdateApprovalRuleTemplateContent(input *codecommit.UpdateApprovalRuleTemplateContentInput) (*codecommit.UpdateApprovalRuleTemplateContentOutput, error) {
    return a.client.UpdateApprovalRuleTemplateContent(input)
}

func (a *CodeCommitActivities) UpdateApprovalRuleTemplateDescription(input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) (*codecommit.UpdateApprovalRuleTemplateDescriptionOutput, error) {
    return a.client.UpdateApprovalRuleTemplateDescription(input)
}

func (a *CodeCommitActivities) UpdateApprovalRuleTemplateName(input *codecommit.UpdateApprovalRuleTemplateNameInput) (*codecommit.UpdateApprovalRuleTemplateNameOutput, error) {
    return a.client.UpdateApprovalRuleTemplateName(input)
}

func (a *CodeCommitActivities) UpdateComment(input *codecommit.UpdateCommentInput) (*codecommit.UpdateCommentOutput, error) {
    return a.client.UpdateComment(input)
}

func (a *CodeCommitActivities) UpdateDefaultBranch(input *codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error) {
    return a.client.UpdateDefaultBranch(input)
}

func (a *CodeCommitActivities) UpdatePullRequestApprovalRuleContent(input *codecommit.UpdatePullRequestApprovalRuleContentInput) (*codecommit.UpdatePullRequestApprovalRuleContentOutput, error) {
    return a.client.UpdatePullRequestApprovalRuleContent(input)
}

func (a *CodeCommitActivities) UpdatePullRequestApprovalState(input *codecommit.UpdatePullRequestApprovalStateInput) (*codecommit.UpdatePullRequestApprovalStateOutput, error) {
    return a.client.UpdatePullRequestApprovalState(input)
}

func (a *CodeCommitActivities) UpdatePullRequestDescription(input *codecommit.UpdatePullRequestDescriptionInput) (*codecommit.UpdatePullRequestDescriptionOutput, error) {
    return a.client.UpdatePullRequestDescription(input)
}

func (a *CodeCommitActivities) UpdatePullRequestStatus(input *codecommit.UpdatePullRequestStatusInput) (*codecommit.UpdatePullRequestStatusOutput, error) {
    return a.client.UpdatePullRequestStatus(input)
}

func (a *CodeCommitActivities) UpdatePullRequestTitle(input *codecommit.UpdatePullRequestTitleInput) (*codecommit.UpdatePullRequestTitleOutput, error) {
    return a.client.UpdatePullRequestTitle(input)
}

func (a *CodeCommitActivities) UpdateRepositoryDescription(input *codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error) {
    return a.client.UpdateRepositoryDescription(input)
}

func (a *CodeCommitActivities) UpdateRepositoryName(input *codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error) {
    return a.client.UpdateRepositoryName(input)
}
