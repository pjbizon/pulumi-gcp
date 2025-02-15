// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This represents the rollout state. This is part of the game server
 * deployment.
 *
 * To get more information about GameServerDeploymentRollout, see:
 *
 * * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/GameServerDeploymentRollout)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/game-servers/docs)
 *
 * ## Example Usage
 * ### Game Service Deployment Rollout Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultGameServerDeployment = new gcp.gameservices.GameServerDeployment("defaultGameServerDeployment", {
 *     deploymentId: "tf-test-deployment",
 *     description: "a deployment description",
 * });
 * const defaultGameServerConfig = new gcp.gameservices.GameServerConfig("defaultGameServerConfig", {
 *     configId: "tf-test-config",
 *     deploymentId: defaultGameServerDeployment.deploymentId,
 *     description: "a config description",
 *     fleetConfigs: [{
 *         name: "some-non-guid",
 *         fleetSpec: JSON.stringify({
 *             replicas: 1,
 *             scheduling: "Packed",
 *             template: {
 *                 metadata: {
 *                     name: "tf-test-game-server-template",
 *                 },
 *                 spec: {
 *                     ports: [{
 *                         name: "default",
 *                         portPolicy: "Dynamic",
 *                         containerPort: 7654,
 *                         protocol: "UDP",
 *                     }],
 *                     template: {
 *                         spec: {
 *                             containers: [{
 *                                 name: "simple-udp-server",
 *                                 image: "gcr.io/agones-images/udp-server:0.14",
 *                             }],
 *                         },
 *                     },
 *                 },
 *             },
 *         }),
 *     }],
 * });
 * const defaultGameServerDeploymentRollout = new gcp.gameservices.GameServerDeploymentRollout("defaultGameServerDeploymentRollout", {
 *     deploymentId: defaultGameServerDeployment.deploymentId,
 *     defaultGameServerConfig: defaultGameServerConfig.name,
 * });
 * ```
 *
 * ## Import
 *
 * GameServerDeploymentRollout can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:gameservices/gameServerDeploymentRollout:GameServerDeploymentRollout default projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gameservices/gameServerDeploymentRollout:GameServerDeploymentRollout default {{project}}/{{deployment_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gameservices/gameServerDeploymentRollout:GameServerDeploymentRollout default {{deployment_id}}
 * ```
 */
export class GameServerDeploymentRollout extends pulumi.CustomResource {
    /**
     * Get an existing GameServerDeploymentRollout resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GameServerDeploymentRolloutState, opts?: pulumi.CustomResourceOptions): GameServerDeploymentRollout {
        return new GameServerDeploymentRollout(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:gameservices/gameServerDeploymentRollout:GameServerDeploymentRollout';

    /**
     * Returns true if the given object is an instance of GameServerDeploymentRollout.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GameServerDeploymentRollout {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GameServerDeploymentRollout.__pulumiType;
    }

    /**
     * This field points to the game server config that is
     * applied by default to all realms and clusters. For example,
     * `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
     */
    public readonly defaultGameServerConfig!: pulumi.Output<string>;
    /**
     * The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.
     */
    public readonly deploymentId!: pulumi.Output<string>;
    /**
     * The gameServerConfigOverrides contains the per game server config
     * overrides. The overrides are processed in the order they are listed. As
     * soon as a match is found for a cluster, the rest of the list is not
     * processed.
     * Structure is documented below.
     */
    public readonly gameServerConfigOverrides!: pulumi.Output<outputs.gameservices.GameServerDeploymentRolloutGameServerConfigOverride[] | undefined>;
    /**
     * The resource id of the game server deployment eg:
     * 'projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a GameServerDeploymentRollout resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GameServerDeploymentRolloutArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GameServerDeploymentRolloutArgs | GameServerDeploymentRolloutState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GameServerDeploymentRolloutState | undefined;
            resourceInputs["defaultGameServerConfig"] = state ? state.defaultGameServerConfig : undefined;
            resourceInputs["deploymentId"] = state ? state.deploymentId : undefined;
            resourceInputs["gameServerConfigOverrides"] = state ? state.gameServerConfigOverrides : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as GameServerDeploymentRolloutArgs | undefined;
            if ((!args || args.defaultGameServerConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultGameServerConfig'");
            }
            if ((!args || args.deploymentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentId'");
            }
            resourceInputs["defaultGameServerConfig"] = args ? args.defaultGameServerConfig : undefined;
            resourceInputs["deploymentId"] = args ? args.deploymentId : undefined;
            resourceInputs["gameServerConfigOverrides"] = args ? args.gameServerConfigOverrides : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GameServerDeploymentRollout.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GameServerDeploymentRollout resources.
 */
export interface GameServerDeploymentRolloutState {
    /**
     * This field points to the game server config that is
     * applied by default to all realms and clusters. For example,
     * `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
     */
    defaultGameServerConfig?: pulumi.Input<string>;
    /**
     * The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.
     */
    deploymentId?: pulumi.Input<string>;
    /**
     * The gameServerConfigOverrides contains the per game server config
     * overrides. The overrides are processed in the order they are listed. As
     * soon as a match is found for a cluster, the rest of the list is not
     * processed.
     * Structure is documented below.
     */
    gameServerConfigOverrides?: pulumi.Input<pulumi.Input<inputs.gameservices.GameServerDeploymentRolloutGameServerConfigOverride>[]>;
    /**
     * The resource id of the game server deployment eg:
     * 'projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout'.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GameServerDeploymentRollout resource.
 */
export interface GameServerDeploymentRolloutArgs {
    /**
     * This field points to the game server config that is
     * applied by default to all realms and clusters. For example,
     * `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
     */
    defaultGameServerConfig: pulumi.Input<string>;
    /**
     * The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.
     */
    deploymentId: pulumi.Input<string>;
    /**
     * The gameServerConfigOverrides contains the per game server config
     * overrides. The overrides are processed in the order they are listed. As
     * soon as a match is found for a cluster, the rest of the list is not
     * processed.
     * Structure is documented below.
     */
    gameServerConfigOverrides?: pulumi.Input<pulumi.Input<inputs.gameservices.GameServerDeploymentRolloutGameServerConfigOverride>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
