// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./trigger";
export * from "./workerPool";

// Import resources to register:
import { Trigger } from "./trigger";
import { WorkerPool } from "./workerPool";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:cloudbuild/trigger:Trigger":
                return new Trigger(name, <any>undefined, { urn })
            case "gcp:cloudbuild/workerPool:WorkerPool":
                return new WorkerPool(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "cloudbuild/trigger", _module)
pulumi.runtime.registerResourceModule("gcp", "cloudbuild/workerPool", _module)
