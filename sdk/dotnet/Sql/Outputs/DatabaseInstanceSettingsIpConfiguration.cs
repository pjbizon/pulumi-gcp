// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Outputs
{

    [OutputType]
    public sealed class DatabaseInstanceSettingsIpConfiguration
    {
        /// <summary>
        /// The name of the allocated ip range for the private ip CloudSQL instance. For example: "google-managed-services-default". If set, the cloned instance ip will be created in the allocated range. The range name must comply with [RFC 1035](https://tools.ietf.org/html/rfc1035). Specifically, the name must be 1-63 characters long and match the regular expression a-z?.
        /// </summary>
        public readonly string? AllocatedIpRange;
        public readonly ImmutableArray<Outputs.DatabaseInstanceSettingsIpConfigurationAuthorizedNetwork> AuthorizedNetworks;
        /// <summary>
        /// Whether this Cloud SQL instance should be assigned
        /// a public IPV4 address. At least `ipv4_enabled` must be enabled or a
        /// `private_network` must be configured.
        /// </summary>
        public readonly bool? Ipv4Enabled;
        /// <summary>
        /// The VPC network from which the Cloud SQL
        /// instance is accessible for private IP. For example, projects/myProject/global/networks/default.
        /// Specifying a network enables private IP.
        /// At least `ipv4_enabled` must be enabled or a `private_network` must be configured.
        /// This setting can be updated, but it cannot be removed after it is set.
        /// </summary>
        public readonly string? PrivateNetwork;
        /// <summary>
        /// Whether SSL connections over IP are enforced or not.
        /// </summary>
        public readonly bool? RequireSsl;

        [OutputConstructor]
        private DatabaseInstanceSettingsIpConfiguration(
            string? allocatedIpRange,

            ImmutableArray<Outputs.DatabaseInstanceSettingsIpConfigurationAuthorizedNetwork> authorizedNetworks,

            bool? ipv4Enabled,

            string? privateNetwork,

            bool? requireSsl)
        {
            AllocatedIpRange = allocatedIpRange;
            AuthorizedNetworks = authorizedNetworks;
            Ipv4Enabled = ipv4Enabled;
            PrivateNetwork = privateNetwork;
            RequireSsl = requireSsl;
        }
    }
}
