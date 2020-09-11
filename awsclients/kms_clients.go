package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kms"
	"go.temporal.io/sdk/workflow"
)

type KMSClient interface {
       CancelKeyDeletion(ctx workflow.Context, input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error)
       CancelKeyDeletionAsync(ctx workflow.Context, input *kms.CancelKeyDeletionInput) *KmsCancelKeyDeletionResult

       ConnectCustomKeyStore(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error)
       ConnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) *KmsConnectCustomKeyStoreResult

       CreateAlias(ctx workflow.Context, input *kms.CreateAliasInput) (*kms.CreateAliasOutput, error)
       CreateAliasAsync(ctx workflow.Context, input *kms.CreateAliasInput) *KmsCreateAliasResult

       CreateCustomKeyStore(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error)
       CreateCustomKeyStoreAsync(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) *KmsCreateCustomKeyStoreResult

       CreateGrant(ctx workflow.Context, input *kms.CreateGrantInput) (*kms.CreateGrantOutput, error)
       CreateGrantAsync(ctx workflow.Context, input *kms.CreateGrantInput) *KmsCreateGrantResult

       CreateKey(ctx workflow.Context, input *kms.CreateKeyInput) (*kms.CreateKeyOutput, error)
       CreateKeyAsync(ctx workflow.Context, input *kms.CreateKeyInput) *KmsCreateKeyResult

       Decrypt(ctx workflow.Context, input *kms.DecryptInput) (*kms.DecryptOutput, error)
       DecryptAsync(ctx workflow.Context, input *kms.DecryptInput) *KmsDecryptResult

       DeleteAlias(ctx workflow.Context, input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error)
       DeleteAliasAsync(ctx workflow.Context, input *kms.DeleteAliasInput) *KmsDeleteAliasResult

       DeleteCustomKeyStore(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error)
       DeleteCustomKeyStoreAsync(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) *KmsDeleteCustomKeyStoreResult

       DeleteImportedKeyMaterial(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error)
       DeleteImportedKeyMaterialAsync(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) *KmsDeleteImportedKeyMaterialResult

       DescribeCustomKeyStores(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error)
       DescribeCustomKeyStoresAsync(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) *KmsDescribeCustomKeyStoresResult

       DescribeKey(ctx workflow.Context, input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)
       DescribeKeyAsync(ctx workflow.Context, input *kms.DescribeKeyInput) *KmsDescribeKeyResult

       DisableKey(ctx workflow.Context, input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error)
       DisableKeyAsync(ctx workflow.Context, input *kms.DisableKeyInput) *KmsDisableKeyResult

       DisableKeyRotation(ctx workflow.Context, input *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error)
       DisableKeyRotationAsync(ctx workflow.Context, input *kms.DisableKeyRotationInput) *KmsDisableKeyRotationResult

       DisconnectCustomKeyStore(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error)
       DisconnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) *KmsDisconnectCustomKeyStoreResult

       EnableKey(ctx workflow.Context, input *kms.EnableKeyInput) (*kms.EnableKeyOutput, error)
       EnableKeyAsync(ctx workflow.Context, input *kms.EnableKeyInput) *KmsEnableKeyResult

       EnableKeyRotation(ctx workflow.Context, input *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error)
       EnableKeyRotationAsync(ctx workflow.Context, input *kms.EnableKeyRotationInput) *KmsEnableKeyRotationResult

       Encrypt(ctx workflow.Context, input *kms.EncryptInput) (*kms.EncryptOutput, error)
       EncryptAsync(ctx workflow.Context, input *kms.EncryptInput) *KmsEncryptResult

       GenerateDataKey(ctx workflow.Context, input *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error)
       GenerateDataKeyAsync(ctx workflow.Context, input *kms.GenerateDataKeyInput) *KmsGenerateDataKeyResult

       GenerateDataKeyPair(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error)
       GenerateDataKeyPairAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) *KmsGenerateDataKeyPairResult

       GenerateDataKeyPairWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error)
       GenerateDataKeyPairWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) *KmsGenerateDataKeyPairWithoutPlaintextResult

       GenerateDataKeyWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)
       GenerateDataKeyWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) *KmsGenerateDataKeyWithoutPlaintextResult

       GenerateRandom(ctx workflow.Context, input *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error)
       GenerateRandomAsync(ctx workflow.Context, input *kms.GenerateRandomInput) *KmsGenerateRandomResult

       GetKeyPolicy(ctx workflow.Context, input *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error)
       GetKeyPolicyAsync(ctx workflow.Context, input *kms.GetKeyPolicyInput) *KmsGetKeyPolicyResult

       GetKeyRotationStatus(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error)
       GetKeyRotationStatusAsync(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) *KmsGetKeyRotationStatusResult

       GetParametersForImport(ctx workflow.Context, input *kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error)
       GetParametersForImportAsync(ctx workflow.Context, input *kms.GetParametersForImportInput) *KmsGetParametersForImportResult

       GetPublicKey(ctx workflow.Context, input *kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error)
       GetPublicKeyAsync(ctx workflow.Context, input *kms.GetPublicKeyInput) *KmsGetPublicKeyResult

       ImportKeyMaterial(ctx workflow.Context, input *kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error)
       ImportKeyMaterialAsync(ctx workflow.Context, input *kms.ImportKeyMaterialInput) *KmsImportKeyMaterialResult

       ListAliases(ctx workflow.Context, input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error)
       ListAliasesAsync(ctx workflow.Context, input *kms.ListAliasesInput) *KmsListAliasesResult

       ListGrants(ctx workflow.Context, input *kms.ListGrantsInput) (*kms.ListGrantsResponse, error)
       ListGrantsAsync(ctx workflow.Context, input *kms.ListGrantsInput) *KmsListGrantsResult

       ListKeyPolicies(ctx workflow.Context, input *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error)
       ListKeyPoliciesAsync(ctx workflow.Context, input *kms.ListKeyPoliciesInput) *KmsListKeyPoliciesResult

       ListKeys(ctx workflow.Context, input *kms.ListKeysInput) (*kms.ListKeysOutput, error)
       ListKeysAsync(ctx workflow.Context, input *kms.ListKeysInput) *KmsListKeysResult

       ListResourceTags(ctx workflow.Context, input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error)
       ListResourceTagsAsync(ctx workflow.Context, input *kms.ListResourceTagsInput) *KmsListResourceTagsResult

       ListRetirableGrants(ctx workflow.Context, input *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error)
       ListRetirableGrantsAsync(ctx workflow.Context, input *kms.ListRetirableGrantsInput) *KmsListRetirableGrantsResult

       PutKeyPolicy(ctx workflow.Context, input *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error)
       PutKeyPolicyAsync(ctx workflow.Context, input *kms.PutKeyPolicyInput) *KmsPutKeyPolicyResult

       ReEncrypt(ctx workflow.Context, input *kms.ReEncryptInput) (*kms.ReEncryptOutput, error)
       ReEncryptAsync(ctx workflow.Context, input *kms.ReEncryptInput) *KmsReEncryptResult

       RetireGrant(ctx workflow.Context, input *kms.RetireGrantInput) (*kms.RetireGrantOutput, error)
       RetireGrantAsync(ctx workflow.Context, input *kms.RetireGrantInput) *KmsRetireGrantResult

       RevokeGrant(ctx workflow.Context, input *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error)
       RevokeGrantAsync(ctx workflow.Context, input *kms.RevokeGrantInput) *KmsRevokeGrantResult

       ScheduleKeyDeletion(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error)
       ScheduleKeyDeletionAsync(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) *KmsScheduleKeyDeletionResult

       Sign(ctx workflow.Context, input *kms.SignInput) (*kms.SignOutput, error)
       SignAsync(ctx workflow.Context, input *kms.SignInput) *KmsSignResult

       TagResource(ctx workflow.Context, input *kms.TagResourceInput) (*kms.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *kms.TagResourceInput) *KmsTagResourceResult

       UntagResource(ctx workflow.Context, input *kms.UntagResourceInput) (*kms.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *kms.UntagResourceInput) *KmsUntagResourceResult

       UpdateAlias(ctx workflow.Context, input *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error)
       UpdateAliasAsync(ctx workflow.Context, input *kms.UpdateAliasInput) *KmsUpdateAliasResult

       UpdateCustomKeyStore(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error)
       UpdateCustomKeyStoreAsync(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) *KmsUpdateCustomKeyStoreResult

       UpdateKeyDescription(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error)
       UpdateKeyDescriptionAsync(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) *KmsUpdateKeyDescriptionResult

       Verify(ctx workflow.Context, input *kms.VerifyInput) (*kms.VerifyOutput, error)
       VerifyAsync(ctx workflow.Context, input *kms.VerifyInput) *KmsVerifyResult
}

