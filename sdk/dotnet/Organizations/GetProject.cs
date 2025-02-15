// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static class GetProject
    {
        /// <summary>
        /// Use this data source to get project details.
        /// For more information see
        /// [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects#Project)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var project = Output.Create(Gcp.Organizations.GetProject.InvokeAsync());
        ///         this.ProjectNumber = project.Apply(project =&gt; project.Number);
        ///     }
        /// 
        ///     [Output("projectNumber")]
        ///     public Output&lt;string&gt; ProjectNumber { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("gcp:organizations/getProject:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get project details.
        /// For more information see
        /// [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects#Project)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var project = Output.Create(Gcp.Organizations.GetProject.InvokeAsync());
        ///         this.ProjectNumber = project.Apply(project =&gt; project.Number);
        ///     }
        /// 
        ///     [Output("projectNumber")]
        ///     public Output&lt;string&gt; ProjectNumber { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectResult>("gcp:organizations/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project ID. If it is not provided, the provider project is used.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetProjectArgs()
        {
        }
    }

    public sealed class GetProjectInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project ID. If it is not provided, the provider project is used.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        public readonly bool AutoCreateNetwork;
        public readonly string BillingAccount;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        /// <summary>
        /// The numeric identifier of the project.
        /// </summary>
        public readonly string Number;
        public readonly string OrgId;
        public readonly string? ProjectId;
        public readonly bool SkipDelete;

        [OutputConstructor]
        private GetProjectResult(
            bool autoCreateNetwork,

            string billingAccount,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string number,

            string orgId,

            string? projectId,

            bool skipDelete)
        {
            AutoCreateNetwork = autoCreateNetwork;
            BillingAccount = billingAccount;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
            Number = number;
            OrgId = orgId;
            ProjectId = projectId;
            SkipDelete = skipDelete;
        }
    }
}
