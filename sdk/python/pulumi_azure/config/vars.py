# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

__config__ = pulumi.Config('azure')

client_certificate_password = __config__.get('clientCertificatePassword')

client_certificate_path = __config__.get('clientCertificatePath')

client_id = __config__.get('clientId') or (utilities.get_env('ARM_CLIENT_ID') or '')

client_secret = __config__.get('clientSecret') or (utilities.get_env('ARM_CLIENT_SECRET') or '')

environment = utilities.require_with_default(lambda: __config__.require('environment'), (utilities.get_env('ARM_ENVIRONMENT') or 'public'))

msi_endpoint = __config__.get('msiEndpoint') or (utilities.get_env('ARM_MSI_ENDPOINT') or '')

partner_id = __config__.get('partnerId')

skip_credentials_validation = __config__.get('skipCredentialsValidation') or (utilities.get_env_bool('ARM_SKIP_CREDENTIALS_VALIDATION') or False)

skip_provider_registration = __config__.get('skipProviderRegistration') or (utilities.get_env_bool('ARM_SKIP_PROVIDER_REGISTRATION') or False)

subscription_id = __config__.get('subscriptionId') or (utilities.get_env('ARM_SUBSCRIPTION_ID') or '')

tenant_id = __config__.get('tenantId') or (utilities.get_env('ARM_TENANT_ID') or '')

use_msi = __config__.get('useMsi') or (utilities.get_env_bool('ARM_USE_MSI') or False)

