package awsclients

import (
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CloudDirectoryClient interface {
    AddFacetToObject(ctx workflow.Context, input *clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error)
    AddFacetToObjectAsync(ctx workflow.Context, input *clouddirectory.AddFacetToObjectInput) *ClouddirectoryAddFacetToObjectResult

    ApplySchema(ctx workflow.Context, input *clouddirectory.ApplySchemaInput) (*clouddirectory.ApplySchemaOutput, error)
    ApplySchemaAsync(ctx workflow.Context, input *clouddirectory.ApplySchemaInput) *ClouddirectoryApplySchemaResult

    AttachObject(ctx workflow.Context, input *clouddirectory.AttachObjectInput) (*clouddirectory.AttachObjectOutput, error)
    AttachObjectAsync(ctx workflow.Context, input *clouddirectory.AttachObjectInput) *ClouddirectoryAttachObjectResult

    AttachPolicy(ctx workflow.Context, input *clouddirectory.AttachPolicyInput) (*clouddirectory.AttachPolicyOutput, error)
    AttachPolicyAsync(ctx workflow.Context, input *clouddirectory.AttachPolicyInput) *ClouddirectoryAttachPolicyResult

    AttachToIndex(ctx workflow.Context, input *clouddirectory.AttachToIndexInput) (*clouddirectory.AttachToIndexOutput, error)
    AttachToIndexAsync(ctx workflow.Context, input *clouddirectory.AttachToIndexInput) *ClouddirectoryAttachToIndexResult

    AttachTypedLink(ctx workflow.Context, input *clouddirectory.AttachTypedLinkInput) (*clouddirectory.AttachTypedLinkOutput, error)
    AttachTypedLinkAsync(ctx workflow.Context, input *clouddirectory.AttachTypedLinkInput) *ClouddirectoryAttachTypedLinkResult

    BatchRead(ctx workflow.Context, input *clouddirectory.BatchReadInput) (*clouddirectory.BatchReadOutput, error)
    BatchReadAsync(ctx workflow.Context, input *clouddirectory.BatchReadInput) *ClouddirectoryBatchReadResult

    BatchWrite(ctx workflow.Context, input *clouddirectory.BatchWriteInput) (*clouddirectory.BatchWriteOutput, error)
    BatchWriteAsync(ctx workflow.Context, input *clouddirectory.BatchWriteInput) *ClouddirectoryBatchWriteResult

    CreateDirectory(ctx workflow.Context, input *clouddirectory.CreateDirectoryInput) (*clouddirectory.CreateDirectoryOutput, error)
    CreateDirectoryAsync(ctx workflow.Context, input *clouddirectory.CreateDirectoryInput) *ClouddirectoryCreateDirectoryResult

    CreateFacet(ctx workflow.Context, input *clouddirectory.CreateFacetInput) (*clouddirectory.CreateFacetOutput, error)
    CreateFacetAsync(ctx workflow.Context, input *clouddirectory.CreateFacetInput) *ClouddirectoryCreateFacetResult

    CreateIndex(ctx workflow.Context, input *clouddirectory.CreateIndexInput) (*clouddirectory.CreateIndexOutput, error)
    CreateIndexAsync(ctx workflow.Context, input *clouddirectory.CreateIndexInput) *ClouddirectoryCreateIndexResult

    CreateObject(ctx workflow.Context, input *clouddirectory.CreateObjectInput) (*clouddirectory.CreateObjectOutput, error)
    CreateObjectAsync(ctx workflow.Context, input *clouddirectory.CreateObjectInput) *ClouddirectoryCreateObjectResult

    CreateSchema(ctx workflow.Context, input *clouddirectory.CreateSchemaInput) (*clouddirectory.CreateSchemaOutput, error)
    CreateSchemaAsync(ctx workflow.Context, input *clouddirectory.CreateSchemaInput) *ClouddirectoryCreateSchemaResult

    CreateTypedLinkFacet(ctx workflow.Context, input *clouddirectory.CreateTypedLinkFacetInput) (*clouddirectory.CreateTypedLinkFacetOutput, error)
    CreateTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.CreateTypedLinkFacetInput) *ClouddirectoryCreateTypedLinkFacetResult

    DeleteDirectory(ctx workflow.Context, input *clouddirectory.DeleteDirectoryInput) (*clouddirectory.DeleteDirectoryOutput, error)
    DeleteDirectoryAsync(ctx workflow.Context, input *clouddirectory.DeleteDirectoryInput) *ClouddirectoryDeleteDirectoryResult

    DeleteFacet(ctx workflow.Context, input *clouddirectory.DeleteFacetInput) (*clouddirectory.DeleteFacetOutput, error)
    DeleteFacetAsync(ctx workflow.Context, input *clouddirectory.DeleteFacetInput) *ClouddirectoryDeleteFacetResult

    DeleteObject(ctx workflow.Context, input *clouddirectory.DeleteObjectInput) (*clouddirectory.DeleteObjectOutput, error)
    DeleteObjectAsync(ctx workflow.Context, input *clouddirectory.DeleteObjectInput) *ClouddirectoryDeleteObjectResult

    DeleteSchema(ctx workflow.Context, input *clouddirectory.DeleteSchemaInput) (*clouddirectory.DeleteSchemaOutput, error)
    DeleteSchemaAsync(ctx workflow.Context, input *clouddirectory.DeleteSchemaInput) *ClouddirectoryDeleteSchemaResult

    DeleteTypedLinkFacet(ctx workflow.Context, input *clouddirectory.DeleteTypedLinkFacetInput) (*clouddirectory.DeleteTypedLinkFacetOutput, error)
    DeleteTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.DeleteTypedLinkFacetInput) *ClouddirectoryDeleteTypedLinkFacetResult

    DetachFromIndex(ctx workflow.Context, input *clouddirectory.DetachFromIndexInput) (*clouddirectory.DetachFromIndexOutput, error)
    DetachFromIndexAsync(ctx workflow.Context, input *clouddirectory.DetachFromIndexInput) *ClouddirectoryDetachFromIndexResult

    DetachObject(ctx workflow.Context, input *clouddirectory.DetachObjectInput) (*clouddirectory.DetachObjectOutput, error)
    DetachObjectAsync(ctx workflow.Context, input *clouddirectory.DetachObjectInput) *ClouddirectoryDetachObjectResult

    DetachPolicy(ctx workflow.Context, input *clouddirectory.DetachPolicyInput) (*clouddirectory.DetachPolicyOutput, error)
    DetachPolicyAsync(ctx workflow.Context, input *clouddirectory.DetachPolicyInput) *ClouddirectoryDetachPolicyResult

    DetachTypedLink(ctx workflow.Context, input *clouddirectory.DetachTypedLinkInput) (*clouddirectory.DetachTypedLinkOutput, error)
    DetachTypedLinkAsync(ctx workflow.Context, input *clouddirectory.DetachTypedLinkInput) *ClouddirectoryDetachTypedLinkResult

    DisableDirectory(ctx workflow.Context, input *clouddirectory.DisableDirectoryInput) (*clouddirectory.DisableDirectoryOutput, error)
    DisableDirectoryAsync(ctx workflow.Context, input *clouddirectory.DisableDirectoryInput) *ClouddirectoryDisableDirectoryResult

    EnableDirectory(ctx workflow.Context, input *clouddirectory.EnableDirectoryInput) (*clouddirectory.EnableDirectoryOutput, error)
    EnableDirectoryAsync(ctx workflow.Context, input *clouddirectory.EnableDirectoryInput) *ClouddirectoryEnableDirectoryResult

    GetAppliedSchemaVersion(ctx workflow.Context, input *clouddirectory.GetAppliedSchemaVersionInput) (*clouddirectory.GetAppliedSchemaVersionOutput, error)
    GetAppliedSchemaVersionAsync(ctx workflow.Context, input *clouddirectory.GetAppliedSchemaVersionInput) *ClouddirectoryGetAppliedSchemaVersionResult

    GetDirectory(ctx workflow.Context, input *clouddirectory.GetDirectoryInput) (*clouddirectory.GetDirectoryOutput, error)
    GetDirectoryAsync(ctx workflow.Context, input *clouddirectory.GetDirectoryInput) *ClouddirectoryGetDirectoryResult

    GetFacet(ctx workflow.Context, input *clouddirectory.GetFacetInput) (*clouddirectory.GetFacetOutput, error)
    GetFacetAsync(ctx workflow.Context, input *clouddirectory.GetFacetInput) *ClouddirectoryGetFacetResult

    GetLinkAttributes(ctx workflow.Context, input *clouddirectory.GetLinkAttributesInput) (*clouddirectory.GetLinkAttributesOutput, error)
    GetLinkAttributesAsync(ctx workflow.Context, input *clouddirectory.GetLinkAttributesInput) *ClouddirectoryGetLinkAttributesResult

    GetObjectAttributes(ctx workflow.Context, input *clouddirectory.GetObjectAttributesInput) (*clouddirectory.GetObjectAttributesOutput, error)
    GetObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.GetObjectAttributesInput) *ClouddirectoryGetObjectAttributesResult

    GetObjectInformation(ctx workflow.Context, input *clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error)
    GetObjectInformationAsync(ctx workflow.Context, input *clouddirectory.GetObjectInformationInput) *ClouddirectoryGetObjectInformationResult

    GetSchemaAsJson(ctx workflow.Context, input *clouddirectory.GetSchemaAsJsonInput) (*clouddirectory.GetSchemaAsJsonOutput, error)
    GetSchemaAsJsonAsync(ctx workflow.Context, input *clouddirectory.GetSchemaAsJsonInput) *ClouddirectoryGetSchemaAsJsonResult

    GetTypedLinkFacetInformation(ctx workflow.Context, input *clouddirectory.GetTypedLinkFacetInformationInput) (*clouddirectory.GetTypedLinkFacetInformationOutput, error)
    GetTypedLinkFacetInformationAsync(ctx workflow.Context, input *clouddirectory.GetTypedLinkFacetInformationInput) *ClouddirectoryGetTypedLinkFacetInformationResult

    ListAppliedSchemaArns(ctx workflow.Context, input *clouddirectory.ListAppliedSchemaArnsInput) (*clouddirectory.ListAppliedSchemaArnsOutput, error)
    ListAppliedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListAppliedSchemaArnsInput) *ClouddirectoryListAppliedSchemaArnsResult

    ListAttachedIndices(ctx workflow.Context, input *clouddirectory.ListAttachedIndicesInput) (*clouddirectory.ListAttachedIndicesOutput, error)
    ListAttachedIndicesAsync(ctx workflow.Context, input *clouddirectory.ListAttachedIndicesInput) *ClouddirectoryListAttachedIndicesResult

    ListDevelopmentSchemaArns(ctx workflow.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error)
    ListDevelopmentSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput) *ClouddirectoryListDevelopmentSchemaArnsResult

    ListDirectories(ctx workflow.Context, input *clouddirectory.ListDirectoriesInput) (*clouddirectory.ListDirectoriesOutput, error)
    ListDirectoriesAsync(ctx workflow.Context, input *clouddirectory.ListDirectoriesInput) *ClouddirectoryListDirectoriesResult

    ListFacetAttributes(ctx workflow.Context, input *clouddirectory.ListFacetAttributesInput) (*clouddirectory.ListFacetAttributesOutput, error)
    ListFacetAttributesAsync(ctx workflow.Context, input *clouddirectory.ListFacetAttributesInput) *ClouddirectoryListFacetAttributesResult

    ListFacetNames(ctx workflow.Context, input *clouddirectory.ListFacetNamesInput) (*clouddirectory.ListFacetNamesOutput, error)
    ListFacetNamesAsync(ctx workflow.Context, input *clouddirectory.ListFacetNamesInput) *ClouddirectoryListFacetNamesResult

    ListIncomingTypedLinks(ctx workflow.Context, input *clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error)
    ListIncomingTypedLinksAsync(ctx workflow.Context, input *clouddirectory.ListIncomingTypedLinksInput) *ClouddirectoryListIncomingTypedLinksResult

    ListIndex(ctx workflow.Context, input *clouddirectory.ListIndexInput) (*clouddirectory.ListIndexOutput, error)
    ListIndexAsync(ctx workflow.Context, input *clouddirectory.ListIndexInput) *ClouddirectoryListIndexResult

    ListManagedSchemaArns(ctx workflow.Context, input *clouddirectory.ListManagedSchemaArnsInput) (*clouddirectory.ListManagedSchemaArnsOutput, error)
    ListManagedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListManagedSchemaArnsInput) *ClouddirectoryListManagedSchemaArnsResult

    ListObjectAttributes(ctx workflow.Context, input *clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error)
    ListObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.ListObjectAttributesInput) *ClouddirectoryListObjectAttributesResult

    ListObjectChildren(ctx workflow.Context, input *clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error)
    ListObjectChildrenAsync(ctx workflow.Context, input *clouddirectory.ListObjectChildrenInput) *ClouddirectoryListObjectChildrenResult

    ListObjectParentPaths(ctx workflow.Context, input *clouddirectory.ListObjectParentPathsInput) (*clouddirectory.ListObjectParentPathsOutput, error)
    ListObjectParentPathsAsync(ctx workflow.Context, input *clouddirectory.ListObjectParentPathsInput) *ClouddirectoryListObjectParentPathsResult

    ListObjectParents(ctx workflow.Context, input *clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error)
    ListObjectParentsAsync(ctx workflow.Context, input *clouddirectory.ListObjectParentsInput) *ClouddirectoryListObjectParentsResult

    ListObjectPolicies(ctx workflow.Context, input *clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error)
    ListObjectPoliciesAsync(ctx workflow.Context, input *clouddirectory.ListObjectPoliciesInput) *ClouddirectoryListObjectPoliciesResult

    ListOutgoingTypedLinks(ctx workflow.Context, input *clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error)
    ListOutgoingTypedLinksAsync(ctx workflow.Context, input *clouddirectory.ListOutgoingTypedLinksInput) *ClouddirectoryListOutgoingTypedLinksResult

    ListPolicyAttachments(ctx workflow.Context, input *clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error)
    ListPolicyAttachmentsAsync(ctx workflow.Context, input *clouddirectory.ListPolicyAttachmentsInput) *ClouddirectoryListPolicyAttachmentsResult

    ListPublishedSchemaArns(ctx workflow.Context, input *clouddirectory.ListPublishedSchemaArnsInput) (*clouddirectory.ListPublishedSchemaArnsOutput, error)
    ListPublishedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListPublishedSchemaArnsInput) *ClouddirectoryListPublishedSchemaArnsResult

    ListTagsForResource(ctx workflow.Context, input *clouddirectory.ListTagsForResourceInput) (*clouddirectory.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *clouddirectory.ListTagsForResourceInput) *ClouddirectoryListTagsForResourceResult

    ListTypedLinkFacetAttributes(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error)
    ListTypedLinkFacetAttributesAsync(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput) *ClouddirectoryListTypedLinkFacetAttributesResult

    ListTypedLinkFacetNames(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetNamesInput) (*clouddirectory.ListTypedLinkFacetNamesOutput, error)
    ListTypedLinkFacetNamesAsync(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetNamesInput) *ClouddirectoryListTypedLinkFacetNamesResult

    LookupPolicy(ctx workflow.Context, input *clouddirectory.LookupPolicyInput) (*clouddirectory.LookupPolicyOutput, error)
    LookupPolicyAsync(ctx workflow.Context, input *clouddirectory.LookupPolicyInput) *ClouddirectoryLookupPolicyResult

    PublishSchema(ctx workflow.Context, input *clouddirectory.PublishSchemaInput) (*clouddirectory.PublishSchemaOutput, error)
    PublishSchemaAsync(ctx workflow.Context, input *clouddirectory.PublishSchemaInput) *ClouddirectoryPublishSchemaResult

    PutSchemaFromJson(ctx workflow.Context, input *clouddirectory.PutSchemaFromJsonInput) (*clouddirectory.PutSchemaFromJsonOutput, error)
    PutSchemaFromJsonAsync(ctx workflow.Context, input *clouddirectory.PutSchemaFromJsonInput) *ClouddirectoryPutSchemaFromJsonResult

    RemoveFacetFromObject(ctx workflow.Context, input *clouddirectory.RemoveFacetFromObjectInput) (*clouddirectory.RemoveFacetFromObjectOutput, error)
    RemoveFacetFromObjectAsync(ctx workflow.Context, input *clouddirectory.RemoveFacetFromObjectInput) *ClouddirectoryRemoveFacetFromObjectResult

    TagResource(ctx workflow.Context, input *clouddirectory.TagResourceInput) (*clouddirectory.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *clouddirectory.TagResourceInput) *ClouddirectoryTagResourceResult

    UntagResource(ctx workflow.Context, input *clouddirectory.UntagResourceInput) (*clouddirectory.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *clouddirectory.UntagResourceInput) *ClouddirectoryUntagResourceResult

    UpdateFacet(ctx workflow.Context, input *clouddirectory.UpdateFacetInput) (*clouddirectory.UpdateFacetOutput, error)
    UpdateFacetAsync(ctx workflow.Context, input *clouddirectory.UpdateFacetInput) *ClouddirectoryUpdateFacetResult

    UpdateLinkAttributes(ctx workflow.Context, input *clouddirectory.UpdateLinkAttributesInput) (*clouddirectory.UpdateLinkAttributesOutput, error)
    UpdateLinkAttributesAsync(ctx workflow.Context, input *clouddirectory.UpdateLinkAttributesInput) *ClouddirectoryUpdateLinkAttributesResult

    UpdateObjectAttributes(ctx workflow.Context, input *clouddirectory.UpdateObjectAttributesInput) (*clouddirectory.UpdateObjectAttributesOutput, error)
    UpdateObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.UpdateObjectAttributesInput) *ClouddirectoryUpdateObjectAttributesResult

    UpdateSchema(ctx workflow.Context, input *clouddirectory.UpdateSchemaInput) (*clouddirectory.UpdateSchemaOutput, error)
    UpdateSchemaAsync(ctx workflow.Context, input *clouddirectory.UpdateSchemaInput) *ClouddirectoryUpdateSchemaResult

    UpdateTypedLinkFacet(ctx workflow.Context, input *clouddirectory.UpdateTypedLinkFacetInput) (*clouddirectory.UpdateTypedLinkFacetOutput, error)
    UpdateTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.UpdateTypedLinkFacetInput) *ClouddirectoryUpdateTypedLinkFacetResult

    UpgradeAppliedSchema(ctx workflow.Context, input *clouddirectory.UpgradeAppliedSchemaInput) (*clouddirectory.UpgradeAppliedSchemaOutput, error)
    UpgradeAppliedSchemaAsync(ctx workflow.Context, input *clouddirectory.UpgradeAppliedSchemaInput) *ClouddirectoryUpgradeAppliedSchemaResult

    UpgradePublishedSchema(ctx workflow.Context, input *clouddirectory.UpgradePublishedSchemaInput) (*clouddirectory.UpgradePublishedSchemaOutput, error)
    UpgradePublishedSchemaAsync(ctx workflow.Context, input *clouddirectory.UpgradePublishedSchemaInput) *ClouddirectoryUpgradePublishedSchemaResult
}

type ClouddirectoryAddFacetToObjectResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryAddFacetToObjectResult) Get(ctx workflow.Context) (*clouddirectory.AddFacetToObjectOutput, error) {
    var output clouddirectory.AddFacetToObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryApplySchemaResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryApplySchemaResult) Get(ctx workflow.Context) (*clouddirectory.ApplySchemaOutput, error) {
    var output clouddirectory.ApplySchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryAttachObjectResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryAttachObjectResult) Get(ctx workflow.Context) (*clouddirectory.AttachObjectOutput, error) {
    var output clouddirectory.AttachObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryAttachPolicyResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryAttachPolicyResult) Get(ctx workflow.Context) (*clouddirectory.AttachPolicyOutput, error) {
    var output clouddirectory.AttachPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryAttachToIndexResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryAttachToIndexResult) Get(ctx workflow.Context) (*clouddirectory.AttachToIndexOutput, error) {
    var output clouddirectory.AttachToIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryAttachTypedLinkResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryAttachTypedLinkResult) Get(ctx workflow.Context) (*clouddirectory.AttachTypedLinkOutput, error) {
    var output clouddirectory.AttachTypedLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryBatchReadResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryBatchReadResult) Get(ctx workflow.Context) (*clouddirectory.BatchReadOutput, error) {
    var output clouddirectory.BatchReadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryBatchWriteResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryBatchWriteResult) Get(ctx workflow.Context) (*clouddirectory.BatchWriteOutput, error) {
    var output clouddirectory.BatchWriteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryCreateDirectoryResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryCreateDirectoryResult) Get(ctx workflow.Context) (*clouddirectory.CreateDirectoryOutput, error) {
    var output clouddirectory.CreateDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryCreateFacetResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryCreateFacetResult) Get(ctx workflow.Context) (*clouddirectory.CreateFacetOutput, error) {
    var output clouddirectory.CreateFacetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryCreateIndexResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryCreateIndexResult) Get(ctx workflow.Context) (*clouddirectory.CreateIndexOutput, error) {
    var output clouddirectory.CreateIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryCreateObjectResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryCreateObjectResult) Get(ctx workflow.Context) (*clouddirectory.CreateObjectOutput, error) {
    var output clouddirectory.CreateObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryCreateSchemaResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryCreateSchemaResult) Get(ctx workflow.Context) (*clouddirectory.CreateSchemaOutput, error) {
    var output clouddirectory.CreateSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryCreateTypedLinkFacetResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryCreateTypedLinkFacetResult) Get(ctx workflow.Context) (*clouddirectory.CreateTypedLinkFacetOutput, error) {
    var output clouddirectory.CreateTypedLinkFacetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDeleteDirectoryResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDeleteDirectoryResult) Get(ctx workflow.Context) (*clouddirectory.DeleteDirectoryOutput, error) {
    var output clouddirectory.DeleteDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDeleteFacetResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDeleteFacetResult) Get(ctx workflow.Context) (*clouddirectory.DeleteFacetOutput, error) {
    var output clouddirectory.DeleteFacetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDeleteObjectResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDeleteObjectResult) Get(ctx workflow.Context) (*clouddirectory.DeleteObjectOutput, error) {
    var output clouddirectory.DeleteObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDeleteSchemaResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDeleteSchemaResult) Get(ctx workflow.Context) (*clouddirectory.DeleteSchemaOutput, error) {
    var output clouddirectory.DeleteSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDeleteTypedLinkFacetResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDeleteTypedLinkFacetResult) Get(ctx workflow.Context) (*clouddirectory.DeleteTypedLinkFacetOutput, error) {
    var output clouddirectory.DeleteTypedLinkFacetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDetachFromIndexResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDetachFromIndexResult) Get(ctx workflow.Context) (*clouddirectory.DetachFromIndexOutput, error) {
    var output clouddirectory.DetachFromIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDetachObjectResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDetachObjectResult) Get(ctx workflow.Context) (*clouddirectory.DetachObjectOutput, error) {
    var output clouddirectory.DetachObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDetachPolicyResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDetachPolicyResult) Get(ctx workflow.Context) (*clouddirectory.DetachPolicyOutput, error) {
    var output clouddirectory.DetachPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDetachTypedLinkResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDetachTypedLinkResult) Get(ctx workflow.Context) (*clouddirectory.DetachTypedLinkOutput, error) {
    var output clouddirectory.DetachTypedLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryDisableDirectoryResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryDisableDirectoryResult) Get(ctx workflow.Context) (*clouddirectory.DisableDirectoryOutput, error) {
    var output clouddirectory.DisableDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryEnableDirectoryResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryEnableDirectoryResult) Get(ctx workflow.Context) (*clouddirectory.EnableDirectoryOutput, error) {
    var output clouddirectory.EnableDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetAppliedSchemaVersionResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetAppliedSchemaVersionResult) Get(ctx workflow.Context) (*clouddirectory.GetAppliedSchemaVersionOutput, error) {
    var output clouddirectory.GetAppliedSchemaVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetDirectoryResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetDirectoryResult) Get(ctx workflow.Context) (*clouddirectory.GetDirectoryOutput, error) {
    var output clouddirectory.GetDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetFacetResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetFacetResult) Get(ctx workflow.Context) (*clouddirectory.GetFacetOutput, error) {
    var output clouddirectory.GetFacetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetLinkAttributesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetLinkAttributesResult) Get(ctx workflow.Context) (*clouddirectory.GetLinkAttributesOutput, error) {
    var output clouddirectory.GetLinkAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetObjectAttributesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetObjectAttributesResult) Get(ctx workflow.Context) (*clouddirectory.GetObjectAttributesOutput, error) {
    var output clouddirectory.GetObjectAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetObjectInformationResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetObjectInformationResult) Get(ctx workflow.Context) (*clouddirectory.GetObjectInformationOutput, error) {
    var output clouddirectory.GetObjectInformationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetSchemaAsJsonResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetSchemaAsJsonResult) Get(ctx workflow.Context) (*clouddirectory.GetSchemaAsJsonOutput, error) {
    var output clouddirectory.GetSchemaAsJsonOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryGetTypedLinkFacetInformationResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryGetTypedLinkFacetInformationResult) Get(ctx workflow.Context) (*clouddirectory.GetTypedLinkFacetInformationOutput, error) {
    var output clouddirectory.GetTypedLinkFacetInformationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListAppliedSchemaArnsResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListAppliedSchemaArnsResult) Get(ctx workflow.Context) (*clouddirectory.ListAppliedSchemaArnsOutput, error) {
    var output clouddirectory.ListAppliedSchemaArnsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListAttachedIndicesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListAttachedIndicesResult) Get(ctx workflow.Context) (*clouddirectory.ListAttachedIndicesOutput, error) {
    var output clouddirectory.ListAttachedIndicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListDevelopmentSchemaArnsResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListDevelopmentSchemaArnsResult) Get(ctx workflow.Context) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error) {
    var output clouddirectory.ListDevelopmentSchemaArnsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListDirectoriesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListDirectoriesResult) Get(ctx workflow.Context) (*clouddirectory.ListDirectoriesOutput, error) {
    var output clouddirectory.ListDirectoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListFacetAttributesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListFacetAttributesResult) Get(ctx workflow.Context) (*clouddirectory.ListFacetAttributesOutput, error) {
    var output clouddirectory.ListFacetAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListFacetNamesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListFacetNamesResult) Get(ctx workflow.Context) (*clouddirectory.ListFacetNamesOutput, error) {
    var output clouddirectory.ListFacetNamesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListIncomingTypedLinksResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListIncomingTypedLinksResult) Get(ctx workflow.Context) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
    var output clouddirectory.ListIncomingTypedLinksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListIndexResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListIndexResult) Get(ctx workflow.Context) (*clouddirectory.ListIndexOutput, error) {
    var output clouddirectory.ListIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListManagedSchemaArnsResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListManagedSchemaArnsResult) Get(ctx workflow.Context) (*clouddirectory.ListManagedSchemaArnsOutput, error) {
    var output clouddirectory.ListManagedSchemaArnsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListObjectAttributesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListObjectAttributesResult) Get(ctx workflow.Context) (*clouddirectory.ListObjectAttributesOutput, error) {
    var output clouddirectory.ListObjectAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListObjectChildrenResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListObjectChildrenResult) Get(ctx workflow.Context) (*clouddirectory.ListObjectChildrenOutput, error) {
    var output clouddirectory.ListObjectChildrenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListObjectParentPathsResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListObjectParentPathsResult) Get(ctx workflow.Context) (*clouddirectory.ListObjectParentPathsOutput, error) {
    var output clouddirectory.ListObjectParentPathsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListObjectParentsResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListObjectParentsResult) Get(ctx workflow.Context) (*clouddirectory.ListObjectParentsOutput, error) {
    var output clouddirectory.ListObjectParentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListObjectPoliciesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListObjectPoliciesResult) Get(ctx workflow.Context) (*clouddirectory.ListObjectPoliciesOutput, error) {
    var output clouddirectory.ListObjectPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListOutgoingTypedLinksResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListOutgoingTypedLinksResult) Get(ctx workflow.Context) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
    var output clouddirectory.ListOutgoingTypedLinksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListPolicyAttachmentsResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListPolicyAttachmentsResult) Get(ctx workflow.Context) (*clouddirectory.ListPolicyAttachmentsOutput, error) {
    var output clouddirectory.ListPolicyAttachmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListPublishedSchemaArnsResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListPublishedSchemaArnsResult) Get(ctx workflow.Context) (*clouddirectory.ListPublishedSchemaArnsOutput, error) {
    var output clouddirectory.ListPublishedSchemaArnsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListTagsForResourceResult) Get(ctx workflow.Context) (*clouddirectory.ListTagsForResourceOutput, error) {
    var output clouddirectory.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListTypedLinkFacetAttributesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListTypedLinkFacetAttributesResult) Get(ctx workflow.Context) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error) {
    var output clouddirectory.ListTypedLinkFacetAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryListTypedLinkFacetNamesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryListTypedLinkFacetNamesResult) Get(ctx workflow.Context) (*clouddirectory.ListTypedLinkFacetNamesOutput, error) {
    var output clouddirectory.ListTypedLinkFacetNamesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryLookupPolicyResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryLookupPolicyResult) Get(ctx workflow.Context) (*clouddirectory.LookupPolicyOutput, error) {
    var output clouddirectory.LookupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryPublishSchemaResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryPublishSchemaResult) Get(ctx workflow.Context) (*clouddirectory.PublishSchemaOutput, error) {
    var output clouddirectory.PublishSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryPutSchemaFromJsonResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryPutSchemaFromJsonResult) Get(ctx workflow.Context) (*clouddirectory.PutSchemaFromJsonOutput, error) {
    var output clouddirectory.PutSchemaFromJsonOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryRemoveFacetFromObjectResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryRemoveFacetFromObjectResult) Get(ctx workflow.Context) (*clouddirectory.RemoveFacetFromObjectOutput, error) {
    var output clouddirectory.RemoveFacetFromObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryTagResourceResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryTagResourceResult) Get(ctx workflow.Context) (*clouddirectory.TagResourceOutput, error) {
    var output clouddirectory.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUntagResourceResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUntagResourceResult) Get(ctx workflow.Context) (*clouddirectory.UntagResourceOutput, error) {
    var output clouddirectory.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUpdateFacetResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUpdateFacetResult) Get(ctx workflow.Context) (*clouddirectory.UpdateFacetOutput, error) {
    var output clouddirectory.UpdateFacetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUpdateLinkAttributesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUpdateLinkAttributesResult) Get(ctx workflow.Context) (*clouddirectory.UpdateLinkAttributesOutput, error) {
    var output clouddirectory.UpdateLinkAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUpdateObjectAttributesResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUpdateObjectAttributesResult) Get(ctx workflow.Context) (*clouddirectory.UpdateObjectAttributesOutput, error) {
    var output clouddirectory.UpdateObjectAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUpdateSchemaResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUpdateSchemaResult) Get(ctx workflow.Context) (*clouddirectory.UpdateSchemaOutput, error) {
    var output clouddirectory.UpdateSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUpdateTypedLinkFacetResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUpdateTypedLinkFacetResult) Get(ctx workflow.Context) (*clouddirectory.UpdateTypedLinkFacetOutput, error) {
    var output clouddirectory.UpdateTypedLinkFacetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUpgradeAppliedSchemaResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUpgradeAppliedSchemaResult) Get(ctx workflow.Context) (*clouddirectory.UpgradeAppliedSchemaOutput, error) {
    var output clouddirectory.UpgradeAppliedSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ClouddirectoryUpgradePublishedSchemaResult struct {
	Result workflow.Future
}

func (r *ClouddirectoryUpgradePublishedSchemaResult) Get(ctx workflow.Context) (*clouddirectory.UpgradePublishedSchemaOutput, error) {
    var output clouddirectory.UpgradePublishedSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudDirectoryStub struct {
    activities awsactivities.CloudDirectoryActivities
}

func NewCloudDirectoryStub() CloudDirectoryClient {
    return &CloudDirectoryStub{}
}

func (a *CloudDirectoryStub) AddFacetToObject(ctx workflow.Context, input *clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error) {
    var output clouddirectory.AddFacetToObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddFacetToObject, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) AddFacetToObjectAsync(ctx workflow.Context, input *clouddirectory.AddFacetToObjectInput) *ClouddirectoryAddFacetToObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddFacetToObject, input)
    return &ClouddirectoryAddFacetToObjectResult{Result: future}
}

