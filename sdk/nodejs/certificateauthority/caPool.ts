// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A CaPool represents a group of CertificateAuthorities that form a trust anchor. A CaPool can be used to manage
 * issuance policies for one or more CertificateAuthority resources and to rotate CA certificates in and out of the
 * trust anchor.
 *
 * ## Example Usage
 * ### Privateca Capool Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultCaPool = new gcp.certificateauthority.CaPool("default", {
 *     labels: {
 *         foo: "bar",
 *     },
 *     location: "us-central1",
 *     publishingOptions: {
 *         publishCaCert: true,
 *         publishCrl: true,
 *     },
 *     tier: "ENTERPRISE",
 * });
 * ```
 * ### Privateca Capool All Fields
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultCaPool = new gcp.certificateauthority.CaPool("default", {
 *     issuancePolicy: {
 *         allowedIssuanceModes: {
 *             allowConfigBasedIssuance: true,
 *             allowCsrBasedIssuance: true,
 *         },
 *         allowedKeyTypes: [
 *             {
 *                 ellipticCurve: {
 *                     signatureAlgorithm: "ECDSA_P256",
 *                 },
 *             },
 *             {
 *                 rsa: {
 *                     maxModulusSize: "10",
 *                     minModulusSize: "5",
 *                 },
 *             },
 *         ],
 *         baselineValues: {
 *             additionalExtensions: [{
 *                 critical: true,
 *                 objectId: {
 *                     objectIdPaths: [
 *                         1,
 *                         7,
 *                     ],
 *                 },
 *                 value: "asdf",
 *             }],
 *             aiaOcspServers: ["example.com"],
 *             caOptions: {
 *                 isCa: true,
 *                 maxIssuerPathLength: 10,
 *             },
 *             keyUsage: {
 *                 baseKeyUsage: {
 *                     certSign: false,
 *                     contentCommitment: true,
 *                     crlSign: true,
 *                     dataEncipherment: true,
 *                     decipherOnly: true,
 *                     digitalSignature: true,
 *                     keyAgreement: true,
 *                     keyEncipherment: false,
 *                 },
 *                 extendedKeyUsage: {
 *                     clientAuth: false,
 *                     codeSigning: true,
 *                     emailProtection: true,
 *                     serverAuth: true,
 *                     timeStamping: true,
 *                 },
 *             },
 *             policyIds: [
 *                 {
 *                     objectIdPaths: [
 *                         1,
 *                         5,
 *                     ],
 *                 },
 *                 {
 *                     objectIdPaths: [
 *                         1,
 *                         5,
 *                         7,
 *                     ],
 *                 },
 *             ],
 *         },
 *         identityConstraints: {
 *             allowSubjectAltNamesPassthrough: true,
 *             allowSubjectPassthrough: true,
 *             celExpression: {
 *                 expression: "subject_alt_names.all(san, san.type == DNS || san.type == EMAIL )",
 *                 title: "My title",
 *             },
 *         },
 *         maximumLifetime: "50000s",
 *     },
 *     labels: {
 *         foo: "bar",
 *     },
 *     location: "us-central1",
 *     publishingOptions: {
 *         publishCaCert: false,
 *         publishCrl: true,
 *     },
 *     tier: "ENTERPRISE",
 * });
 * ```
 *
 * ## Import
 *
 * CaPool can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/caPool:CaPool default projects/{{project}}/locations/{{location}}/caPools/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{location}}/{{name}}
 * ```
 */
export class CaPool extends pulumi.CustomResource {
    /**
     * Get an existing CaPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CaPoolState, opts?: pulumi.CustomResourceOptions): CaPool {
        return new CaPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:certificateauthority/caPool:CaPool';

    /**
     * Returns true if the given object is an instance of CaPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CaPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CaPool.__pulumiType;
    }

    /**
     * The IssuancePolicy to control how Certificates will be issued from this CaPool.
     * Structure is documented below.
     */
    public readonly issuancePolicy!: pulumi.Output<outputs.certificateauthority.CaPoolIssuancePolicy | undefined>;
    /**
     * Labels with user-defined metadata.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
     * "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name for this CaPool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
     * Structure is documented below.
     */
    public readonly publishingOptions!: pulumi.Output<outputs.certificateauthority.CaPoolPublishingOptions | undefined>;
    /**
     * The Tier of this CaPool.
     * Possible values are `ENTERPRISE` and `DEVOPS`.
     */
    public readonly tier!: pulumi.Output<string>;

    /**
     * Create a CaPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CaPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CaPoolArgs | CaPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CaPoolState | undefined;
            resourceInputs["issuancePolicy"] = state ? state.issuancePolicy : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["publishingOptions"] = state ? state.publishingOptions : undefined;
            resourceInputs["tier"] = state ? state.tier : undefined;
        } else {
            const args = argsOrState as CaPoolArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.tier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tier'");
            }
            resourceInputs["issuancePolicy"] = args ? args.issuancePolicy : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["publishingOptions"] = args ? args.publishingOptions : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CaPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CaPool resources.
 */
export interface CaPoolState {
    /**
     * The IssuancePolicy to control how Certificates will be issued from this CaPool.
     * Structure is documented below.
     */
    issuancePolicy?: pulumi.Input<inputs.certificateauthority.CaPoolIssuancePolicy>;
    /**
     * Labels with user-defined metadata.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
     * "1.3kg", "count": "3" }.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
     */
    location?: pulumi.Input<string>;
    /**
     * The name for this CaPool.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
     * Structure is documented below.
     */
    publishingOptions?: pulumi.Input<inputs.certificateauthority.CaPoolPublishingOptions>;
    /**
     * The Tier of this CaPool.
     * Possible values are `ENTERPRISE` and `DEVOPS`.
     */
    tier?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CaPool resource.
 */
export interface CaPoolArgs {
    /**
     * The IssuancePolicy to control how Certificates will be issued from this CaPool.
     * Structure is documented below.
     */
    issuancePolicy?: pulumi.Input<inputs.certificateauthority.CaPoolIssuancePolicy>;
    /**
     * Labels with user-defined metadata.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
     * "1.3kg", "count": "3" }.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
     */
    location: pulumi.Input<string>;
    /**
     * The name for this CaPool.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
     * Structure is documented below.
     */
    publishingOptions?: pulumi.Input<inputs.certificateauthority.CaPoolPublishingOptions>;
    /**
     * The Tier of this CaPool.
     * Possible values are `ENTERPRISE` and `DEVOPS`.
     */
    tier: pulumi.Input<string>;
}
