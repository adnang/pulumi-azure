// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Outputs
{

    [OutputType]
    public sealed class AutoscaleSettingProfileCapacity
    {
        /// <summary>
        /// The number of instances that are available for scaling if metrics are not available for evaluation. The default is only used if the current instance count is lower than the default. Valid values are between `0` and `1000`.
        /// </summary>
        public readonly int Default;
        /// <summary>
        /// The maximum number of instances for this resource. Valid values are between `0` and `1000`.
        /// </summary>
        public readonly int Maximum;
        /// <summary>
        /// The minimum number of instances for this resource. Valid values are between `0` and `1000`.
        /// </summary>
        public readonly int Minimum;

        [OutputConstructor]
        private AutoscaleSettingProfileCapacity(
            int @default,

            int maximum,

            int minimum)
        {
            Default = @default;
            Maximum = maximum;
            Minimum = minimum;
        }
    }
}