func (a *CloudDirectoryStub) ApplySchema(ctx workflow.Context, input *clouddirectory.ApplySchemaInput) (*clouddirectory.ApplySchemaOutput, error) {
    var output clouddirectory.ApplySchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ApplySchema, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ApplySchemaAsync(ctx workflow.Context, input *clouddirectory.ApplySchemaInput) *ClouddirectoryApplySchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ApplySchema, input)
    return &ClouddirectoryApplySchemaResult{Result: future}
}

func (a *CloudDirectoryStub) AttachObject(ctx workflow.Context, input *clouddirectory.AttachObjectInput) (*clouddirectory.AttachObjectOutput, error) {
    var output clouddirectory.AttachObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachObject, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) AttachObjectAsync(ctx workflow.Context, input *clouddirectory.AttachObjectInput) *ClouddirectoryAttachObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachObject, input)
    return &ClouddirectoryAttachObjectResult{Result: future}
}

func (a *CloudDirectoryStub) AttachPolicy(ctx workflow.Context, input *clouddirectory.AttachPolicyInput) (*clouddirectory.AttachPolicyOutput, error) {
    var output clouddirectory.AttachPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) AttachPolicyAsync(ctx workflow.Context, input *clouddirectory.AttachPolicyInput) *ClouddirectoryAttachPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachPolicy, input)
    return &ClouddirectoryAttachPolicyResult{Result: future}
}

