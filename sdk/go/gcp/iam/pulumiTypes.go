// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadIdentityPoolProviderAws struct {
	// The AWS account ID.
	AccountId string `pulumi:"accountId"`
}

// WorkloadIdentityPoolProviderAwsInput is an input type that accepts WorkloadIdentityPoolProviderAwsArgs and WorkloadIdentityPoolProviderAwsOutput values.
// You can construct a concrete instance of `WorkloadIdentityPoolProviderAwsInput` via:
//
//          WorkloadIdentityPoolProviderAwsArgs{...}
type WorkloadIdentityPoolProviderAwsInput interface {
	pulumi.Input

	ToWorkloadIdentityPoolProviderAwsOutput() WorkloadIdentityPoolProviderAwsOutput
	ToWorkloadIdentityPoolProviderAwsOutputWithContext(context.Context) WorkloadIdentityPoolProviderAwsOutput
}

type WorkloadIdentityPoolProviderAwsArgs struct {
	// The AWS account ID.
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (WorkloadIdentityPoolProviderAwsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPoolProviderAws)(nil)).Elem()
}

func (i WorkloadIdentityPoolProviderAwsArgs) ToWorkloadIdentityPoolProviderAwsOutput() WorkloadIdentityPoolProviderAwsOutput {
	return i.ToWorkloadIdentityPoolProviderAwsOutputWithContext(context.Background())
}

func (i WorkloadIdentityPoolProviderAwsArgs) ToWorkloadIdentityPoolProviderAwsOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderAwsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolProviderAwsOutput)
}

func (i WorkloadIdentityPoolProviderAwsArgs) ToWorkloadIdentityPoolProviderAwsPtrOutput() WorkloadIdentityPoolProviderAwsPtrOutput {
	return i.ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(context.Background())
}

func (i WorkloadIdentityPoolProviderAwsArgs) ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderAwsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolProviderAwsOutput).ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(ctx)
}

// WorkloadIdentityPoolProviderAwsPtrInput is an input type that accepts WorkloadIdentityPoolProviderAwsArgs, WorkloadIdentityPoolProviderAwsPtr and WorkloadIdentityPoolProviderAwsPtrOutput values.
// You can construct a concrete instance of `WorkloadIdentityPoolProviderAwsPtrInput` via:
//
//          WorkloadIdentityPoolProviderAwsArgs{...}
//
//  or:
//
//          nil
type WorkloadIdentityPoolProviderAwsPtrInput interface {
	pulumi.Input

	ToWorkloadIdentityPoolProviderAwsPtrOutput() WorkloadIdentityPoolProviderAwsPtrOutput
	ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(context.Context) WorkloadIdentityPoolProviderAwsPtrOutput
}

type workloadIdentityPoolProviderAwsPtrType WorkloadIdentityPoolProviderAwsArgs

func WorkloadIdentityPoolProviderAwsPtr(v *WorkloadIdentityPoolProviderAwsArgs) WorkloadIdentityPoolProviderAwsPtrInput {
	return (*workloadIdentityPoolProviderAwsPtrType)(v)
}

func (*workloadIdentityPoolProviderAwsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadIdentityPoolProviderAws)(nil)).Elem()
}

func (i *workloadIdentityPoolProviderAwsPtrType) ToWorkloadIdentityPoolProviderAwsPtrOutput() WorkloadIdentityPoolProviderAwsPtrOutput {
	return i.ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(context.Background())
}

func (i *workloadIdentityPoolProviderAwsPtrType) ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderAwsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolProviderAwsPtrOutput)
}

type WorkloadIdentityPoolProviderAwsOutput struct{ *pulumi.OutputState }

func (WorkloadIdentityPoolProviderAwsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPoolProviderAws)(nil)).Elem()
}

func (o WorkloadIdentityPoolProviderAwsOutput) ToWorkloadIdentityPoolProviderAwsOutput() WorkloadIdentityPoolProviderAwsOutput {
	return o
}

