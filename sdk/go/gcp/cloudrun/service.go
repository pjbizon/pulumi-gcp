// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service acts as a top-level container that manages a set of Routes and
// Configurations which implement a network service. Service exists to provide a
// singular abstraction which can be access controlled, reasoned about, and
// which encapsulates software lifecycle decisions such as rollout policy and
// team resource ownership. Service acts only as an orchestrator of the
// underlying Routes and Configurations (much as a kubernetes Deployment
// orchestrates ReplicaSets).
//
// The Service's controller will track the statuses of its owned Configuration
// and Route, reflecting their statuses and conditions as its own.
//
// See also:
// https://github.com/knative/specs/blob/main/specs/serving/overview.md
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/run/docs/)
//
// > **Warning:** `googleCloudrunService` creates a Managed Google Cloud Run Service. If you need to create
// a Cloud Run Service on Anthos(GKE/VMWare) then you will need to create it using the kubernetes alpha provider.
// Have a look at the Cloud Run Anthos example below.
//
// ## Example Usage
// ### Cloud Run Service Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
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
// 					LatestRevision: pulumi.Bool(true),
// 					Percent:        pulumi.Int(100),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloud Run Service Sql
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance, err := sql.NewDatabaseInstance(ctx, "instance", &sql.DatabaseInstanceArgs{
// 			Region:          pulumi.String("us-east1"),
// 			DatabaseVersion: pulumi.String("MYSQL_5_7"),
// 			Settings: &sql.DatabaseInstanceSettingsArgs{
// 				Tier: pulumi.String("db-f1-micro"),
// 			},
// 			DeletionProtection: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
// 			Location: pulumi.String("us-central1"),
// 			Template: &cloudrun.ServiceTemplateArgs{
// 				Spec: &cloudrun.ServiceTemplateSpecArgs{
// 					Containers: cloudrun.ServiceTemplateSpecContainerArray{
// 						&cloudrun.ServiceTemplateSpecContainerArgs{
// 							Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
// 						},
// 					},
// 				},
// 				Metadata: &cloudrun.ServiceTemplateMetadataArgs{
// 					Annotations: pulumi.StringMap{
// 						"autoscaling.knative.dev/maxScale":      pulumi.String("1000"),
// 						"run.googleapis.com/cloudsql-instances": instance.ConnectionName,
// 						"run.googleapis.com/client-name":        pulumi.String("demo"),
// 					},
// 				},
// 			},
// 			AutogenerateRevisionName: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloud Run Service Noauth
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
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
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		noauthIAMPolicy, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/run.invoker",
// 					Members: []string{
// 						"allUsers",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudrun.NewIamPolicy(ctx, "noauthIamPolicy", &cloudrun.IamPolicyArgs{
// 			Location:   _default.Location,
// 			Project:    _default.Project,
// 			Service:    _default.Name,
// 			PolicyData: pulumi.String(noauthIAMPolicy.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloud Run Service Multiple Environment Variables
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
// 			Location: pulumi.String("us-central1"),
// 			Template: &cloudrun.ServiceTemplateArgs{
// 				Spec: &cloudrun.ServiceTemplateSpecArgs{
// 					Containers: cloudrun.ServiceTemplateSpecContainerArray{
// 						&cloudrun.ServiceTemplateSpecContainerArgs{
// 							Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
// 							Envs: cloudrun.ServiceTemplateSpecContainerEnvArray{
// 								&cloudrun.ServiceTemplateSpecContainerEnvArgs{
// 									Name:  pulumi.String("SOURCE"),
// 									Value: pulumi.String("remote"),
// 								},
// 								&cloudrun.ServiceTemplateSpecContainerEnvArgs{
// 									Name:  pulumi.String("TARGET"),
// 									Value: pulumi.String("home"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Metadata: &cloudrun.ServiceMetadataArgs{
// 				Annotations: pulumi.StringMap{
// 					"generated-by": pulumi.String("magic-modules"),
// 				},
// 			},
// 			Traffics: cloudrun.ServiceTrafficArray{
// 				&cloudrun.ServiceTrafficArgs{
// 					Percent:        pulumi.Int(100),
// 					LatestRevision: pulumi.Bool(true),
// 				},
// 			},
// 			AutogenerateRevisionName: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloud Run Service Traffic Split
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
// 			Location: pulumi.String("us-central1"),
// 			Template: &cloudrun.ServiceTemplateArgs{
// 				Metadata: &cloudrun.ServiceTemplateMetadataArgs{
// 					Name: pulumi.String("cloudrun-srv-green"),
// 				},
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
// 					Percent:      pulumi.Int(25),
// 					RevisionName: pulumi.String("cloudrun-srv-green"),
// 				},
// 				&cloudrun.ServiceTrafficArgs{
// 					Percent:      pulumi.Int(75),
// 					RevisionName: pulumi.String("cloudrun-srv-blue"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloud Run Service Secret Environment Variables
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		secret, err := secretmanager.NewSecret(ctx, "secret", &secretmanager.SecretArgs{
// 			SecretId: pulumi.String("secret"),
// 			Replication: &secretmanager.SecretReplicationArgs{
// 				Automatic: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = secretmanager.NewSecretVersion(ctx, "secret-version-data", &secretmanager.SecretVersionArgs{
// 			Secret:     secret.Name,
// 			SecretData: pulumi.String("secret-data"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = secretmanager.NewSecretIamMember(ctx, "secret-access", &secretmanager.SecretIamMemberArgs{
// 			SecretId: secret.ID(),
// 			Role:     pulumi.String("roles/secretmanager.secretAccessor"),
// 			Member:   pulumi.String(fmt.Sprintf("%v%v%v", "serviceAccount:", project.Number, "-compute@developer.gserviceaccount.com")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			secret,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
// 			Location: pulumi.String("us-central1"),
// 			Template: &cloudrun.ServiceTemplateArgs{
// 				Spec: &cloudrun.ServiceTemplateSpecArgs{
// 					Containers: cloudrun.ServiceTemplateSpecContainerArray{
// 						&cloudrun.ServiceTemplateSpecContainerArgs{
// 							Image: pulumi.String("gcr.io/cloudrun/hello"),
// 							Envs: cloudrun.ServiceTemplateSpecContainerEnvArray{
// 								&cloudrun.ServiceTemplateSpecContainerEnvArgs{
// 									Name: pulumi.String("SECRET_ENV_VAR"),
// 									ValueFrom: &cloudrun.ServiceTemplateSpecContainerEnvValueFromArgs{
// 										SecretKeyRef: &cloudrun.ServiceTemplateSpecContainerEnvValueFromSecretKeyRefArgs{
// 											Name: secret.SecretId,
// 											Key:  pulumi.String("1"),
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Metadata: &cloudrun.ServiceMetadataArgs{
// 				Annotations: pulumi.StringMap{
// 					"generated-by": pulumi.String("magic-modules"),
// 				},
// 			},
// 			Traffics: cloudrun.ServiceTrafficArray{
// 				&cloudrun.ServiceTrafficArgs{
// 					Percent:        pulumi.Int(100),
// 					LatestRevision: pulumi.Bool(true),
// 				},
// 			},
// 			AutogenerateRevisionName: pulumi.Bool(true),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			secret_version_data,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloud Run Service Secret Volumes
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		secret, err := secretmanager.NewSecret(ctx, "secret", &secretmanager.SecretArgs{
// 			SecretId: pulumi.String("secret"),
// 			Replication: &secretmanager.SecretReplicationArgs{
// 				Automatic: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = secretmanager.NewSecretVersion(ctx, "secret-version-data", &secretmanager.SecretVersionArgs{
// 			Secret:     secret.Name,
// 			SecretData: pulumi.String("secret-data"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = secretmanager.NewSecretIamMember(ctx, "secret-access", &secretmanager.SecretIamMemberArgs{
// 			SecretId: secret.ID(),
// 			Role:     pulumi.String("roles/secretmanager.secretAccessor"),
// 			Member:   pulumi.String(fmt.Sprintf("%v%v%v", "serviceAccount:", project.Number, "-compute@developer.gserviceaccount.com")),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			secret,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
// 			Location: pulumi.String("us-central1"),
// 			Template: &cloudrun.ServiceTemplateArgs{
// 				Spec: &cloudrun.ServiceTemplateSpecArgs{
// 					Containers: cloudrun.ServiceTemplateSpecContainerArray{
// 						&cloudrun.ServiceTemplateSpecContainerArgs{
// 							Image: pulumi.String("gcr.io/cloudrun/hello"),
// 							VolumeMounts: cloudrun.ServiceTemplateSpecContainerVolumeMountArray{
// 								&cloudrun.ServiceTemplateSpecContainerVolumeMountArgs{
// 									Name:      pulumi.String("a-volume"),
// 									MountPath: pulumi.String("/secrets"),
// 								},
// 							},
// 						},
// 					},
// 					Volumes: cloudrun.ServiceTemplateSpecVolumeArray{
// 						&cloudrun.ServiceTemplateSpecVolumeArgs{
// 							Name: pulumi.String("a-volume"),
// 							Secret: &cloudrun.ServiceTemplateSpecVolumeSecretArgs{
// 								SecretName:  secret.SecretId,
// 								DefaultMode: pulumi.Int(292),
// 								Items: cloudrun.ServiceTemplateSpecVolumeSecretItemArray{
// 									&cloudrun.ServiceTemplateSpecVolumeSecretItemArgs{
// 										Key:  pulumi.String("1"),
// 										Path: pulumi.String("my-secret"),
// 										Mode: pulumi.Int(256),
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Metadata: &cloudrun.ServiceMetadataArgs{
// 				Annotations: pulumi.StringMap{
// 					"generated-by": pulumi.String("magic-modules"),
// 				},
// 			},
// 			Traffics: cloudrun.ServiceTrafficArray{
// 				&cloudrun.ServiceTrafficArgs{
// 					Percent:        pulumi.Int(100),
// 					LatestRevision: pulumi.Bool(true),
// 				},
// 			},
// 			AutogenerateRevisionName: pulumi.Bool(true),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			secret_version_data,
// 		}))
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
// Service can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudrun/service:Service default locations/{{location}}/namespaces/{{project}}/services/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudrun/service:Service default {{location}}/{{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudrun/service:Service default {{location}}/{{name}}
// ```
type Service struct {
	pulumi.CustomResourceState

	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName pulumi.BoolPtrOutput `pulumi:"autogenerateRevisionName"`
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata ServiceMetadataOutput `pulumi:"metadata"`
	// Volume's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The current status of the Service.
	Statuses ServiceStatusArrayOutput `pulumi:"statuses"`
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template ServiceTemplatePtrOutput `pulumi:"template"`
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics ServiceTrafficArrayOutput `pulumi:"traffics"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Service
	err := ctx.RegisterResource("gcp:cloudrun/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("gcp:cloudrun/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName *bool `pulumi:"autogenerateRevisionName"`
	// The location of the cloud run instance. eg us-central1
	Location *string `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata *ServiceMetadata `pulumi:"metadata"`
	// Volume's name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The current status of the Service.
	Statuses []ServiceStatus `pulumi:"statuses"`
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template *ServiceTemplate `pulumi:"template"`
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics []ServiceTraffic `pulumi:"traffics"`
}

type ServiceState struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName pulumi.BoolPtrInput
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringPtrInput
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata ServiceMetadataPtrInput
	// Volume's name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The current status of the Service.
	Statuses ServiceStatusArrayInput
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template ServiceTemplatePtrInput
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics ServiceTrafficArrayInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName *bool `pulumi:"autogenerateRevisionName"`
	// The location of the cloud run instance. eg us-central1
	Location string `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata *ServiceMetadata `pulumi:"metadata"`
	// Volume's name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template *ServiceTemplate `pulumi:"template"`
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics []ServiceTraffic `pulumi:"traffics"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// If set to `true`, the revision name (template.metadata.name) will be omitted and
	// autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
	// is also set.
	// (For legacy support, if `template.metadata.name` is unset in state while
	// this field is set to false, the revision name will still autogenerate.)
	AutogenerateRevisionName pulumi.BoolPtrInput
	// The location of the cloud run instance. eg us-central1
	Location pulumi.StringInput
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// Structure is documented below.
	Metadata ServiceMetadataPtrInput
	// Volume's name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// Structure is documented below.
	Template ServiceTemplatePtrInput
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// Structure is documented below.
	Traffics ServiceTrafficArrayInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//          ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//          ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
