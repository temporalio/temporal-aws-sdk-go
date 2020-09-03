package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/aws/aws-sdk-go/service/clouddirectory/clouddirectoryiface"
)

type CloudDirectoryActivities struct {
	client clouddirectoryiface.CloudDirectoryAPI
}

func NewCloudDirectoryActivities(client clouddirectoryiface.CloudDirectoryAPI) *CloudDirectoryActivities {
    return &CloudDirectoryActivities{client: client}
}

func (a *CloudDirectoryActivities) AddFacetToObject(input *clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error) {
    return a.client.AddFacetToObject(input)
}

func (a *CloudDirectoryActivities) ApplySchema(input *clouddirectory.ApplySchemaInput) (*clouddirectory.ApplySchemaOutput, error) {
    return a.client.ApplySchema(input)
}

func (a *CloudDirectoryActivities) AttachObject(input *clouddirectory.AttachObjectInput) (*clouddirectory.AttachObjectOutput, error) {
    return a.client.AttachObject(input)
}

func (a *CloudDirectoryActivities) AttachPolicy(input *clouddirectory.AttachPolicyInput) (*clouddirectory.AttachPolicyOutput, error) {
    return a.client.AttachPolicy(input)
}

func (a *CloudDirectoryActivities) AttachToIndex(input *clouddirectory.AttachToIndexInput) (*clouddirectory.AttachToIndexOutput, error) {
    return a.client.AttachToIndex(input)
}

func (a *CloudDirectoryActivities) AttachTypedLink(input *clouddirectory.AttachTypedLinkInput) (*clouddirectory.AttachTypedLinkOutput, error) {
    return a.client.AttachTypedLink(input)
}

func (a *CloudDirectoryActivities) BatchRead(input *clouddirectory.BatchReadInput) (*clouddirectory.BatchReadOutput, error) {
    return a.client.BatchRead(input)
}

func (a *CloudDirectoryActivities) BatchWrite(input *clouddirectory.BatchWriteInput) (*clouddirectory.BatchWriteOutput, error) {
    return a.client.BatchWrite(input)
}

func (a *CloudDirectoryActivities) CreateDirectory(input *clouddirectory.CreateDirectoryInput) (*clouddirectory.CreateDirectoryOutput, error) {
    return a.client.CreateDirectory(input)
}

func (a *CloudDirectoryActivities) CreateFacet(input *clouddirectory.CreateFacetInput) (*clouddirectory.CreateFacetOutput, error) {
    return a.client.CreateFacet(input)
}

func (a *CloudDirectoryActivities) CreateIndex(input *clouddirectory.CreateIndexInput) (*clouddirectory.CreateIndexOutput, error) {
    return a.client.CreateIndex(input)
}

func (a *CloudDirectoryActivities) CreateObject(input *clouddirectory.CreateObjectInput) (*clouddirectory.CreateObjectOutput, error) {
    return a.client.CreateObject(input)
}

func (a *CloudDirectoryActivities) CreateSchema(input *clouddirectory.CreateSchemaInput) (*clouddirectory.CreateSchemaOutput, error) {
    return a.client.CreateSchema(input)
}

func (a *CloudDirectoryActivities) CreateTypedLinkFacet(input *clouddirectory.CreateTypedLinkFacetInput) (*clouddirectory.CreateTypedLinkFacetOutput, error) {
    return a.client.CreateTypedLinkFacet(input)
}

func (a *CloudDirectoryActivities) DeleteDirectory(input *clouddirectory.DeleteDirectoryInput) (*clouddirectory.DeleteDirectoryOutput, error) {
    return a.client.DeleteDirectory(input)
}

func (a *CloudDirectoryActivities) DeleteFacet(input *clouddirectory.DeleteFacetInput) (*clouddirectory.DeleteFacetOutput, error) {
    return a.client.DeleteFacet(input)
}

func (a *CloudDirectoryActivities) DeleteObject(input *clouddirectory.DeleteObjectInput) (*clouddirectory.DeleteObjectOutput, error) {
    return a.client.DeleteObject(input)
}

func (a *CloudDirectoryActivities) DeleteSchema(input *clouddirectory.DeleteSchemaInput) (*clouddirectory.DeleteSchemaOutput, error) {
    return a.client.DeleteSchema(input)
}

func (a *CloudDirectoryActivities) DeleteTypedLinkFacet(input *clouddirectory.DeleteTypedLinkFacetInput) (*clouddirectory.DeleteTypedLinkFacetOutput, error) {
    return a.client.DeleteTypedLinkFacet(input)
}

func (a *CloudDirectoryActivities) DetachFromIndex(input *clouddirectory.DetachFromIndexInput) (*clouddirectory.DetachFromIndexOutput, error) {
    return a.client.DetachFromIndex(input)
}

func (a *CloudDirectoryActivities) DetachObject(input *clouddirectory.DetachObjectInput) (*clouddirectory.DetachObjectOutput, error) {
    return a.client.DetachObject(input)
}