func (a *CloudDirectoryStub) AttachToIndex(ctx workflow.Context, input *clouddirectory.AttachToIndexInput) (*clouddirectory.AttachToIndexOutput, error) {
    var output clouddirectory.AttachToIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachToIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) AttachToIndexAsync(ctx workflow.Context, input *clouddirectory.AttachToIndexInput) *ClouddirectoryAttachToIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachToIndex, input)
    return &ClouddirectoryAttachToIndexResult{Result: future}
}

func (a *CloudDirectoryStub) AttachTypedLink(ctx workflow.Context, input *clouddirectory.AttachTypedLinkInput) (*clouddirectory.AttachTypedLinkOutput, error) {
    var output clouddirectory.AttachTypedLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachTypedLink, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) AttachTypedLinkAsync(ctx workflow.Context, input *clouddirectory.AttachTypedLinkInput) *ClouddirectoryAttachTypedLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachTypedLink, input)
    return &ClouddirectoryAttachTypedLinkResult{Result: future}
}

func (a *CloudDirectoryStub) BatchRead(ctx workflow.Context, input *clouddirectory.BatchReadInput) (*clouddirectory.BatchReadOutput, error) {
    var output clouddirectory.BatchReadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchRead, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) BatchReadAsync(ctx workflow.Context, input *clouddirectory.BatchReadInput) *ClouddirectoryBatchReadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchRead, input)
    return &ClouddirectoryBatchReadResult{Result: future}
}

