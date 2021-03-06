// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class ScaleSetNetworkProfileIpConfigurationPublicIpAddressConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name label for the dns settings.
        /// </summary>
        [Input("domainNameLabel", required: true)]
        public Input<string> DomainNameLabel { get; set; } = null!;

        /// <summary>
        /// The idle timeout in minutes. This value must be between 4 and 30.
        /// </summary>
        [Input("idleTimeout", required: true)]
        public Input<int> IdleTimeout { get; set; } = null!;

        /// <summary>
        /// The name of the public ip address configuration
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ScaleSetNetworkProfileIpConfigurationPublicIpAddressConfigurationGetArgs()
        {
        }
    }
}
