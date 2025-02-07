// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudIdentity
{
    public static class GetGroupMemberships
    {
        /// <summary>
        /// Use this data source to get list of the Cloud Identity Group Memberships within a given Group.
        /// 
        /// https://cloud.google.com/identity/docs/concepts/overview#memberships
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
        ///         var members = Output.Create(Gcp.CloudIdentity.GetGroupMemberships.InvokeAsync(new Gcp.CloudIdentity.GetGroupMembershipsArgs
        ///         {
        ///             Group = "groups/123eab45c6defghi",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupMembershipsResult> InvokeAsync(GetGroupMembershipsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupMembershipsResult>("gcp:cloudidentity/getGroupMemberships:getGroupMemberships", args ?? new GetGroupMembershipsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get list of the Cloud Identity Group Memberships within a given Group.
        /// 
        /// https://cloud.google.com/identity/docs/concepts/overview#memberships
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
        ///         var members = Output.Create(Gcp.CloudIdentity.GetGroupMemberships.InvokeAsync(new Gcp.CloudIdentity.GetGroupMembershipsArgs
        ///         {
        ///             Group = "groups/123eab45c6defghi",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupMembershipsResult> Invoke(GetGroupMembershipsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupMembershipsResult>("gcp:cloudidentity/getGroupMemberships:getGroupMemberships", args ?? new GetGroupMembershipsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupMembershipsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}.
        /// </summary>
        [Input("group", required: true)]
        public string Group { get; set; } = null!;

        public GetGroupMembershipsArgs()
        {
        }
    }

    public sealed class GetGroupMembershipsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        public GetGroupMembershipsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupMembershipsResult
    {
        public readonly string Group;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of memberships under the given group. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMembershipsMembershipResult> Memberships;

        [OutputConstructor]
        private GetGroupMembershipsResult(
            string group,

            string id,

            ImmutableArray<Outputs.GetGroupMembershipsMembershipResult> memberships)
        {
            Group = group;
            Id = id;
            Memberships = memberships;
        }
    }
}