func (a *CloudDirectoryStub) BatchWrite(ctx workflow.Context, input *clouddirectory.BatchWriteInput) (*clouddirectory.BatchWriteOutput, error) {
    var output clouddirectory.BatchWriteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchWrite, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) BatchWriteAsync(ctx workflow.Context, input *clouddirectory.BatchWriteInput) *ClouddirectoryBatchWriteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchWrite, input)
    return &ClouddirectoryBatchWriteResult{Result: future}
}

func (a *CloudDirectoryStub) CreateDirectory(ctx workflow.Context, input *clouddirectory.CreateDirectoryInput) (*clouddirectory.CreateDirectoryOutput, error) {
    var output clouddirectory.CreateDirectoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDirectory, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) CreateDirectoryAsync(ctx workflow.Context, input *clouddirectory.CreateDirectoryInput) *ClouddirectoryCreateDirectoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDirectory, input)
    return &ClouddirectoryCreateDirectoryResult{Result: future}
}

func (a *CloudDirectoryStub) CreateFacet(ctx workflow.Context, input *clouddirectory.CreateFacetInput) (*clouddirectory.CreateFacetOutput, error) {
    var output clouddirectory.CreateFacetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFacet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) CreateFacetAsync(ctx workflow.Context, input *clouddirectory.CreateFacetInput) *ClouddirectoryCreateFacetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFacet, input)
    return &ClouddirectoryCreateFacetResult{Result: future}
}

func (a *CloudDirectoryStub) CreateIndex(ctx workflow.Context, input *clouddirectory.CreateIndexInput) (*clouddirectory.CreateIndexOutput, error) {
    var output clouddirectory.CreateIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) CreateIndexAsync(ctx workflow.Context, input *clouddirectory.CreateIndexInput) *ClouddirectoryCreateIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateIndex, input)
    return &ClouddirectoryCreateIndexResult{Result: future}
}

func (a *CloudDirectoryStub) CreateObject(ctx workflow.Context, input *clouddirectory.CreateObjectInput) (*clouddirectory.CreateObjectOutput, error) {
    var output clouddirectory.CreateObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateObject, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) CreateObjectAsync(ctx workflow.Context, input *clouddirectory.CreateObjectInput) *ClouddirectoryCreateObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateObject, input)
    return &ClouddirectoryCreateObjectResult{Result: future}
}

func (a *CloudDirectoryStub) CreateSchema(ctx workflow.Context, input *clouddirectory.CreateSchemaInput) (*clouddirectory.CreateSchemaOutput, error) {
    var output clouddirectory.CreateSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) CreateSchemaAsync(ctx workflow.Context, input *clouddirectory.CreateSchemaInput) *ClouddirectoryCreateSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSchema, input)
    return &ClouddirectoryCreateSchemaResult{Result: future}
}

func (a *CloudDirectoryStub) CreateTypedLinkFacet(ctx workflow.Context, input *clouddirectory.CreateTypedLinkFacetInput) (*clouddirectory.CreateTypedLinkFacetOutput, error) {
    var output clouddirectory.CreateTypedLinkFacetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTypedLinkFacet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) CreateTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.CreateTypedLinkFacetInput) *ClouddirectoryCreateTypedLinkFacetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTypedLinkFacet, input)
    return &ClouddirectoryCreateTypedLinkFacetResult{Result: future}
}

func (a *CloudDirectoryStub) DeleteDirectory(ctx workflow.Context, input *clouddirectory.DeleteDirectoryInput) (*clouddirectory.DeleteDirectoryOutput, error) {
    var output clouddirectory.DeleteDirectoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectory, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DeleteDirectoryAsync(ctx workflow.Context, input *clouddirectory.DeleteDirectoryInput) *ClouddirectoryDeleteDirectoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectory, input)
    return &ClouddirectoryDeleteDirectoryResult{Result: future}
}

func (a *CloudDirectoryStub) DeleteFacet(ctx workflow.Context, input *clouddirectory.DeleteFacetInput) (*clouddirectory.DeleteFacetOutput, error) {
    var output clouddirectory.DeleteFacetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFacet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DeleteFacetAsync(ctx workflow.Context, input *clouddirectory.DeleteFacetInput) *ClouddirectoryDeleteFacetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFacet, input)
    return &ClouddirectoryDeleteFacetResult{Result: future}
}

func (a *CloudDirectoryStub) DeleteObject(ctx workflow.Context, input *clouddirectory.DeleteObjectInput) (*clouddirectory.DeleteObjectOutput, error) {
    var output clouddirectory.DeleteObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteObject, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DeleteObjectAsync(ctx workflow.Context, input *clouddirectory.DeleteObjectInput) *ClouddirectoryDeleteObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteObject, input)
    return &ClouddirectoryDeleteObjectResult{Result: future}
}

