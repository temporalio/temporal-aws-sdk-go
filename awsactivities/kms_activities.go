package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

type KMSActivities struct {
	client kmsiface.KMSAPI
}

func NewKMSActivities(client kmsiface.KMSAPI) *KMSActivities {
    return &KMSActivities{client: client}
}


func (a *KMSActivities) CancelKeyDeletion(input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error) {
    return a.client.CancelKeyDeletion(input)
}



func (a *KMSActivities) ConnectCustomKeyStore(input *kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error) {
    return a.client.ConnectCustomKeyStore(input)
}



func (a *KMSActivities) CreateAlias(input *kms.CreateAliasInput) (*kms.CreateAliasOutput, error) {
    return a.client.CreateAlias(input)
}



func (a *KMSActivities) CreateCustomKeyStore(input *kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error) {
    return a.client.CreateCustomKeyStore(input)
}



func (a *KMSActivities) CreateGrant(input *kms.CreateGrantInput) (*kms.CreateGrantOutput, error) {
    return a.client.CreateGrant(input)
}



func (a *KMSActivities) CreateKey(input *kms.CreateKeyInput) (*kms.CreateKeyOutput, error) {
    return a.client.CreateKey(input)
}



func (a *KMSActivities) Decrypt(input *kms.DecryptInput) (*kms.DecryptOutput, error) {
    return a.client.Decrypt(input)
}



func (a *KMSActivities) DeleteAlias(input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error) {
    return a.client.DeleteAlias(input)
}



func (a *KMSActivities) DeleteCustomKeyStore(input *kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error) {
    return a.client.DeleteCustomKeyStore(input)
}



func (a *KMSActivities) DeleteImportedKeyMaterial(input *kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error) {
    return a.client.DeleteImportedKeyMaterial(input)
}



func (a *KMSActivities) DescribeCustomKeyStores(input *kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error) {
    return a.client.DescribeCustomKeyStores(input)
}



func (a *KMSActivities) DescribeKey(input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error) {
    return a.client.DescribeKey(input)
}



func (a *KMSActivities) DisableKey(input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error) {
    return a.client.DisableKey(input)
}



func (a *KMSActivities) DisableKeyRotation(input *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error) {
    return a.client.DisableKeyRotation(input)
}



func (a *KMSActivities) DisconnectCustomKeyStore(input *kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error) {
    return a.client.DisconnectCustomKeyStore(input)
}



func (a *KMSActivities) EnableKey(input *kms.EnableKeyInput) (*kms.EnableKeyOutput, error) {
    return a.client.EnableKey(input)
}



func (a *KMSActivities) EnableKeyRotation(input *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error) {
    return a.client.EnableKeyRotation(input)
}



func (a *KMSActivities) Encrypt(input *kms.EncryptInput) (*kms.EncryptOutput, error) {
    return a.client.Encrypt(input)
}



func (a *KMSActivities) GenerateDataKey(input *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error) {
    return a.client.GenerateDataKey(input)
}



func (a *KMSActivities) GenerateDataKeyPair(input *kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error) {
    return a.client.GenerateDataKeyPair(input)
}



func (a *KMSActivities) GenerateDataKeyPairWithoutPlaintext(input *kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error) {
    return a.client.GenerateDataKeyPairWithoutPlaintext(input)
}



func (a *KMSActivities) GenerateDataKeyWithoutPlaintext(input *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
    return a.client.GenerateDataKeyWithoutPlaintext(input)
}



func (a *KMSActivities) GenerateRandom(input *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error) {
    return a.client.GenerateRandom(input)
}



func (a *KMSActivities) GetKeyPolicy(input *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error) {
    return a.client.GetKeyPolicy(input)
}



func (a *KMSActivities) GetKeyRotationStatus(input *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error) {
    return a.client.GetKeyRotationStatus(input)
}



func (a *KMSActivities) GetParametersForImport(input *kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error) {
    return a.client.GetParametersForImport(input)
}



func (a *KMSActivities) GetPublicKey(input *kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error) {
    return a.client.GetPublicKey(input)
}



func (a *KMSActivities) ImportKeyMaterial(input *kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error) {
    return a.client.ImportKeyMaterial(input)
}



func (a *KMSActivities) ListAliases(input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error) {
    return a.client.ListAliases(input)
}



func (a *KMSActivities) ListGrants(input *kms.ListGrantsInput) (*kms.ListGrantsResponse, error) {
    return a.client.ListGrants(input)
}



func (a *KMSActivities) ListKeyPolicies(input *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error) {
    return a.client.ListKeyPolicies(input)
}



func (a *KMSActivities) ListKeys(input *kms.ListKeysInput) (*kms.ListKeysOutput, error) {
    return a.client.ListKeys(input)
}



func (a *KMSActivities) ListResourceTags(input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error) {
    return a.client.ListResourceTags(input)
}



func (a *KMSActivities) ListRetirableGrants(input *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error) {
    return a.client.ListRetirableGrants(input)
}



func (a *KMSActivities) PutKeyPolicy(input *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error) {
    return a.client.PutKeyPolicy(input)
}



func (a *KMSActivities) ReEncrypt(input *kms.ReEncryptInput) (*kms.ReEncryptOutput, error) {
    return a.client.ReEncrypt(input)
}



func (a *KMSActivities) RetireGrant(input *kms.RetireGrantInput) (*kms.RetireGrantOutput, error) {
    return a.client.RetireGrant(input)
}



func (a *KMSActivities) RevokeGrant(input *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error) {
    return a.client.RevokeGrant(input)
}



func (a *KMSActivities) ScheduleKeyDeletion(input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error) {
    return a.client.ScheduleKeyDeletion(input)
}



func (a *KMSActivities) Sign(input *kms.SignInput) (*kms.SignOutput, error) {
    return a.client.Sign(input)
}



func (a *KMSActivities) TagResource(input *kms.TagResourceInput) (*kms.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *KMSActivities) UntagResource(input *kms.UntagResourceInput) (*kms.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *KMSActivities) UpdateAlias(input *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error) {
    return a.client.UpdateAlias(input)
}



func (a *KMSActivities) UpdateCustomKeyStore(input *kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error) {
    return a.client.UpdateCustomKeyStore(input)
}



func (a *KMSActivities) UpdateKeyDescription(input *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error) {
    return a.client.UpdateKeyDescription(input)
}



func (a *KMSActivities) Verify(input *kms.VerifyInput) (*kms.VerifyOutput, error) {
    return a.client.Verify(input)
}

