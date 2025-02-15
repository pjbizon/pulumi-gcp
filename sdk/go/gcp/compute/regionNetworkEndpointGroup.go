// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A regional NEG that can support Serverless Products.
//
// To get more information about RegionNetworkEndpointGroup, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/regionNetworkEndpointGroups)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/serverless-neg-concepts)
//
// ## Example Usage
// ### Region Network Endpoint Group Functions
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
// 			Location: pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		archive, err := storage.NewBucketObject(ctx, "archive", &storage.BucketObjectArgs{
// 			Bucket: bucket.Name,
// 			Source: pulumi.NewFileAsset("path/to/index.zip"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		functionNegFunction, err := cloudfunctions.NewFunction(ctx, "functionNegFunction", &cloudfunctions.FunctionArgs{
// 			Description:         pulumi.String("My function"),
// 			Runtime:             pulumi.String("nodejs10"),
// 			AvailableMemoryMb:   pulumi.Int(128),
// 			SourceArchiveBucket: bucket.Name,
// 			SourceArchiveObject: archive.Name,
// 			TriggerHttp:         pulumi.Bool(true),
// 			Timeout:             pulumi.Int(60),
// 			EntryPoint:          pulumi.String("helloGET"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionNetworkEndpointGroup(ctx, "functionNegRegionNetworkEndpointGroup", &compute.RegionNetworkEndpointGroupArgs{
// 			NetworkEndpointType: pulumi.String("SERVERLESS"),
// 			Region:              pulumi.String("us-central1"),
// 			CloudFunction: &compute.RegionNetworkEndpointGroupCloudFunctionArgs{
// 				Function: functionNegFunction.Name,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Region Network Endpoint Group Cloudrun
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cloudrunNegService, err := cloudrun.NewService(ctx, "cloudrunNegService", &cloudrun.ServiceArgs{
// 			Location: pulumi.String("us-central1"),
// 			Template: &cloudrun.ServiceTemplateArgs{
// 				Spec: &cloudrun.ServiceTemplateSpecArgs{
// 					Containers: cloudrun.ServiceTemplateSpecContainerArray{
// 						&cloudrun.ServiceTemplateSpecContainerArgs{
// 							Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
// 						},
// 					},
// 				},
// 			},
// 			Traffics: cloudrun.ServiceTrafficArray{
// 				&cloudrun.ServiceTrafficArgs{
// 					Percent:        pulumi.Int(100),
// 					LatestRevision: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionNetworkEndpointGroup(ctx, "cloudrunNegRegionNetworkEndpointGroup", &compute.RegionNetworkEndpointGroupArgs{
// 			NetworkEndpointType: pulumi.String("SERVERLESS"),
// 			Region:              pulumi.String("us-central1"),
// 			CloudRun: &compute.RegionNetworkEndpointGroupCloudRunArgs{
// 				Service: cloudrunNegService.Name,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Region Network Endpoint Group Appengine
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/appengine"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		appengineNegBucket, err := storage.NewBucket(ctx, "appengineNegBucket", &storage.BucketArgs{
// 			Location: pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		appengineNegBucketObject, err := storage.NewBucketObject(ctx, "appengineNegBucketObject", &storage.BucketObjectArgs{
// 			Bucket: appengineNegBucket.Name,
// 			Source: pulumi.NewFileAsset("./test-fixtures/appengine/hello-world.zip"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		appengineNegFlexibleAppVersion, err := appengine.NewFlexibleAppVersion(ctx, "appengineNegFlexibleAppVersion", &appengine.FlexibleAppVersionArgs{
// 			VersionId: pulumi.String("v1"),
// 			Service:   pulumi.String("appengine-network-endpoint-group"),
// 			Runtime:   pulumi.String("nodejs"),
// 			Entrypoint: &appengine.FlexibleAppVersionEntrypointArgs{
// 				Shell: pulumi.String("node ./app.js"),
// 			},
// 			Deployment: &appengine.FlexibleAppVersionDeploymentArgs{
// 				Zip: &appengine.FlexibleAppVersionDeploymentZipArgs{
// 					SourceUrl: pulumi.All(appengineNegBucket.Name, appengineNegBucketObject.Name).ApplyT(func(_args []interface{}) (string, error) {
// 						appengineNegBucketName := _args[0].(string)
// 						appengineNegBucketObjectName := _args[1].(string)
// 						return fmt.Sprintf("%v%v%v%v", "https://storage.googleapis.com/", appengineNegBucketName, "/", appengineNegBucketObjectName), nil
// 					}).(pulumi.StringOutput),
// 				},
// 			},
// 			LivenessCheck: &appengine.FlexibleAppVersionLivenessCheckArgs{
// 				Path: pulumi.String("/"),
// 			},
// 			ReadinessCheck: &appengine.FlexibleAppVersionReadinessCheckArgs{
// 				Path: pulumi.String("/"),
// 			},
// 			EnvVariables: pulumi.StringMap{
// 				"port": pulumi.String("8080"),
// 			},
// 			Handlers: appengine.FlexibleAppVersionHandlerArray{
// 				&appengine.FlexibleAppVersionHandlerArgs{
// 					UrlRegex:       pulumi.String(".*\\/my-path\\/*"),
// 					SecurityLevel:  pulumi.String("SECURE_ALWAYS"),
// 					Login:          pulumi.String("LOGIN_REQUIRED"),
// 					AuthFailAction: pulumi.String("AUTH_FAIL_ACTION_REDIRECT"),
// 					StaticFiles: &appengine.FlexibleAppVersionHandlerStaticFilesArgs{
// 						Path:            pulumi.String("my-other-path"),
// 						UploadPathRegex: pulumi.String(".*\\/my-path\\/*"),
// 					},
// 				},
// 			},
// 			AutomaticScaling: &appengine.FlexibleAppVersionAutomaticScalingArgs{
// 				CoolDownPeriod: pulumi.String("120s"),
// 				CpuUtilization: &appengine.FlexibleAppVersionAutomaticScalingCpuUtilizationArgs{
// 					TargetUtilization: pulumi.Float64(0.5),
// 				},
// 			},
// 			NoopOnDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRegionNetworkEndpointGroup(ctx, "appengineNegRegionNetworkEndpointGroup", &compute.RegionNetworkEndpointGroupArgs{
// 			NetworkEndpointType: pulumi.String("SERVERLESS"),
// 			Region:              pulumi.String("us-central1"),
// 			AppEngine: &compute.RegionNetworkEndpointGroupAppEngineArgs{
// 				Service: appengineNegFlexibleAppVersion.Service,
// 				Version: appengineNegFlexibleAppVersion.VersionId,
// 			},
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
// RegionNetworkEndpointGroup can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default {{name}}
// ```
type RegionNetworkEndpointGroup struct {
	pulumi.CustomResourceState

	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	AppEngine RegionNetworkEndpointGroupAppEnginePtrOutput `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudFunction RegionNetworkEndpointGroupCloudFunctionPtrOutput `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudRun RegionNetworkEndpointGroupCloudRunPtrOutput `pulumi:"cloudRun"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
	// Default value is `SERVERLESS`.
	// Possible values are `SERVERLESS`.
	NetworkEndpointType pulumi.StringPtrOutput `pulumi:"networkEndpointType"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the Serverless NEGs Reside.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or
	// serverlessDeployment may be set.
	ServerlessDeployment RegionNetworkEndpointGroupServerlessDeploymentPtrOutput `pulumi:"serverlessDeployment"`
}

// NewRegionNetworkEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewRegionNetworkEndpointGroup(ctx *pulumi.Context,
	name string, args *RegionNetworkEndpointGroupArgs, opts ...pulumi.ResourceOption) (*RegionNetworkEndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionNetworkEndpointGroup
	err := ctx.RegisterResource("gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionNetworkEndpointGroup gets an existing RegionNetworkEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionNetworkEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionNetworkEndpointGroupState, opts ...pulumi.ResourceOption) (*RegionNetworkEndpointGroup, error) {
	var resource RegionNetworkEndpointGroup
	err := ctx.ReadResource("gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionNetworkEndpointGroup resources.
type regionNetworkEndpointGroupState struct {
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	AppEngine *RegionNetworkEndpointGroupAppEngine `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudFunction *RegionNetworkEndpointGroupCloudFunction `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudRun *RegionNetworkEndpointGroupCloudRun `pulumi:"cloudRun"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
	// Default value is `SERVERLESS`.
	// Possible values are `SERVERLESS`.
	NetworkEndpointType *string `pulumi:"networkEndpointType"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the Serverless NEGs Reside.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or
	// serverlessDeployment may be set.
	ServerlessDeployment *RegionNetworkEndpointGroupServerlessDeployment `pulumi:"serverlessDeployment"`
}

type RegionNetworkEndpointGroupState struct {
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	AppEngine RegionNetworkEndpointGroupAppEnginePtrInput
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudFunction RegionNetworkEndpointGroupCloudFunctionPtrInput
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudRun RegionNetworkEndpointGroupCloudRunPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
	// Default value is `SERVERLESS`.
	// Possible values are `SERVERLESS`.
	NetworkEndpointType pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the Serverless NEGs Reside.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or
	// serverlessDeployment may be set.
	ServerlessDeployment RegionNetworkEndpointGroupServerlessDeploymentPtrInput
}

func (RegionNetworkEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkEndpointGroupState)(nil)).Elem()
}

type regionNetworkEndpointGroupArgs struct {
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	AppEngine *RegionNetworkEndpointGroupAppEngine `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudFunction *RegionNetworkEndpointGroupCloudFunction `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudRun *RegionNetworkEndpointGroupCloudRun `pulumi:"cloudRun"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
	// Default value is `SERVERLESS`.
	// Possible values are `SERVERLESS`.
	NetworkEndpointType *string `pulumi:"networkEndpointType"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the Serverless NEGs Reside.
	Region string `pulumi:"region"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or
	// serverlessDeployment may be set.
	ServerlessDeployment *RegionNetworkEndpointGroupServerlessDeployment `pulumi:"serverlessDeployment"`
}

// The set of arguments for constructing a RegionNetworkEndpointGroup resource.
type RegionNetworkEndpointGroupArgs struct {
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	AppEngine RegionNetworkEndpointGroupAppEnginePtrInput
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudFunction RegionNetworkEndpointGroupCloudFunctionPtrInput
	// Only valid when networkEndpointType is "SERVERLESS".
	// Only one of cloud_run, app_engine, cloudFunction or serverlessDeployment may be set.
	// Structure is documented below.
	CloudRun RegionNetworkEndpointGroupCloudRunPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
	// Default value is `SERVERLESS`.
	// Possible values are `SERVERLESS`.
	NetworkEndpointType pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the Serverless NEGs Reside.
	Region pulumi.StringInput
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or
	// serverlessDeployment may be set.
	ServerlessDeployment RegionNetworkEndpointGroupServerlessDeploymentPtrInput
}

func (RegionNetworkEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkEndpointGroupArgs)(nil)).Elem()
}

type RegionNetworkEndpointGroupInput interface {
	pulumi.Input

	ToRegionNetworkEndpointGroupOutput() RegionNetworkEndpointGroupOutput
	ToRegionNetworkEndpointGroupOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupOutput
}

func (*RegionNetworkEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNetworkEndpointGroup)(nil)).Elem()
}

func (i *RegionNetworkEndpointGroup) ToRegionNetworkEndpointGroupOutput() RegionNetworkEndpointGroupOutput {
	return i.ToRegionNetworkEndpointGroupOutputWithContext(context.Background())
}

func (i *RegionNetworkEndpointGroup) ToRegionNetworkEndpointGroupOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkEndpointGroupOutput)
}

// RegionNetworkEndpointGroupArrayInput is an input type that accepts RegionNetworkEndpointGroupArray and RegionNetworkEndpointGroupArrayOutput values.
// You can construct a concrete instance of `RegionNetworkEndpointGroupArrayInput` via:
//
//          RegionNetworkEndpointGroupArray{ RegionNetworkEndpointGroupArgs{...} }
type RegionNetworkEndpointGroupArrayInput interface {
	pulumi.Input

	ToRegionNetworkEndpointGroupArrayOutput() RegionNetworkEndpointGroupArrayOutput
	ToRegionNetworkEndpointGroupArrayOutputWithContext(context.Context) RegionNetworkEndpointGroupArrayOutput
}

type RegionNetworkEndpointGroupArray []RegionNetworkEndpointGroupInput

func (RegionNetworkEndpointGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionNetworkEndpointGroup)(nil)).Elem()
}

func (i RegionNetworkEndpointGroupArray) ToRegionNetworkEndpointGroupArrayOutput() RegionNetworkEndpointGroupArrayOutput {
	return i.ToRegionNetworkEndpointGroupArrayOutputWithContext(context.Background())
}

func (i RegionNetworkEndpointGroupArray) ToRegionNetworkEndpointGroupArrayOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkEndpointGroupArrayOutput)
}

// RegionNetworkEndpointGroupMapInput is an input type that accepts RegionNetworkEndpointGroupMap and RegionNetworkEndpointGroupMapOutput values.
// You can construct a concrete instance of `RegionNetworkEndpointGroupMapInput` via:
//
//          RegionNetworkEndpointGroupMap{ "key": RegionNetworkEndpointGroupArgs{...} }
type RegionNetworkEndpointGroupMapInput interface {
	pulumi.Input

	ToRegionNetworkEndpointGroupMapOutput() RegionNetworkEndpointGroupMapOutput
	ToRegionNetworkEndpointGroupMapOutputWithContext(context.Context) RegionNetworkEndpointGroupMapOutput
}

type RegionNetworkEndpointGroupMap map[string]RegionNetworkEndpointGroupInput

func (RegionNetworkEndpointGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionNetworkEndpointGroup)(nil)).Elem()
}

func (i RegionNetworkEndpointGroupMap) ToRegionNetworkEndpointGroupMapOutput() RegionNetworkEndpointGroupMapOutput {
	return i.ToRegionNetworkEndpointGroupMapOutputWithContext(context.Background())
}

func (i RegionNetworkEndpointGroupMap) ToRegionNetworkEndpointGroupMapOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkEndpointGroupMapOutput)
}

type RegionNetworkEndpointGroupOutput struct{ *pulumi.OutputState }

func (RegionNetworkEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNetworkEndpointGroup)(nil)).Elem()
}

func (o RegionNetworkEndpointGroupOutput) ToRegionNetworkEndpointGroupOutput() RegionNetworkEndpointGroupOutput {
	return o
}

func (o RegionNetworkEndpointGroupOutput) ToRegionNetworkEndpointGroupOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupOutput {
	return o
}

type RegionNetworkEndpointGroupArrayOutput struct{ *pulumi.OutputState }

func (RegionNetworkEndpointGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionNetworkEndpointGroup)(nil)).Elem()
}

func (o RegionNetworkEndpointGroupArrayOutput) ToRegionNetworkEndpointGroupArrayOutput() RegionNetworkEndpointGroupArrayOutput {
	return o
}

func (o RegionNetworkEndpointGroupArrayOutput) ToRegionNetworkEndpointGroupArrayOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupArrayOutput {
	return o
}

func (o RegionNetworkEndpointGroupArrayOutput) Index(i pulumi.IntInput) RegionNetworkEndpointGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionNetworkEndpointGroup {
		return vs[0].([]*RegionNetworkEndpointGroup)[vs[1].(int)]
	}).(RegionNetworkEndpointGroupOutput)
}

type RegionNetworkEndpointGroupMapOutput struct{ *pulumi.OutputState }

func (RegionNetworkEndpointGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionNetworkEndpointGroup)(nil)).Elem()
}

func (o RegionNetworkEndpointGroupMapOutput) ToRegionNetworkEndpointGroupMapOutput() RegionNetworkEndpointGroupMapOutput {
	return o
}

func (o RegionNetworkEndpointGroupMapOutput) ToRegionNetworkEndpointGroupMapOutputWithContext(ctx context.Context) RegionNetworkEndpointGroupMapOutput {
	return o
}

func (o RegionNetworkEndpointGroupMapOutput) MapIndex(k pulumi.StringInput) RegionNetworkEndpointGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionNetworkEndpointGroup {
		return vs[0].(map[string]*RegionNetworkEndpointGroup)[vs[1].(string)]
	}).(RegionNetworkEndpointGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNetworkEndpointGroupInput)(nil)).Elem(), &RegionNetworkEndpointGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNetworkEndpointGroupArrayInput)(nil)).Elem(), RegionNetworkEndpointGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNetworkEndpointGroupMapInput)(nil)).Elem(), RegionNetworkEndpointGroupMap{})
	pulumi.RegisterOutputType(RegionNetworkEndpointGroupOutput{})
	pulumi.RegisterOutputType(RegionNetworkEndpointGroupArrayOutput{})
	pulumi.RegisterOutputType(RegionNetworkEndpointGroupMapOutput{})
}