func (a *CloudDirectoryStub) DeleteSchema(ctx workflow.Context, input *clouddirectory.DeleteSchemaInput) (*clouddirectory.DeleteSchemaOutput, error) {
    var output clouddirectory.DeleteSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DeleteSchemaAsync(ctx workflow.Context, input *clouddirectory.DeleteSchemaInput) *ClouddirectoryDeleteSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSchema, input)
    return &ClouddirectoryDeleteSchemaResult{Result: future}
}

func (a *CloudDirectoryStub) DeleteTypedLinkFacet(ctx workflow.Context, input *clouddirectory.DeleteTypedLinkFacetInput) (*clouddirectory.DeleteTypedLinkFacetOutput, error) {
    var output clouddirectory.DeleteTypedLinkFacetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTypedLinkFacet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DeleteTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.DeleteTypedLinkFacetInput) *ClouddirectoryDeleteTypedLinkFacetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTypedLinkFacet, input)
    return &ClouddirectoryDeleteTypedLinkFacetResult{Result: future}
}

func (a *CloudDirectoryStub) DetachFromIndex(ctx workflow.Context, input *clouddirectory.DetachFromIndexInput) (*clouddirectory.DetachFromIndexOutput, error) {
    var output clouddirectory.DetachFromIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachFromIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DetachFromIndexAsync(ctx workflow.Context, input *clouddirectory.DetachFromIndexInput) *ClouddirectoryDetachFromIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachFromIndex, input)
    return &ClouddirectoryDetachFromIndexResult{Result: future}
}

func (a *CloudDirectoryStub) DetachObject(ctx workflow.Context, input *clouddirectory.DetachObjectInput) (*clouddirectory.DetachObjectOutput, error) {
    var output clouddirectory.DetachObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachObject, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DetachObjectAsync(ctx workflow.Context, input *clouddirectory.DetachObjectInput) *ClouddirectoryDetachObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachObject, input)
    return &ClouddirectoryDetachObjectResult{Result: future}
}

func (a *CloudDirectoryStub) DetachPolicy(ctx workflow.Context, input *clouddirectory.DetachPolicyInput) (*clouddirectory.DetachPolicyOutput, error) {
    var output clouddirectory.DetachPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DetachPolicyAsync(ctx workflow.Context, input *clouddirectory.DetachPolicyInput) *ClouddirectoryDetachPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachPolicy, input)
    return &ClouddirectoryDetachPolicyResult{Result: future}
}

func (a *CloudDirectoryStub) DetachTypedLink(ctx workflow.Context, input *clouddirectory.DetachTypedLinkInput) (*clouddirectory.DetachTypedLinkOutput, error) {
    var output clouddirectory.DetachTypedLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachTypedLink, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DetachTypedLinkAsync(ctx workflow.Context, input *clouddirectory.DetachTypedLinkInput) *ClouddirectoryDetachTypedLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachTypedLink, input)
    return &ClouddirectoryDetachTypedLinkResult{Result: future}
}

func (a *CloudDirectoryStub) DisableDirectory(ctx workflow.Context, input *clouddirectory.DisableDirectoryInput) (*clouddirectory.DisableDirectoryOutput, error) {
    var output clouddirectory.DisableDirectoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableDirectory, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) DisableDirectoryAsync(ctx workflow.Context, input *clouddirectory.DisableDirectoryInput) *ClouddirectoryDisableDirectoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableDirectory, input)
    return &ClouddirectoryDisableDirectoryResult{Result: future}
}

func (a *CloudDirectoryStub) EnableDirectory(ctx workflow.Context, input *clouddirectory.EnableDirectoryInput) (*clouddirectory.EnableDirectoryOutput, error) {
    var output clouddirectory.EnableDirectoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableDirectory, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) EnableDirectoryAsync(ctx workflow.Context, input *clouddirectory.EnableDirectoryInput) *ClouddirectoryEnableDirectoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableDirectory, input)
    return &ClouddirectoryEnableDirectoryResult{Result: future}
}

func (a *CloudDirectoryStub) GetAppliedSchemaVersion(ctx workflow.Context, input *clouddirectory.GetAppliedSchemaVersionInput) (*clouddirectory.GetAppliedSchemaVersionOutput, error) {
    var output clouddirectory.GetAppliedSchemaVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAppliedSchemaVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetAppliedSchemaVersionAsync(ctx workflow.Context, input *clouddirectory.GetAppliedSchemaVersionInput) *ClouddirectoryGetAppliedSchemaVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAppliedSchemaVersion, input)
    return &ClouddirectoryGetAppliedSchemaVersionResult{Result: future}
}

func (a *CloudDirectoryStub) GetDirectory(ctx workflow.Context, input *clouddirectory.GetDirectoryInput) (*clouddirectory.GetDirectoryOutput, error) {
    var output clouddirectory.GetDirectoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDirectory, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetDirectoryAsync(ctx workflow.Context, input *clouddirectory.GetDirectoryInput) *ClouddirectoryGetDirectoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDirectory, input)
    return &ClouddirectoryGetDirectoryResult{Result: future}
}

func (a *CloudDirectoryStub) GetFacet(ctx workflow.Context, input *clouddirectory.GetFacetInput) (*clouddirectory.GetFacetOutput, error) {
    var output clouddirectory.GetFacetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFacet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetFacetAsync(ctx workflow.Context, input *clouddirectory.GetFacetInput) *ClouddirectoryGetFacetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFacet, input)
    return &ClouddirectoryGetFacetResult{Result: future}
}

func (a *CloudDirectoryStub) GetLinkAttributes(ctx workflow.Context, input *clouddirectory.GetLinkAttributesInput) (*clouddirectory.GetLinkAttributesOutput, error) {
    var output clouddirectory.GetLinkAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLinkAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetLinkAttributesAsync(ctx workflow.Context, input *clouddirectory.GetLinkAttributesInput) *ClouddirectoryGetLinkAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLinkAttributes, input)
    return &ClouddirectoryGetLinkAttributesResult{Result: future}
}

func (a *CloudDirectoryStub) GetObjectAttributes(ctx workflow.Context, input *clouddirectory.GetObjectAttributesInput) (*clouddirectory.GetObjectAttributesOutput, error) {
    var output clouddirectory.GetObjectAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetObjectAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.GetObjectAttributesInput) *ClouddirectoryGetObjectAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetObjectAttributes, input)
    return &ClouddirectoryGetObjectAttributesResult{Result: future}
}

func (a *CloudDirectoryStub) GetObjectInformation(ctx workflow.Context, input *clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error) {
    var output clouddirectory.GetObjectInformationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetObjectInformation, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetObjectInformationAsync(ctx workflow.Context, input *clouddirectory.GetObjectInformationInput) *ClouddirectoryGetObjectInformationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetObjectInformation, input)
    return &ClouddirectoryGetObjectInformationResult{Result: future}
}

func (a *CloudDirectoryStub) GetSchemaAsJson(ctx workflow.Context, input *clouddirectory.GetSchemaAsJsonInput) (*clouddirectory.GetSchemaAsJsonOutput, error) {
    var output clouddirectory.GetSchemaAsJsonOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSchemaAsJson, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetSchemaAsJsonAsync(ctx workflow.Context, input *clouddirectory.GetSchemaAsJsonInput) *ClouddirectoryGetSchemaAsJsonResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSchemaAsJson, input)
    return &ClouddirectoryGetSchemaAsJsonResult{Result: future}
}

func (a *CloudDirectoryStub) GetTypedLinkFacetInformation(ctx workflow.Context, input *clouddirectory.GetTypedLinkFacetInformationInput) (*clouddirectory.GetTypedLinkFacetInformationOutput, error) {
    var output clouddirectory.GetTypedLinkFacetInformationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTypedLinkFacetInformation, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) GetTypedLinkFacetInformationAsync(ctx workflow.Context, input *clouddirectory.GetTypedLinkFacetInformationInput) *ClouddirectoryGetTypedLinkFacetInformationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTypedLinkFacetInformation, input)
    return &ClouddirectoryGetTypedLinkFacetInformationResult{Result: future}
}