func (o WorkloadIdentityPoolProviderAwsOutput) ToWorkloadIdentityPoolProviderAwsOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderAwsOutput {
	return o
}

func (o WorkloadIdentityPoolProviderAwsOutput) ToWorkloadIdentityPoolProviderAwsPtrOutput() WorkloadIdentityPoolProviderAwsPtrOutput {
	return o.ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(context.Background())
}

func (o WorkloadIdentityPoolProviderAwsOutput) ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderAwsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadIdentityPoolProviderAws) *WorkloadIdentityPoolProviderAws {
		return &v
	}).(WorkloadIdentityPoolProviderAwsPtrOutput)
}

// The AWS account ID.
func (o WorkloadIdentityPoolProviderAwsOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadIdentityPoolProviderAws) string { return v.AccountId }).(pulumi.StringOutput)
}

type WorkloadIdentityPoolProviderAwsPtrOutput struct{ *pulumi.OutputState }

func (WorkloadIdentityPoolProviderAwsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadIdentityPoolProviderAws)(nil)).Elem()
}

func (o WorkloadIdentityPoolProviderAwsPtrOutput) ToWorkloadIdentityPoolProviderAwsPtrOutput() WorkloadIdentityPoolProviderAwsPtrOutput {
	return o
}

func (o WorkloadIdentityPoolProviderAwsPtrOutput) ToWorkloadIdentityPoolProviderAwsPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderAwsPtrOutput {
	return o
}

func (o WorkloadIdentityPoolProviderAwsPtrOutput) Elem() WorkloadIdentityPoolProviderAwsOutput {
	return o.ApplyT(func(v *WorkloadIdentityPoolProviderAws) WorkloadIdentityPoolProviderAws {
		if v != nil {
			return *v
		}
		var ret WorkloadIdentityPoolProviderAws
		return ret
	}).(WorkloadIdentityPoolProviderAwsOutput)
}

// The AWS account ID.
func (o WorkloadIdentityPoolProviderAwsPtrOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadIdentityPoolProviderAws) *string {
		if v == nil {
			return nil
		}
		return &v.AccountId
	}).(pulumi.StringPtrOutput)
}

type WorkloadIdentityPoolProviderOidc struct {
	// Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange
	// requests are rejected if the token audience does not match one of the configured
	// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	// be configured.
	// If this list is empty, the OIDC token audience must be equal to the full canonical
	// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	// For example:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AllowedAudiences []string `pulumi:"allowedAudiences"`
	// The OIDC issuer URL.
	IssuerUri string `pulumi:"issuerUri"`
}

// WorkloadIdentityPoolProviderOidcInput is an input type that accepts WorkloadIdentityPoolProviderOidcArgs and WorkloadIdentityPoolProviderOidcOutput values.
// You can construct a concrete instance of `WorkloadIdentityPoolProviderOidcInput` via:
//
//          WorkloadIdentityPoolProviderOidcArgs{...}
type WorkloadIdentityPoolProviderOidcInput interface {
	pulumi.Input

	ToWorkloadIdentityPoolProviderOidcOutput() WorkloadIdentityPoolProviderOidcOutput
	ToWorkloadIdentityPoolProviderOidcOutputWithContext(context.Context) WorkloadIdentityPoolProviderOidcOutput
}

type WorkloadIdentityPoolProviderOidcArgs struct {
	// Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange
	// requests are rejected if the token audience does not match one of the configured
	// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	// be configured.
	// If this list is empty, the OIDC token audience must be equal to the full canonical
	// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	// For example:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AllowedAudiences pulumi.StringArrayInput `pulumi:"allowedAudiences"`
	// The OIDC issuer URL.
	IssuerUri pulumi.StringInput `pulumi:"issuerUri"`
}

func (WorkloadIdentityPoolProviderOidcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (i WorkloadIdentityPoolProviderOidcArgs) ToWorkloadIdentityPoolProviderOidcOutput() WorkloadIdentityPoolProviderOidcOutput {
	return i.ToWorkloadIdentityPoolProviderOidcOutputWithContext(context.Background())
}

func (i WorkloadIdentityPoolProviderOidcArgs) ToWorkloadIdentityPoolProviderOidcOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOidcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolProviderOidcOutput)
}

func (i WorkloadIdentityPoolProviderOidcArgs) ToWorkloadIdentityPoolProviderOidcPtrOutput() WorkloadIdentityPoolProviderOidcPtrOutput {
	return i.ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(context.Background())
}

func (i WorkloadIdentityPoolProviderOidcArgs) ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOidcPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolProviderOidcOutput).ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(ctx)
}

// WorkloadIdentityPoolProviderOidcPtrInput is an input type that accepts WorkloadIdentityPoolProviderOidcArgs, WorkloadIdentityPoolProviderOidcPtr and WorkloadIdentityPoolProviderOidcPtrOutput values.
// You can construct a concrete instance of `WorkloadIdentityPoolProviderOidcPtrInput` via:
//
//          WorkloadIdentityPoolProviderOidcArgs{...}
//
//  or:
//
//          nil
type WorkloadIdentityPoolProviderOidcPtrInput interface {
	pulumi.Input

	ToWorkloadIdentityPoolProviderOidcPtrOutput() WorkloadIdentityPoolProviderOidcPtrOutput
	ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(context.Context) WorkloadIdentityPoolProviderOidcPtrOutput
}

type workloadIdentityPoolProviderOidcPtrType WorkloadIdentityPoolProviderOidcArgs

func WorkloadIdentityPoolProviderOidcPtr(v *WorkloadIdentityPoolProviderOidcArgs) WorkloadIdentityPoolProviderOidcPtrInput {
	return (*workloadIdentityPoolProviderOidcPtrType)(v)
}

func (*workloadIdentityPoolProviderOidcPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (i *workloadIdentityPoolProviderOidcPtrType) ToWorkloadIdentityPoolProviderOidcPtrOutput() WorkloadIdentityPoolProviderOidcPtrOutput {
	return i.ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(context.Background())
}

func (i *workloadIdentityPoolProviderOidcPtrType) ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOidcPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolProviderOidcPtrOutput)
}

type WorkloadIdentityPoolProviderOidcOutput struct{ *pulumi.OutputState }

func (WorkloadIdentityPoolProviderOidcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (o WorkloadIdentityPoolProviderOidcOutput) ToWorkloadIdentityPoolProviderOidcOutput() WorkloadIdentityPoolProviderOidcOutput {
	return o
}

func (o WorkloadIdentityPoolProviderOidcOutput) ToWorkloadIdentityPoolProviderOidcOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOidcOutput {
	return o
}

func (o WorkloadIdentityPoolProviderOidcOutput) ToWorkloadIdentityPoolProviderOidcPtrOutput() WorkloadIdentityPoolProviderOidcPtrOutput {
	return o.ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(context.Background())
}

func (o WorkloadIdentityPoolProviderOidcOutput) ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOidcPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadIdentityPoolProviderOidc) *WorkloadIdentityPoolProviderOidc {
		return &v
	}).(WorkloadIdentityPoolProviderOidcPtrOutput)
}

// Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange
// requests are rejected if the token audience does not match one of the configured
// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
// be configured.
// If this list is empty, the OIDC token audience must be equal to the full canonical
// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
// For example:
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
func (o WorkloadIdentityPoolProviderOidcOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadIdentityPoolProviderOidc) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

// The OIDC issuer URL.
func (o WorkloadIdentityPoolProviderOidcOutput) IssuerUri() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadIdentityPoolProviderOidc) string { return v.IssuerUri }).(pulumi.StringOutput)
}

type WorkloadIdentityPoolProviderOidcPtrOutput struct{ *pulumi.OutputState }

func (WorkloadIdentityPoolProviderOidcPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (o WorkloadIdentityPoolProviderOidcPtrOutput) ToWorkloadIdentityPoolProviderOidcPtrOutput() WorkloadIdentityPoolProviderOidcPtrOutput {
	return o
}

func (o WorkloadIdentityPoolProviderOidcPtrOutput) ToWorkloadIdentityPoolProviderOidcPtrOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOidcPtrOutput {
	return o
}

func (o WorkloadIdentityPoolProviderOidcPtrOutput) Elem() WorkloadIdentityPoolProviderOidcOutput {
	return o.ApplyT(func(v *WorkloadIdentityPoolProviderOidc) WorkloadIdentityPoolProviderOidc {
		if v != nil {
			return *v
		}
		var ret WorkloadIdentityPoolProviderOidc
		return ret
	}).(WorkloadIdentityPoolProviderOidcOutput)
}

// Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange
// requests are rejected if the token audience does not match one of the configured
// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
// be configured.
// If this list is empty, the OIDC token audience must be equal to the full canonical
// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
// For example:
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
func (o WorkloadIdentityPoolProviderOidcPtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadIdentityPoolProviderOidc) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

// The OIDC issuer URL.
func (o WorkloadIdentityPoolProviderOidcPtrOutput) IssuerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadIdentityPoolProviderOidc) *string {
		if v == nil {
			return nil
		}
		return &v.IssuerUri
	}).(pulumi.StringPtrOutput)
}

type GetTestablePermissionsPermission struct {
	// Whether the corresponding API has been enabled for the resource.
	ApiDisabled bool `pulumi:"apiDisabled"`
	// The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`
	CustomSupportLevel string `pulumi:"customSupportLevel"`
	// Name of the permission.
	Name string `pulumi:"name"`
	// Release stage of the permission.
	Stage string `pulumi:"stage"`
	// Human readable title of the permission.
	Title string `pulumi:"title"`
}

// GetTestablePermissionsPermissionInput is an input type that accepts GetTestablePermissionsPermissionArgs and GetTestablePermissionsPermissionOutput values.
// You can construct a concrete instance of `GetTestablePermissionsPermissionInput` via:
//
//          GetTestablePermissionsPermissionArgs{...}
type GetTestablePermissionsPermissionInput interface {
	pulumi.Input

	ToGetTestablePermissionsPermissionOutput() GetTestablePermissionsPermissionOutput
	ToGetTestablePermissionsPermissionOutputWithContext(context.Context) GetTestablePermissionsPermissionOutput
}

type GetTestablePermissionsPermissionArgs struct {
	// Whether the corresponding API has been enabled for the resource.
	ApiDisabled pulumi.BoolInput `pulumi:"apiDisabled"`
	// The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`
	CustomSupportLevel pulumi.StringInput `pulumi:"customSupportLevel"`
	// Name of the permission.
	Name pulumi.StringInput `pulumi:"name"`
	// Release stage of the permission.
	Stage pulumi.StringInput `pulumi:"stage"`
	// Human readable title of the permission.
	Title pulumi.StringInput `pulumi:"title"`
}

func (GetTestablePermissionsPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestablePermissionsPermission)(nil)).Elem()
}

func (i GetTestablePermissionsPermissionArgs) ToGetTestablePermissionsPermissionOutput() GetTestablePermissionsPermissionOutput {
	return i.ToGetTestablePermissionsPermissionOutputWithContext(context.Background())
}

func (i GetTestablePermissionsPermissionArgs) ToGetTestablePermissionsPermissionOutputWithContext(ctx context.Context) GetTestablePermissionsPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTestablePermissionsPermissionOutput)
}

