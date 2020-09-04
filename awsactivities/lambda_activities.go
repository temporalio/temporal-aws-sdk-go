
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

type LambdaActivities struct {
	client lambdaiface.LambdaAPI
}

func NewLambdaActivities(client lambdaiface.LambdaAPI) *LambdaActivities {
    return &LambdaActivities{client: client}
}

func (a *LambdaActivities) AddLayerVersionPermission(input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error) {
    return a.client.AddLayerVersionPermission(input)
}

func (a *LambdaActivities) AddPermission(input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {
    return a.client.AddPermission(input)
}

func (a *LambdaActivities) CreateAlias(input *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error) {
    return a.client.CreateAlias(input)
}

func (a *LambdaActivities) CreateEventSourceMapping(input *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    return a.client.CreateEventSourceMapping(input)
}

func (a *LambdaActivities) CreateFunction(input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error) {
    return a.client.CreateFunction(input)
}

func (a *LambdaActivities) DeleteAlias(input *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error) {
    return a.client.DeleteAlias(input)
}

func (a *LambdaActivities) DeleteEventSourceMapping(input *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    return a.client.DeleteEventSourceMapping(input)
}

func (a *LambdaActivities) DeleteFunction(input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error) {
    return a.client.DeleteFunction(input)
}

func (a *LambdaActivities) DeleteFunctionConcurrency(input *lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error) {
    return a.client.DeleteFunctionConcurrency(input)
}

func (a *LambdaActivities) DeleteFunctionEventInvokeConfig(input *lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error) {
    return a.client.DeleteFunctionEventInvokeConfig(input)
}

func (a *LambdaActivities) DeleteLayerVersion(input *lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error) {
    return a.client.DeleteLayerVersion(input)
}

func (a *LambdaActivities) DeleteProvisionedConcurrencyConfig(input *lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error) {
    return a.client.DeleteProvisionedConcurrencyConfig(input)
}

func (a *LambdaActivities) GetAccountSettings(input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error) {
    return a.client.GetAccountSettings(input)
}

func (a *LambdaActivities) GetAlias(input *lambda.GetAliasInput) (*lambda.AliasConfiguration, error) {
    return a.client.GetAlias(input)
}

func (a *LambdaActivities) GetEventSourceMapping(input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    return a.client.GetEventSourceMapping(input)
}

func (a *LambdaActivities) GetFunction(input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
    return a.client.GetFunction(input)
}

func (a *LambdaActivities) GetFunctionConcurrency(input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error) {
    return a.client.GetFunctionConcurrency(input)
}

func (a *LambdaActivities) GetFunctionConfiguration(input *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
    return a.client.GetFunctionConfiguration(input)
}

func (a *LambdaActivities) GetFunctionEventInvokeConfig(input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
    return a.client.GetFunctionEventInvokeConfig(input)
}

func (a *LambdaActivities) GetLayerVersion(input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error) {
    return a.client.GetLayerVersion(input)
}

func (a *LambdaActivities) GetLayerVersionByArn(input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error) {
    return a.client.GetLayerVersionByArn(input)
}

func (a *LambdaActivities) GetLayerVersionPolicy(input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error) {
    return a.client.GetLayerVersionPolicy(input)
}

func (a *LambdaActivities) GetPolicy(input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error) {
    return a.client.GetPolicy(input)
}

func (a *LambdaActivities) GetProvisionedConcurrencyConfig(input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
    return a.client.GetProvisionedConcurrencyConfig(input)
}

func (a *LambdaActivities) Invoke(input *lambda.InvokeInput) (*lambda.InvokeOutput, error) {
    return a.client.Invoke(input)
}

func (a *LambdaActivities) ListAliases(input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error) {
    return a.client.ListAliases(input)
}

func (a *LambdaActivities) ListEventSourceMappings(input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error) {
    return a.client.ListEventSourceMappings(input)
}

func (a *LambdaActivities) ListFunctionEventInvokeConfigs(input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
    return a.client.ListFunctionEventInvokeConfigs(input)
}

func (a *LambdaActivities) ListFunctions(input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error) {
    return a.client.ListFunctions(input)
}

func (a *LambdaActivities) ListLayerVersions(input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error) {
    return a.client.ListLayerVersions(input)
}

func (a *LambdaActivities) ListLayers(input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error) {
    return a.client.ListLayers(input)
}

func (a *LambdaActivities) ListProvisionedConcurrencyConfigs(input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
    return a.client.ListProvisionedConcurrencyConfigs(input)
}

func (a *LambdaActivities) ListTags(input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *LambdaActivities) ListVersionsByFunction(input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error) {
    return a.client.ListVersionsByFunction(input)
}

func (a *LambdaActivities) PublishLayerVersion(input *lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error) {
    return a.client.PublishLayerVersion(input)
}

func (a *LambdaActivities) PublishVersion(input *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error) {
    return a.client.PublishVersion(input)
}

func (a *LambdaActivities) PutFunctionConcurrency(input *lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error) {
    return a.client.PutFunctionConcurrency(input)
}

func (a *LambdaActivities) PutFunctionEventInvokeConfig(input *lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error) {
    return a.client.PutFunctionEventInvokeConfig(input)
}

func (a *LambdaActivities) PutProvisionedConcurrencyConfig(input *lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error) {
    return a.client.PutProvisionedConcurrencyConfig(input)
}

func (a *LambdaActivities) RemoveLayerVersionPermission(input *lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error) {
    return a.client.RemoveLayerVersionPermission(input)
}

func (a *LambdaActivities) RemovePermission(input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error) {
    return a.client.RemovePermission(input)
}

func (a *LambdaActivities) TagResource(input *lambda.TagResourceInput) (*lambda.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *LambdaActivities) UntagResource(input *lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *LambdaActivities) UpdateAlias(input *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error) {
    return a.client.UpdateAlias(input)
}

func (a *LambdaActivities) UpdateEventSourceMapping(input *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    return a.client.UpdateEventSourceMapping(input)
}

func (a *LambdaActivities) UpdateFunctionCode(input *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error) {
    return a.client.UpdateFunctionCode(input)
}

func (a *LambdaActivities) UpdateFunctionConfiguration(input *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
    return a.client.UpdateFunctionConfiguration(input)
}

func (a *LambdaActivities) UpdateFunctionEventInvokeConfig(input *lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error) {
    return a.client.UpdateFunctionEventInvokeConfig(input)
}

func (a *LambdaActivities) WaitUntilFunctionActive(input *lambda.GetFunctionConfigurationInput) error {
    return a.client.WaitUntilFunctionActive(input)
}

func (a *LambdaActivities) WaitUntilFunctionExists(input *lambda.GetFunctionInput) error {
    return a.client.WaitUntilFunctionExists(input)
}

func (a *LambdaActivities) WaitUntilFunctionUpdated(input *lambda.GetFunctionConfigurationInput) error {
    return a.client.WaitUntilFunctionUpdated(input)
}