func (a *CloudDirectoryStub) ListAppliedSchemaArns(ctx workflow.Context, input *clouddirectory.ListAppliedSchemaArnsInput) (*clouddirectory.ListAppliedSchemaArnsOutput, error) {
    var output clouddirectory.ListAppliedSchemaArnsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAppliedSchemaArns, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListAppliedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListAppliedSchemaArnsInput) *ClouddirectoryListAppliedSchemaArnsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAppliedSchemaArns, input)
    return &ClouddirectoryListAppliedSchemaArnsResult{Result: future}
}

func (a *CloudDirectoryStub) ListAttachedIndices(ctx workflow.Context, input *clouddirectory.ListAttachedIndicesInput) (*clouddirectory.ListAttachedIndicesOutput, error) {
    var output clouddirectory.ListAttachedIndicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttachedIndices, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListAttachedIndicesAsync(ctx workflow.Context, input *clouddirectory.ListAttachedIndicesInput) *ClouddirectoryListAttachedIndicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAttachedIndices, input)
    return &ClouddirectoryListAttachedIndicesResult{Result: future}
}

func (a *CloudDirectoryStub) ListDevelopmentSchemaArns(ctx workflow.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error) {
    var output clouddirectory.ListDevelopmentSchemaArnsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDevelopmentSchemaArns, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListDevelopmentSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput) *ClouddirectoryListDevelopmentSchemaArnsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDevelopmentSchemaArns, input)
    return &ClouddirectoryListDevelopmentSchemaArnsResult{Result: future}
}

func (a *CloudDirectoryStub) ListDirectories(ctx workflow.Context, input *clouddirectory.ListDirectoriesInput) (*clouddirectory.ListDirectoriesOutput, error) {
    var output clouddirectory.ListDirectoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDirectories, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListDirectoriesAsync(ctx workflow.Context, input *clouddirectory.ListDirectoriesInput) *ClouddirectoryListDirectoriesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDirectories, input)
    return &ClouddirectoryListDirectoriesResult{Result: future}
}

func (a *CloudDirectoryStub) ListFacetAttributes(ctx workflow.Context, input *clouddirectory.ListFacetAttributesInput) (*clouddirectory.ListFacetAttributesOutput, error) {
    var output clouddirectory.ListFacetAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFacetAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListFacetAttributesAsync(ctx workflow.Context, input *clouddirectory.ListFacetAttributesInput) *ClouddirectoryListFacetAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFacetAttributes, input)
    return &ClouddirectoryListFacetAttributesResult{Result: future}
}

func (a *CloudDirectoryStub) ListFacetNames(ctx workflow.Context, input *clouddirectory.ListFacetNamesInput) (*clouddirectory.ListFacetNamesOutput, error) {
    var output clouddirectory.ListFacetNamesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFacetNames, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListFacetNamesAsync(ctx workflow.Context, input *clouddirectory.ListFacetNamesInput) *ClouddirectoryListFacetNamesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFacetNames, input)
    return &ClouddirectoryListFacetNamesResult{Result: future}
}

func (a *CloudDirectoryStub) ListIncomingTypedLinks(ctx workflow.Context, input *clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
    var output clouddirectory.ListIncomingTypedLinksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIncomingTypedLinks, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListIncomingTypedLinksAsync(ctx workflow.Context, input *clouddirectory.ListIncomingTypedLinksInput) *ClouddirectoryListIncomingTypedLinksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListIncomingTypedLinks, input)
    return &ClouddirectoryListIncomingTypedLinksResult{Result: future}
}

func (a *CloudDirectoryStub) ListIndex(ctx workflow.Context, input *clouddirectory.ListIndexInput) (*clouddirectory.ListIndexOutput, error) {
    var output clouddirectory.ListIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListIndexAsync(ctx workflow.Context, input *clouddirectory.ListIndexInput) *ClouddirectoryListIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListIndex, input)
    return &ClouddirectoryListIndexResult{Result: future}
}

func (a *CloudDirectoryStub) ListManagedSchemaArns(ctx workflow.Context, input *clouddirectory.ListManagedSchemaArnsInput) (*clouddirectory.ListManagedSchemaArnsOutput, error) {
    var output clouddirectory.ListManagedSchemaArnsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListManagedSchemaArns, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListManagedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListManagedSchemaArnsInput) *ClouddirectoryListManagedSchemaArnsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListManagedSchemaArns, input)
    return &ClouddirectoryListManagedSchemaArnsResult{Result: future}
}

func (a *CloudDirectoryStub) ListObjectAttributes(ctx workflow.Context, input *clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error) {
    var output clouddirectory.ListObjectAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListObjectAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.ListObjectAttributesInput) *ClouddirectoryListObjectAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListObjectAttributes, input)
    return &ClouddirectoryListObjectAttributesResult{Result: future}
}

func (a *CloudDirectoryStub) ListObjectChildren(ctx workflow.Context, input *clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error) {
    var output clouddirectory.ListObjectChildrenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListObjectChildren, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListObjectChildrenAsync(ctx workflow.Context, input *clouddirectory.ListObjectChildrenInput) *ClouddirectoryListObjectChildrenResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListObjectChildren, input)
    return &ClouddirectoryListObjectChildrenResult{Result: future}
}

func (a *CloudDirectoryStub) ListObjectParentPaths(ctx workflow.Context, input *clouddirectory.ListObjectParentPathsInput) (*clouddirectory.ListObjectParentPathsOutput, error) {
    var output clouddirectory.ListObjectParentPathsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListObjectParentPaths, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListObjectParentPathsAsync(ctx workflow.Context, input *clouddirectory.ListObjectParentPathsInput) *ClouddirectoryListObjectParentPathsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListObjectParentPaths, input)
    return &ClouddirectoryListObjectParentPathsResult{Result: future}
}

func (a *CloudDirectoryStub) ListObjectParents(ctx workflow.Context, input *clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error) {
    var output clouddirectory.ListObjectParentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListObjectParents, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListObjectParentsAsync(ctx workflow.Context, input *clouddirectory.ListObjectParentsInput) *ClouddirectoryListObjectParentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListObjectParents, input)
    return &ClouddirectoryListObjectParentsResult{Result: future}
}

func (a *CloudDirectoryStub) ListObjectPolicies(ctx workflow.Context, input *clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error) {
    var output clouddirectory.ListObjectPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListObjectPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListObjectPoliciesAsync(ctx workflow.Context, input *clouddirectory.ListObjectPoliciesInput) *ClouddirectoryListObjectPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListObjectPolicies, input)
    return &ClouddirectoryListObjectPoliciesResult{Result: future}
}

func (a *CloudDirectoryStub) ListOutgoingTypedLinks(ctx workflow.Context, input *clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
    var output clouddirectory.ListOutgoingTypedLinksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOutgoingTypedLinks, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListOutgoingTypedLinksAsync(ctx workflow.Context, input *clouddirectory.ListOutgoingTypedLinksInput) *ClouddirectoryListOutgoingTypedLinksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOutgoingTypedLinks, input)
    return &ClouddirectoryListOutgoingTypedLinksResult{Result: future}
}

func (a *CloudDirectoryStub) ListPolicyAttachments(ctx workflow.Context, input *clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error) {
    var output clouddirectory.ListPolicyAttachmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicyAttachments, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListPolicyAttachmentsAsync(ctx workflow.Context, input *clouddirectory.ListPolicyAttachmentsInput) *ClouddirectoryListPolicyAttachmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPolicyAttachments, input)
    return &ClouddirectoryListPolicyAttachmentsResult{Result: future}
}

func (a *CloudDirectoryStub) ListPublishedSchemaArns(ctx workflow.Context, input *clouddirectory.ListPublishedSchemaArnsInput) (*clouddirectory.ListPublishedSchemaArnsOutput, error) {
    var output clouddirectory.ListPublishedSchemaArnsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPublishedSchemaArns, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListPublishedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListPublishedSchemaArnsInput) *ClouddirectoryListPublishedSchemaArnsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPublishedSchemaArns, input)
    return &ClouddirectoryListPublishedSchemaArnsResult{Result: future}
}

