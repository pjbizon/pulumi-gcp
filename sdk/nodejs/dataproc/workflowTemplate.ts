// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A Workflow Template is a reusable workflow configuration. It defines a graph of jobs with information on where to run those jobs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const template = new gcp.dataproc.WorkflowTemplate("template", {
 *     jobs: [
 *         {
 *             sparkJob: {
 *                 mainClass: "SomeClass",
 *             },
 *             stepId: "someJob",
 *         },
 *         {
 *             prerequisiteStepIds: ["someJob"],
 *             prestoJob: {
 *                 queryFileUri: "someuri",
 *             },
 *             stepId: "otherJob",
 *         },
 *     ],
 *     location: "us-central1",
 *     placement: {
 *         managedCluster: {
 *             clusterName: "my-cluster",
 *             config: {
 *                 gceClusterConfig: {
 *                     tags: [
 *                         "foo",
 *                         "bar",
 *                     ],
 *                     zone: "us-central1-a",
 *                 },
 *                 masterConfig: {
 *                     diskConfig: {
 *                         bootDiskSizeGb: 15,
 *                         bootDiskType: "pd-ssd",
 *                     },
 *                     machineType: "n1-standard-1",
 *                     numInstances: 1,
 *                 },
 *                 secondaryWorkerConfig: {
 *                     numInstances: 2,
 *                 },
 *                 softwareConfig: {
 *                     imageVersion: "1.3.7-deb9",
 *                 },
 *                 workerConfig: {
 *                     diskConfig: {
 *                         bootDiskSizeGb: 10,
 *                         numLocalSsds: 2,
 *                     },
 *                     machineType: "n1-standard-2",
 *                     numInstances: 3,
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * WorkflowTemplate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default {{location}}/{{name}}
 * ```
 */
export class WorkflowTemplate extends pulumi.CustomResource {
    /**
     * Get an existing WorkflowTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowTemplateState, opts?: pulumi.CustomResourceOptions): WorkflowTemplate {
        return new WorkflowTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/workflowTemplate:WorkflowTemplate';

    /**
     * Returns true if the given object is an instance of WorkflowTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkflowTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkflowTemplate.__pulumiType;
    }

    /**
     * Output only. The time template was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
     */
    public readonly dagTimeout!: pulumi.Output<string | undefined>;
    /**
     * Required. The Directed Acyclic Graph of Jobs to submit.
     */
    public readonly jobs!: pulumi.Output<outputs.dataproc.WorkflowTemplateJob[]>;
    /**
     * Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
     */
    public readonly parameters!: pulumi.Output<outputs.dataproc.WorkflowTemplateParameter[] | undefined>;
    /**
     * Required. WorkflowTemplate scheduling information.
     */
    public readonly placement!: pulumi.Output<outputs.dataproc.WorkflowTemplatePlacement>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. The time template was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
     *
     * @deprecated version is not useful as a configurable field, and will be removed in the future.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a WorkflowTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowTemplateArgs | WorkflowTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkflowTemplateState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dagTimeout"] = state ? state.dagTimeout : undefined;
            resourceInputs["jobs"] = state ? state.jobs : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["placement"] = state ? state.placement : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as WorkflowTemplateArgs | undefined;
            if ((!args || args.jobs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobs'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.placement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'placement'");
            }
            resourceInputs["dagTimeout"] = args ? args.dagTimeout : undefined;
            resourceInputs["jobs"] = args ? args.jobs : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["placement"] = args ? args.placement : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkflowTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkflowTemplate resources.
 */
export interface WorkflowTemplateState {
    /**
     * Output only. The time template was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
     */
    dagTimeout?: pulumi.Input<string>;
    /**
     * Required. The Directed Acyclic Graph of Jobs to submit.
     */
    jobs?: pulumi.Input<pulumi.Input<inputs.dataproc.WorkflowTemplateJob>[]>;
    /**
     * Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.dataproc.WorkflowTemplateParameter>[]>;
    /**
     * Required. WorkflowTemplate scheduling information.
     */
    placement?: pulumi.Input<inputs.dataproc.WorkflowTemplatePlacement>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. The time template was last updated.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
     *
     * @deprecated version is not useful as a configurable field, and will be removed in the future.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a WorkflowTemplate resource.
 */
export interface WorkflowTemplateArgs {
    /**
     * (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
     */
    dagTimeout?: pulumi.Input<string>;
    /**
     * Required. The Directed Acyclic Graph of Jobs to submit.
     */
    jobs: pulumi.Input<pulumi.Input<inputs.dataproc.WorkflowTemplateJob>[]>;
    /**
     * Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.dataproc.WorkflowTemplateParameter>[]>;
    /**
     * Required. WorkflowTemplate scheduling information.
     */
    placement: pulumi.Input<inputs.dataproc.WorkflowTemplatePlacement>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
     *
     * @deprecated version is not useful as a configurable field, and will be removed in the future.
     */
    version?: pulumi.Input<number>;
}
