// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Recaptcha
{
    /// <summary>
    /// The RecaptchaEnterprise Key resource
    /// 
    /// ## Example Usage
    /// ### Android_key
    /// A basic test of recaptcha enterprise key that can be used by Android apps
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Gcp.Recaptcha.EnterpriseKey("primary", new Gcp.Recaptcha.EnterpriseKeyArgs
    ///         {
    ///             AndroidSettings = new Gcp.Recaptcha.Inputs.EnterpriseKeyAndroidSettingsArgs
    ///             {
    ///                 AllowAllPackageNames = true,
    ///                 AllowedPackageNames = {},
    ///             },
    ///             DisplayName = "display-name-one",
    ///             Labels = 
    ///             {
    ///                 { "label-one", "value-one" },
    ///             },
    ///             Project = "my-project-name",
    ///             TestingOptions = new Gcp.Recaptcha.Inputs.EnterpriseKeyTestingOptionsArgs
    ///             {
    ///                 TestingScore = 0.8,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Ios_key
    /// A basic test of recaptcha enterprise key that can be used by iOS apps
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Gcp.Recaptcha.EnterpriseKey("primary", new Gcp.Recaptcha.EnterpriseKeyArgs
    ///         {
    ///             DisplayName = "display-name-one",
    ///             IosSettings = new Gcp.Recaptcha.Inputs.EnterpriseKeyIosSettingsArgs
    ///             {
    ///                 AllowAllBundleIds = true,
    ///                 AllowedBundleIds = {},
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "label-one", "value-one" },
    ///             },
    ///             Project = "my-project-name",
    ///             TestingOptions = new Gcp.Recaptcha.Inputs.EnterpriseKeyTestingOptionsArgs
    ///             {
    ///                 TestingScore = 1,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Minimal_key
    /// A minimal test of recaptcha enterprise key
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Gcp.Recaptcha.EnterpriseKey("primary", new Gcp.Recaptcha.EnterpriseKeyArgs
    ///         {
    ///             DisplayName = "display-name-one",
    ///             Labels = ,
    ///             Project = "my-project-name",
    ///             WebSettings = new Gcp.Recaptcha.Inputs.EnterpriseKeyWebSettingsArgs
    ///             {
    ///                 AllowAllDomains = true,
    ///                 IntegrationType = "SCORE",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Web_key
    /// A basic test of recaptcha enterprise key that can be used by websites
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Gcp.Recaptcha.EnterpriseKey("primary", new Gcp.Recaptcha.EnterpriseKeyArgs
    ///         {
    ///             DisplayName = "display-name-one",
    ///             Labels = 
    ///             {
    ///                 { "label-one", "value-one" },
    ///             },
    ///             Project = "my-project-name",
    ///             TestingOptions = new Gcp.Recaptcha.Inputs.EnterpriseKeyTestingOptionsArgs
    ///             {
    ///                 TestingChallenge = "NOCAPTCHA",
    ///                 TestingScore = 0.5,
    ///             },
    ///             WebSettings = new Gcp.Recaptcha.Inputs.EnterpriseKeyWebSettingsArgs
    ///             {
    ///                 AllowAllDomains = true,
    ///                 AllowedDomains = {},
    ///                 ChallengeSecurityPreference = "USABILITY",
    ///                 IntegrationType = "CHECKBOX",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Web_score_key
    /// A basic test of recaptcha enterprise key with score integration type that can be used by websites
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Gcp.Recaptcha.EnterpriseKey("primary", new Gcp.Recaptcha.EnterpriseKeyArgs
    ///         {
    ///             DisplayName = "display-name-one",
    ///             Labels = 
    ///             {
    ///                 { "label-one", "value-one" },
    ///             },
    ///             Project = "my-project-name",
    ///             TestingOptions = new Gcp.Recaptcha.Inputs.EnterpriseKeyTestingOptionsArgs
    ///             {
    ///                 TestingScore = 0.5,
    ///             },
    ///             WebSettings = new Gcp.Recaptcha.Inputs.EnterpriseKeyWebSettingsArgs
    ///             {
    ///                 AllowAllDomains = true,
    ///                 AllowAmpTraffic = false,
    ///                 AllowedDomains = {},
    ///                 IntegrationType = "SCORE",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Key can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:recaptcha/enterpriseKey:EnterpriseKey default projects/{{project}}/keys/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:recaptcha/enterpriseKey:EnterpriseKey default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:recaptcha/enterpriseKey:EnterpriseKey default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:recaptcha/enterpriseKey:EnterpriseKey")]
    public partial class EnterpriseKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Settings for keys that can be used by Android apps.
        /// </summary>
        [Output("androidSettings")]
        public Output<Outputs.EnterpriseKeyAndroidSettings?> AndroidSettings { get; private set; } = null!;

        /// <summary>
        /// The timestamp corresponding to the creation of this Key.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Human-readable display name of this key. Modifiable by user.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Settings for keys that can be used by iOS apps.
        /// </summary>
        [Output("iosSettings")]
        public Output<Outputs.EnterpriseKeyIosSettings?> IosSettings { get; private set; } = null!;

        /// <summary>
        /// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name for the Key in the format "projects/{project}/keys/{key}".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Options for user acceptance testing.
        /// </summary>
        [Output("testingOptions")]
        public Output<Outputs.EnterpriseKeyTestingOptions?> TestingOptions { get; private set; } = null!;

        /// <summary>
        /// Settings for keys that can be used by websites.
        /// </summary>
        [Output("webSettings")]
        public Output<Outputs.EnterpriseKeyWebSettings?> WebSettings { get; private set; } = null!;


        /// <summary>
        /// Create a EnterpriseKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnterpriseKey(string name, EnterpriseKeyArgs args, CustomResourceOptions? options = null)
            : base("gcp:recaptcha/enterpriseKey:EnterpriseKey", name, args ?? new EnterpriseKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnterpriseKey(string name, Input<string> id, EnterpriseKeyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:recaptcha/enterpriseKey:EnterpriseKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EnterpriseKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnterpriseKey Get(string name, Input<string> id, EnterpriseKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new EnterpriseKey(name, id, state, options);
        }
    }

    public sealed class EnterpriseKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for keys that can be used by Android apps.
        /// </summary>
        [Input("androidSettings")]
        public Input<Inputs.EnterpriseKeyAndroidSettingsArgs>? AndroidSettings { get; set; }

        /// <summary>
        /// Human-readable display name of this key. Modifiable by user.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Settings for keys that can be used by iOS apps.
        /// </summary>
        [Input("iosSettings")]
        public Input<Inputs.EnterpriseKeyIosSettingsArgs>? IosSettings { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Options for user acceptance testing.
        /// </summary>
        [Input("testingOptions")]
        public Input<Inputs.EnterpriseKeyTestingOptionsArgs>? TestingOptions { get; set; }

        /// <summary>
        /// Settings for keys that can be used by websites.
        /// </summary>
        [Input("webSettings")]
        public Input<Inputs.EnterpriseKeyWebSettingsArgs>? WebSettings { get; set; }

        public EnterpriseKeyArgs()
        {
        }
    }

    public sealed class EnterpriseKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for keys that can be used by Android apps.
        /// </summary>
        [Input("androidSettings")]
        public Input<Inputs.EnterpriseKeyAndroidSettingsGetArgs>? AndroidSettings { get; set; }

        /// <summary>
        /// The timestamp corresponding to the creation of this Key.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Human-readable display name of this key. Modifiable by user.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Settings for keys that can be used by iOS apps.
        /// </summary>
        [Input("iosSettings")]
        public Input<Inputs.EnterpriseKeyIosSettingsGetArgs>? IosSettings { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the Key in the format "projects/{project}/keys/{key}".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Options for user acceptance testing.
        /// </summary>
        [Input("testingOptions")]
        public Input<Inputs.EnterpriseKeyTestingOptionsGetArgs>? TestingOptions { get; set; }

        /// <summary>
        /// Settings for keys that can be used by websites.
        /// </summary>
        [Input("webSettings")]
        public Input<Inputs.EnterpriseKeyWebSettingsGetArgs>? WebSettings { get; set; }

        public EnterpriseKeyState()
        {
        }
    }
}
