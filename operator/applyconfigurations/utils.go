// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/operator/v1"
	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	operatorv1alpha1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=operator.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("AccessLogging"):
		return &operatorv1.AccessLoggingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AdditionalNetworkDefinition"):
		return &operatorv1.AdditionalNetworkDefinitionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AddPage"):
		return &operatorv1.AddPageApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Authentication"):
		return &operatorv1.AuthenticationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AuthenticationSpec"):
		return &operatorv1.AuthenticationSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AuthenticationStatus"):
		return &operatorv1.AuthenticationStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSClassicLoadBalancerParameters"):
		return &operatorv1.AWSClassicLoadBalancerParametersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSCSIDriverConfigSpec"):
		return &operatorv1.AWSCSIDriverConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AWSLoadBalancerParameters"):
		return &operatorv1.AWSLoadBalancerParametersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AzureCSIDriverConfigSpec"):
		return &operatorv1.AzureCSIDriverConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AzureDiskEncryptionSet"):
		return &operatorv1.AzureDiskEncryptionSetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClientTLS"):
		return &operatorv1.ClientTLSApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CloudCredential"):
		return &operatorv1.CloudCredentialApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CloudCredentialSpec"):
		return &operatorv1.CloudCredentialSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CloudCredentialStatus"):
		return &operatorv1.CloudCredentialStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterCSIDriver"):
		return &operatorv1.ClusterCSIDriverApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterCSIDriverSpec"):
		return &operatorv1.ClusterCSIDriverSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterCSIDriverStatus"):
		return &operatorv1.ClusterCSIDriverStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterNetworkEntry"):
		return &operatorv1.ClusterNetworkEntryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Config"):
		return &operatorv1.ConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConfigSpec"):
		return &operatorv1.ConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConfigStatus"):
		return &operatorv1.ConfigStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Console"):
		return &operatorv1.ConsoleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleConfigRoute"):
		return &operatorv1.ConsoleConfigRouteApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleCustomization"):
		return &operatorv1.ConsoleCustomizationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleProviders"):
		return &operatorv1.ConsoleProvidersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleSpec"):
		return &operatorv1.ConsoleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConsoleStatus"):
		return &operatorv1.ConsoleStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ContainerLoggingDestinationParameters"):
		return &operatorv1.ContainerLoggingDestinationParametersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CSIDriverConfigSpec"):
		return &operatorv1.CSIDriverConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CSISnapshotController"):
		return &operatorv1.CSISnapshotControllerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CSISnapshotControllerSpec"):
		return &operatorv1.CSISnapshotControllerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CSISnapshotControllerStatus"):
		return &operatorv1.CSISnapshotControllerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DefaultNetworkDefinition"):
		return &operatorv1.DefaultNetworkDefinitionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeveloperConsoleCatalogCategory"):
		return &operatorv1.DeveloperConsoleCatalogCategoryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeveloperConsoleCatalogCategoryMeta"):
		return &operatorv1.DeveloperConsoleCatalogCategoryMetaApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeveloperConsoleCatalogCustomization"):
		return &operatorv1.DeveloperConsoleCatalogCustomizationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeveloperConsoleCatalogTypes"):
		return &operatorv1.DeveloperConsoleCatalogTypesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNS"):
		return &operatorv1.DNSApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSCache"):
		return &operatorv1.DNSCacheApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSNodePlacement"):
		return &operatorv1.DNSNodePlacementApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSOverTLSConfig"):
		return &operatorv1.DNSOverTLSConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSSpec"):
		return &operatorv1.DNSSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSStatus"):
		return &operatorv1.DNSStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DNSTransportConfig"):
		return &operatorv1.DNSTransportConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EgressIPConfig"):
		return &operatorv1.EgressIPConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EndpointPublishingStrategy"):
		return &operatorv1.EndpointPublishingStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Etcd"):
		return &operatorv1.EtcdApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EtcdSpec"):
		return &operatorv1.EtcdSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EtcdStatus"):
		return &operatorv1.EtcdStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExportNetworkFlows"):
		return &operatorv1.ExportNetworkFlowsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FeaturesMigration"):
		return &operatorv1.FeaturesMigrationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ForwardPlugin"):
		return &operatorv1.ForwardPluginApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatewayConfig"):
		return &operatorv1.GatewayConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GathererStatus"):
		return &operatorv1.GathererStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GatherStatus"):
		return &operatorv1.GatherStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GCPCSIDriverConfigSpec"):
		return &operatorv1.GCPCSIDriverConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GCPKMSKeyReference"):
		return &operatorv1.GCPKMSKeyReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GCPLoadBalancerParameters"):
		return &operatorv1.GCPLoadBalancerParametersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GenerationStatus"):
		return &operatorv1.GenerationStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HealthCheck"):
		return &operatorv1.HealthCheckApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HostNetworkStrategy"):
		return &operatorv1.HostNetworkStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPCompressionPolicy"):
		return &operatorv1.HTTPCompressionPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HybridOverlayConfig"):
		return &operatorv1.HybridOverlayConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IBMLoadBalancerParameters"):
		return &operatorv1.IBMLoadBalancerParametersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressController"):
		return &operatorv1.IngressControllerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerCaptureHTTPCookie"):
		return &operatorv1.IngressControllerCaptureHTTPCookieApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerCaptureHTTPCookieUnion"):
		return &operatorv1.IngressControllerCaptureHTTPCookieUnionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerCaptureHTTPHeader"):
		return &operatorv1.IngressControllerCaptureHTTPHeaderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerCaptureHTTPHeaders"):
		return &operatorv1.IngressControllerCaptureHTTPHeadersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerHTTPHeaders"):
		return &operatorv1.IngressControllerHTTPHeadersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerHTTPUniqueIdHeaderPolicy"):
		return &operatorv1.IngressControllerHTTPUniqueIdHeaderPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerLogging"):
		return &operatorv1.IngressControllerLoggingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerSpec"):
		return &operatorv1.IngressControllerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerStatus"):
		return &operatorv1.IngressControllerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IngressControllerTuningOptions"):
		return &operatorv1.IngressControllerTuningOptionsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("InsightsOperator"):
		return &operatorv1.InsightsOperatorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("InsightsOperatorSpec"):
		return &operatorv1.InsightsOperatorSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("InsightsOperatorStatus"):
		return &operatorv1.InsightsOperatorStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("InsightsReport"):
		return &operatorv1.InsightsReportApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IPAMConfig"):
		return &operatorv1.IPAMConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IPFIXConfig"):
		return &operatorv1.IPFIXConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeAPIServer"):
		return &operatorv1.KubeAPIServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeAPIServerSpec"):
		return &operatorv1.KubeAPIServerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeAPIServerStatus"):
		return &operatorv1.KubeAPIServerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeControllerManager"):
		return &operatorv1.KubeControllerManagerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeControllerManagerSpec"):
		return &operatorv1.KubeControllerManagerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeControllerManagerStatus"):
		return &operatorv1.KubeControllerManagerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeScheduler"):
		return &operatorv1.KubeSchedulerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeSchedulerSpec"):
		return &operatorv1.KubeSchedulerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeSchedulerStatus"):
		return &operatorv1.KubeSchedulerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeStorageVersionMigrator"):
		return &operatorv1.KubeStorageVersionMigratorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeStorageVersionMigratorSpec"):
		return &operatorv1.KubeStorageVersionMigratorSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeStorageVersionMigratorStatus"):
		return &operatorv1.KubeStorageVersionMigratorStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KuryrConfig"):
		return &operatorv1.KuryrConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LoadBalancerStrategy"):
		return &operatorv1.LoadBalancerStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LoggingDestination"):
		return &operatorv1.LoggingDestinationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MTUMigration"):
		return &operatorv1.MTUMigrationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MTUMigrationValues"):
		return &operatorv1.MTUMigrationValuesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetFlowConfig"):
		return &operatorv1.NetFlowConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Network"):
		return &operatorv1.NetworkApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkMigration"):
		return &operatorv1.NetworkMigrationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkSpec"):
		return &operatorv1.NetworkSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkStatus"):
		return &operatorv1.NetworkStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodePlacement"):
		return &operatorv1.NodePlacementApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodePortStrategy"):
		return &operatorv1.NodePortStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NodeStatus"):
		return &operatorv1.NodeStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OAuthAPIServerStatus"):
		return &operatorv1.OAuthAPIServerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenShiftAPIServer"):
		return &operatorv1.OpenShiftAPIServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenShiftAPIServerSpec"):
		return &operatorv1.OpenShiftAPIServerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenShiftAPIServerStatus"):
		return &operatorv1.OpenShiftAPIServerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenShiftControllerManager"):
		return &operatorv1.OpenShiftControllerManagerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenShiftControllerManagerSpec"):
		return &operatorv1.OpenShiftControllerManagerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenShiftControllerManagerStatus"):
		return &operatorv1.OpenShiftControllerManagerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OpenShiftSDNConfig"):
		return &operatorv1.OpenShiftSDNConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OperatorCondition"):
		return &operatorv1.OperatorConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OperatorSpec"):
		return &operatorv1.OperatorSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OperatorStatus"):
		return &operatorv1.OperatorStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OVNKubernetesConfig"):
		return &operatorv1.OVNKubernetesConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Perspective"):
		return &operatorv1.PerspectiveApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PerspectiveVisibility"):
		return &operatorv1.PerspectiveVisibilityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PinnedResourceReference"):
		return &operatorv1.PinnedResourceReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicyAuditConfig"):
		return &operatorv1.PolicyAuditConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrivateStrategy"):
		return &operatorv1.PrivateStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProjectAccess"):
		return &operatorv1.ProjectAccessApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProviderLoadBalancerParameters"):
		return &operatorv1.ProviderLoadBalancerParametersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProxyConfig"):
		return &operatorv1.ProxyConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QuickStarts"):
		return &operatorv1.QuickStartsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ResourceAttributesAccessReview"):
		return &operatorv1.ResourceAttributesAccessReviewApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RouteAdmissionPolicy"):
		return &operatorv1.RouteAdmissionPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Server"):
		return &operatorv1.ServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceAccountIssuerStatus"):
		return &operatorv1.ServiceAccountIssuerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCA"):
		return &operatorv1.ServiceCAApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCASpec"):
		return &operatorv1.ServiceCASpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCAStatus"):
		return &operatorv1.ServiceCAStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCatalogAPIServer"):
		return &operatorv1.ServiceCatalogAPIServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCatalogAPIServerSpec"):
		return &operatorv1.ServiceCatalogAPIServerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCatalogAPIServerStatus"):
		return &operatorv1.ServiceCatalogAPIServerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCatalogControllerManager"):
		return &operatorv1.ServiceCatalogControllerManagerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCatalogControllerManagerSpec"):
		return &operatorv1.ServiceCatalogControllerManagerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCatalogControllerManagerStatus"):
		return &operatorv1.ServiceCatalogControllerManagerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SFlowConfig"):
		return &operatorv1.SFlowConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SimpleMacvlanConfig"):
		return &operatorv1.SimpleMacvlanConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticIPAMAddresses"):
		return &operatorv1.StaticIPAMAddressesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticIPAMConfig"):
		return &operatorv1.StaticIPAMConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticIPAMDNS"):
		return &operatorv1.StaticIPAMDNSApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticIPAMRoutes"):
		return &operatorv1.StaticIPAMRoutesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticPodOperatorSpec"):
		return &operatorv1.StaticPodOperatorSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticPodOperatorStatus"):
		return &operatorv1.StaticPodOperatorStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StatuspageProvider"):
		return &operatorv1.StatuspageProviderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Storage"):
		return &operatorv1.StorageApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StorageSpec"):
		return &operatorv1.StorageSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StorageStatus"):
		return &operatorv1.StorageStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SyslogLoggingDestinationParameters"):
		return &operatorv1.SyslogLoggingDestinationParametersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Upstream"):
		return &operatorv1.UpstreamApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("UpstreamResolvers"):
		return &operatorv1.UpstreamResolversApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("VSphereCSIDriverConfigSpec"):
		return &operatorv1.VSphereCSIDriverConfigSpecApplyConfiguration{}

		// Group=operator.openshift.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("BackupJobReference"):
		return &operatorv1alpha1.BackupJobReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EtcdBackup"):
		return &operatorv1alpha1.EtcdBackupApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EtcdBackupSpec"):
		return &operatorv1alpha1.EtcdBackupSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EtcdBackupStatus"):
		return &operatorv1alpha1.EtcdBackupStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageContentSourcePolicy"):
		return &operatorv1alpha1.ImageContentSourcePolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ImageContentSourcePolicySpec"):
		return &operatorv1alpha1.ImageContentSourcePolicySpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OLM"):
		return &operatorv1alpha1.OLMApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OLMSpec"):
		return &operatorv1alpha1.OLMSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OLMStatus"):
		return &operatorv1alpha1.OLMStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RepositoryDigestMirrors"):
		return &operatorv1alpha1.RepositoryDigestMirrorsApplyConfiguration{}

	}
	return nil
}
