// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices.Outputs
{

    [OutputType]
    public sealed class EdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatch
    {
        /// <summary>
        /// The queryParameterMatch matches if the value of the parameter exactly matches the contents of exactMatch.
        /// </summary>
        public readonly string? ExactMatch;
        /// <summary>
        /// Headers to remove from the response prior to sending it back to the client.
        /// Response headers are only sent to the client, and do not have an effect on the cache serving the response.
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// If set to false (default), the headerMatch is considered a match if the match criteria above are met.
        /// If set to true, the headerMatch is considered a match if the match criteria above are NOT met.
        /// </summary>
        public readonly bool? InvertMatch;
        /// <summary>
        /// The value of the header must start with the contents of prefixMatch.
        /// </summary>
        public readonly string? PrefixMatch;
        /// <summary>
        /// Specifies that the queryParameterMatch matches if the request contains the query parameter, irrespective of whether the parameter has a value or not.
        /// </summary>
        public readonly bool? PresentMatch;
        /// <summary>
        /// The value of the header must end with the contents of suffixMatch.
        /// </summary>
        public readonly string? SuffixMatch;

        [OutputConstructor]
        private EdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatch(
            string? exactMatch,

            string headerName,

            bool? invertMatch,

            string? prefixMatch,

            bool? presentMatch,

            string? suffixMatch)
        {
            ExactMatch = exactMatch;
            HeaderName = headerName;
            InvertMatch = invertMatch;
            PrefixMatch = prefixMatch;
            PresentMatch = presentMatch;
            SuffixMatch = suffixMatch;
        }
    }
}