type KMSStub struct {
}

func NewKMSStub() KMSClient {
    return &KMSStub{}
}

type KmsCancelKeyDeletionResult struct {
	Result workflow.Future
}

func (r *KmsCancelKeyDeletionResult) Get(ctx workflow.Context) (*kms.CancelKeyDeletionOutput, error) {
    var output kms.CancelKeyDeletionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsConnectCustomKeyStoreResult struct {
	Result workflow.Future
}

func (r *KmsConnectCustomKeyStoreResult) Get(ctx workflow.Context) (*kms.ConnectCustomKeyStoreOutput, error) {
    var output kms.ConnectCustomKeyStoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsCreateAliasResult struct {
	Result workflow.Future
}

func (r *KmsCreateAliasResult) Get(ctx workflow.Context) (*kms.CreateAliasOutput, error) {
    var output kms.CreateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsCreateCustomKeyStoreResult struct {
	Result workflow.Future
}

func (r *KmsCreateCustomKeyStoreResult) Get(ctx workflow.Context) (*kms.CreateCustomKeyStoreOutput, error) {
    var output kms.CreateCustomKeyStoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsCreateGrantResult struct {
	Result workflow.Future
}

func (r *KmsCreateGrantResult) Get(ctx workflow.Context) (*kms.CreateGrantOutput, error) {
    var output kms.CreateGrantOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsCreateKeyResult struct {
	Result workflow.Future
}

func (r *KmsCreateKeyResult) Get(ctx workflow.Context) (*kms.CreateKeyOutput, error) {
    var output kms.CreateKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDecryptResult struct {
	Result workflow.Future
}

func (r *KmsDecryptResult) Get(ctx workflow.Context) (*kms.DecryptOutput, error) {
    var output kms.DecryptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDeleteAliasResult struct {
	Result workflow.Future
}

func (r *KmsDeleteAliasResult) Get(ctx workflow.Context) (*kms.DeleteAliasOutput, error) {
    var output kms.DeleteAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDeleteCustomKeyStoreResult struct {
	Result workflow.Future
}

func (r *KmsDeleteCustomKeyStoreResult) Get(ctx workflow.Context) (*kms.DeleteCustomKeyStoreOutput, error) {
    var output kms.DeleteCustomKeyStoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDeleteImportedKeyMaterialResult struct {
	Result workflow.Future
}

func (r *KmsDeleteImportedKeyMaterialResult) Get(ctx workflow.Context) (*kms.DeleteImportedKeyMaterialOutput, error) {
    var output kms.DeleteImportedKeyMaterialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDescribeCustomKeyStoresResult struct {
	Result workflow.Future
}

func (r *KmsDescribeCustomKeyStoresResult) Get(ctx workflow.Context) (*kms.DescribeCustomKeyStoresOutput, error) {
    var output kms.DescribeCustomKeyStoresOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDescribeKeyResult struct {
	Result workflow.Future
}

func (r *KmsDescribeKeyResult) Get(ctx workflow.Context) (*kms.DescribeKeyOutput, error) {
    var output kms.DescribeKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDisableKeyResult struct {
	Result workflow.Future
}

func (r *KmsDisableKeyResult) Get(ctx workflow.Context) (*kms.DisableKeyOutput, error) {
    var output kms.DisableKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDisableKeyRotationResult struct {
	Result workflow.Future
}

func (r *KmsDisableKeyRotationResult) Get(ctx workflow.Context) (*kms.DisableKeyRotationOutput, error) {
    var output kms.DisableKeyRotationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsDisconnectCustomKeyStoreResult struct {
	Result workflow.Future
}

func (r *KmsDisconnectCustomKeyStoreResult) Get(ctx workflow.Context) (*kms.DisconnectCustomKeyStoreOutput, error) {
    var output kms.DisconnectCustomKeyStoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsEnableKeyResult struct {
	Result workflow.Future
}

func (r *KmsEnableKeyResult) Get(ctx workflow.Context) (*kms.EnableKeyOutput, error) {
    var output kms.EnableKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsEnableKeyRotationResult struct {
	Result workflow.Future
}

func (r *KmsEnableKeyRotationResult) Get(ctx workflow.Context) (*kms.EnableKeyRotationOutput, error) {
    var output kms.EnableKeyRotationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsEncryptResult struct {
	Result workflow.Future
}

func (r *KmsEncryptResult) Get(ctx workflow.Context) (*kms.EncryptOutput, error) {
    var output kms.EncryptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGenerateDataKeyResult struct {
	Result workflow.Future
}

func (r *KmsGenerateDataKeyResult) Get(ctx workflow.Context) (*kms.GenerateDataKeyOutput, error) {
    var output kms.GenerateDataKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGenerateDataKeyPairResult struct {
	Result workflow.Future
}

func (r *KmsGenerateDataKeyPairResult) Get(ctx workflow.Context) (*kms.GenerateDataKeyPairOutput, error) {
    var output kms.GenerateDataKeyPairOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGenerateDataKeyPairWithoutPlaintextResult struct {
	Result workflow.Future
}

func (r *KmsGenerateDataKeyPairWithoutPlaintextResult) Get(ctx workflow.Context) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error) {
    var output kms.GenerateDataKeyPairWithoutPlaintextOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGenerateDataKeyWithoutPlaintextResult struct {
	Result workflow.Future
}

func (r *KmsGenerateDataKeyWithoutPlaintextResult) Get(ctx workflow.Context) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
    var output kms.GenerateDataKeyWithoutPlaintextOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGenerateRandomResult struct {
	Result workflow.Future
}

func (r *KmsGenerateRandomResult) Get(ctx workflow.Context) (*kms.GenerateRandomOutput, error) {
    var output kms.GenerateRandomOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGetKeyPolicyResult struct {
	Result workflow.Future
}

func (r *KmsGetKeyPolicyResult) Get(ctx workflow.Context) (*kms.GetKeyPolicyOutput, error) {
    var output kms.GetKeyPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGetKeyRotationStatusResult struct {
	Result workflow.Future
}

func (r *KmsGetKeyRotationStatusResult) Get(ctx workflow.Context) (*kms.GetKeyRotationStatusOutput, error) {
    var output kms.GetKeyRotationStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGetParametersForImportResult struct {
	Result workflow.Future
}

func (r *KmsGetParametersForImportResult) Get(ctx workflow.Context) (*kms.GetParametersForImportOutput, error) {
    var output kms.GetParametersForImportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsGetPublicKeyResult struct {
	Result workflow.Future
}

func (r *KmsGetPublicKeyResult) Get(ctx workflow.Context) (*kms.GetPublicKeyOutput, error) {
    var output kms.GetPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsImportKeyMaterialResult struct {
	Result workflow.Future
}

func (r *KmsImportKeyMaterialResult) Get(ctx workflow.Context) (*kms.ImportKeyMaterialOutput, error) {
    var output kms.ImportKeyMaterialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsListAliasesResult struct {
	Result workflow.Future
}

func (r *KmsListAliasesResult) Get(ctx workflow.Context) (*kms.ListAliasesOutput, error) {
    var output kms.ListAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsListGrantsResult struct {
	Result workflow.Future
}

func (r *KmsListGrantsResult) Get(ctx workflow.Context) (*kms.ListGrantsResponse, error) {
    var output kms.ListGrantsResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsListKeyPoliciesResult struct {
	Result workflow.Future
}

func (r *KmsListKeyPoliciesResult) Get(ctx workflow.Context) (*kms.ListKeyPoliciesOutput, error) {
    var output kms.ListKeyPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsListKeysResult struct {
	Result workflow.Future
}

func (r *KmsListKeysResult) Get(ctx workflow.Context) (*kms.ListKeysOutput, error) {
    var output kms.ListKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsListResourceTagsResult struct {
	Result workflow.Future
}

func (r *KmsListResourceTagsResult) Get(ctx workflow.Context) (*kms.ListResourceTagsOutput, error) {
    var output kms.ListResourceTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsListRetirableGrantsResult struct {
	Result workflow.Future
}

func (r *KmsListRetirableGrantsResult) Get(ctx workflow.Context) (*kms.ListGrantsResponse, error) {
    var output kms.ListGrantsResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsPutKeyPolicyResult struct {
	Result workflow.Future
}

func (r *KmsPutKeyPolicyResult) Get(ctx workflow.Context) (*kms.PutKeyPolicyOutput, error) {
    var output kms.PutKeyPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsReEncryptResult struct {
	Result workflow.Future
}

func (r *KmsReEncryptResult) Get(ctx workflow.Context) (*kms.ReEncryptOutput, error) {
    var output kms.ReEncryptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsRetireGrantResult struct {
	Result workflow.Future
}

func (r *KmsRetireGrantResult) Get(ctx workflow.Context) (*kms.RetireGrantOutput, error) {
    var output kms.RetireGrantOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsRevokeGrantResult struct {
	Result workflow.Future
}

func (r *KmsRevokeGrantResult) Get(ctx workflow.Context) (*kms.RevokeGrantOutput, error) {
    var output kms.RevokeGrantOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsScheduleKeyDeletionResult struct {
	Result workflow.Future
}

func (r *KmsScheduleKeyDeletionResult) Get(ctx workflow.Context) (*kms.ScheduleKeyDeletionOutput, error) {
    var output kms.ScheduleKeyDeletionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsSignResult struct {
	Result workflow.Future
}

func (r *KmsSignResult) Get(ctx workflow.Context) (*kms.SignOutput, error) {
    var output kms.SignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsTagResourceResult struct {
	Result workflow.Future
}

func (r *KmsTagResourceResult) Get(ctx workflow.Context) (*kms.TagResourceOutput, error) {
    var output kms.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsUntagResourceResult struct {
	Result workflow.Future
}

func (r *KmsUntagResourceResult) Get(ctx workflow.Context) (*kms.UntagResourceOutput, error) {
    var output kms.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsUpdateAliasResult struct {
	Result workflow.Future
}

func (r *KmsUpdateAliasResult) Get(ctx workflow.Context) (*kms.UpdateAliasOutput, error) {
    var output kms.UpdateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsUpdateCustomKeyStoreResult struct {
	Result workflow.Future
}

func (r *KmsUpdateCustomKeyStoreResult) Get(ctx workflow.Context) (*kms.UpdateCustomKeyStoreOutput, error) {
    var output kms.UpdateCustomKeyStoreOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsUpdateKeyDescriptionResult struct {
	Result workflow.Future
}

func (r *KmsUpdateKeyDescriptionResult) Get(ctx workflow.Context) (*kms.UpdateKeyDescriptionOutput, error) {
    var output kms.UpdateKeyDescriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type KmsVerifyResult struct {
	Result workflow.Future
}

func (r *KmsVerifyResult) Get(ctx workflow.Context) (*kms.VerifyOutput, error) {
    var output kms.VerifyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) CancelKeyDeletion(ctx workflow.Context, input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error) {
    var output kms.CancelKeyDeletionOutput
    err := workflow.ExecuteActivity(ctx, "KMS.CancelKeyDeletion", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) CancelKeyDeletionAsync(ctx workflow.Context, input *kms.CancelKeyDeletionInput) *KmsCancelKeyDeletionResult {
    future := workflow.ExecuteActivity(ctx, "KMS.CancelKeyDeletion", input)
    return &KmsCancelKeyDeletionResult{Result: future}
}

func (a *KMSStub) ConnectCustomKeyStore(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error) {
    var output kms.ConnectCustomKeyStoreOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ConnectCustomKeyStore", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ConnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) *KmsConnectCustomKeyStoreResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ConnectCustomKeyStore", input)
    return &KmsConnectCustomKeyStoreResult{Result: future}
}

func (a *KMSStub) CreateAlias(ctx workflow.Context, input *kms.CreateAliasInput) (*kms.CreateAliasOutput, error) {
    var output kms.CreateAliasOutput
    err := workflow.ExecuteActivity(ctx, "KMS.CreateAlias", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) CreateAliasAsync(ctx workflow.Context, input *kms.CreateAliasInput) *KmsCreateAliasResult {
    future := workflow.ExecuteActivity(ctx, "KMS.CreateAlias", input)
    return &KmsCreateAliasResult{Result: future}
}

func (a *KMSStub) CreateCustomKeyStore(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error) {
    var output kms.CreateCustomKeyStoreOutput
    err := workflow.ExecuteActivity(ctx, "KMS.CreateCustomKeyStore", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) CreateCustomKeyStoreAsync(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) *KmsCreateCustomKeyStoreResult {
    future := workflow.ExecuteActivity(ctx, "KMS.CreateCustomKeyStore", input)
    return &KmsCreateCustomKeyStoreResult{Result: future}
}

func (a *KMSStub) CreateGrant(ctx workflow.Context, input *kms.CreateGrantInput) (*kms.CreateGrantOutput, error) {
    var output kms.CreateGrantOutput
    err := workflow.ExecuteActivity(ctx, "KMS.CreateGrant", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) CreateGrantAsync(ctx workflow.Context, input *kms.CreateGrantInput) *KmsCreateGrantResult {
    future := workflow.ExecuteActivity(ctx, "KMS.CreateGrant", input)
    return &KmsCreateGrantResult{Result: future}
}

func (a *KMSStub) CreateKey(ctx workflow.Context, input *kms.CreateKeyInput) (*kms.CreateKeyOutput, error) {
    var output kms.CreateKeyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.CreateKey", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) CreateKeyAsync(ctx workflow.Context, input *kms.CreateKeyInput) *KmsCreateKeyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.CreateKey", input)
    return &KmsCreateKeyResult{Result: future}
}

func (a *KMSStub) Decrypt(ctx workflow.Context, input *kms.DecryptInput) (*kms.DecryptOutput, error) {
    var output kms.DecryptOutput
    err := workflow.ExecuteActivity(ctx, "KMS.Decrypt", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DecryptAsync(ctx workflow.Context, input *kms.DecryptInput) *KmsDecryptResult {
    future := workflow.ExecuteActivity(ctx, "KMS.Decrypt", input)
    return &KmsDecryptResult{Result: future}
}

func (a *KMSStub) DeleteAlias(ctx workflow.Context, input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error) {
    var output kms.DeleteAliasOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DeleteAlias", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DeleteAliasAsync(ctx workflow.Context, input *kms.DeleteAliasInput) *KmsDeleteAliasResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DeleteAlias", input)
    return &KmsDeleteAliasResult{Result: future}
}

func (a *KMSStub) DeleteCustomKeyStore(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error) {
    var output kms.DeleteCustomKeyStoreOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DeleteCustomKeyStore", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DeleteCustomKeyStoreAsync(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) *KmsDeleteCustomKeyStoreResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DeleteCustomKeyStore", input)
    return &KmsDeleteCustomKeyStoreResult{Result: future}
}

func (a *KMSStub) DeleteImportedKeyMaterial(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error) {
    var output kms.DeleteImportedKeyMaterialOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DeleteImportedKeyMaterial", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DeleteImportedKeyMaterialAsync(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) *KmsDeleteImportedKeyMaterialResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DeleteImportedKeyMaterial", input)
    return &KmsDeleteImportedKeyMaterialResult{Result: future}
}

func (a *KMSStub) DescribeCustomKeyStores(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error) {
    var output kms.DescribeCustomKeyStoresOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DescribeCustomKeyStores", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DescribeCustomKeyStoresAsync(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) *KmsDescribeCustomKeyStoresResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DescribeCustomKeyStores", input)
    return &KmsDescribeCustomKeyStoresResult{Result: future}
}

func (a *KMSStub) DescribeKey(ctx workflow.Context, input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error) {
    var output kms.DescribeKeyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DescribeKey", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DescribeKeyAsync(ctx workflow.Context, input *kms.DescribeKeyInput) *KmsDescribeKeyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DescribeKey", input)
    return &KmsDescribeKeyResult{Result: future}
}

func (a *KMSStub) DisableKey(ctx workflow.Context, input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error) {
    var output kms.DisableKeyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DisableKey", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DisableKeyAsync(ctx workflow.Context, input *kms.DisableKeyInput) *KmsDisableKeyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DisableKey", input)
    return &KmsDisableKeyResult{Result: future}
}

func (a *KMSStub) DisableKeyRotation(ctx workflow.Context, input *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error) {
    var output kms.DisableKeyRotationOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DisableKeyRotation", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DisableKeyRotationAsync(ctx workflow.Context, input *kms.DisableKeyRotationInput) *KmsDisableKeyRotationResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DisableKeyRotation", input)
    return &KmsDisableKeyRotationResult{Result: future}
}

func (a *KMSStub) DisconnectCustomKeyStore(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error) {
    var output kms.DisconnectCustomKeyStoreOutput
    err := workflow.ExecuteActivity(ctx, "KMS.DisconnectCustomKeyStore", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) DisconnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) *KmsDisconnectCustomKeyStoreResult {
    future := workflow.ExecuteActivity(ctx, "KMS.DisconnectCustomKeyStore", input)
    return &KmsDisconnectCustomKeyStoreResult{Result: future}
}

func (a *KMSStub) EnableKey(ctx workflow.Context, input *kms.EnableKeyInput) (*kms.EnableKeyOutput, error) {
    var output kms.EnableKeyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.EnableKey", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) EnableKeyAsync(ctx workflow.Context, input *kms.EnableKeyInput) *KmsEnableKeyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.EnableKey", input)
    return &KmsEnableKeyResult{Result: future}
}

func (a *KMSStub) EnableKeyRotation(ctx workflow.Context, input *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error) {
    var output kms.EnableKeyRotationOutput
    err := workflow.ExecuteActivity(ctx, "KMS.EnableKeyRotation", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) EnableKeyRotationAsync(ctx workflow.Context, input *kms.EnableKeyRotationInput) *KmsEnableKeyRotationResult {
    future := workflow.ExecuteActivity(ctx, "KMS.EnableKeyRotation", input)
    return &KmsEnableKeyRotationResult{Result: future}
}

func (a *KMSStub) Encrypt(ctx workflow.Context, input *kms.EncryptInput) (*kms.EncryptOutput, error) {
    var output kms.EncryptOutput
    err := workflow.ExecuteActivity(ctx, "KMS.Encrypt", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) EncryptAsync(ctx workflow.Context, input *kms.EncryptInput) *KmsEncryptResult {
    future := workflow.ExecuteActivity(ctx, "KMS.Encrypt", input)
    return &KmsEncryptResult{Result: future}
}

func (a *KMSStub) GenerateDataKey(ctx workflow.Context, input *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error) {
    var output kms.GenerateDataKeyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKey", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GenerateDataKeyAsync(ctx workflow.Context, input *kms.GenerateDataKeyInput) *KmsGenerateDataKeyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKey", input)
    return &KmsGenerateDataKeyResult{Result: future}
}

func (a *KMSStub) GenerateDataKeyPair(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error) {
    var output kms.GenerateDataKeyPairOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKeyPair", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GenerateDataKeyPairAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) *KmsGenerateDataKeyPairResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKeyPair", input)
    return &KmsGenerateDataKeyPairResult{Result: future}
}

func (a *KMSStub) GenerateDataKeyPairWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error) {
    var output kms.GenerateDataKeyPairWithoutPlaintextOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKeyPairWithoutPlaintext", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GenerateDataKeyPairWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) *KmsGenerateDataKeyPairWithoutPlaintextResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKeyPairWithoutPlaintext", input)
    return &KmsGenerateDataKeyPairWithoutPlaintextResult{Result: future}
}

func (a *KMSStub) GenerateDataKeyWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
    var output kms.GenerateDataKeyWithoutPlaintextOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKeyWithoutPlaintext", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GenerateDataKeyWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) *KmsGenerateDataKeyWithoutPlaintextResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GenerateDataKeyWithoutPlaintext", input)
    return &KmsGenerateDataKeyWithoutPlaintextResult{Result: future}
}

func (a *KMSStub) GenerateRandom(ctx workflow.Context, input *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error) {
    var output kms.GenerateRandomOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GenerateRandom", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GenerateRandomAsync(ctx workflow.Context, input *kms.GenerateRandomInput) *KmsGenerateRandomResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GenerateRandom", input)
    return &KmsGenerateRandomResult{Result: future}
}

func (a *KMSStub) GetKeyPolicy(ctx workflow.Context, input *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error) {
    var output kms.GetKeyPolicyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GetKeyPolicy", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GetKeyPolicyAsync(ctx workflow.Context, input *kms.GetKeyPolicyInput) *KmsGetKeyPolicyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GetKeyPolicy", input)
    return &KmsGetKeyPolicyResult{Result: future}
}

func (a *KMSStub) GetKeyRotationStatus(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error) {
    var output kms.GetKeyRotationStatusOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GetKeyRotationStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GetKeyRotationStatusAsync(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) *KmsGetKeyRotationStatusResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GetKeyRotationStatus", input)
    return &KmsGetKeyRotationStatusResult{Result: future}
}

func (a *KMSStub) GetParametersForImport(ctx workflow.Context, input *kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error) {
    var output kms.GetParametersForImportOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GetParametersForImport", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GetParametersForImportAsync(ctx workflow.Context, input *kms.GetParametersForImportInput) *KmsGetParametersForImportResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GetParametersForImport", input)
    return &KmsGetParametersForImportResult{Result: future}
}

func (a *KMSStub) GetPublicKey(ctx workflow.Context, input *kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error) {
    var output kms.GetPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.GetPublicKey", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) GetPublicKeyAsync(ctx workflow.Context, input *kms.GetPublicKeyInput) *KmsGetPublicKeyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.GetPublicKey", input)
    return &KmsGetPublicKeyResult{Result: future}
}

func (a *KMSStub) ImportKeyMaterial(ctx workflow.Context, input *kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error) {
    var output kms.ImportKeyMaterialOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ImportKeyMaterial", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ImportKeyMaterialAsync(ctx workflow.Context, input *kms.ImportKeyMaterialInput) *KmsImportKeyMaterialResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ImportKeyMaterial", input)
    return &KmsImportKeyMaterialResult{Result: future}
}

func (a *KMSStub) ListAliases(ctx workflow.Context, input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error) {
    var output kms.ListAliasesOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ListAliases", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ListAliasesAsync(ctx workflow.Context, input *kms.ListAliasesInput) *KmsListAliasesResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ListAliases", input)
    return &KmsListAliasesResult{Result: future}
}

func (a *KMSStub) ListGrants(ctx workflow.Context, input *kms.ListGrantsInput) (*kms.ListGrantsResponse, error) {
    var output kms.ListGrantsResponse
    err := workflow.ExecuteActivity(ctx, "KMS.ListGrants", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ListGrantsAsync(ctx workflow.Context, input *kms.ListGrantsInput) *KmsListGrantsResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ListGrants", input)
    return &KmsListGrantsResult{Result: future}
}

func (a *KMSStub) ListKeyPolicies(ctx workflow.Context, input *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error) {
    var output kms.ListKeyPoliciesOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ListKeyPolicies", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ListKeyPoliciesAsync(ctx workflow.Context, input *kms.ListKeyPoliciesInput) *KmsListKeyPoliciesResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ListKeyPolicies", input)
    return &KmsListKeyPoliciesResult{Result: future}
}

func (a *KMSStub) ListKeys(ctx workflow.Context, input *kms.ListKeysInput) (*kms.ListKeysOutput, error) {
    var output kms.ListKeysOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ListKeys", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ListKeysAsync(ctx workflow.Context, input *kms.ListKeysInput) *KmsListKeysResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ListKeys", input)
    return &KmsListKeysResult{Result: future}
}

func (a *KMSStub) ListResourceTags(ctx workflow.Context, input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error) {
    var output kms.ListResourceTagsOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ListResourceTags", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ListResourceTagsAsync(ctx workflow.Context, input *kms.ListResourceTagsInput) *KmsListResourceTagsResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ListResourceTags", input)
    return &KmsListResourceTagsResult{Result: future}
}

func (a *KMSStub) ListRetirableGrants(ctx workflow.Context, input *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error) {
    var output kms.ListGrantsResponse
    err := workflow.ExecuteActivity(ctx, "KMS.ListRetirableGrants", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ListRetirableGrantsAsync(ctx workflow.Context, input *kms.ListRetirableGrantsInput) *KmsListRetirableGrantsResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ListRetirableGrants", input)
    return &KmsListRetirableGrantsResult{Result: future}
}

func (a *KMSStub) PutKeyPolicy(ctx workflow.Context, input *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error) {
    var output kms.PutKeyPolicyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.PutKeyPolicy", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) PutKeyPolicyAsync(ctx workflow.Context, input *kms.PutKeyPolicyInput) *KmsPutKeyPolicyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.PutKeyPolicy", input)
    return &KmsPutKeyPolicyResult{Result: future}
}

func (a *KMSStub) ReEncrypt(ctx workflow.Context, input *kms.ReEncryptInput) (*kms.ReEncryptOutput, error) {
    var output kms.ReEncryptOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ReEncrypt", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ReEncryptAsync(ctx workflow.Context, input *kms.ReEncryptInput) *KmsReEncryptResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ReEncrypt", input)
    return &KmsReEncryptResult{Result: future}
}

func (a *KMSStub) RetireGrant(ctx workflow.Context, input *kms.RetireGrantInput) (*kms.RetireGrantOutput, error) {
    var output kms.RetireGrantOutput
    err := workflow.ExecuteActivity(ctx, "KMS.RetireGrant", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) RetireGrantAsync(ctx workflow.Context, input *kms.RetireGrantInput) *KmsRetireGrantResult {
    future := workflow.ExecuteActivity(ctx, "KMS.RetireGrant", input)
    return &KmsRetireGrantResult{Result: future}
}

func (a *KMSStub) RevokeGrant(ctx workflow.Context, input *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error) {
    var output kms.RevokeGrantOutput
    err := workflow.ExecuteActivity(ctx, "KMS.RevokeGrant", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) RevokeGrantAsync(ctx workflow.Context, input *kms.RevokeGrantInput) *KmsRevokeGrantResult {
    future := workflow.ExecuteActivity(ctx, "KMS.RevokeGrant", input)
    return &KmsRevokeGrantResult{Result: future}
}

func (a *KMSStub) ScheduleKeyDeletion(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error) {
    var output kms.ScheduleKeyDeletionOutput
    err := workflow.ExecuteActivity(ctx, "KMS.ScheduleKeyDeletion", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) ScheduleKeyDeletionAsync(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) *KmsScheduleKeyDeletionResult {
    future := workflow.ExecuteActivity(ctx, "KMS.ScheduleKeyDeletion", input)
    return &KmsScheduleKeyDeletionResult{Result: future}
}

func (a *KMSStub) Sign(ctx workflow.Context, input *kms.SignInput) (*kms.SignOutput, error) {
    var output kms.SignOutput
    err := workflow.ExecuteActivity(ctx, "KMS.Sign", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) SignAsync(ctx workflow.Context, input *kms.SignInput) *KmsSignResult {
    future := workflow.ExecuteActivity(ctx, "KMS.Sign", input)
    return &KmsSignResult{Result: future}
}

func (a *KMSStub) TagResource(ctx workflow.Context, input *kms.TagResourceInput) (*kms.TagResourceOutput, error) {
    var output kms.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "KMS.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) TagResourceAsync(ctx workflow.Context, input *kms.TagResourceInput) *KmsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "KMS.TagResource", input)
    return &KmsTagResourceResult{Result: future}
}

func (a *KMSStub) UntagResource(ctx workflow.Context, input *kms.UntagResourceInput) (*kms.UntagResourceOutput, error) {
    var output kms.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "KMS.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) UntagResourceAsync(ctx workflow.Context, input *kms.UntagResourceInput) *KmsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "KMS.UntagResource", input)
    return &KmsUntagResourceResult{Result: future}
}

func (a *KMSStub) UpdateAlias(ctx workflow.Context, input *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error) {
    var output kms.UpdateAliasOutput
    err := workflow.ExecuteActivity(ctx, "KMS.UpdateAlias", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) UpdateAliasAsync(ctx workflow.Context, input *kms.UpdateAliasInput) *KmsUpdateAliasResult {
    future := workflow.ExecuteActivity(ctx, "KMS.UpdateAlias", input)
    return &KmsUpdateAliasResult{Result: future}
}

func (a *KMSStub) UpdateCustomKeyStore(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error) {
    var output kms.UpdateCustomKeyStoreOutput
    err := workflow.ExecuteActivity(ctx, "KMS.UpdateCustomKeyStore", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) UpdateCustomKeyStoreAsync(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) *KmsUpdateCustomKeyStoreResult {
    future := workflow.ExecuteActivity(ctx, "KMS.UpdateCustomKeyStore", input)
    return &KmsUpdateCustomKeyStoreResult{Result: future}
}

func (a *KMSStub) UpdateKeyDescription(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error) {
    var output kms.UpdateKeyDescriptionOutput
    err := workflow.ExecuteActivity(ctx, "KMS.UpdateKeyDescription", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) UpdateKeyDescriptionAsync(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) *KmsUpdateKeyDescriptionResult {
    future := workflow.ExecuteActivity(ctx, "KMS.UpdateKeyDescription", input)
    return &KmsUpdateKeyDescriptionResult{Result: future}
}

func (a *KMSStub) Verify(ctx workflow.Context, input *kms.VerifyInput) (*kms.VerifyOutput, error) {
    var output kms.VerifyOutput
    err := workflow.ExecuteActivity(ctx, "KMS.Verify", input).Get(ctx, &output)
    return &output, err
}

func (a *KMSStub) VerifyAsync(ctx workflow.Context, input *kms.VerifyInput) *KmsVerifyResult {
    future := workflow.ExecuteActivity(ctx, "KMS.Verify", input)
    return &KmsVerifyResult{Result: future}
}
