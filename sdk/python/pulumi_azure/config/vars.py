# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('azure')

auxiliary_tenant_ids = __config__.get('auxiliaryTenantIds')

client_certificate_password = __config__.get('clientCertificatePassword') or (utilities.get_env('AZURE_CLIENT_CERTIFICATE_PASSWORD', 'ARM_CLIENT_CERTIFICATE_PASSWORD') or '')
"""
The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client
Certificate
"""

client_certificate_path = __config__.get('clientCertificatePath') or (utilities.get_env('AZURE_CLIENT_CERTIFICATE_PATH', 'ARM_CLIENT_CERTIFICATE_PATH') or '')
"""
The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
Principal using a Client Certificate.
"""

client_id = __config__.get('clientId') or (utilities.get_env('AZURE_CLIENT_ID', 'ARM_CLIENT_ID') or '')
"""
The Client ID which should be used.
"""

client_secret = __config__.get('clientSecret') or (utilities.get_env('AZURE_CLIENT_SECRET', 'ARM_CLIENT_SECRET') or '')
"""
The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
"""

disable_correlation_request_id = __config__.get('disableCorrelationRequestId')
"""
This will disable the x-ms-correlation-request-id header.
"""

disable_terraform_partner_id = __config__.get('disableTerraformPartnerId') or (utilities.get_env_bool('ARM_DISABLE_TERRAFORM_PARTNER_ID') or True)
"""
This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
"""

environment = __config__.get('environment') or (utilities.get_env('AZURE_ENVIRONMENT', 'ARM_ENVIRONMENT') or 'public')
"""
The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to
public.
"""

features = __config__.get('features')

location = __config__.get('location') or utilities.get_env('ARM_LOCATION')

msi_endpoint = __config__.get('msiEndpoint') or (utilities.get_env('ARM_MSI_ENDPOINT') or '')
"""
The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
automatically.
"""

partner_id = __config__.get('partnerId') or (utilities.get_env('ARM_PARTNER_ID') or '')
"""
A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
"""

skip_credentials_validation = __config__.get('skipCredentialsValidation') or (utilities.get_env_bool('ARM_SKIP_CREDENTIALS_VALIDATION') or False)
"""
This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
"""

skip_provider_registration = __config__.get('skipProviderRegistration') or (utilities.get_env_bool('ARM_SKIP_PROVIDER_REGISTRATION') or False)
"""
Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already
registered?
"""

storage_use_azuread = __config__.get('storageUseAzuread') or (utilities.get_env_bool('ARM_STORAGE_USE_AZUREAD') or False)
"""
Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
"""

subscription_id = __config__.get('subscriptionId') or (utilities.get_env('ARM_SUBSCRIPTION_ID') or '')
"""
The Subscription ID which should be used.
"""

tenant_id = __config__.get('tenantId') or (utilities.get_env('AZURE_TENANT_ID', 'ARM_TENANT_ID') or '')
"""
The Tenant ID which should be used.
"""

use_msi = __config__.get('useMsi') or (utilities.get_env_bool('ARM_USE_MSI') or False)
"""
Allowed Managed Service Identity be used for Authentication.
"""

