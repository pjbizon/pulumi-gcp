// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage.Inputs
{

    public sealed class BucketObjectCustomerEncryptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encryption algorithm. Default: AES256
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        /// <summary>
        /// Base64 encoded Customer-Supplied Encryption Key.
        /// </summary>
        [Input("encryptionKey", required: true)]
        public Input<string> EncryptionKey { get; set; } = null!;

        public BucketObjectCustomerEncryptionGetArgs()
        {
        }
    }
}
