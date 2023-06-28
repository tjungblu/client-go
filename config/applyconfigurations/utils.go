// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/config/v1"
	v1alpha1 "github.com/openshift/api/config/v1alpha1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	configv1alpha1 "github.com/openshift/client-go/config/applyconfigurations/config/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=config.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("AlibabaCloudPlatformStatus"):
		return &configv1.AlibabaCloudPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlibabaCloudResourceTag"):
		return &configv1.AlibabaCloudResourceTagApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("APIServer"):
		return &configv1.APIServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("APIServerEncryption"):
		return &configv1.APIServerEncryptionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("APIServerNamedServingCert"):
		return &configv1.APIServerNamedServingCertApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("APIServerServingCerts"):
		return &configv1.APIServerServingCertsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("APIServerSpec"):
		return &configv1.APIServerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Audit"):
		return &configv1.AuditApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AuditCustomRule"):
		return &configv1.AuditCustomRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Authentication"):
		return &configv1.AuthenticationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AuthenticationSpec"):
		return &configv1.AuthenticationSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AuthenticationStatus"):
		return &configv1.AuthenticationStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSDNSSpec"):
		return &configv1.AWSDNSSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSIngressSpec"):
		return &configv1.AWSIngressSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSPlatformSpec"):
		return &configv1.AWSPlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSPlatformStatus"):
		return &configv1.AWSPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSResourceTag"):
		return &configv1.AWSResourceTagApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSServiceEndpoint"):
		return &configv1.AWSServiceEndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AzurePlatformStatus"):
		return &configv1.AzurePlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AzureResourceTag"):
		return &configv1.AzureResourceTagApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BareMetalPlatformLoadBalancer"):
		return &configv1.BareMetalPlatformLoadBalancerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BareMetalPlatformStatus"):
		return &configv1.BareMetalPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BasicAuthIdentityProvider"):
		return &configv1.BasicAuthIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Build"):
		return &configv1.BuildApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildDefaults"):
		return &configv1.BuildDefaultsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildOverrides"):
		return &configv1.BuildOverridesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildSpec"):
		return &configv1.BuildSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CloudControllerManagerStatus"):
		return &configv1.CloudControllerManagerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterCondition"):
		return &configv1.ClusterConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterNetworkEntry"):
		return &configv1.ClusterNetworkEntryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterOperator"):
		return &configv1.ClusterOperatorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterOperatorStatus"):
		return &configv1.ClusterOperatorStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterOperatorStatusCondition"):
		return &configv1.ClusterOperatorStatusConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterVersion"):
		return &configv1.ClusterVersionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterVersionCapabilitiesSpec"):
		return &configv1.ClusterVersionCapabilitiesSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterVersionCapabilitiesStatus"):
		return &configv1.ClusterVersionCapabilitiesStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterVersionSpec"):
		return &configv1.ClusterVersionSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterVersionStatus"):
		return &configv1.ClusterVersionStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ComponentOverride"):
		return &configv1.ComponentOverrideApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ComponentRouteSpec"):
		return &configv1.ComponentRouteSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ComponentRouteStatus"):
		return &configv1.ComponentRouteStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConditionalUpdate"):
		return &configv1.ConditionalUpdateApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConditionalUpdateRisk"):
		return &configv1.ConditionalUpdateRiskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConfigMapFileReference"):
		return &configv1.ConfigMapFileReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConfigMapNameReference"):
		return &configv1.ConfigMapNameReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Console"):
		return &configv1.ConsoleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleAuthentication"):
		return &configv1.ConsoleAuthenticationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSpec"):
		return &configv1.ConsoleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleStatus"):
		return &configv1.ConsoleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomFeatureGates"):
		return &configv1.CustomFeatureGatesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomTLSProfile"):
		return &configv1.CustomTLSProfileApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeprecatedWebhookTokenAuthenticator"):
		return &configv1.DeprecatedWebhookTokenAuthenticatorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNS"):
		return &configv1.DNSApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSPlatformSpec"):
		return &configv1.DNSPlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSSpec"):
		return &configv1.DNSSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSZone"):
		return &configv1.DNSZoneApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EquinixMetalPlatformStatus"):
		return &configv1.EquinixMetalPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternalIPConfig"):
		return &configv1.ExternalIPConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternalIPPolicy"):
		return &configv1.ExternalIPPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternalPlatformSpec"):
		return &configv1.ExternalPlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternalPlatformStatus"):
		return &configv1.ExternalPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FeatureGate"):
		return &configv1.FeatureGateApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FeatureGateAttributes"):
		return &configv1.FeatureGateAttributesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FeatureGateDetails"):
		return &configv1.FeatureGateDetailsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FeatureGateSelection"):
		return &configv1.FeatureGateSelectionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FeatureGateSpec"):
		return &configv1.FeatureGateSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FeatureGateStatus"):
		return &configv1.FeatureGateStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GCPPlatformStatus"):
		return &configv1.GCPPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GitHubIdentityProvider"):
		return &configv1.GitHubIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GitLabIdentityProvider"):
		return &configv1.GitLabIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GoogleIdentityProvider"):
		return &configv1.GoogleIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTPasswdIdentityProvider"):
		return &configv1.HTPasswdIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HubSource"):
		return &configv1.HubSourceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HubSourceStatus"):
		return &configv1.HubSourceStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IBMCloudPlatformStatus"):
		return &configv1.IBMCloudPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IdentityProvider"):
		return &configv1.IdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IdentityProviderConfig"):
		return &configv1.IdentityProviderConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Image"):
		return &configv1.ImageApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageContentPolicy"):
		return &configv1.ImageContentPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageContentPolicySpec"):
		return &configv1.ImageContentPolicySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageDigestMirrors"):
		return &configv1.ImageDigestMirrorsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageDigestMirrorSet"):
		return &configv1.ImageDigestMirrorSetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageDigestMirrorSetSpec"):
		return &configv1.ImageDigestMirrorSetSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageLabel"):
		return &configv1.ImageLabelApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageSpec"):
		return &configv1.ImageSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageStatus"):
		return &configv1.ImageStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageTagMirrors"):
		return &configv1.ImageTagMirrorsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageTagMirrorSet"):
		return &configv1.ImageTagMirrorSetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageTagMirrorSetSpec"):
		return &configv1.ImageTagMirrorSetSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Infrastructure"):
		return &configv1.InfrastructureApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("InfrastructureSpec"):
		return &configv1.InfrastructureSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("InfrastructureStatus"):
		return &configv1.InfrastructureStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Ingress"):
		return &configv1.IngressApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressPlatformSpec"):
		return &configv1.IngressPlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressSpec"):
		return &configv1.IngressSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressStatus"):
		return &configv1.IngressStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KeystoneIdentityProvider"):
		return &configv1.KeystoneIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubevirtPlatformStatus"):
		return &configv1.KubevirtPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LDAPAttributeMapping"):
		return &configv1.LDAPAttributeMappingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LDAPIdentityProvider"):
		return &configv1.LDAPIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LoadBalancer"):
		return &configv1.LoadBalancerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MaxAgePolicy"):
		return &configv1.MaxAgePolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MTUMigration"):
		return &configv1.MTUMigrationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MTUMigrationValues"):
		return &configv1.MTUMigrationValuesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Network"):
		return &configv1.NetworkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkMigration"):
		return &configv1.NetworkMigrationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkSpec"):
		return &configv1.NetworkSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkStatus"):
		return &configv1.NetworkStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Node"):
		return &configv1.NodeApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeSpec"):
		return &configv1.NodeSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NutanixPlatformLoadBalancer"):
		return &configv1.NutanixPlatformLoadBalancerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NutanixPlatformSpec"):
		return &configv1.NutanixPlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NutanixPlatformStatus"):
		return &configv1.NutanixPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NutanixPrismElementEndpoint"):
		return &configv1.NutanixPrismElementEndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NutanixPrismEndpoint"):
		return &configv1.NutanixPrismEndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OAuth"):
		return &configv1.OAuthApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OAuthRemoteConnectionInfo"):
		return &configv1.OAuthRemoteConnectionInfoApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OAuthSpec"):
		return &configv1.OAuthSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OAuthTemplates"):
		return &configv1.OAuthTemplatesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjectReference"):
		return &configv1.ObjectReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenIDClaims"):
		return &configv1.OpenIDClaimsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenIDIdentityProvider"):
		return &configv1.OpenIDIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenStackPlatformLoadBalancer"):
		return &configv1.OpenStackPlatformLoadBalancerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenStackPlatformStatus"):
		return &configv1.OpenStackPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OperandVersion"):
		return &configv1.OperandVersionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OperatorHub"):
		return &configv1.OperatorHubApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OperatorHubSpec"):
		return &configv1.OperatorHubSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OperatorHubStatus"):
		return &configv1.OperatorHubStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvirtPlatformLoadBalancer"):
		return &configv1.OvirtPlatformLoadBalancerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OvirtPlatformStatus"):
		return &configv1.OvirtPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PlatformSpec"):
		return &configv1.PlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PlatformStatus"):
		return &configv1.PlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PowerVSPlatformSpec"):
		return &configv1.PowerVSPlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PowerVSPlatformStatus"):
		return &configv1.PowerVSPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PowerVSServiceEndpoint"):
		return &configv1.PowerVSServiceEndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Project"):
		return &configv1.ProjectApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProjectSpec"):
		return &configv1.ProjectSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PromQLClusterCondition"):
		return &configv1.PromQLClusterConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Proxy"):
		return &configv1.ProxyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProxySpec"):
		return &configv1.ProxySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProxyStatus"):
		return &configv1.ProxyStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RegistryLocation"):
		return &configv1.RegistryLocationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RegistrySources"):
		return &configv1.RegistrySourcesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Release"):
		return &configv1.ReleaseApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RepositoryDigestMirrors"):
		return &configv1.RepositoryDigestMirrorsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RequestHeaderIdentityProvider"):
		return &configv1.RequestHeaderIdentityProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RequiredHSTSPolicy"):
		return &configv1.RequiredHSTSPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Scheduler"):
		return &configv1.SchedulerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SchedulerSpec"):
		return &configv1.SchedulerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SecretNameReference"):
		return &configv1.SecretNameReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TemplateReference"):
		return &configv1.TemplateReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TLSProfileSpec"):
		return &configv1.TLSProfileSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TLSSecurityProfile"):
		return &configv1.TLSSecurityProfileApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TokenConfig"):
		return &configv1.TokenConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Update"):
		return &configv1.UpdateApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("UpdateHistory"):
		return &configv1.UpdateHistoryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformFailureDomainSpec"):
		return &configv1.VSpherePlatformFailureDomainSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformLoadBalancer"):
		return &configv1.VSpherePlatformLoadBalancerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformNodeNetworking"):
		return &configv1.VSpherePlatformNodeNetworkingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformNodeNetworkingSpec"):
		return &configv1.VSpherePlatformNodeNetworkingSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformSpec"):
		return &configv1.VSpherePlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformStatus"):
		return &configv1.VSpherePlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformTopology"):
		return &configv1.VSpherePlatformTopologyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSpherePlatformVCenterSpec"):
		return &configv1.VSpherePlatformVCenterSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebhookTokenAuthenticator"):
		return &configv1.WebhookTokenAuthenticatorApplyConfiguration{}

		// Group=config.openshift.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("Backup"):
		return &configv1alpha1.BackupApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BackupSpec"):
		return &configv1alpha1.BackupSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EtcdBackupSpec"):
		return &configv1alpha1.EtcdBackupSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GatherConfig"):
		return &configv1alpha1.GatherConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InsightsDataGather"):
		return &configv1alpha1.InsightsDataGatherApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InsightsDataGatherSpec"):
		return &configv1alpha1.InsightsDataGatherSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RetentionCountConfig"):
		return &configv1alpha1.RetentionCountConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RetentionPolicy"):
		return &configv1alpha1.RetentionPolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RetentionSizeConfig"):
		return &configv1alpha1.RetentionSizeConfigApplyConfiguration{}

	}
	return nil
}