func (a *CloudDirectoryStub) ListTagsForResource(ctx workflow.Context, input *clouddirectory.ListTagsForResourceInput) (*clouddirectory.ListTagsForResourceOutput, error) {
    var output clouddirectory.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListTagsForResourceAsync(ctx workflow.Context, input *clouddirectory.ListTagsForResourceInput) *ClouddirectoryListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &ClouddirectoryListTagsForResourceResult{Result: future}
}

func (a *CloudDirectoryStub) ListTypedLinkFacetAttributes(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error) {
    var output clouddirectory.ListTypedLinkFacetAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTypedLinkFacetAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListTypedLinkFacetAttributesAsync(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput) *ClouddirectoryListTypedLinkFacetAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTypedLinkFacetAttributes, input)
    return &ClouddirectoryListTypedLinkFacetAttributesResult{Result: future}
}

func (a *CloudDirectoryStub) ListTypedLinkFacetNames(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetNamesInput) (*clouddirectory.ListTypedLinkFacetNamesOutput, error) {
    var output clouddirectory.ListTypedLinkFacetNamesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTypedLinkFacetNames, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) ListTypedLinkFacetNamesAsync(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetNamesInput) *ClouddirectoryListTypedLinkFacetNamesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTypedLinkFacetNames, input)
    return &ClouddirectoryListTypedLinkFacetNamesResult{Result: future}
}

func (a *CloudDirectoryStub) LookupPolicy(ctx workflow.Context, input *clouddirectory.LookupPolicyInput) (*clouddirectory.LookupPolicyOutput, error) {
    var output clouddirectory.LookupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.LookupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) LookupPolicyAsync(ctx workflow.Context, input *clouddirectory.LookupPolicyInput) *ClouddirectoryLookupPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.LookupPolicy, input)
    return &ClouddirectoryLookupPolicyResult{Result: future}
}

func (a *CloudDirectoryStub) PublishSchema(ctx workflow.Context, input *clouddirectory.PublishSchemaInput) (*clouddirectory.PublishSchemaOutput, error) {
    var output clouddirectory.PublishSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PublishSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) PublishSchemaAsync(ctx workflow.Context, input *clouddirectory.PublishSchemaInput) *ClouddirectoryPublishSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PublishSchema, input)
    return &ClouddirectoryPublishSchemaResult{Result: future}
}

func (a *CloudDirectoryStub) PutSchemaFromJson(ctx workflow.Context, input *clouddirectory.PutSchemaFromJsonInput) (*clouddirectory.PutSchemaFromJsonOutput, error) {
    var output clouddirectory.PutSchemaFromJsonOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSchemaFromJson, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) PutSchemaFromJsonAsync(ctx workflow.Context, input *clouddirectory.PutSchemaFromJsonInput) *ClouddirectoryPutSchemaFromJsonResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutSchemaFromJson, input)
    return &ClouddirectoryPutSchemaFromJsonResult{Result: future}
}

func (a *CloudDirectoryStub) RemoveFacetFromObject(ctx workflow.Context, input *clouddirectory.RemoveFacetFromObjectInput) (*clouddirectory.RemoveFacetFromObjectOutput, error) {
    var output clouddirectory.RemoveFacetFromObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveFacetFromObject, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) RemoveFacetFromObjectAsync(ctx workflow.Context, input *clouddirectory.RemoveFacetFromObjectInput) *ClouddirectoryRemoveFacetFromObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveFacetFromObject, input)
    return &ClouddirectoryRemoveFacetFromObjectResult{Result: future}
}

func (a *CloudDirectoryStub) TagResource(ctx workflow.Context, input *clouddirectory.TagResourceInput) (*clouddirectory.TagResourceOutput, error) {
    var output clouddirectory.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) TagResourceAsync(ctx workflow.Context, input *clouddirectory.TagResourceInput) *ClouddirectoryTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &ClouddirectoryTagResourceResult{Result: future}
}

func (a *CloudDirectoryStub) UntagResource(ctx workflow.Context, input *clouddirectory.UntagResourceInput) (*clouddirectory.UntagResourceOutput, error) {
    var output clouddirectory.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UntagResourceAsync(ctx workflow.Context, input *clouddirectory.UntagResourceInput) *ClouddirectoryUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &ClouddirectoryUntagResourceResult{Result: future}
}

func (a *CloudDirectoryStub) UpdateFacet(ctx workflow.Context, input *clouddirectory.UpdateFacetInput) (*clouddirectory.UpdateFacetOutput, error) {
    var output clouddirectory.UpdateFacetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFacet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UpdateFacetAsync(ctx workflow.Context, input *clouddirectory.UpdateFacetInput) *ClouddirectoryUpdateFacetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFacet, input)
    return &ClouddirectoryUpdateFacetResult{Result: future}
}

func (a *CloudDirectoryStub) UpdateLinkAttributes(ctx workflow.Context, input *clouddirectory.UpdateLinkAttributesInput) (*clouddirectory.UpdateLinkAttributesOutput, error) {
    var output clouddirectory.UpdateLinkAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLinkAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UpdateLinkAttributesAsync(ctx workflow.Context, input *clouddirectory.UpdateLinkAttributesInput) *ClouddirectoryUpdateLinkAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateLinkAttributes, input)
    return &ClouddirectoryUpdateLinkAttributesResult{Result: future}
}

func (a *CloudDirectoryStub) UpdateObjectAttributes(ctx workflow.Context, input *clouddirectory.UpdateObjectAttributesInput) (*clouddirectory.UpdateObjectAttributesOutput, error) {
    var output clouddirectory.UpdateObjectAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateObjectAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UpdateObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.UpdateObjectAttributesInput) *ClouddirectoryUpdateObjectAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateObjectAttributes, input)
    return &ClouddirectoryUpdateObjectAttributesResult{Result: future}
}

func (a *CloudDirectoryStub) UpdateSchema(ctx workflow.Context, input *clouddirectory.UpdateSchemaInput) (*clouddirectory.UpdateSchemaOutput, error) {
    var output clouddirectory.UpdateSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UpdateSchemaAsync(ctx workflow.Context, input *clouddirectory.UpdateSchemaInput) *ClouddirectoryUpdateSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSchema, input)
    return &ClouddirectoryUpdateSchemaResult{Result: future}
}

func (a *CloudDirectoryStub) UpdateTypedLinkFacet(ctx workflow.Context, input *clouddirectory.UpdateTypedLinkFacetInput) (*clouddirectory.UpdateTypedLinkFacetOutput, error) {
    var output clouddirectory.UpdateTypedLinkFacetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTypedLinkFacet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UpdateTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.UpdateTypedLinkFacetInput) *ClouddirectoryUpdateTypedLinkFacetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTypedLinkFacet, input)
    return &ClouddirectoryUpdateTypedLinkFacetResult{Result: future}
}

func (a *CloudDirectoryStub) UpgradeAppliedSchema(ctx workflow.Context, input *clouddirectory.UpgradeAppliedSchemaInput) (*clouddirectory.UpgradeAppliedSchemaOutput, error) {
    var output clouddirectory.UpgradeAppliedSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpgradeAppliedSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UpgradeAppliedSchemaAsync(ctx workflow.Context, input *clouddirectory.UpgradeAppliedSchemaInput) *ClouddirectoryUpgradeAppliedSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpgradeAppliedSchema, input)
    return &ClouddirectoryUpgradeAppliedSchemaResult{Result: future}
}

func (a *CloudDirectoryStub) UpgradePublishedSchema(ctx workflow.Context, input *clouddirectory.UpgradePublishedSchemaInput) (*clouddirectory.UpgradePublishedSchemaOutput, error) {
    var output clouddirectory.UpgradePublishedSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpgradePublishedSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudDirectoryStub) UpgradePublishedSchemaAsync(ctx workflow.Context, input *clouddirectory.UpgradePublishedSchemaInput) *ClouddirectoryUpgradePublishedSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpgradePublishedSchema, input)
    return &ClouddirectoryUpgradePublishedSchemaResult{Result: future}
}
