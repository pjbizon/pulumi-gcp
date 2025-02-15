// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AssuredWorkloads
{
    /// <summary>
    /// The AssuredWorkloads Workload resource
    /// 
    /// ## Example Usage
    /// ### Basic_workload
    /// A basic test of a assuredworkloads api
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Gcp.AssuredWorkloads.Workload("primary", new Gcp.AssuredWorkloads.WorkloadArgs
    ///         {
    ///             BillingAccount = "billingAccounts/000000-0000000-0000000-000000",
    ///             ComplianceRegime = "FEDRAMP_MODERATE",
    ///             DisplayName = "Workload Example",
    ///             KmsSettings = new Gcp.AssuredWorkloads.Inputs.WorkloadKmsSettingsArgs
    ///             {
    ///                 NextRotationTime = "9999-10-02T15:01:23Z",
    ///                 RotationPeriod = "10368000s",
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "label-one", "value-one" },
    ///             },
    ///             Location = "us-west1",
    ///             Organization = "123456789",
    ///             ProvisionedResourcesParent = "folders/519620126891",
    ///             ResourceSettings = 
    ///             {
    ///                 new Gcp.AssuredWorkloads.Inputs.WorkloadResourceSettingArgs
    ///                 {
    ///                     ResourceType = "CONSUMER_PROJECT",
    ///                 },
    ///                 new Gcp.AssuredWorkloads.Inputs.WorkloadResourceSettingArgs
    ///                 {
    ///                     ResourceType = "ENCRYPTION_KEYS_PROJECT",
    ///                 },
    ///                 new Gcp.AssuredWorkloads.Inputs.WorkloadResourceSettingArgs
    ///                 {
    ///                     ResourceId = "ring",
    ///                     ResourceType = "KEYRING",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Workload can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:assuredworkloads/workload:Workload default organizations/{{organization}}/locations/{{location}}/workloads/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:assuredworkloads/workload:Workload default {{organization}}/{{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:assuredworkloads/workload:Workload")]
    public partial class Workload : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
        /// </summary>
        [Output("billingAccount")]
        public Output<string> BillingAccount { get; private set; } = null!;

        /// <summary>
        /// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
        /// </summary>
        [Output("complianceRegime")]
        public Output<string> ComplianceRegime { get; private set; } = null!;

        /// <summary>
        /// Output only. Immutable. The Workload creation timestamp.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
        /// </summary>
        [Output("kmsSettings")]
        public Output<Outputs.WorkloadKmsSettings?> KmsSettings { get; private set; } = null!;

        /// <summary>
        /// Optional. Labels applied to the workload.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Output only. The resource name of the workload.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization for the resource
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
        /// </summary>
        [Output("provisionedResourcesParent")]
        public Output<string?> ProvisionedResourcesParent { get; private set; } = null!;

        /// <summary>
        /// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
        /// </summary>
        [Output("resourceSettings")]
        public Output<ImmutableArray<Outputs.WorkloadResourceSetting>> ResourceSettings { get; private set; } = null!;

        /// <summary>
        /// Output only. The resources associated with this workload. These resources will be created when creating the workload. If
        /// any of the projects already exist, the workload creation will fail. Always read only.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.WorkloadResource>> Resources { get; private set; } = null!;


        /// <summary>
        /// Create a Workload resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workload(string name, WorkloadArgs args, CustomResourceOptions? options = null)
            : base("gcp:assuredworkloads/workload:Workload", name, args ?? new WorkloadArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workload(string name, Input<string> id, WorkloadState? state = null, CustomResourceOptions? options = null)
            : base("gcp:assuredworkloads/workload:Workload", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Workload resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workload Get(string name, Input<string> id, WorkloadState? state = null, CustomResourceOptions? options = null)
        {
            return new Workload(name, id, state, options);
        }
    }

    public sealed class WorkloadArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
        /// </summary>
        [Input("billingAccount", required: true)]
        public Input<string> BillingAccount { get; set; } = null!;

        /// <summary>
        /// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
        /// </summary>
        [Input("complianceRegime", required: true)]
        public Input<string> ComplianceRegime { get; set; } = null!;

        /// <summary>
        /// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
        /// </summary>
        [Input("kmsSettings")]
        public Input<Inputs.WorkloadKmsSettingsArgs>? KmsSettings { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Labels applied to the workload.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The organization for the resource
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
        /// </summary>
        [Input("provisionedResourcesParent")]
        public Input<string>? ProvisionedResourcesParent { get; set; }

        [Input("resourceSettings")]
        private InputList<Inputs.WorkloadResourceSettingArgs>? _resourceSettings;

        /// <summary>
        /// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
        /// </summary>
        public InputList<Inputs.WorkloadResourceSettingArgs> ResourceSettings
        {
            get => _resourceSettings ?? (_resourceSettings = new InputList<Inputs.WorkloadResourceSettingArgs>());
            set => _resourceSettings = value;
        }

        public WorkloadArgs()
        {
        }
    }

    public sealed class WorkloadState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
        /// </summary>
        [Input("complianceRegime")]
        public Input<string>? ComplianceRegime { get; set; }

        /// <summary>
        /// Output only. Immutable. The Workload creation timestamp.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
        /// </summary>
        [Input("kmsSettings")]
        public Input<Inputs.WorkloadKmsSettingsGetArgs>? KmsSettings { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Labels applied to the workload.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Output only. The resource name of the workload.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization for the resource
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
        /// </summary>
        [Input("provisionedResourcesParent")]
        public Input<string>? ProvisionedResourcesParent { get; set; }

        [Input("resourceSettings")]
        private InputList<Inputs.WorkloadResourceSettingGetArgs>? _resourceSettings;

        /// <summary>
        /// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
        /// </summary>
        public InputList<Inputs.WorkloadResourceSettingGetArgs> ResourceSettings
        {
            get => _resourceSettings ?? (_resourceSettings = new InputList<Inputs.WorkloadResourceSettingGetArgs>());
            set => _resourceSettings = value;
        }

        [Input("resources")]
        private InputList<Inputs.WorkloadResourceGetArgs>? _resources;

        /// <summary>
        /// Output only. The resources associated with this workload. These resources will be created when creating the workload. If
        /// any of the projects already exist, the workload creation will fail. Always read only.
        /// </summary>
        public InputList<Inputs.WorkloadResourceGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.WorkloadResourceGetArgs>());
            set => _resources = value;
        }

        public WorkloadState()
        {
        }
    }
}
