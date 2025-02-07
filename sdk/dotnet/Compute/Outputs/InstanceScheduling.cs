// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class InstanceScheduling
    {
        /// <summary>
        /// Specifies if the instance should be
        /// restarted if it was terminated by Compute Engine (not a user).
        /// Defaults to true.
        /// </summary>
        public readonly bool? AutomaticRestart;
        /// <summary>
        /// The minimum number of virtual CPUs this instance will consume when running on a sole-tenant node.
        /// </summary>
        public readonly int? MinNodeCpus;
        /// <summary>
        /// Specifies node affinities or anti-affinities
        /// to determine which sole-tenant nodes your instances and managed instance
        /// groups will use as host systems. Read more on sole-tenant node creation
        /// [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
        /// Structure documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceSchedulingNodeAffinity> NodeAffinities;
        /// <summary>
        /// Describes maintenance behavior for the
        /// instance. Can be MIGRATE or TERMINATE, for more info, read
        /// [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).
        /// </summary>
        public readonly string? OnHostMaintenance;
        /// <summary>
        /// Specifies if the instance is preemptible.
        /// If this field is set to true, then `automatic_restart` must be
        /// set to false.  Defaults to false.
        /// </summary>
        public readonly bool? Preemptible;

        [OutputConstructor]
        private InstanceScheduling(
            bool? automaticRestart,

            int? minNodeCpus,

            ImmutableArray<Outputs.InstanceSchedulingNodeAffinity> nodeAffinities,

            string? onHostMaintenance,

            bool? preemptible)
        {
            AutomaticRestart = automaticRestart;
            MinNodeCpus = minNodeCpus;
            NodeAffinities = nodeAffinities;
            OnHostMaintenance = onHostMaintenance;
            Preemptible = preemptible;
        }
    }
}
