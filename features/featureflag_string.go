// Code generated by "stringer -type=FeatureFlag"; DO NOT EDIT.

package features

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[unused-0]
	_ = x[WriteIssuedNamesPrecert-1]
	_ = x[CAAValidationMethods-2]
	_ = x[CAAAccountURI-3]
	_ = x[HeadNonceStatusOK-4]
	_ = x[EnforceMultiVA-5]
	_ = x[MultiVAFullResults-6]
	_ = x[RemoveWFE2AccountID-7]
	_ = x[CheckRenewalFirst-8]
	_ = x[MandatoryPOSTAsGET-9]
	_ = x[AllowV1Registration-10]
	_ = x[ParallelCheckFailedValidation-11]
	_ = x[DeleteUnusedChallenges-12]
	_ = x[V1DisableNewValidations-13]
	_ = x[PrecertificateRevocation-14]
	_ = x[StripDefaultSchemePort-15]
	_ = x[StoreIssuerInfo-16]
}

const _FeatureFlag_name = "unusedWriteIssuedNamesPrecertCAAValidationMethodsCAAAccountURIHeadNonceStatusOKEnforceMultiVAMultiVAFullResultsRemoveWFE2AccountIDCheckRenewalFirstMandatoryPOSTAsGETAllowV1RegistrationParallelCheckFailedValidationDeleteUnusedChallengesV1DisableNewValidationsPrecertificateRevocationStripDefaultSchemePortStoreIssuerInfo"

var _FeatureFlag_index = [...]uint16{0, 6, 29, 49, 62, 79, 93, 111, 130, 147, 165, 184, 213, 235, 258, 282, 304, 319}

func (i FeatureFlag) String() string {
	if i < 0 || i >= FeatureFlag(len(_FeatureFlag_index)-1) {
		return "FeatureFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureFlag_name[_FeatureFlag_index[i]:_FeatureFlag_index[i+1]]
}
