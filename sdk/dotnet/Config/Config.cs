// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Gcp
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("gcp");

        private static readonly __Value<string?> _accessApprovalCustomEndpoint = new __Value<string?>(() => __config.Get("accessApprovalCustomEndpoint"));
        public static string? AccessApprovalCustomEndpoint
        {
            get => _accessApprovalCustomEndpoint.Get();
            set => _accessApprovalCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _accessContextManagerCustomEndpoint = new __Value<string?>(() => __config.Get("accessContextManagerCustomEndpoint"));
        public static string? AccessContextManagerCustomEndpoint
        {
            get => _accessContextManagerCustomEndpoint.Get();
            set => _accessContextManagerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _accessToken = new __Value<string?>(() => __config.Get("accessToken"));
        public static string? AccessToken
        {
            get => _accessToken.Get();
            set => _accessToken.Set(value);
        }

        private static readonly __Value<string?> _activeDirectoryCustomEndpoint = new __Value<string?>(() => __config.Get("activeDirectoryCustomEndpoint"));
        public static string? ActiveDirectoryCustomEndpoint
        {
            get => _activeDirectoryCustomEndpoint.Get();
            set => _activeDirectoryCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _apiGatewayCustomEndpoint = new __Value<string?>(() => __config.Get("apiGatewayCustomEndpoint"));
        public static string? ApiGatewayCustomEndpoint
        {
            get => _apiGatewayCustomEndpoint.Get();
            set => _apiGatewayCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _apigeeCustomEndpoint = new __Value<string?>(() => __config.Get("apigeeCustomEndpoint"));
        public static string? ApigeeCustomEndpoint
        {
            get => _apigeeCustomEndpoint.Get();
            set => _apigeeCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _appEngineCustomEndpoint = new __Value<string?>(() => __config.Get("appEngineCustomEndpoint"));
        public static string? AppEngineCustomEndpoint
        {
            get => _appEngineCustomEndpoint.Get();
            set => _appEngineCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _artifactRegistryCustomEndpoint = new __Value<string?>(() => __config.Get("artifactRegistryCustomEndpoint"));
        public static string? ArtifactRegistryCustomEndpoint
        {
            get => _artifactRegistryCustomEndpoint.Get();
            set => _artifactRegistryCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _assuredWorkloadsCustomEndpoint = new __Value<string?>(() => __config.Get("assuredWorkloadsCustomEndpoint"));
        public static string? AssuredWorkloadsCustomEndpoint
        {
            get => _assuredWorkloadsCustomEndpoint.Get();
            set => _assuredWorkloadsCustomEndpoint.Set(value);
        }

        private static readonly __Value<Pulumi.Gcp.Config.Types.Batching?> _batching = new __Value<Pulumi.Gcp.Config.Types.Batching?>(() => __config.GetObject<Pulumi.Gcp.Config.Types.Batching>("batching"));
        public static Pulumi.Gcp.Config.Types.Batching? Batching
        {
            get => _batching.Get();
            set => _batching.Set(value);
        }

        private static readonly __Value<string?> _bigQueryCustomEndpoint = new __Value<string?>(() => __config.Get("bigQueryCustomEndpoint"));
        public static string? BigQueryCustomEndpoint
        {
            get => _bigQueryCustomEndpoint.Get();
            set => _bigQueryCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _bigqueryConnectionCustomEndpoint = new __Value<string?>(() => __config.Get("bigqueryConnectionCustomEndpoint"));
        public static string? BigqueryConnectionCustomEndpoint
        {
            get => _bigqueryConnectionCustomEndpoint.Get();
            set => _bigqueryConnectionCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _bigqueryDataTransferCustomEndpoint = new __Value<string?>(() => __config.Get("bigqueryDataTransferCustomEndpoint"));
        public static string? BigqueryDataTransferCustomEndpoint
        {
            get => _bigqueryDataTransferCustomEndpoint.Get();
            set => _bigqueryDataTransferCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _bigqueryReservationCustomEndpoint = new __Value<string?>(() => __config.Get("bigqueryReservationCustomEndpoint"));
        public static string? BigqueryReservationCustomEndpoint
        {
            get => _bigqueryReservationCustomEndpoint.Get();
            set => _bigqueryReservationCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _bigtableCustomEndpoint = new __Value<string?>(() => __config.Get("bigtableCustomEndpoint"));
        public static string? BigtableCustomEndpoint
        {
            get => _bigtableCustomEndpoint.Get();
            set => _bigtableCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _billingCustomEndpoint = new __Value<string?>(() => __config.Get("billingCustomEndpoint"));
        public static string? BillingCustomEndpoint
        {
            get => _billingCustomEndpoint.Get();
            set => _billingCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _billingProject = new __Value<string?>(() => __config.Get("billingProject"));
        public static string? BillingProject
        {
            get => _billingProject.Get();
            set => _billingProject.Set(value);
        }

        private static readonly __Value<string?> _binaryAuthorizationCustomEndpoint = new __Value<string?>(() => __config.Get("binaryAuthorizationCustomEndpoint"));
        public static string? BinaryAuthorizationCustomEndpoint
        {
            get => _binaryAuthorizationCustomEndpoint.Get();
            set => _binaryAuthorizationCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudAssetCustomEndpoint = new __Value<string?>(() => __config.Get("cloudAssetCustomEndpoint"));
        public static string? CloudAssetCustomEndpoint
        {
            get => _cloudAssetCustomEndpoint.Get();
            set => _cloudAssetCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudBillingCustomEndpoint = new __Value<string?>(() => __config.Get("cloudBillingCustomEndpoint"));
        public static string? CloudBillingCustomEndpoint
        {
            get => _cloudBillingCustomEndpoint.Get();
            set => _cloudBillingCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudBuildCustomEndpoint = new __Value<string?>(() => __config.Get("cloudBuildCustomEndpoint"));
        public static string? CloudBuildCustomEndpoint
        {
            get => _cloudBuildCustomEndpoint.Get();
            set => _cloudBuildCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudBuildWorkerPoolCustomEndpoint = new __Value<string?>(() => __config.Get("cloudBuildWorkerPoolCustomEndpoint"));
        public static string? CloudBuildWorkerPoolCustomEndpoint
        {
            get => _cloudBuildWorkerPoolCustomEndpoint.Get();
            set => _cloudBuildWorkerPoolCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudFunctionsCustomEndpoint = new __Value<string?>(() => __config.Get("cloudFunctionsCustomEndpoint"));
        public static string? CloudFunctionsCustomEndpoint
        {
            get => _cloudFunctionsCustomEndpoint.Get();
            set => _cloudFunctionsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudIdentityCustomEndpoint = new __Value<string?>(() => __config.Get("cloudIdentityCustomEndpoint"));
        public static string? CloudIdentityCustomEndpoint
        {
            get => _cloudIdentityCustomEndpoint.Get();
            set => _cloudIdentityCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudIotCustomEndpoint = new __Value<string?>(() => __config.Get("cloudIotCustomEndpoint"));
        public static string? CloudIotCustomEndpoint
        {
            get => _cloudIotCustomEndpoint.Get();
            set => _cloudIotCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudResourceManagerCustomEndpoint = new __Value<string?>(() => __config.Get("cloudResourceManagerCustomEndpoint"));
        public static string? CloudResourceManagerCustomEndpoint
        {
            get => _cloudResourceManagerCustomEndpoint.Get();
            set => _cloudResourceManagerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudRunCustomEndpoint = new __Value<string?>(() => __config.Get("cloudRunCustomEndpoint"));
        public static string? CloudRunCustomEndpoint
        {
            get => _cloudRunCustomEndpoint.Get();
            set => _cloudRunCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudSchedulerCustomEndpoint = new __Value<string?>(() => __config.Get("cloudSchedulerCustomEndpoint"));
        public static string? CloudSchedulerCustomEndpoint
        {
            get => _cloudSchedulerCustomEndpoint.Get();
            set => _cloudSchedulerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _cloudTasksCustomEndpoint = new __Value<string?>(() => __config.Get("cloudTasksCustomEndpoint"));
        public static string? CloudTasksCustomEndpoint
        {
            get => _cloudTasksCustomEndpoint.Get();
            set => _cloudTasksCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _composerCustomEndpoint = new __Value<string?>(() => __config.Get("composerCustomEndpoint"));
        public static string? ComposerCustomEndpoint
        {
            get => _composerCustomEndpoint.Get();
            set => _composerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _computeCustomEndpoint = new __Value<string?>(() => __config.Get("computeCustomEndpoint"));
        public static string? ComputeCustomEndpoint
        {
            get => _computeCustomEndpoint.Get();
            set => _computeCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _containerAnalysisCustomEndpoint = new __Value<string?>(() => __config.Get("containerAnalysisCustomEndpoint"));
        public static string? ContainerAnalysisCustomEndpoint
        {
            get => _containerAnalysisCustomEndpoint.Get();
            set => _containerAnalysisCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _containerAwsCustomEndpoint = new __Value<string?>(() => __config.Get("containerAwsCustomEndpoint"));
        public static string? ContainerAwsCustomEndpoint
        {
            get => _containerAwsCustomEndpoint.Get();
            set => _containerAwsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _containerAzureCustomEndpoint = new __Value<string?>(() => __config.Get("containerAzureCustomEndpoint"));
        public static string? ContainerAzureCustomEndpoint
        {
            get => _containerAzureCustomEndpoint.Get();
            set => _containerAzureCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _containerCustomEndpoint = new __Value<string?>(() => __config.Get("containerCustomEndpoint"));
        public static string? ContainerCustomEndpoint
        {
            get => _containerCustomEndpoint.Get();
            set => _containerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _credentials = new __Value<string?>(() => __config.Get("credentials"));
        public static string? Credentials
        {
            get => _credentials.Get();
            set => _credentials.Set(value);
        }

        private static readonly __Value<string?> _dataCatalogCustomEndpoint = new __Value<string?>(() => __config.Get("dataCatalogCustomEndpoint"));
        public static string? DataCatalogCustomEndpoint
        {
            get => _dataCatalogCustomEndpoint.Get();
            set => _dataCatalogCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _dataFusionCustomEndpoint = new __Value<string?>(() => __config.Get("dataFusionCustomEndpoint"));
        public static string? DataFusionCustomEndpoint
        {
            get => _dataFusionCustomEndpoint.Get();
            set => _dataFusionCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _dataLossPreventionCustomEndpoint = new __Value<string?>(() => __config.Get("dataLossPreventionCustomEndpoint"));
        public static string? DataLossPreventionCustomEndpoint
        {
            get => _dataLossPreventionCustomEndpoint.Get();
            set => _dataLossPreventionCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _dataflowCustomEndpoint = new __Value<string?>(() => __config.Get("dataflowCustomEndpoint"));
        public static string? DataflowCustomEndpoint
        {
            get => _dataflowCustomEndpoint.Get();
            set => _dataflowCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _dataprocCustomEndpoint = new __Value<string?>(() => __config.Get("dataprocCustomEndpoint"));
        public static string? DataprocCustomEndpoint
        {
            get => _dataprocCustomEndpoint.Get();
            set => _dataprocCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _dataprocMetastoreCustomEndpoint = new __Value<string?>(() => __config.Get("dataprocMetastoreCustomEndpoint"));
        public static string? DataprocMetastoreCustomEndpoint
        {
            get => _dataprocMetastoreCustomEndpoint.Get();
            set => _dataprocMetastoreCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _datastoreCustomEndpoint = new __Value<string?>(() => __config.Get("datastoreCustomEndpoint"));
        public static string? DatastoreCustomEndpoint
        {
            get => _datastoreCustomEndpoint.Get();
            set => _datastoreCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _deploymentManagerCustomEndpoint = new __Value<string?>(() => __config.Get("deploymentManagerCustomEndpoint"));
        public static string? DeploymentManagerCustomEndpoint
        {
            get => _deploymentManagerCustomEndpoint.Get();
            set => _deploymentManagerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _dialogflowCustomEndpoint = new __Value<string?>(() => __config.Get("dialogflowCustomEndpoint"));
        public static string? DialogflowCustomEndpoint
        {
            get => _dialogflowCustomEndpoint.Get();
            set => _dialogflowCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _dialogflowCxCustomEndpoint = new __Value<string?>(() => __config.Get("dialogflowCxCustomEndpoint"));
        public static string? DialogflowCxCustomEndpoint
        {
            get => _dialogflowCxCustomEndpoint.Get();
            set => _dialogflowCxCustomEndpoint.Set(value);
        }

        private static readonly __Value<bool?> _disableGooglePartnerName = new __Value<bool?>(() => __config.GetBoolean("disableGooglePartnerName"));
        public static bool? DisableGooglePartnerName
        {
            get => _disableGooglePartnerName.Get();
            set => _disableGooglePartnerName.Set(value);
        }

        private static readonly __Value<string?> _dnsCustomEndpoint = new __Value<string?>(() => __config.Get("dnsCustomEndpoint"));
        public static string? DnsCustomEndpoint
        {
            get => _dnsCustomEndpoint.Get();
            set => _dnsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _essentialContactsCustomEndpoint = new __Value<string?>(() => __config.Get("essentialContactsCustomEndpoint"));
        public static string? EssentialContactsCustomEndpoint
        {
            get => _essentialContactsCustomEndpoint.Get();
            set => _essentialContactsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _eventarcCustomEndpoint = new __Value<string?>(() => __config.Get("eventarcCustomEndpoint"));
        public static string? EventarcCustomEndpoint
        {
            get => _eventarcCustomEndpoint.Get();
            set => _eventarcCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _filestoreCustomEndpoint = new __Value<string?>(() => __config.Get("filestoreCustomEndpoint"));
        public static string? FilestoreCustomEndpoint
        {
            get => _filestoreCustomEndpoint.Get();
            set => _filestoreCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _firebaseCustomEndpoint = new __Value<string?>(() => __config.Get("firebaseCustomEndpoint"));
        public static string? FirebaseCustomEndpoint
        {
            get => _firebaseCustomEndpoint.Get();
            set => _firebaseCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _firestoreCustomEndpoint = new __Value<string?>(() => __config.Get("firestoreCustomEndpoint"));
        public static string? FirestoreCustomEndpoint
        {
            get => _firestoreCustomEndpoint.Get();
            set => _firestoreCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _gameServicesCustomEndpoint = new __Value<string?>(() => __config.Get("gameServicesCustomEndpoint"));
        public static string? GameServicesCustomEndpoint
        {
            get => _gameServicesCustomEndpoint.Get();
            set => _gameServicesCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _gkeHubCustomEndpoint = new __Value<string?>(() => __config.Get("gkeHubCustomEndpoint"));
        public static string? GkeHubCustomEndpoint
        {
            get => _gkeHubCustomEndpoint.Get();
            set => _gkeHubCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _gkehubFeatureCustomEndpoint = new __Value<string?>(() => __config.Get("gkehubFeatureCustomEndpoint"));
        public static string? GkehubFeatureCustomEndpoint
        {
            get => _gkehubFeatureCustomEndpoint.Get();
            set => _gkehubFeatureCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _googlePartnerName = new __Value<string?>(() => __config.Get("googlePartnerName"));
        public static string? GooglePartnerName
        {
            get => _googlePartnerName.Get();
            set => _googlePartnerName.Set(value);
        }

        private static readonly __Value<string?> _healthcareCustomEndpoint = new __Value<string?>(() => __config.Get("healthcareCustomEndpoint"));
        public static string? HealthcareCustomEndpoint
        {
            get => _healthcareCustomEndpoint.Get();
            set => _healthcareCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _iamBetaCustomEndpoint = new __Value<string?>(() => __config.Get("iamBetaCustomEndpoint"));
        public static string? IamBetaCustomEndpoint
        {
            get => _iamBetaCustomEndpoint.Get();
            set => _iamBetaCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _iamCredentialsCustomEndpoint = new __Value<string?>(() => __config.Get("iamCredentialsCustomEndpoint"));
        public static string? IamCredentialsCustomEndpoint
        {
            get => _iamCredentialsCustomEndpoint.Get();
            set => _iamCredentialsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _iamCustomEndpoint = new __Value<string?>(() => __config.Get("iamCustomEndpoint"));
        public static string? IamCustomEndpoint
        {
            get => _iamCustomEndpoint.Get();
            set => _iamCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _iapCustomEndpoint = new __Value<string?>(() => __config.Get("iapCustomEndpoint"));
        public static string? IapCustomEndpoint
        {
            get => _iapCustomEndpoint.Get();
            set => _iapCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _identityPlatformCustomEndpoint = new __Value<string?>(() => __config.Get("identityPlatformCustomEndpoint"));
        public static string? IdentityPlatformCustomEndpoint
        {
            get => _identityPlatformCustomEndpoint.Get();
            set => _identityPlatformCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _impersonateServiceAccount = new __Value<string?>(() => __config.Get("impersonateServiceAccount"));
        public static string? ImpersonateServiceAccount
        {
            get => _impersonateServiceAccount.Get();
            set => _impersonateServiceAccount.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _impersonateServiceAccountDelegates = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("impersonateServiceAccountDelegates"));
        public static ImmutableArray<string> ImpersonateServiceAccountDelegates
        {
            get => _impersonateServiceAccountDelegates.Get();
            set => _impersonateServiceAccountDelegates.Set(value);
        }

        private static readonly __Value<string?> _kmsCustomEndpoint = new __Value<string?>(() => __config.Get("kmsCustomEndpoint"));
        public static string? KmsCustomEndpoint
        {
            get => _kmsCustomEndpoint.Get();
            set => _kmsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _loggingCustomEndpoint = new __Value<string?>(() => __config.Get("loggingCustomEndpoint"));
        public static string? LoggingCustomEndpoint
        {
            get => _loggingCustomEndpoint.Get();
            set => _loggingCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _memcacheCustomEndpoint = new __Value<string?>(() => __config.Get("memcacheCustomEndpoint"));
        public static string? MemcacheCustomEndpoint
        {
            get => _memcacheCustomEndpoint.Get();
            set => _memcacheCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _mlEngineCustomEndpoint = new __Value<string?>(() => __config.Get("mlEngineCustomEndpoint"));
        public static string? MlEngineCustomEndpoint
        {
            get => _mlEngineCustomEndpoint.Get();
            set => _mlEngineCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _monitoringCustomEndpoint = new __Value<string?>(() => __config.Get("monitoringCustomEndpoint"));
        public static string? MonitoringCustomEndpoint
        {
            get => _monitoringCustomEndpoint.Get();
            set => _monitoringCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _networkConnectivityCustomEndpoint = new __Value<string?>(() => __config.Get("networkConnectivityCustomEndpoint"));
        public static string? NetworkConnectivityCustomEndpoint
        {
            get => _networkConnectivityCustomEndpoint.Get();
            set => _networkConnectivityCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _networkManagementCustomEndpoint = new __Value<string?>(() => __config.Get("networkManagementCustomEndpoint"));
        public static string? NetworkManagementCustomEndpoint
        {
            get => _networkManagementCustomEndpoint.Get();
            set => _networkManagementCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _networkServicesCustomEndpoint = new __Value<string?>(() => __config.Get("networkServicesCustomEndpoint"));
        public static string? NetworkServicesCustomEndpoint
        {
            get => _networkServicesCustomEndpoint.Get();
            set => _networkServicesCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _notebooksCustomEndpoint = new __Value<string?>(() => __config.Get("notebooksCustomEndpoint"));
        public static string? NotebooksCustomEndpoint
        {
            get => _notebooksCustomEndpoint.Get();
            set => _notebooksCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _orgPolicyCustomEndpoint = new __Value<string?>(() => __config.Get("orgPolicyCustomEndpoint"));
        public static string? OrgPolicyCustomEndpoint
        {
            get => _orgPolicyCustomEndpoint.Get();
            set => _orgPolicyCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _osConfigCustomEndpoint = new __Value<string?>(() => __config.Get("osConfigCustomEndpoint"));
        public static string? OsConfigCustomEndpoint
        {
            get => _osConfigCustomEndpoint.Get();
            set => _osConfigCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _osLoginCustomEndpoint = new __Value<string?>(() => __config.Get("osLoginCustomEndpoint"));
        public static string? OsLoginCustomEndpoint
        {
            get => _osLoginCustomEndpoint.Get();
            set => _osLoginCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _privatecaCustomEndpoint = new __Value<string?>(() => __config.Get("privatecaCustomEndpoint"));
        public static string? PrivatecaCustomEndpoint
        {
            get => _privatecaCustomEndpoint.Get();
            set => _privatecaCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _project = new __Value<string?>(() => __config.Get("project") ?? Utilities.GetEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT"));
        public static string? Project
        {
            get => _project.Get();
            set => _project.Set(value);
        }

        private static readonly __Value<string?> _pubsubCustomEndpoint = new __Value<string?>(() => __config.Get("pubsubCustomEndpoint"));
        public static string? PubsubCustomEndpoint
        {
            get => _pubsubCustomEndpoint.Get();
            set => _pubsubCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _pubsubLiteCustomEndpoint = new __Value<string?>(() => __config.Get("pubsubLiteCustomEndpoint"));
        public static string? PubsubLiteCustomEndpoint
        {
            get => _pubsubLiteCustomEndpoint.Get();
            set => _pubsubLiteCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _recaptchaEnterpriseCustomEndpoint = new __Value<string?>(() => __config.Get("recaptchaEnterpriseCustomEndpoint"));
        public static string? RecaptchaEnterpriseCustomEndpoint
        {
            get => _recaptchaEnterpriseCustomEndpoint.Get();
            set => _recaptchaEnterpriseCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _redisCustomEndpoint = new __Value<string?>(() => __config.Get("redisCustomEndpoint"));
        public static string? RedisCustomEndpoint
        {
            get => _redisCustomEndpoint.Get();
            set => _redisCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region") ?? Utilities.GetEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION"));
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _requestReason = new __Value<string?>(() => __config.Get("requestReason"));
        public static string? RequestReason
        {
            get => _requestReason.Get();
            set => _requestReason.Set(value);
        }

        private static readonly __Value<string?> _requestTimeout = new __Value<string?>(() => __config.Get("requestTimeout"));
        public static string? RequestTimeout
        {
            get => _requestTimeout.Get();
            set => _requestTimeout.Set(value);
        }

        private static readonly __Value<string?> _resourceManagerCustomEndpoint = new __Value<string?>(() => __config.Get("resourceManagerCustomEndpoint"));
        public static string? ResourceManagerCustomEndpoint
        {
            get => _resourceManagerCustomEndpoint.Get();
            set => _resourceManagerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _resourceManagerV2CustomEndpoint = new __Value<string?>(() => __config.Get("resourceManagerV2CustomEndpoint"));
        public static string? ResourceManagerV2CustomEndpoint
        {
            get => _resourceManagerV2CustomEndpoint.Get();
            set => _resourceManagerV2CustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _runtimeConfigCustomEndpoint = new __Value<string?>(() => __config.Get("runtimeConfigCustomEndpoint"));
        public static string? RuntimeConfigCustomEndpoint
        {
            get => _runtimeConfigCustomEndpoint.Get();
            set => _runtimeConfigCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _runtimeconfigCustomEndpoint = new __Value<string?>(() => __config.Get("runtimeconfigCustomEndpoint"));
        public static string? RuntimeconfigCustomEndpoint
        {
            get => _runtimeconfigCustomEndpoint.Get();
            set => _runtimeconfigCustomEndpoint.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _scopes = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("scopes"));
        public static ImmutableArray<string> Scopes
        {
            get => _scopes.Get();
            set => _scopes.Set(value);
        }

        private static readonly __Value<string?> _secretManagerCustomEndpoint = new __Value<string?>(() => __config.Get("secretManagerCustomEndpoint"));
        public static string? SecretManagerCustomEndpoint
        {
            get => _secretManagerCustomEndpoint.Get();
            set => _secretManagerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _securityCenterCustomEndpoint = new __Value<string?>(() => __config.Get("securityCenterCustomEndpoint"));
        public static string? SecurityCenterCustomEndpoint
        {
            get => _securityCenterCustomEndpoint.Get();
            set => _securityCenterCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _securityScannerCustomEndpoint = new __Value<string?>(() => __config.Get("securityScannerCustomEndpoint"));
        public static string? SecurityScannerCustomEndpoint
        {
            get => _securityScannerCustomEndpoint.Get();
            set => _securityScannerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _serviceDirectoryCustomEndpoint = new __Value<string?>(() => __config.Get("serviceDirectoryCustomEndpoint"));
        public static string? ServiceDirectoryCustomEndpoint
        {
            get => _serviceDirectoryCustomEndpoint.Get();
            set => _serviceDirectoryCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _serviceManagementCustomEndpoint = new __Value<string?>(() => __config.Get("serviceManagementCustomEndpoint"));
        public static string? ServiceManagementCustomEndpoint
        {
            get => _serviceManagementCustomEndpoint.Get();
            set => _serviceManagementCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _serviceNetworkingCustomEndpoint = new __Value<string?>(() => __config.Get("serviceNetworkingCustomEndpoint"));
        public static string? ServiceNetworkingCustomEndpoint
        {
            get => _serviceNetworkingCustomEndpoint.Get();
            set => _serviceNetworkingCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _serviceUsageCustomEndpoint = new __Value<string?>(() => __config.Get("serviceUsageCustomEndpoint"));
        public static string? ServiceUsageCustomEndpoint
        {
            get => _serviceUsageCustomEndpoint.Get();
            set => _serviceUsageCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _sourceRepoCustomEndpoint = new __Value<string?>(() => __config.Get("sourceRepoCustomEndpoint"));
        public static string? SourceRepoCustomEndpoint
        {
            get => _sourceRepoCustomEndpoint.Get();
            set => _sourceRepoCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _spannerCustomEndpoint = new __Value<string?>(() => __config.Get("spannerCustomEndpoint"));
        public static string? SpannerCustomEndpoint
        {
            get => _spannerCustomEndpoint.Get();
            set => _spannerCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _sqlCustomEndpoint = new __Value<string?>(() => __config.Get("sqlCustomEndpoint"));
        public static string? SqlCustomEndpoint
        {
            get => _sqlCustomEndpoint.Get();
            set => _sqlCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _storageCustomEndpoint = new __Value<string?>(() => __config.Get("storageCustomEndpoint"));
        public static string? StorageCustomEndpoint
        {
            get => _storageCustomEndpoint.Get();
            set => _storageCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _storageTransferCustomEndpoint = new __Value<string?>(() => __config.Get("storageTransferCustomEndpoint"));
        public static string? StorageTransferCustomEndpoint
        {
            get => _storageTransferCustomEndpoint.Get();
            set => _storageTransferCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _tagsCustomEndpoint = new __Value<string?>(() => __config.Get("tagsCustomEndpoint"));
        public static string? TagsCustomEndpoint
        {
            get => _tagsCustomEndpoint.Get();
            set => _tagsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _tpuCustomEndpoint = new __Value<string?>(() => __config.Get("tpuCustomEndpoint"));
        public static string? TpuCustomEndpoint
        {
            get => _tpuCustomEndpoint.Get();
            set => _tpuCustomEndpoint.Set(value);
        }

        private static readonly __Value<bool?> _userProjectOverride = new __Value<bool?>(() => __config.GetBoolean("userProjectOverride"));
        public static bool? UserProjectOverride
        {
            get => _userProjectOverride.Get();
            set => _userProjectOverride.Set(value);
        }

        private static readonly __Value<string?> _vertexAiCustomEndpoint = new __Value<string?>(() => __config.Get("vertexAiCustomEndpoint"));
        public static string? VertexAiCustomEndpoint
        {
            get => _vertexAiCustomEndpoint.Get();
            set => _vertexAiCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _vpcAccessCustomEndpoint = new __Value<string?>(() => __config.Get("vpcAccessCustomEndpoint"));
        public static string? VpcAccessCustomEndpoint
        {
            get => _vpcAccessCustomEndpoint.Get();
            set => _vpcAccessCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _workflowsCustomEndpoint = new __Value<string?>(() => __config.Get("workflowsCustomEndpoint"));
        public static string? WorkflowsCustomEndpoint
        {
            get => _workflowsCustomEndpoint.Get();
            set => _workflowsCustomEndpoint.Set(value);
        }

        private static readonly __Value<string?> _zone = new __Value<string?>(() => __config.Get("zone") ?? Utilities.GetEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE"));
        public static string? Zone
        {
            get => _zone.Get();
            set => _zone.Set(value);
        }

        public static class Types
        {

             public class Batching
             {
                public bool? EnableBatching { get; set; }
                public string? SendAfter { get; set; } = null!;
            }
        }
    }
}
