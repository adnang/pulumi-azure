// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayBackendHttpSettingConnectionDrainingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds connection draining is active. Acceptable values are from `1` second to `3600` seconds.
        /// </summary>
        [Input("drainTimeoutSec", required: true)]
        public Input<int> DrainTimeoutSec { get; set; } = null!;

        /// <summary>
        /// If connection draining is enabled or not.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public ApplicationGatewayBackendHttpSettingConnectionDrainingGetArgs()
        {
        }
    }
}