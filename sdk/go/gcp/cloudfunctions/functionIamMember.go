// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FunctionIamMember struct {
	pulumi.CustomResourceState

	CloudFunction pulumi.StringOutput                 `pulumi:"cloudFunction"`
	Condition     FunctionIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag          pulumi.StringOutput                 `pulumi:"etag"`
	Member        pulumi.StringOutput                 `pulumi:"member"`
	Project       pulumi.StringOutput                 `pulumi:"project"`
	Region        pulumi.StringOutput                 `pulumi:"region"`
	Role          pulumi.StringOutput                 `pulumi:"role"`
}

// NewFunctionIamMember registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamMember(ctx *pulumi.Context,
	name string, args *FunctionIamMemberArgs, opts ...pulumi.ResourceOption) (*FunctionIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudFunction == nil {
		return nil, errors.New("invalid value for required argument 'CloudFunction'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource FunctionIamMember
	err := ctx.RegisterResource("gcp:cloudfunctions/functionIamMember:FunctionIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamMember gets an existing FunctionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamMemberState, opts ...pulumi.ResourceOption) (*FunctionIamMember, error) {
	var resource FunctionIamMember
	err := ctx.ReadResource("gcp:cloudfunctions/functionIamMember:FunctionIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamMember resources.
type functionIamMemberState struct {
	CloudFunction *string                     `pulumi:"cloudFunction"`
	Condition     *FunctionIamMemberCondition `pulumi:"condition"`
	Etag          *string                     `pulumi:"etag"`
	Member        *string                     `pulumi:"member"`
	Project       *string                     `pulumi:"project"`
	Region        *string                     `pulumi:"region"`
	Role          *string                     `pulumi:"role"`
}

type FunctionIamMemberState struct {
	CloudFunction pulumi.StringPtrInput
	Condition     FunctionIamMemberConditionPtrInput
	Etag          pulumi.StringPtrInput
	Member        pulumi.StringPtrInput
	Project       pulumi.StringPtrInput
	Region        pulumi.StringPtrInput
	Role          pulumi.StringPtrInput
}

func (FunctionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamMemberState)(nil)).Elem()
}

type functionIamMemberArgs struct {
	CloudFunction string                      `pulumi:"cloudFunction"`
	Condition     *FunctionIamMemberCondition `pulumi:"condition"`
	Member        string                      `pulumi:"member"`
	Project       *string                     `pulumi:"project"`
	Region        *string                     `pulumi:"region"`
	Role          string                      `pulumi:"role"`
}

// The set of arguments for constructing a FunctionIamMember resource.
type FunctionIamMemberArgs struct {
	CloudFunction pulumi.StringInput
	Condition     FunctionIamMemberConditionPtrInput
	Member        pulumi.StringInput
	Project       pulumi.StringPtrInput
	Region        pulumi.StringPtrInput
	Role          pulumi.StringInput
}

func (FunctionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamMemberArgs)(nil)).Elem()
}

type FunctionIamMemberInput interface {
	pulumi.Input

	ToFunctionIamMemberOutput() FunctionIamMemberOutput
	ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput
}

func (*FunctionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamMember)(nil)).Elem()
}

func (i *FunctionIamMember) ToFunctionIamMemberOutput() FunctionIamMemberOutput {
	return i.ToFunctionIamMemberOutputWithContext(context.Background())
}

func (i *FunctionIamMember) ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamMemberOutput)
}

// FunctionIamMemberArrayInput is an input type that accepts FunctionIamMemberArray and FunctionIamMemberArrayOutput values.
// You can construct a concrete instance of `FunctionIamMemberArrayInput` via:
//
//          FunctionIamMemberArray{ FunctionIamMemberArgs{...} }
type FunctionIamMemberArrayInput interface {
	pulumi.Input

	ToFunctionIamMemberArrayOutput() FunctionIamMemberArrayOutput
	ToFunctionIamMemberArrayOutputWithContext(context.Context) FunctionIamMemberArrayOutput
}

type FunctionIamMemberArray []FunctionIamMemberInput

func (FunctionIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamMember)(nil)).Elem()
}

func (i FunctionIamMemberArray) ToFunctionIamMemberArrayOutput() FunctionIamMemberArrayOutput {
	return i.ToFunctionIamMemberArrayOutputWithContext(context.Background())
}

func (i FunctionIamMemberArray) ToFunctionIamMemberArrayOutputWithContext(ctx context.Context) FunctionIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamMemberArrayOutput)
}

// FunctionIamMemberMapInput is an input type that accepts FunctionIamMemberMap and FunctionIamMemberMapOutput values.
// You can construct a concrete instance of `FunctionIamMemberMapInput` via:
//
//          FunctionIamMemberMap{ "key": FunctionIamMemberArgs{...} }
type FunctionIamMemberMapInput interface {
	pulumi.Input

	ToFunctionIamMemberMapOutput() FunctionIamMemberMapOutput
	ToFunctionIamMemberMapOutputWithContext(context.Context) FunctionIamMemberMapOutput
}

type FunctionIamMemberMap map[string]FunctionIamMemberInput

func (FunctionIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamMember)(nil)).Elem()
}

func (i FunctionIamMemberMap) ToFunctionIamMemberMapOutput() FunctionIamMemberMapOutput {
	return i.ToFunctionIamMemberMapOutputWithContext(context.Background())
}

func (i FunctionIamMemberMap) ToFunctionIamMemberMapOutputWithContext(ctx context.Context) FunctionIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamMemberMapOutput)
}

type FunctionIamMemberOutput struct{ *pulumi.OutputState }

func (FunctionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamMember)(nil)).Elem()
}

func (o FunctionIamMemberOutput) ToFunctionIamMemberOutput() FunctionIamMemberOutput {
	return o
}

func (o FunctionIamMemberOutput) ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput {
	return o
}

type FunctionIamMemberArrayOutput struct{ *pulumi.OutputState }

func (FunctionIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamMember)(nil)).Elem()
}

func (o FunctionIamMemberArrayOutput) ToFunctionIamMemberArrayOutput() FunctionIamMemberArrayOutput {
	return o
}

func (o FunctionIamMemberArrayOutput) ToFunctionIamMemberArrayOutputWithContext(ctx context.Context) FunctionIamMemberArrayOutput {
	return o
}

func (o FunctionIamMemberArrayOutput) Index(i pulumi.IntInput) FunctionIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionIamMember {
		return vs[0].([]*FunctionIamMember)[vs[1].(int)]
	}).(FunctionIamMemberOutput)
}

type FunctionIamMemberMapOutput struct{ *pulumi.OutputState }

func (FunctionIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamMember)(nil)).Elem()
}

func (o FunctionIamMemberMapOutput) ToFunctionIamMemberMapOutput() FunctionIamMemberMapOutput {
	return o
}

func (o FunctionIamMemberMapOutput) ToFunctionIamMemberMapOutputWithContext(ctx context.Context) FunctionIamMemberMapOutput {
	return o
}

func (o FunctionIamMemberMapOutput) MapIndex(k pulumi.StringInput) FunctionIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionIamMember {
		return vs[0].(map[string]*FunctionIamMember)[vs[1].(string)]
	}).(FunctionIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamMemberInput)(nil)).Elem(), &FunctionIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamMemberArrayInput)(nil)).Elem(), FunctionIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamMemberMapInput)(nil)).Elem(), FunctionIamMemberMap{})
	pulumi.RegisterOutputType(FunctionIamMemberOutput{})
	pulumi.RegisterOutputType(FunctionIamMemberArrayOutput{})
	pulumi.RegisterOutputType(FunctionIamMemberMapOutput{})
}
