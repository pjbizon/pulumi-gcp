// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Environment can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:composer/environment:Environment default projects/{{project}}/locations/{{region}}/environments/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:composer/environment:Environment default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:composer/environment:Environment default {{name}}
 * ```
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentState, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:composer/environment:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * Configuration parameters for this environment.
     */
    public readonly config!: pulumi.Output<outputs.composer.EnvironmentConfig>;
    /**
     * User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
     * are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
     * must conform to the following regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be between 0 and 63
     * characters long and must conform to the regular expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 64 labels can be
     * associated with a given environment. Both keys and values must be <= 128 bytes in size.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the environment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The location or Compute Engine region for the environment.
     */
    public readonly region!: pulumi.Output<string | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentArgs | EnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnvironmentState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as EnvironmentArgs | undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Environment resources.
 */
export interface EnvironmentState {
    /**
     * Configuration parameters for this environment.
     */
    config?: pulumi.Input<inputs.composer.EnvironmentConfig>;
    /**
     * User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
     * are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
     * must conform to the following regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be between 0 and 63
     * characters long and must conform to the regular expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 64 labels can be
     * associated with a given environment. Both keys and values must be <= 128 bytes in size.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The location or Compute Engine region for the environment.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Configuration parameters for this environment.
     */
    config?: pulumi.Input<inputs.composer.EnvironmentConfig>;
    /**
     * User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
     * are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
     * must conform to the following regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be between 0 and 63
     * characters long and must conform to the regular expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 64 labels can be
     * associated with a given environment. Both keys and values must be <= 128 bytes in size.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The location or Compute Engine region for the environment.
     */
    region?: pulumi.Input<string>;
}
