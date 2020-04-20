// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayProbeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Hostname used for this Probe. If the Application Gateway is configured for a single site, by default the Host name should be specified as ‘127.0.0.1’, unless otherwise configured in custom probe. Cannot be set if `pick_host_name_from_backend_http_settings` is set to `true`.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The Interval between two consecutive probes in seconds. Possible values range from 1 second to a maximum of 86,400 seconds.
        /// </summary>
        [Input("interval", required: true)]
        public Input<int> Interval { get; set; } = null!;

        /// <summary>
        /// A `match` block as defined above.
        /// </summary>
        [Input("match")]
        public Input<Inputs.ApplicationGatewayProbeMatchArgs>? Match { get; set; }

        /// <summary>
        /// The minimum number of servers that are always marked as healthy. Defaults to `0`.
        /// </summary>
        [Input("minimumServers")]
        public Input<int>? MinimumServers { get; set; }

        /// <summary>
        /// The Name of the Probe.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Path used for this Probe.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Whether the host header should be picked from the backend http settings. Defaults to `false`.
        /// </summary>
        [Input("pickHostNameFromBackendHttpSettings")]
        public Input<bool>? PickHostNameFromBackendHttpSettings { get; set; }

        /// <summary>
        /// The Protocol used for this Probe. Possible values are `Http` and `Https`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The Timeout used for this Probe, which indicates when a probe becomes unhealthy. Possible values range from 1 second to a maximum of 86,400 seconds.
        /// </summary>
        [Input("timeout", required: true)]
        public Input<int> Timeout { get; set; } = null!;

        /// <summary>
        /// The Unhealthy Threshold for this Probe, which indicates the amount of retries which should be attempted before a node is deemed unhealthy. Possible values are from 1 - 20 seconds.
        /// </summary>
        [Input("unhealthyThreshold", required: true)]
        public Input<int> UnhealthyThreshold { get; set; } = null!;

        public ApplicationGatewayProbeArgs()
        {
        }
    }
}