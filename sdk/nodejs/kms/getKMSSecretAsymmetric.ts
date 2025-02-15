// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getKMSSecretAsymmetric(args: GetKMSSecretAsymmetricArgs, opts?: pulumi.InvokeOptions): Promise<GetKMSSecretAsymmetricResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gcp:kms/getKMSSecretAsymmetric:getKMSSecretAsymmetric", {
        "ciphertext": args.ciphertext,
        "crc32": args.crc32,
        "cryptoKeyVersion": args.cryptoKeyVersion,
    }, opts);
}

/**
 * A collection of arguments for invoking getKMSSecretAsymmetric.
 */
export interface GetKMSSecretAsymmetricArgs {
    /**
     * The ciphertext to be decrypted, encoded in base64
     */
    ciphertext: string;
    /**
     * The crc32 checksum of the `ciphertext` in hexadecimal notation. If not specified, it will be computed.
     */
    crc32?: string;
    /**
     * The id of the CryptoKey version that will be used to
     * decrypt the provided ciphertext. This is represented by the format
     * `projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}/cryptoKeyVersions/{version}`.
     */
    cryptoKeyVersion: string;
}

/**
 * A collection of values returned by getKMSSecretAsymmetric.
 */
export interface GetKMSSecretAsymmetricResult {
    readonly ciphertext: string;
    /**
     * Contains the crc32 checksum of the provided ciphertext.
     */
    readonly crc32?: string;
    readonly cryptoKeyVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Contains the result of decrypting the provided ciphertext.
     */
    readonly plaintext: string;
}

export function getKMSSecretAsymmetricOutput(args: GetKMSSecretAsymmetricOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKMSSecretAsymmetricResult> {
    return pulumi.output(args).apply(a => getKMSSecretAsymmetric(a, opts))
}

/**
 * A collection of arguments for invoking getKMSSecretAsymmetric.
 */
export interface GetKMSSecretAsymmetricOutputArgs {
    /**
     * The ciphertext to be decrypted, encoded in base64
     */
    ciphertext: pulumi.Input<string>;
    /**
     * The crc32 checksum of the `ciphertext` in hexadecimal notation. If not specified, it will be computed.
     */
    crc32?: pulumi.Input<string>;
    /**
     * The id of the CryptoKey version that will be used to
     * decrypt the provided ciphertext. This is represented by the format
     * `projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}/cryptoKeyVersions/{version}`.
     */
    cryptoKeyVersion: pulumi.Input<string>;
}
