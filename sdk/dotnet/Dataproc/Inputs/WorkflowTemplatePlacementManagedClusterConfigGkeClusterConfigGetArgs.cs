// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. A target for the deployment.
        /// </summary>
        [Input("namespacedGkeDeploymentTarget")]
        public Input<Inputs.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetGetArgs>? NamespacedGkeDeploymentTarget { get; set; }

        public WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigGetArgs()
        {
        }
    }
}