// GetTestablePermissionsPermissionArrayInput is an input type that accepts GetTestablePermissionsPermissionArray and GetTestablePermissionsPermissionArrayOutput values.
// You can construct a concrete instance of `GetTestablePermissionsPermissionArrayInput` via:
//
//          GetTestablePermissionsPermissionArray{ GetTestablePermissionsPermissionArgs{...} }
type GetTestablePermissionsPermissionArrayInput interface {
	pulumi.Input

	ToGetTestablePermissionsPermissionArrayOutput() GetTestablePermissionsPermissionArrayOutput
	ToGetTestablePermissionsPermissionArrayOutputWithContext(context.Context) GetTestablePermissionsPermissionArrayOutput
}

type GetTestablePermissionsPermissionArray []GetTestablePermissionsPermissionInput

func (GetTestablePermissionsPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTestablePermissionsPermission)(nil)).Elem()
}

func (i GetTestablePermissionsPermissionArray) ToGetTestablePermissionsPermissionArrayOutput() GetTestablePermissionsPermissionArrayOutput {
	return i.ToGetTestablePermissionsPermissionArrayOutputWithContext(context.Background())
}

func (i GetTestablePermissionsPermissionArray) ToGetTestablePermissionsPermissionArrayOutputWithContext(ctx context.Context) GetTestablePermissionsPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTestablePermissionsPermissionArrayOutput)
}

type GetTestablePermissionsPermissionOutput struct{ *pulumi.OutputState }

func (GetTestablePermissionsPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestablePermissionsPermission)(nil)).Elem()
}

func (o GetTestablePermissionsPermissionOutput) ToGetTestablePermissionsPermissionOutput() GetTestablePermissionsPermissionOutput {
	return o
}

func (o GetTestablePermissionsPermissionOutput) ToGetTestablePermissionsPermissionOutputWithContext(ctx context.Context) GetTestablePermissionsPermissionOutput {
	return o
}

// Whether the corresponding API has been enabled for the resource.
func (o GetTestablePermissionsPermissionOutput) ApiDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetTestablePermissionsPermission) bool { return v.ApiDisabled }).(pulumi.BoolOutput)
}

// The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`
func (o GetTestablePermissionsPermissionOutput) CustomSupportLevel() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestablePermissionsPermission) string { return v.CustomSupportLevel }).(pulumi.StringOutput)
}

// Name of the permission.
func (o GetTestablePermissionsPermissionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestablePermissionsPermission) string { return v.Name }).(pulumi.StringOutput)
}

// Release stage of the permission.
func (o GetTestablePermissionsPermissionOutput) Stage() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestablePermissionsPermission) string { return v.Stage }).(pulumi.StringOutput)
}

// Human readable title of the permission.
func (o GetTestablePermissionsPermissionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetTestablePermissionsPermission) string { return v.Title }).(pulumi.StringOutput)
}

type GetTestablePermissionsPermissionArrayOutput struct{ *pulumi.OutputState }

func (GetTestablePermissionsPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTestablePermissionsPermission)(nil)).Elem()
}

func (o GetTestablePermissionsPermissionArrayOutput) ToGetTestablePermissionsPermissionArrayOutput() GetTestablePermissionsPermissionArrayOutput {
	return o
}

func (o GetTestablePermissionsPermissionArrayOutput) ToGetTestablePermissionsPermissionArrayOutputWithContext(ctx context.Context) GetTestablePermissionsPermissionArrayOutput {
	return o
}

func (o GetTestablePermissionsPermissionArrayOutput) Index(i pulumi.IntInput) GetTestablePermissionsPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTestablePermissionsPermission {
		return vs[0].([]GetTestablePermissionsPermission)[vs[1].(int)]
	}).(GetTestablePermissionsPermissionOutput)
}

type GetWorkloadIdentityPoolProviderAw struct {
	AccountId string `pulumi:"accountId"`
}

