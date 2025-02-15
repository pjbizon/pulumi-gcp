// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FunctionIamBinding struct {
	pulumi.CustomResourceState

	CloudFunction pulumi.StringOutput                  `pulumi:"cloudFunction"`
	Condition     FunctionIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag          pulumi.StringOutput                  `pulumi:"etag"`
	Members       pulumi.StringArrayOutput             `pulumi:"members"`
	Project       pulumi.StringOutput                  `pulumi:"project"`
	Region        pulumi.StringOutput                  `pulumi:"region"`
	Role          pulumi.StringOutput                  `pulumi:"role"`
}

// NewFunctionIamBinding registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamBinding(ctx *pulumi.Context,
	name string, args *FunctionIamBindingArgs, opts ...pulumi.ResourceOption) (*FunctionIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudFunction == nil {
		return nil, errors.New("invalid value for required argument 'CloudFunction'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource FunctionIamBinding
	err := ctx.RegisterResource("gcp:cloudfunctions/functionIamBinding:FunctionIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamBinding gets an existing FunctionIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamBindingState, opts ...pulumi.ResourceOption) (*FunctionIamBinding, error) {
	var resource FunctionIamBinding
	err := ctx.ReadResource("gcp:cloudfunctions/functionIamBinding:FunctionIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamBinding resources.
type functionIamBindingState struct {
	CloudFunction *string                      `pulumi:"cloudFunction"`
	Condition     *FunctionIamBindingCondition `pulumi:"condition"`
	Etag          *string                      `pulumi:"etag"`
	Members       []string                     `pulumi:"members"`
	Project       *string                      `pulumi:"project"`
	Region        *string                      `pulumi:"region"`
	Role          *string                      `pulumi:"role"`
}

type FunctionIamBindingState struct {
	CloudFunction pulumi.StringPtrInput
	Condition     FunctionIamBindingConditionPtrInput
	Etag          pulumi.StringPtrInput
	Members       pulumi.StringArrayInput
	Project       pulumi.StringPtrInput
	Region        pulumi.StringPtrInput
	Role          pulumi.StringPtrInput
}

func (FunctionIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamBindingState)(nil)).Elem()
}

type functionIamBindingArgs struct {
	CloudFunction string                       `pulumi:"cloudFunction"`
	Condition     *FunctionIamBindingCondition `pulumi:"condition"`
	Members       []string                     `pulumi:"members"`
	Project       *string                      `pulumi:"project"`
	Region        *string                      `pulumi:"region"`
	Role          string                       `pulumi:"role"`
}

// The set of arguments for constructing a FunctionIamBinding resource.
type FunctionIamBindingArgs struct {
	CloudFunction pulumi.StringInput
	Condition     FunctionIamBindingConditionPtrInput
	Members       pulumi.StringArrayInput
	Project       pulumi.StringPtrInput
	Region        pulumi.StringPtrInput
	Role          pulumi.StringInput
}

func (FunctionIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamBindingArgs)(nil)).Elem()
}

type FunctionIamBindingInput interface {
	pulumi.Input

	ToFunctionIamBindingOutput() FunctionIamBindingOutput
	ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput
}

func (*FunctionIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamBinding)(nil)).Elem()
}

func (i *FunctionIamBinding) ToFunctionIamBindingOutput() FunctionIamBindingOutput {
	return i.ToFunctionIamBindingOutputWithContext(context.Background())
}

func (i *FunctionIamBinding) ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamBindingOutput)
}

// FunctionIamBindingArrayInput is an input type that accepts FunctionIamBindingArray and FunctionIamBindingArrayOutput values.
// You can construct a concrete instance of `FunctionIamBindingArrayInput` via:
//
//          FunctionIamBindingArray{ FunctionIamBindingArgs{...} }
type FunctionIamBindingArrayInput interface {
	pulumi.Input

	ToFunctionIamBindingArrayOutput() FunctionIamBindingArrayOutput
	ToFunctionIamBindingArrayOutputWithContext(context.Context) FunctionIamBindingArrayOutput
}

