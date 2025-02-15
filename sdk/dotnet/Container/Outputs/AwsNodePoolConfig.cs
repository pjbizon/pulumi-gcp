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
    public sealed class AwsNodePoolConfig
    {
        /// <summary>
        /// Required. The ARN of the AWS KMS key used to encrypt node pool configuration.
        /// </summary>
        public readonly Outputs.AwsNodePoolConfigConfigEncryption ConfigEncryption;
        /// <summary>
        /// Required. The name of the AWS IAM role assigned to nodes in the pool.
        /// </summary>
        public readonly string IamInstanceProfile;
        /// <summary>
        /// Optional. The AWS instance type. When unspecified, it defaults to `t3.medium`.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// Optional. The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Optional. Template for the root volume provisioned for node pool nodes. Volumes will be provisioned in the availability zone assigned to the node pool subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.
        /// </summary>
        public readonly Outputs.AwsNodePoolConfigRootVolume? RootVolume;
        /// <summary>
        /// Optional. The IDs of additional security groups to add to nodes in this pool. The manager will automatically create security groups with minimum rules needed for a functioning cluster.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// Optional. The SSH configuration.
        /// </summary>
        public readonly Outputs.AwsNodePoolConfigSshConfig? SshConfig;
        /// <summary>
        /// Optional. Key/value metadata to assign to each underlying AWS resource. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Optional. The initial taints assigned to nodes of this node pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.AwsNodePoolConfigTaint> Taints;

        [OutputConstructor]
        private AwsNodePoolConfig(
            Outputs.AwsNodePoolConfigConfigEncryption configEncryption,

            string iamInstanceProfile,

            string? instanceType,

            ImmutableDictionary<string, string>? labels,

            Outputs.AwsNodePoolConfigRootVolume? rootVolume,

            ImmutableArray<string> securityGroupIds,

            Outputs.AwsNodePoolConfigSshConfig? sshConfig,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.AwsNodePoolConfigTaint> taints)
        {
            ConfigEncryption = configEncryption;
            IamInstanceProfile = iamInstanceProfile;
            InstanceType = instanceType;
            Labels = labels;
            RootVolume = rootVolume;
            SecurityGroupIds = securityGroupIds;
            SshConfig = sshConfig;
            Tags = tags;
            Taints = taints;
        }
    }
}
