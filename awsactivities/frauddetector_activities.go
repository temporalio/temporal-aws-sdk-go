package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"github.com/aws/aws-sdk-go/service/frauddetector/frauddetectoriface"
)

type FraudDetectorActivities struct {
	client frauddetectoriface.FraudDetectorAPI
}

func NewFraudDetectorActivities(client frauddetectoriface.FraudDetectorAPI) *FraudDetectorActivities {
    return &FraudDetectorActivities{client: client}
}

func (a *FraudDetectorActivities) BatchCreateVariable(input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
    return a.client.BatchCreateVariable(input)
}

func (a *FraudDetectorActivities) BatchGetVariable(input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error) {
    return a.client.BatchGetVariable(input)
}

func (a *FraudDetectorActivities) CreateDetectorVersion(input *frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error) {
    return a.client.CreateDetectorVersion(input)
}

func (a *FraudDetectorActivities) CreateModel(input *frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error) {
    return a.client.CreateModel(input)
}

func (a *FraudDetectorActivities) CreateModelVersion(input *frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error) {
    return a.client.CreateModelVersion(input)
}

func (a *FraudDetectorActivities) CreateRule(input *frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error) {
    return a.client.CreateRule(input)
}

func (a *FraudDetectorActivities) CreateVariable(input *frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error) {
    return a.client.CreateVariable(input)
}

func (a *FraudDetectorActivities) DeleteDetector(input *frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error) {
    return a.client.DeleteDetector(input)
}

func (a *FraudDetectorActivities) DeleteDetectorVersion(input *frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error) {
    return a.client.DeleteDetectorVersion(input)
}

func (a *FraudDetectorActivities) DeleteEvent(input *frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error) {
    return a.client.DeleteEvent(input)
}

func (a *FraudDetectorActivities) DeleteRule(input *frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error) {
    return a.client.DeleteRule(input)
}

func (a *FraudDetectorActivities) DescribeDetector(input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error) {
    return a.client.DescribeDetector(input)
}

func (a *FraudDetectorActivities) DescribeModelVersions(input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error) {
    return a.client.DescribeModelVersions(input)
}

func (a *FraudDetectorActivities) GetDetectorVersion(input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error) {
    return a.client.GetDetectorVersion(input)
}

func (a *FraudDetectorActivities) GetDetectors(input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error) {
    return a.client.GetDetectors(input)
}

func (a *FraudDetectorActivities) GetEntityTypes(input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error) {
    return a.client.GetEntityTypes(input)
}

func (a *FraudDetectorActivities) GetEventPrediction(input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error) {
    return a.client.GetEventPrediction(input)
}

func (a *FraudDetectorActivities) GetEventTypes(input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error) {
    return a.client.GetEventTypes(input)
}

func (a *FraudDetectorActivities) GetExternalModels(input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error) {
    return a.client.GetExternalModels(input)
}

func (a *FraudDetectorActivities) GetKMSEncryptionKey(input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
    return a.client.GetKMSEncryptionKey(input)
}

func (a *FraudDetectorActivities) GetLabels(input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error) {
    return a.client.GetLabels(input)
}

func (a *FraudDetectorActivities) GetModelVersion(input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error) {
    return a.client.GetModelVersion(input)
}

func (a *FraudDetectorActivities) GetModels(input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error) {
    return a.client.GetModels(input)
}

func (a *FraudDetectorActivities) GetOutcomes(input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error) {
    return a.client.GetOutcomes(input)
}

func (a *FraudDetectorActivities) GetRules(input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error) {
    return a.client.GetRules(input)
}

func (a *FraudDetectorActivities) GetVariables(input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error) {
    return a.client.GetVariables(input)
}

func (a *FraudDetectorActivities) ListTagsForResource(input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *FraudDetectorActivities) PutDetector(input *frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error) {
    return a.client.PutDetector(input)
}

func (a *FraudDetectorActivities) PutEntityType(input *frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error) {
    return a.client.PutEntityType(input)
}

func (a *FraudDetectorActivities) PutEventType(input *frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error) {
    return a.client.PutEventType(input)
}

func (a *FraudDetectorActivities) PutExternalModel(input *frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error) {
    return a.client.PutExternalModel(input)
}

func (a *FraudDetectorActivities) PutKMSEncryptionKey(input *frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error) {
    return a.client.PutKMSEncryptionKey(input)
}

func (a *FraudDetectorActivities) PutLabel(input *frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error) {
    return a.client.PutLabel(input)
}

func (a *FraudDetectorActivities) PutOutcome(input *frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error) {
    return a.client.PutOutcome(input)
}

func (a *FraudDetectorActivities) TagResource(input *frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *FraudDetectorActivities) UntagResource(input *frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersion(input *frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error) {
    return a.client.UpdateDetectorVersion(input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersionMetadata(input *frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error) {
    return a.client.UpdateDetectorVersionMetadata(input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersionStatus(input *frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error) {
    return a.client.UpdateDetectorVersionStatus(input)
}

func (a *FraudDetectorActivities) UpdateModel(input *frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error) {
    return a.client.UpdateModel(input)
}

func (a *FraudDetectorActivities) UpdateModelVersion(input *frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error) {
    return a.client.UpdateModelVersion(input)
}

func (a *FraudDetectorActivities) UpdateModelVersionStatus(input *frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error) {
    return a.client.UpdateModelVersionStatus(input)
}

func (a *FraudDetectorActivities) UpdateRuleMetadata(input *frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error) {
    return a.client.UpdateRuleMetadata(input)
}

func (a *FraudDetectorActivities) UpdateRuleVersion(input *frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error) {
    return a.client.UpdateRuleVersion(input)
}

func (a *FraudDetectorActivities) UpdateVariable(input *frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error) {
    return a.client.UpdateVariable(input)
}