type FunctionIamBindingArray []FunctionIamBindingInput

func (FunctionIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamBinding)(nil)).Elem()
}

func (i FunctionIamBindingArray) ToFunctionIamBindingArrayOutput() FunctionIamBindingArrayOutput {
	return i.ToFunctionIamBindingArrayOutputWithContext(context.Background())
}

func (i FunctionIamBindingArray) ToFunctionIamBindingArrayOutputWithContext(ctx context.Context) FunctionIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamBindingArrayOutput)
}

// FunctionIamBindingMapInput is an input type that accepts FunctionIamBindingMap and FunctionIamBindingMapOutput values.
// You can construct a concrete instance of `FunctionIamBindingMapInput` via:
//
//          FunctionIamBindingMap{ "key": FunctionIamBindingArgs{...} }
type FunctionIamBindingMapInput interface {
	pulumi.Input

	ToFunctionIamBindingMapOutput() FunctionIamBindingMapOutput
	ToFunctionIamBindingMapOutputWithContext(context.Context) FunctionIamBindingMapOutput
}

type FunctionIamBindingMap map[string]FunctionIamBindingInput

func (FunctionIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamBinding)(nil)).Elem()
}

func (i FunctionIamBindingMap) ToFunctionIamBindingMapOutput() FunctionIamBindingMapOutput {
	return i.ToFunctionIamBindingMapOutputWithContext(context.Background())
}

func (i FunctionIamBindingMap) ToFunctionIamBindingMapOutputWithContext(ctx context.Context) FunctionIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamBindingMapOutput)
}

type FunctionIamBindingOutput struct{ *pulumi.OutputState }

func (FunctionIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionIamBinding)(nil)).Elem()
}

func (o FunctionIamBindingOutput) ToFunctionIamBindingOutput() FunctionIamBindingOutput {
	return o
}

func (o FunctionIamBindingOutput) ToFunctionIamBindingOutputWithContext(ctx context.Context) FunctionIamBindingOutput {
	return o
}

type FunctionIamBindingArrayOutput struct{ *pulumi.OutputState }

func (FunctionIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionIamBinding)(nil)).Elem()
}

func (o FunctionIamBindingArrayOutput) ToFunctionIamBindingArrayOutput() FunctionIamBindingArrayOutput {
	return o
}

func (o FunctionIamBindingArrayOutput) ToFunctionIamBindingArrayOutputWithContext(ctx context.Context) FunctionIamBindingArrayOutput {
	return o
}

func (o FunctionIamBindingArrayOutput) Index(i pulumi.IntInput) FunctionIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionIamBinding {
		return vs[0].([]*FunctionIamBinding)[vs[1].(int)]
	}).(FunctionIamBindingOutput)
}

type FunctionIamBindingMapOutput struct{ *pulumi.OutputState }

func (FunctionIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionIamBinding)(nil)).Elem()
}

func (o FunctionIamBindingMapOutput) ToFunctionIamBindingMapOutput() FunctionIamBindingMapOutput {
	return o
}

func (o FunctionIamBindingMapOutput) ToFunctionIamBindingMapOutputWithContext(ctx context.Context) FunctionIamBindingMapOutput {
	return o
}

func (o FunctionIamBindingMapOutput) MapIndex(k pulumi.StringInput) FunctionIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionIamBinding {
		return vs[0].(map[string]*FunctionIamBinding)[vs[1].(string)]
	}).(FunctionIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamBindingInput)(nil)).Elem(), &FunctionIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamBindingArrayInput)(nil)).Elem(), FunctionIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionIamBindingMapInput)(nil)).Elem(), FunctionIamBindingMap{})
	pulumi.RegisterOutputType(FunctionIamBindingOutput{})
	pulumi.RegisterOutputType(FunctionIamBindingArrayOutput{})
	pulumi.RegisterOutputType(FunctionIamBindingMapOutput{})
}