func (a *CloudDirectoryActivities) DetachPolicy(input *clouddirectory.DetachPolicyInput) (*clouddirectory.DetachPolicyOutput, error) {
    return a.client.DetachPolicy(input)
}

func (a *CloudDirectoryActivities) DetachTypedLink(input *clouddirectory.DetachTypedLinkInput) (*clouddirectory.DetachTypedLinkOutput, error) {
    return a.client.DetachTypedLink(input)
}

func (a *CloudDirectoryActivities) DisableDirectory(input *clouddirectory.DisableDirectoryInput) (*clouddirectory.DisableDirectoryOutput, error) {
    return a.client.DisableDirectory(input)
}

func (a *CloudDirectoryActivities) EnableDirectory(input *clouddirectory.EnableDirectoryInput) (*clouddirectory.EnableDirectoryOutput, error) {
    return a.client.EnableDirectory(input)
}

func (a *CloudDirectoryActivities) GetAppliedSchemaVersion(input *clouddirectory.GetAppliedSchemaVersionInput) (*clouddirectory.GetAppliedSchemaVersionOutput, error) {
    return a.client.GetAppliedSchemaVersion(input)
}

func (a *CloudDirectoryActivities) GetDirectory(input *clouddirectory.GetDirectoryInput) (*clouddirectory.GetDirectoryOutput, error) {
    return a.client.GetDirectory(input)
}

func (a *CloudDirectoryActivities) GetFacet(input *clouddirectory.GetFacetInput) (*clouddirectory.GetFacetOutput, error) {
    return a.client.GetFacet(input)
}

func (a *CloudDirectoryActivities) GetLinkAttributes(input *clouddirectory.GetLinkAttributesInput) (*clouddirectory.GetLinkAttributesOutput, error) {
    return a.client.GetLinkAttributes(input)
}

func (a *CloudDirectoryActivities) GetObjectAttributes(input *clouddirectory.GetObjectAttributesInput) (*clouddirectory.GetObjectAttributesOutput, error) {
    return a.client.GetObjectAttributes(input)
}

func (a *CloudDirectoryActivities) GetObjectInformation(input *clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error) {
    return a.client.GetObjectInformation(input)
}

func (a *CloudDirectoryActivities) GetSchemaAsJson(input *clouddirectory.GetSchemaAsJsonInput) (*clouddirectory.GetSchemaAsJsonOutput, error) {
    return a.client.GetSchemaAsJson(input)
}

func (a *CloudDirectoryActivities) GetTypedLinkFacetInformation(input *clouddirectory.GetTypedLinkFacetInformationInput) (*clouddirectory.GetTypedLinkFacetInformationOutput, error) {
    return a.client.GetTypedLinkFacetInformation(input)
}

func (a *CloudDirectoryActivities) ListAppliedSchemaArns(input *clouddirectory.ListAppliedSchemaArnsInput) (*clouddirectory.ListAppliedSchemaArnsOutput, error) {
    return a.client.ListAppliedSchemaArns(input)
}

func (a *CloudDirectoryActivities) ListAttachedIndices(input *clouddirectory.ListAttachedIndicesInput) (*clouddirectory.ListAttachedIndicesOutput, error) {
    return a.client.ListAttachedIndices(input)
}

func (a *CloudDirectoryActivities) ListDevelopmentSchemaArns(input *clouddirectory.ListDevelopmentSchemaArnsInput) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error) {
    return a.client.ListDevelopmentSchemaArns(input)
}

func (a *CloudDirectoryActivities) ListDirectories(input *clouddirectory.ListDirectoriesInput) (*clouddirectory.ListDirectoriesOutput, error) {
    return a.client.ListDirectories(input)
}

func (a *CloudDirectoryActivities) ListFacetAttributes(input *clouddirectory.ListFacetAttributesInput) (*clouddirectory.ListFacetAttributesOutput, error) {
    return a.client.ListFacetAttributes(input)
}

func (a *CloudDirectoryActivities) ListFacetNames(input *clouddirectory.ListFacetNamesInput) (*clouddirectory.ListFacetNamesOutput, error) {
    return a.client.ListFacetNames(input)
}

func (a *CloudDirectoryActivities) ListIncomingTypedLinks(input *clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
    return a.client.ListIncomingTypedLinks(input)
}

func (a *CloudDirectoryActivities) ListIndex(input *clouddirectory.ListIndexInput) (*clouddirectory.ListIndexOutput, error) {
    return a.client.ListIndex(input)
}

func (a *CloudDirectoryActivities) ListManagedSchemaArns(input *clouddirectory.ListManagedSchemaArnsInput) (*clouddirectory.ListManagedSchemaArnsOutput, error) {
    return a.client.ListManagedSchemaArns(input)
}