// GetWorkloadIdentityPoolProviderAwInput is an input type that accepts GetWorkloadIdentityPoolProviderAwArgs and GetWorkloadIdentityPoolProviderAwOutput values.
// You can construct a concrete instance of `GetWorkloadIdentityPoolProviderAwInput` via:
//
//          GetWorkloadIdentityPoolProviderAwArgs{...}
type GetWorkloadIdentityPoolProviderAwInput interface {
	pulumi.Input

	ToGetWorkloadIdentityPoolProviderAwOutput() GetWorkloadIdentityPoolProviderAwOutput
	ToGetWorkloadIdentityPoolProviderAwOutputWithContext(context.Context) GetWorkloadIdentityPoolProviderAwOutput
}

type GetWorkloadIdentityPoolProviderAwArgs struct {
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (GetWorkloadIdentityPoolProviderAwArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkloadIdentityPoolProviderAw)(nil)).Elem()
}

func (i GetWorkloadIdentityPoolProviderAwArgs) ToGetWorkloadIdentityPoolProviderAwOutput() GetWorkloadIdentityPoolProviderAwOutput {
	return i.ToGetWorkloadIdentityPoolProviderAwOutputWithContext(context.Background())
}

func (i GetWorkloadIdentityPoolProviderAwArgs) ToGetWorkloadIdentityPoolProviderAwOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderAwOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkloadIdentityPoolProviderAwOutput)
}

// GetWorkloadIdentityPoolProviderAwArrayInput is an input type that accepts GetWorkloadIdentityPoolProviderAwArray and GetWorkloadIdentityPoolProviderAwArrayOutput values.
// You can construct a concrete instance of `GetWorkloadIdentityPoolProviderAwArrayInput` via:
//
//          GetWorkloadIdentityPoolProviderAwArray{ GetWorkloadIdentityPoolProviderAwArgs{...} }
type GetWorkloadIdentityPoolProviderAwArrayInput interface {
	pulumi.Input

	ToGetWorkloadIdentityPoolProviderAwArrayOutput() GetWorkloadIdentityPoolProviderAwArrayOutput
	ToGetWorkloadIdentityPoolProviderAwArrayOutputWithContext(context.Context) GetWorkloadIdentityPoolProviderAwArrayOutput
}

type GetWorkloadIdentityPoolProviderAwArray []GetWorkloadIdentityPoolProviderAwInput

func (GetWorkloadIdentityPoolProviderAwArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkloadIdentityPoolProviderAw)(nil)).Elem()
}

func (i GetWorkloadIdentityPoolProviderAwArray) ToGetWorkloadIdentityPoolProviderAwArrayOutput() GetWorkloadIdentityPoolProviderAwArrayOutput {
	return i.ToGetWorkloadIdentityPoolProviderAwArrayOutputWithContext(context.Background())
}

func (i GetWorkloadIdentityPoolProviderAwArray) ToGetWorkloadIdentityPoolProviderAwArrayOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderAwArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkloadIdentityPoolProviderAwArrayOutput)
}

type GetWorkloadIdentityPoolProviderAwOutput struct{ *pulumi.OutputState }

func (GetWorkloadIdentityPoolProviderAwOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkloadIdentityPoolProviderAw)(nil)).Elem()
}

func (o GetWorkloadIdentityPoolProviderAwOutput) ToGetWorkloadIdentityPoolProviderAwOutput() GetWorkloadIdentityPoolProviderAwOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderAwOutput) ToGetWorkloadIdentityPoolProviderAwOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderAwOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderAwOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkloadIdentityPoolProviderAw) string { return v.AccountId }).(pulumi.StringOutput)
}

type GetWorkloadIdentityPoolProviderAwArrayOutput struct{ *pulumi.OutputState }

func (GetWorkloadIdentityPoolProviderAwArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkloadIdentityPoolProviderAw)(nil)).Elem()
}

