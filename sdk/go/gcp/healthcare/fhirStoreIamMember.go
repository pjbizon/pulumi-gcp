// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
//
// * `healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
// * `healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
// * `healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
//
// > **Note:** `healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `healthcare.FhirStoreIamBinding` and `healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_healthcare\_fhir\_store\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewFhirStoreIamPolicy(ctx, "fhirStore", &healthcare.FhirStoreIamPolicyArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
// 			PolicyData:  pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_healthcare\_fhir\_store\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewFhirStoreIamBinding(ctx, "fhirStore", &healthcare.FhirStoreIamBindingArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_healthcare\_fhir\_store\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthcare.NewFhirStoreIamMember(ctx, "fhirStore", &healthcare.FhirStoreIamMemberArgs{
// 			FhirStoreId: pulumi.String("your-fhir-store-id"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 			Role:        pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `fhir_store_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `fhir_store_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name
// ```
type FhirStoreIamMember struct {
	pulumi.CustomResourceState

	Condition FhirStoreIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringOutput `pulumi:"fhirStoreId"`
	Member      pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFhirStoreIamMember registers a new resource with the given unique name, arguments, and options.
func NewFhirStoreIamMember(ctx *pulumi.Context,
	name string, args *FhirStoreIamMemberArgs, opts ...pulumi.ResourceOption) (*FhirStoreIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FhirStoreId == nil {
		return nil, errors.New("invalid value for required argument 'FhirStoreId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource FhirStoreIamMember
	err := ctx.RegisterResource("gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirStoreIamMember gets an existing FhirStoreIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStoreIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirStoreIamMemberState, opts ...pulumi.ResourceOption) (*FhirStoreIamMember, error) {
	var resource FhirStoreIamMember
	err := ctx.ReadResource("gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirStoreIamMember resources.
type fhirStoreIamMemberState struct {
	Condition *FhirStoreIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId *string `pulumi:"fhirStoreId"`
	Member      *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FhirStoreIamMemberState struct {
	Condition FhirStoreIamMemberConditionPtrInput
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringPtrInput
	Member      pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FhirStoreIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamMemberState)(nil)).Elem()
}

type fhirStoreIamMemberArgs struct {
	Condition *FhirStoreIamMemberCondition `pulumi:"condition"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId string `pulumi:"fhirStoreId"`
	Member      string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FhirStoreIamMember resource.
type FhirStoreIamMemberArgs struct {
	Condition FhirStoreIamMemberConditionPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringInput
	Member      pulumi.StringInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FhirStoreIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamMemberArgs)(nil)).Elem()
}

type FhirStoreIamMemberInput interface {
	pulumi.Input

	ToFhirStoreIamMemberOutput() FhirStoreIamMemberOutput
	ToFhirStoreIamMemberOutputWithContext(ctx context.Context) FhirStoreIamMemberOutput
}

func (*FhirStoreIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStoreIamMember)(nil)).Elem()
}

func (i *FhirStoreIamMember) ToFhirStoreIamMemberOutput() FhirStoreIamMemberOutput {
	return i.ToFhirStoreIamMemberOutputWithContext(context.Background())
}

func (i *FhirStoreIamMember) ToFhirStoreIamMemberOutputWithContext(ctx context.Context) FhirStoreIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamMemberOutput)
}

// FhirStoreIamMemberArrayInput is an input type that accepts FhirStoreIamMemberArray and FhirStoreIamMemberArrayOutput values.
// You can construct a concrete instance of `FhirStoreIamMemberArrayInput` via:
//
//          FhirStoreIamMemberArray{ FhirStoreIamMemberArgs{...} }
type FhirStoreIamMemberArrayInput interface {
	pulumi.Input

	ToFhirStoreIamMemberArrayOutput() FhirStoreIamMemberArrayOutput
	ToFhirStoreIamMemberArrayOutputWithContext(context.Context) FhirStoreIamMemberArrayOutput
}

type FhirStoreIamMemberArray []FhirStoreIamMemberInput

func (FhirStoreIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FhirStoreIamMember)(nil)).Elem()
}

func (i FhirStoreIamMemberArray) ToFhirStoreIamMemberArrayOutput() FhirStoreIamMemberArrayOutput {
	return i.ToFhirStoreIamMemberArrayOutputWithContext(context.Background())
}

func (i FhirStoreIamMemberArray) ToFhirStoreIamMemberArrayOutputWithContext(ctx context.Context) FhirStoreIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamMemberArrayOutput)
}

// FhirStoreIamMemberMapInput is an input type that accepts FhirStoreIamMemberMap and FhirStoreIamMemberMapOutput values.
// You can construct a concrete instance of `FhirStoreIamMemberMapInput` via:
//
//          FhirStoreIamMemberMap{ "key": FhirStoreIamMemberArgs{...} }
type FhirStoreIamMemberMapInput interface {
	pulumi.Input

	ToFhirStoreIamMemberMapOutput() FhirStoreIamMemberMapOutput
	ToFhirStoreIamMemberMapOutputWithContext(context.Context) FhirStoreIamMemberMapOutput
}

type FhirStoreIamMemberMap map[string]FhirStoreIamMemberInput

func (FhirStoreIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FhirStoreIamMember)(nil)).Elem()
}

func (i FhirStoreIamMemberMap) ToFhirStoreIamMemberMapOutput() FhirStoreIamMemberMapOutput {
	return i.ToFhirStoreIamMemberMapOutputWithContext(context.Background())
}

func (i FhirStoreIamMemberMap) ToFhirStoreIamMemberMapOutputWithContext(ctx context.Context) FhirStoreIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirStoreIamMemberMapOutput)
}

type FhirStoreIamMemberOutput struct{ *pulumi.OutputState }

func (FhirStoreIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirStoreIamMember)(nil)).Elem()
}

func (o FhirStoreIamMemberOutput) ToFhirStoreIamMemberOutput() FhirStoreIamMemberOutput {
	return o
}

func (o FhirStoreIamMemberOutput) ToFhirStoreIamMemberOutputWithContext(ctx context.Context) FhirStoreIamMemberOutput {
	return o
}

type FhirStoreIamMemberArrayOutput struct{ *pulumi.OutputState }

func (FhirStoreIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FhirStoreIamMember)(nil)).Elem()
}

func (o FhirStoreIamMemberArrayOutput) ToFhirStoreIamMemberArrayOutput() FhirStoreIamMemberArrayOutput {
	return o
}

func (o FhirStoreIamMemberArrayOutput) ToFhirStoreIamMemberArrayOutputWithContext(ctx context.Context) FhirStoreIamMemberArrayOutput {
	return o
}

func (o FhirStoreIamMemberArrayOutput) Index(i pulumi.IntInput) FhirStoreIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FhirStoreIamMember {
		return vs[0].([]*FhirStoreIamMember)[vs[1].(int)]
	}).(FhirStoreIamMemberOutput)
}

type FhirStoreIamMemberMapOutput struct{ *pulumi.OutputState }

func (FhirStoreIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FhirStoreIamMember)(nil)).Elem()
}

func (o FhirStoreIamMemberMapOutput) ToFhirStoreIamMemberMapOutput() FhirStoreIamMemberMapOutput {
	return o
}

func (o FhirStoreIamMemberMapOutput) ToFhirStoreIamMemberMapOutputWithContext(ctx context.Context) FhirStoreIamMemberMapOutput {
	return o
}

func (o FhirStoreIamMemberMapOutput) MapIndex(k pulumi.StringInput) FhirStoreIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FhirStoreIamMember {
		return vs[0].(map[string]*FhirStoreIamMember)[vs[1].(string)]
	}).(FhirStoreIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreIamMemberInput)(nil)).Elem(), &FhirStoreIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreIamMemberArrayInput)(nil)).Elem(), FhirStoreIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FhirStoreIamMemberMapInput)(nil)).Elem(), FhirStoreIamMemberMap{})
	pulumi.RegisterOutputType(FhirStoreIamMemberOutput{})
	pulumi.RegisterOutputType(FhirStoreIamMemberArrayOutput{})
	pulumi.RegisterOutputType(FhirStoreIamMemberMapOutput{})
}