func (a *CloudDirectoryActivities) ListObjectAttributes(input *clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error) {
    return a.client.ListObjectAttributes(input)
}

func (a *CloudDirectoryActivities) ListObjectChildren(input *clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error) {
    return a.client.ListObjectChildren(input)
}

func (a *CloudDirectoryActivities) ListObjectParentPaths(input *clouddirectory.ListObjectParentPathsInput) (*clouddirectory.ListObjectParentPathsOutput, error) {
    return a.client.ListObjectParentPaths(input)
}

func (a *CloudDirectoryActivities) ListObjectParents(input *clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error) {
    return a.client.ListObjectParents(input)
}

func (a *CloudDirectoryActivities) ListObjectPolicies(input *clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error) {
    return a.client.ListObjectPolicies(input)
}

func (a *CloudDirectoryActivities) ListOutgoingTypedLinks(input *clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
    return a.client.ListOutgoingTypedLinks(input)
}

func (a *CloudDirectoryActivities) ListPolicyAttachments(input *clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error) {
    return a.client.ListPolicyAttachments(input)
}

func (a *CloudDirectoryActivities) ListPublishedSchemaArns(input *clouddirectory.ListPublishedSchemaArnsInput) (*clouddirectory.ListPublishedSchemaArnsOutput, error) {
    return a.client.ListPublishedSchemaArns(input)
}

func (a *CloudDirectoryActivities) ListTagsForResource(input *clouddirectory.ListTagsForResourceInput) (*clouddirectory.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CloudDirectoryActivities) ListTypedLinkFacetAttributes(input *clouddirectory.ListTypedLinkFacetAttributesInput) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error) {
    return a.client.ListTypedLinkFacetAttributes(input)
}

func (a *CloudDirectoryActivities) ListTypedLinkFacetNames(input *clouddirectory.ListTypedLinkFacetNamesInput) (*clouddirectory.ListTypedLinkFacetNamesOutput, error) {
    return a.client.ListTypedLinkFacetNames(input)
}

func (a *CloudDirectoryActivities) LookupPolicy(input *clouddirectory.LookupPolicyInput) (*clouddirectory.LookupPolicyOutput, error) {
    return a.client.LookupPolicy(input)
}

func (a *CloudDirectoryActivities) PublishSchema(input *clouddirectory.PublishSchemaInput) (*clouddirectory.PublishSchemaOutput, error) {
    return a.client.PublishSchema(input)
}

func (a *CloudDirectoryActivities) PutSchemaFromJson(input *clouddirectory.PutSchemaFromJsonInput) (*clouddirectory.PutSchemaFromJsonOutput, error) {
    return a.client.PutSchemaFromJson(input)
}

func (a *CloudDirectoryActivities) RemoveFacetFromObject(input *clouddirectory.RemoveFacetFromObjectInput) (*clouddirectory.RemoveFacetFromObjectOutput, error) {
    return a.client.RemoveFacetFromObject(input)
}

func (a *CloudDirectoryActivities) TagResource(input *clouddirectory.TagResourceInput) (*clouddirectory.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CloudDirectoryActivities) UntagResource(input *clouddirectory.UntagResourceInput) (*clouddirectory.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CloudDirectoryActivities) UpdateFacet(input *clouddirectory.UpdateFacetInput) (*clouddirectory.UpdateFacetOutput, error) {
    return a.client.UpdateFacet(input)
}

func (a *CloudDirectoryActivities) UpdateLinkAttributes(input *clouddirectory.UpdateLinkAttributesInput) (*clouddirectory.UpdateLinkAttributesOutput, error) {
    return a.client.UpdateLinkAttributes(input)
}

func (a *CloudDirectoryActivities) UpdateObjectAttributes(input *clouddirectory.UpdateObjectAttributesInput) (*clouddirectory.UpdateObjectAttributesOutput, error) {
    return a.client.UpdateObjectAttributes(input)
}

func (a *CloudDirectoryActivities) UpdateSchema(input *clouddirectory.UpdateSchemaInput) (*clouddirectory.UpdateSchemaOutput, error) {
    return a.client.UpdateSchema(input)
}

func (a *CloudDirectoryActivities) UpdateTypedLinkFacet(input *clouddirectory.UpdateTypedLinkFacetInput) (*clouddirectory.UpdateTypedLinkFacetOutput, error) {
    return a.client.UpdateTypedLinkFacet(input)
}

func (a *CloudDirectoryActivities) UpgradeAppliedSchema(input *clouddirectory.UpgradeAppliedSchemaInput) (*clouddirectory.UpgradeAppliedSchemaOutput, error) {
    return a.client.UpgradeAppliedSchema(input)
}

func (a *CloudDirectoryActivities) UpgradePublishedSchema(input *clouddirectory.UpgradePublishedSchemaInput) (*clouddirectory.UpgradePublishedSchemaOutput, error) {
    return a.client.UpgradePublishedSchema(input)
}
