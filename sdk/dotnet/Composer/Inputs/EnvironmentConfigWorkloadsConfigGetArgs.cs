// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Inputs
{

    public sealed class EnvironmentConfigWorkloadsConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("scheduler")]
        public Input<Inputs.EnvironmentConfigWorkloadsConfigSchedulerGetArgs>? Scheduler { get; set; }

        [Input("webServer")]
        public Input<Inputs.EnvironmentConfigWorkloadsConfigWebServerGetArgs>? WebServer { get; set; }

        [Input("worker")]
        public Input<Inputs.EnvironmentConfigWorkloadsConfigWorkerGetArgs>? Worker { get; set; }

        public EnvironmentConfigWorkloadsConfigGetArgs()
        {
        }
    }
}
