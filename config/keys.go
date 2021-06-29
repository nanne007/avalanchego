// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

const (
	FetchOnlyKey                              = "fetch-only"
	ConfigFileKey                             = "config-file"
	VersionKey                                = "version"
	GenesisConfigFileKey                      = "genesis"
	NetworkNameKey                            = "network-id"
	TxFeeKey                                  = "tx-fee"
	CreationTxFeeKey                          = "creation-tx-fee"
	UptimeRequirementKey                      = "uptime-requirement"
	MinValidatorStakeKey                      = "min-validator-stake"
	MaxValidatorStakeKey                      = "max-validator-stake"
	MinDelegatorStakeKey                      = "min-delegator-stake"
	MinDelegatorFeeKey                        = "min-delegation-fee"
	MinStakeDurationKey                       = "min-stake-duration"
	MaxStakeDurationKey                       = "max-stake-duration"
	StakeMintingPeriodKey                     = "stake-minting-period"
	AssertionsEnabledKey                      = "assertions-enabled"
	SignatureVerificationEnabledKey           = "signature-verification-enabled"
	DBTypeKey                                 = "db-type"
	DBPathKey                                 = "db-dir"
	PublicIPKey                               = "public-ip"
	DynamicUpdateDurationKey                  = "dynamic-update-duration"
	DynamicPublicIPResolverKey                = "dynamic-public-ip"
	ConnMeterResetDurationKey                 = "conn-meter-reset-duration"
	ConnMeterMaxConnsKey                      = "conn-meter-max-conns"
	OutboundConnectionThrottlingRps           = "outbound-connection-throttling-rps"
	OutboundConnectionTimeout                 = "outbound-connection-timeout"
	HTTPHostKey                               = "http-host"
	HTTPPortKey                               = "http-port"
	HTTPSEnabledKey                           = "http-tls-enabled"
	HTTPSKeyFileKey                           = "http-tls-key-file"
	HTTPSCertFileKey                          = "http-tls-cert-file"
	HTTPAllowedOrigins                        = "http-allowed-origins"
	APIAuthRequiredKey                        = "api-auth-required"
	APIAuthPasswordFileKey                    = "api-auth-password-file" // #nosec G101
	BootstrapIPsKey                           = "bootstrap-ips"
	BootstrapIDsKey                           = "bootstrap-ids"
	StakingPortKey                            = "staking-port"
	StakingEnabledKey                         = "staking-enabled"
	StakingEphemeralCertEnabledKey            = "staking-ephemeral-cert-enabled"
	StakingKeyPathKey                         = "staking-tls-key-file"
	StakingCertPathKey                        = "staking-tls-cert-file"
	StakingDisabledWeightKey                  = "staking-disabled-weight"
	MaxNonStakerPendingMsgsKey                = "max-non-staker-pending-msgs"
	StakerCPUReservedKey                      = "staker-cpu-reserved"
	NetworkInitialTimeoutKey                  = "network-initial-timeout"
	NetworkMinimumTimeoutKey                  = "network-minimum-timeout"
	NetworkMaximumTimeoutKey                  = "network-maximum-timeout"
	NetworkTimeoutHalflifeKey                 = "network-timeout-halflife"
	NetworkTimeoutCoefficientKey              = "network-timeout-coefficient"
	NetworkHealthMinPeersKey                  = "network-health-min-conn-peers"
	NetworkHealthMaxTimeSinceMsgReceivedKey   = "network-health-max-time-since-msg-received"
	NetworkHealthMaxTimeSinceMsgSentKey       = "network-health-max-time-since-msg-sent"
	NetworkHealthMaxPortionSendQueueFillKey   = "network-health-max-portion-send-queue-full"
	NetworkHealthMaxSendFailRateKey           = "network-health-max-send-fail-rate"
	NetworkHealthMaxOutstandingDurationKey    = "network-health-max-outstanding-request-duration"
	NetworkPeerListSizeKey                    = "network-peer-list-size"
	NetworkPeerListGossipSizeKey              = "network-peer-list-gossip-size"
	NetworkPeerListGossipFreqKey              = "network-peer-list-gossip-frequency"
	SendQueueSizeKey                          = "send-queue-size"
	BenchlistFailThresholdKey                 = "benchlist-fail-threshold"
	BenchlistPeerSummaryEnabledKey            = "benchlist-peer-summary-enabled"
	BenchlistDurationKey                      = "benchlist-duration"
	BenchlistMinFailingDurationKey            = "benchlist-min-failing-duration"
	BuildDirKey                               = "build-dir"
	LogsDirKey                                = "log-dir"
	LogLevelKey                               = "log-level"
	LogDisplayLevelKey                        = "log-display-level"
	LogDisplayHighlightKey                    = "log-display-highlight"
	SnowSampleSizeKey                         = "snow-sample-size"
	SnowQuorumSizeKey                         = "snow-quorum-size"
	SnowVirtuousCommitThresholdKey            = "snow-virtuous-commit-threshold"
	SnowRogueCommitThresholdKey               = "snow-rogue-commit-threshold"
	SnowAvalancheNumParentsKey                = "snow-avalanche-num-parents"
	SnowAvalancheBatchSizeKey                 = "snow-avalanche-batch-size"
	SnowConcurrentRepollsKey                  = "snow-concurrent-repolls"
	SnowOptimalProcessingKey                  = "snow-optimal-processing"
	SnowMaxProcessingKey                      = "snow-max-processing"
	SnowMaxTimeProcessingKey                  = "snow-max-time-processing"
	SnowEpochFirstTransition                  = "snow-epoch-first-transition"
	SnowEpochDuration                         = "snow-epoch-duration"
	WhitelistedSubnetsKey                     = "whitelisted-subnets"
	AdminAPIEnabledKey                        = "api-admin-enabled"
	InfoAPIEnabledKey                         = "api-info-enabled"
	KeystoreAPIEnabledKey                     = "api-keystore-enabled"
	MetricsAPIEnabledKey                      = "api-metrics-enabled"
	HealthAPIEnabledKey                       = "api-health-enabled"
	IpcAPIEnabledKey                          = "api-ipcs-enabled"
	IpcsChainIDsKey                           = "ipcs-chain-ids"
	IpcsPathKey                               = "ipcs-path"
	MeterVMsEnabledKey                        = "meter-vms-enabled"
	ConsensusGossipFrequencyKey               = "consensus-gossip-frequency"
	ConsensusGossipAcceptedFrontierSizeKey    = "consensus-accepted-frontier-gossip-size"
	ConsensusGossipOnAcceptSizeKey            = "consensus-on-accept-gossip-size"
	ConsensusShutdownTimeoutKey               = "consensus-shutdown-timeout"
	FdLimitKey                                = "fd-limit"
	CorethConfigKey                           = "coreth-config"
	IndexContainersEnabled                    = "index-containers-enabled"
	IndexAllowIncompleteKey                   = "index-allow-incomplete"
	RouterHealthMaxDropRateKey                = "router-health-max-drop-rate"
	RouterHealthMaxOutstandingRequestsKey     = "router-health-max-outstanding-requests"
	HealthCheckFreqKey                        = "health-check-frequency"
	HealthCheckAveragerHalflifeKey            = "health-check-averager-halflife"
	RetryBootstrapKey                         = "bootstrap-retry-enabled"
	RetryBootstrapMaxAttemptsKey              = "bootstrap-retry-max-attempts"
	PeerAliasTimeoutKey                       = "peer-alias-timeout"
	PluginModeKey                             = "plugin-mode-enabled"
	BootstrapBeaconConnectionTimeoutKey       = "bootstrap-beacon-connection-timeout"
	BootstrapMaxTimeGetAncestorsKey           = "boostrap-max-time-get-ancestors"
	BootstrapMultiputMaxContainersSentKey     = "bootstrap-multiput-max-containers-sent"
	BootstrapMultiputMaxContainersReceivedKey = "bootstrap-multiput-max-containers-received"
	ChainConfigDirKey                         = "chain-config-dir"
	ProfileDirKey                             = "profile-dir"
	ProfileContinuousEnabledKey               = "profile-continuous-enabled"
	ProfileContinuousFreqKey                  = "profile-continuous-freq"
	ProfileContinuousMaxFilesKey              = "profile-continuous-max-files"
	ThrottlingAtLargeAllocSizeKey             = "throttling-at-large-alloc-size"
	ThrottlingVdrAllocSizeKey                 = "throttling-validator-alloc-size"
	ThrottlingNodeMaxAtLargeBytesKey          = "throttling-node-max-at-large-bytes"
	VMAliasesFileKey                          = "vm-aliases-file"
)
