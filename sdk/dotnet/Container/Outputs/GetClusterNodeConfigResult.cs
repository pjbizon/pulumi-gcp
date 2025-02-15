// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class GetClusterNodeConfigResult
    {
        public readonly string BootDiskKmsKey;
        public readonly int DiskSizeGb;
        public readonly string DiskType;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigEphemeralStorageConfigResult> EphemeralStorageConfigs;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigGcfsConfigResult> GcfsConfigs;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigGuestAcceleratorResult> GuestAccelerators;
        public readonly string ImageType;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigKubeletConfigResult> KubeletConfigs;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigLinuxNodeConfigResult> LinuxNodeConfigs;
        public readonly int LocalSsdCount;
        public readonly string MachineType;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string MinCpuPlatform;
        public readonly string NodeGroup;
        public readonly ImmutableArray<string> OauthScopes;
        public readonly bool Preemptible;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigSandboxConfigResult> SandboxConfigs;
        public readonly string ServiceAccount;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigShieldedInstanceConfigResult> ShieldedInstanceConfigs;
        public readonly bool Spot;
        public readonly ImmutableArray<string> Tags;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigTaintResult> Taints;
        public readonly ImmutableArray<Outputs.GetClusterNodeConfigWorkloadMetadataConfigResult> WorkloadMetadataConfigs;

        [OutputConstructor]
        private GetClusterNodeConfigResult(
            string bootDiskKmsKey,

            int diskSizeGb,

            string diskType,

            ImmutableArray<Outputs.GetClusterNodeConfigEphemeralStorageConfigResult> ephemeralStorageConfigs,

            ImmutableArray<Outputs.GetClusterNodeConfigGcfsConfigResult> gcfsConfigs,

            ImmutableArray<Outputs.GetClusterNodeConfigGuestAcceleratorResult> guestAccelerators,

            string imageType,

            ImmutableArray<Outputs.GetClusterNodeConfigKubeletConfigResult> kubeletConfigs,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetClusterNodeConfigLinuxNodeConfigResult> linuxNodeConfigs,

            int localSsdCount,

            string machineType,

            ImmutableDictionary<string, string> metadata,

            string minCpuPlatform,

            string nodeGroup,

            ImmutableArray<string> oauthScopes,

            bool preemptible,

            ImmutableArray<Outputs.GetClusterNodeConfigSandboxConfigResult> sandboxConfigs,

            string serviceAccount,

            ImmutableArray<Outputs.GetClusterNodeConfigShieldedInstanceConfigResult> shieldedInstanceConfigs,

            bool spot,

            ImmutableArray<string> tags,

            ImmutableArray<Outputs.GetClusterNodeConfigTaintResult> taints,

            ImmutableArray<Outputs.GetClusterNodeConfigWorkloadMetadataConfigResult> workloadMetadataConfigs)
        {
            BootDiskKmsKey = bootDiskKmsKey;
            DiskSizeGb = diskSizeGb;
            DiskType = diskType;
            EphemeralStorageConfigs = ephemeralStorageConfigs;
            GcfsConfigs = gcfsConfigs;
            GuestAccelerators = guestAccelerators;
            ImageType = imageType;
            KubeletConfigs = kubeletConfigs;
            Labels = labels;
            LinuxNodeConfigs = linuxNodeConfigs;
            LocalSsdCount = localSsdCount;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            NodeGroup = nodeGroup;
            OauthScopes = oauthScopes;
            Preemptible = preemptible;
            SandboxConfigs = sandboxConfigs;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfigs = shieldedInstanceConfigs;
            Spot = spot;
            Tags = tags;
            Taints = taints;
            WorkloadMetadataConfigs = workloadMetadataConfigs;
        }
    }
}