func (o GetWorkloadIdentityPoolProviderAwArrayOutput) ToGetWorkloadIdentityPoolProviderAwArrayOutput() GetWorkloadIdentityPoolProviderAwArrayOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderAwArrayOutput) ToGetWorkloadIdentityPoolProviderAwArrayOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderAwArrayOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderAwArrayOutput) Index(i pulumi.IntInput) GetWorkloadIdentityPoolProviderAwOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetWorkloadIdentityPoolProviderAw {
		return vs[0].([]GetWorkloadIdentityPoolProviderAw)[vs[1].(int)]
	}).(GetWorkloadIdentityPoolProviderAwOutput)
}

type GetWorkloadIdentityPoolProviderOidc struct {
	AllowedAudiences []string `pulumi:"allowedAudiences"`
	IssuerUri        string   `pulumi:"issuerUri"`
}

// GetWorkloadIdentityPoolProviderOidcInput is an input type that accepts GetWorkloadIdentityPoolProviderOidcArgs and GetWorkloadIdentityPoolProviderOidcOutput values.
// You can construct a concrete instance of `GetWorkloadIdentityPoolProviderOidcInput` via:
//
//          GetWorkloadIdentityPoolProviderOidcArgs{...}
type GetWorkloadIdentityPoolProviderOidcInput interface {
	pulumi.Input

	ToGetWorkloadIdentityPoolProviderOidcOutput() GetWorkloadIdentityPoolProviderOidcOutput
	ToGetWorkloadIdentityPoolProviderOidcOutputWithContext(context.Context) GetWorkloadIdentityPoolProviderOidcOutput
}

type GetWorkloadIdentityPoolProviderOidcArgs struct {
	AllowedAudiences pulumi.StringArrayInput `pulumi:"allowedAudiences"`
	IssuerUri        pulumi.StringInput      `pulumi:"issuerUri"`
}

func (GetWorkloadIdentityPoolProviderOidcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (i GetWorkloadIdentityPoolProviderOidcArgs) ToGetWorkloadIdentityPoolProviderOidcOutput() GetWorkloadIdentityPoolProviderOidcOutput {
	return i.ToGetWorkloadIdentityPoolProviderOidcOutputWithContext(context.Background())
}

func (i GetWorkloadIdentityPoolProviderOidcArgs) ToGetWorkloadIdentityPoolProviderOidcOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderOidcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkloadIdentityPoolProviderOidcOutput)
}

// GetWorkloadIdentityPoolProviderOidcArrayInput is an input type that accepts GetWorkloadIdentityPoolProviderOidcArray and GetWorkloadIdentityPoolProviderOidcArrayOutput values.
// You can construct a concrete instance of `GetWorkloadIdentityPoolProviderOidcArrayInput` via:
//
//          GetWorkloadIdentityPoolProviderOidcArray{ GetWorkloadIdentityPoolProviderOidcArgs{...} }
type GetWorkloadIdentityPoolProviderOidcArrayInput interface {
	pulumi.Input

	ToGetWorkloadIdentityPoolProviderOidcArrayOutput() GetWorkloadIdentityPoolProviderOidcArrayOutput
	ToGetWorkloadIdentityPoolProviderOidcArrayOutputWithContext(context.Context) GetWorkloadIdentityPoolProviderOidcArrayOutput
}

type GetWorkloadIdentityPoolProviderOidcArray []GetWorkloadIdentityPoolProviderOidcInput

func (GetWorkloadIdentityPoolProviderOidcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (i GetWorkloadIdentityPoolProviderOidcArray) ToGetWorkloadIdentityPoolProviderOidcArrayOutput() GetWorkloadIdentityPoolProviderOidcArrayOutput {
	return i.ToGetWorkloadIdentityPoolProviderOidcArrayOutputWithContext(context.Background())
}

func (i GetWorkloadIdentityPoolProviderOidcArray) ToGetWorkloadIdentityPoolProviderOidcArrayOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderOidcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetWorkloadIdentityPoolProviderOidcArrayOutput)
}

type GetWorkloadIdentityPoolProviderOidcOutput struct{ *pulumi.OutputState }

