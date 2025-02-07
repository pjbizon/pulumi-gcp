// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a google `oauth2` `accessToken` for a different service account than the one initially running the script.
//
// For more information see
// [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)
func GetAccountAccessToken(ctx *pulumi.Context, args *GetAccountAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetAccountAccessTokenResult, error) {
	var rv GetAccountAccessTokenResult
	err := ctx.Invoke("gcp:serviceAccount/getAccountAccessToken:getAccountAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountAccessToken.
type GetAccountAccessTokenArgs struct {
	// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
	Delegates []string `pulumi:"delegates"`
	// Lifetime of the impersonated token (defaults to its max: `3600s`).
	Lifetime *string `pulumi:"lifetime"`
	// The scopes the new credential should have (e.g. `["cloud-platform"]`)
	Scopes []string `pulumi:"scopes"`
	// The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
	TargetServiceAccount string `pulumi:"targetServiceAccount"`
}

// A collection of values returned by getAccountAccessToken.
type GetAccountAccessTokenResult struct {
	// The `accessToken` representing the new generated identity.
	AccessToken string   `pulumi:"accessToken"`
	Delegates   []string `pulumi:"delegates"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string   `pulumi:"id"`
	Lifetime             *string  `pulumi:"lifetime"`
	Scopes               []string `pulumi:"scopes"`
	TargetServiceAccount string   `pulumi:"targetServiceAccount"`
}

func GetAccountAccessTokenOutput(ctx *pulumi.Context, args GetAccountAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetAccountAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountAccessTokenResult, error) {
			args := v.(GetAccountAccessTokenArgs)
			r, err := GetAccountAccessToken(ctx, &args, opts...)
			return *r, err
		}).(GetAccountAccessTokenResultOutput)
}

// A collection of arguments for invoking getAccountAccessToken.
type GetAccountAccessTokenOutputArgs struct {
	// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
	Delegates pulumi.StringArrayInput `pulumi:"delegates"`
	// Lifetime of the impersonated token (defaults to its max: `3600s`).
	Lifetime pulumi.StringPtrInput `pulumi:"lifetime"`
	// The scopes the new credential should have (e.g. `["cloud-platform"]`)
	Scopes pulumi.StringArrayInput `pulumi:"scopes"`
	// The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
	TargetServiceAccount pulumi.StringInput `pulumi:"targetServiceAccount"`
}

func (GetAccountAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountAccessTokenArgs)(nil)).Elem()
}

// A collection of values returned by getAccountAccessToken.
type GetAccountAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetAccountAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountAccessTokenResult)(nil)).Elem()
}

func (o GetAccountAccessTokenResultOutput) ToGetAccountAccessTokenResultOutput() GetAccountAccessTokenResultOutput {
	return o
}

func (o GetAccountAccessTokenResultOutput) ToGetAccountAccessTokenResultOutputWithContext(ctx context.Context) GetAccountAccessTokenResultOutput {
	return o
}

// The `accessToken` representing the new generated identity.
func (o GetAccountAccessTokenResultOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountAccessTokenResult) string { return v.AccessToken }).(pulumi.StringOutput)
}

func (o GetAccountAccessTokenResultOutput) Delegates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountAccessTokenResult) []string { return v.Delegates }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountAccessTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountAccessTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccountAccessTokenResultOutput) Lifetime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountAccessTokenResult) *string { return v.Lifetime }).(pulumi.StringPtrOutput)
}

func (o GetAccountAccessTokenResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountAccessTokenResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o GetAccountAccessTokenResultOutput) TargetServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountAccessTokenResult) string { return v.TargetServiceAccount }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountAccessTokenResultOutput{})
}
