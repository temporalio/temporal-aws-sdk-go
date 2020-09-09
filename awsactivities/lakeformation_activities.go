
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/aws/aws-sdk-go/service/lakeformation/lakeformationiface"
)

type LakeFormationActivities struct {
	client lakeformationiface.LakeFormationAPI
}

func NewLakeFormationActivities(session *session.Session, config... *aws.Config) *LakeFormationActivities {
    client := lakeformation.New(session, config...)
    return &LakeFormationActivities{client: client}
}

func (a *LakeFormationActivities) BatchGrantPermissions(input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error) {
    return a.client.BatchGrantPermissions(input)
}

func (a *LakeFormationActivities) BatchRevokePermissions(input *lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error) {
    return a.client.BatchRevokePermissions(input)
}

func (a *LakeFormationActivities) DeregisterResource(input *lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error) {
    return a.client.DeregisterResource(input)
}

func (a *LakeFormationActivities) DescribeResource(input *lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error) {
    return a.client.DescribeResource(input)
}

func (a *LakeFormationActivities) GetDataLakeSettings(input *lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error) {
    return a.client.GetDataLakeSettings(input)
}

func (a *LakeFormationActivities) GetEffectivePermissionsForPath(input *lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
    return a.client.GetEffectivePermissionsForPath(input)
}

func (a *LakeFormationActivities) GrantPermissions(input *lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error) {
    return a.client.GrantPermissions(input)
}

func (a *LakeFormationActivities) ListPermissions(input *lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error) {
    return a.client.ListPermissions(input)
}

func (a *LakeFormationActivities) ListResources(input *lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error) {
    return a.client.ListResources(input)
}

func (a *LakeFormationActivities) PutDataLakeSettings(input *lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error) {
    return a.client.PutDataLakeSettings(input)
}

func (a *LakeFormationActivities) RegisterResource(input *lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error) {
    return a.client.RegisterResource(input)
}

func (a *LakeFormationActivities) RevokePermissions(input *lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error) {
    return a.client.RevokePermissions(input)
}

func (a *LakeFormationActivities) UpdateResource(input *lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error) {
    return a.client.UpdateResource(input)
}