func (GetWorkloadIdentityPoolProviderOidcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (o GetWorkloadIdentityPoolProviderOidcOutput) ToGetWorkloadIdentityPoolProviderOidcOutput() GetWorkloadIdentityPoolProviderOidcOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderOidcOutput) ToGetWorkloadIdentityPoolProviderOidcOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderOidcOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderOidcOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetWorkloadIdentityPoolProviderOidc) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o GetWorkloadIdentityPoolProviderOidcOutput) IssuerUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetWorkloadIdentityPoolProviderOidc) string { return v.IssuerUri }).(pulumi.StringOutput)
}

type GetWorkloadIdentityPoolProviderOidcArrayOutput struct{ *pulumi.OutputState }

func (GetWorkloadIdentityPoolProviderOidcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetWorkloadIdentityPoolProviderOidc)(nil)).Elem()
}

func (o GetWorkloadIdentityPoolProviderOidcArrayOutput) ToGetWorkloadIdentityPoolProviderOidcArrayOutput() GetWorkloadIdentityPoolProviderOidcArrayOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderOidcArrayOutput) ToGetWorkloadIdentityPoolProviderOidcArrayOutputWithContext(ctx context.Context) GetWorkloadIdentityPoolProviderOidcArrayOutput {
	return o
}

func (o GetWorkloadIdentityPoolProviderOidcArrayOutput) Index(i pulumi.IntInput) GetWorkloadIdentityPoolProviderOidcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetWorkloadIdentityPoolProviderOidc {
		return vs[0].([]GetWorkloadIdentityPoolProviderOidc)[vs[1].(int)]
	}).(GetWorkloadIdentityPoolProviderOidcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadIdentityPoolProviderAwsInput)(nil)).Elem(), WorkloadIdentityPoolProviderAwsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadIdentityPoolProviderAwsPtrInput)(nil)).Elem(), WorkloadIdentityPoolProviderAwsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadIdentityPoolProviderOidcInput)(nil)).Elem(), WorkloadIdentityPoolProviderOidcArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadIdentityPoolProviderOidcPtrInput)(nil)).Elem(), WorkloadIdentityPoolProviderOidcArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTestablePermissionsPermissionInput)(nil)).Elem(), GetTestablePermissionsPermissionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTestablePermissionsPermissionArrayInput)(nil)).Elem(), GetTestablePermissionsPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetWorkloadIdentityPoolProviderAwInput)(nil)).Elem(), GetWorkloadIdentityPoolProviderAwArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetWorkloadIdentityPoolProviderAwArrayInput)(nil)).Elem(), GetWorkloadIdentityPoolProviderAwArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetWorkloadIdentityPoolProviderOidcInput)(nil)).Elem(), GetWorkloadIdentityPoolProviderOidcArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetWorkloadIdentityPoolProviderOidcArrayInput)(nil)).Elem(), GetWorkloadIdentityPoolProviderOidcArray{})
	pulumi.RegisterOutputType(WorkloadIdentityPoolProviderAwsOutput{})
	pulumi.RegisterOutputType(WorkloadIdentityPoolProviderAwsPtrOutput{})
	pulumi.RegisterOutputType(WorkloadIdentityPoolProviderOidcOutput{})
	pulumi.RegisterOutputType(WorkloadIdentityPoolProviderOidcPtrOutput{})
	pulumi.RegisterOutputType(GetTestablePermissionsPermissionOutput{})
	pulumi.RegisterOutputType(GetTestablePermissionsPermissionArrayOutput{})
	pulumi.RegisterOutputType(GetWorkloadIdentityPoolProviderAwOutput{})
	pulumi.RegisterOutputType(GetWorkloadIdentityPoolProviderAwArrayOutput{})
	pulumi.RegisterOutputType(GetWorkloadIdentityPoolProviderOidcOutput{})
	pulumi.RegisterOutputType(GetWorkloadIdentityPoolProviderOidcArrayOutput{})
}
