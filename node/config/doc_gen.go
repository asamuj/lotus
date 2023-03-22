// Code generated by github.com/filecoin-project/lotus/node/config/cfgdocgen. DO NOT EDIT.

package config

type DocField struct {
	Name    string
	Type    string
	Comment string
}

var Doc = map[string][]DocField{
	"API": []DocField{
		{
			Name: "ListenAddress",
			Type: "string",

			Comment: `Binding address for the Lotus API`,
		},
		{
			Name: "RemoteListenAddress",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "Timeout",
			Type: "Duration",

			Comment: ``,
		},
	},
	"Backup": []DocField{
		{
			Name: "DisableMetadataLog",
			Type: "bool",

			Comment: `When set to true disables metadata log (.lotus/kvlog). This can save disk
space by reducing metadata redundancy.

Note that in case of metadata corruption it might be much harder to recover
your node if metadata log is disabled`,
		},
	},
	"BatchFeeConfig": []DocField{
		{
			Name: "Base",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "PerSector",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"Chainstore": []DocField{
		{
			Name: "EnableSplitstore",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "Splitstore",
			Type: "Splitstore",

			Comment: ``,
		},
	},
	"Client": []DocField{
		{
			Name: "UseIpfs",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "IpfsOnlineMode",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "IpfsMAddr",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "IpfsUseForRetrieval",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers between the client
and storage providers for storage deals`,
		},
		{
			Name: "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers between the client
and storage providers for retrieval deals`,
		},
		{
			Name: "OffChainRetrieval",
			Type: "bool",

			Comment: `Require that retrievals perform no on-chain operations. Paid retrievals
without existing payment channels with available funds will fail instead
of automatically performing on-chain operations.`,
		},
	},
	"Common": []DocField{
		{
			Name: "API",
			Type: "API",

			Comment: ``,
		},
		{
			Name: "Backup",
			Type: "Backup",

			Comment: ``,
		},
		{
			Name: "Logging",
			Type: "Logging",

			Comment: ``,
		},
		{
			Name: "Libp2p",
			Type: "Libp2p",

			Comment: ``,
		},
		{
			Name: "Pubsub",
			Type: "Pubsub",

			Comment: ``,
		},
	},
	"DAGStoreConfig": []DocField{
		{
			Name: "RootDir",
			Type: "string",

			Comment: `Path to the dagstore root directory. This directory contains three
subdirectories, which can be symlinked to alternative locations if
need be:
- ./transients: caches unsealed deals that have been fetched from the
storage subsystem for serving retrievals.
- ./indices: stores shard indices.
- ./datastore: holds the KV store tracking the state of every shard
known to the DAG store.
Default value: <LOTUS_MARKETS_PATH>/dagstore (split deployment) or
<LOTUS_MINER_PATH>/dagstore (monolith deployment)`,
		},
		{
			Name: "MaxConcurrentIndex",
			Type: "int",

			Comment: `The maximum amount of indexing jobs that can run simultaneously.
0 means unlimited.
Default value: 5.`,
		},
		{
			Name: "MaxConcurrentReadyFetches",
			Type: "int",

			Comment: `The maximum amount of unsealed deals that can be fetched simultaneously
from the storage subsystem. 0 means unlimited.
Default value: 0 (unlimited).`,
		},
		{
			Name: "MaxConcurrentUnseals",
			Type: "int",

			Comment: `The maximum amount of unseals that can be processed simultaneously
from the storage subsystem. 0 means unlimited.
Default value: 0 (unlimited).`,
		},
		{
			Name: "MaxConcurrencyStorageCalls",
			Type: "int",

			Comment: `The maximum number of simultaneous inflight API calls to the storage
subsystem.
Default value: 100.`,
		},
		{
			Name: "GCInterval",
			Type: "Duration",

			Comment: `The time between calls to periodic dagstore GC, in time.Duration string
representation, e.g. 1m, 5m, 1h.
Default value: 1 minute.`,
		},
	},
	"DealmakingConfig": []DocField{
		{
			Name: "ConsiderOnlineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept online deals`,
		},
		{
			Name: "ConsiderOfflineStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline deals`,
		},
		{
			Name: "ConsiderOnlineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept retrieval deals`,
		},
		{
			Name: "ConsiderOfflineRetrievalDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept offline retrieval deals`,
		},
		{
			Name: "ConsiderVerifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept verified deals`,
		},
		{
			Name: "ConsiderUnverifiedStorageDeals",
			Type: "bool",

			Comment: `When enabled, the miner can accept unverified deals`,
		},
		{
			Name: "PieceCidBlocklist",
			Type: "[]cid.Cid",

			Comment: `A list of Data CIDs to reject when making deals`,
		},
		{
			Name: "ExpectedSealDuration",
			Type: "Duration",

			Comment: `Maximum expected amount of time getting the deal into a sealed sector will take
This includes the time the deal will need to get transferred and published
before being assigned to a sector`,
		},
		{
			Name: "MaxDealStartDelay",
			Type: "Duration",

			Comment: `Maximum amount of time proposed deal StartEpoch can be in future`,
		},
		{
			Name: "PublishMsgPeriod",
			Type: "Duration",

			Comment: `When a deal is ready to publish, the amount of time to wait for more
deals to be ready to publish before publishing them all as a batch`,
		},
		{
			Name: "MaxDealsPerPublishMsg",
			Type: "uint64",

			Comment: `The maximum number of deals to include in a single PublishStorageDeals
message`,
		},
		{
			Name: "MaxProviderCollateralMultiplier",
			Type: "uint64",

			Comment: `The maximum collateral that the provider will put up against a deal,
as a multiplier of the minimum collateral bound`,
		},
		{
			Name: "MaxStagingDealsBytes",
			Type: "int64",

			Comment: `The maximum allowed disk usage size in bytes of staging deals not yet
passed to the sealing node by the markets service. 0 is unlimited.`,
		},
		{
			Name: "SimultaneousTransfersForStorage",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for storage deals`,
		},
		{
			Name: "SimultaneousTransfersForStoragePerClient",
			Type: "uint64",

			Comment: `The maximum number of simultaneous data transfers from any single client
for storage deals.
Unset by default (0), and values higher than SimultaneousTransfersForStorage
will have no effect; i.e. the total number of simultaneous data transfers
across all storage clients is bound by SimultaneousTransfersForStorage
regardless of this number.`,
		},
		{
			Name: "SimultaneousTransfersForRetrieval",
			Type: "uint64",

			Comment: `The maximum number of parallel online data transfers for retrieval deals`,
		},
		{
			Name: "StartEpochSealingBuffer",
			Type: "uint64",

			Comment: `Minimum start epoch buffer to give time for sealing of sector with deal.`,
		},
		{
			Name: "Filter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of storage deals
see https://lotus.filecoin.io/storage-providers/advanced-configurations/market/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalFilter",
			Type: "string",

			Comment: `A command used for fine-grained evaluation of retrieval deals
see https://lotus.filecoin.io/storage-providers/advanced-configurations/market/#using-filters-for-fine-grained-storage-and-retrieval-deal-acceptance for more details`,
		},
		{
			Name: "RetrievalPricing",
			Type: "*RetrievalPricing",

			Comment: ``,
		},
	},
	"Events": []DocField{
		{
			Name: "DisableRealTimeFilterAPI",
			Type: "bool",

			Comment: `EnableEthRPC enables APIs that
DisableRealTimeFilterAPI will disable the RealTimeFilterAPI that can create and query filters for actor events as they are emitted.
The API is enabled when EnableEthRPC is true, but can be disabled selectively with this flag.`,
		},
		{
			Name: "DisableHistoricFilterAPI",
			Type: "bool",

			Comment: `DisableHistoricFilterAPI will disable the HistoricFilterAPI that can create and query filters for actor events
that occurred in the past. HistoricFilterAPI maintains a queryable index of events.
The API is enabled when EnableEthRPC is true, but can be disabled selectively with this flag.`,
		},
		{
			Name: "FilterTTL",
			Type: "Duration",

			Comment: `FilterTTL specifies the time to live for actor event filters. Filters that haven't been accessed longer than
this time become eligible for automatic deletion.`,
		},
		{
			Name: "MaxFilters",
			Type: "int",

			Comment: `MaxFilters specifies the maximum number of filters that may exist at any one time.`,
		},
		{
			Name: "MaxFilterResults",
			Type: "int",

			Comment: `MaxFilterResults specifies the maximum number of results that can be accumulated by an actor event filter.`,
		},
		{
			Name: "MaxFilterHeightRange",
			Type: "uint64",

			Comment: `MaxFilterHeightRange specifies the maximum range of heights that can be used in a filter (to avoid querying
the entire chain)`,
		},
		{
			Name: "DatabasePath",
			Type: "string",

			Comment: `DatabasePath is the full path to a sqlite database that will be used to index actor events to
support the historic filter APIs. If the database does not exist it will be created. The directory containing
the database must already exist and be writeable. If a relative path is provided here, sqlite treats it as
relative to the CWD (current working directory).`,
		},
	},
	"FeeConfig": []DocField{
		{
			Name: "DefaultMaxFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"FevmConfig": []DocField{
		{
			Name: "EnableEthRPC",
			Type: "bool",

			Comment: `EnableEthRPC enables eth_ rpc, and enables storing a mapping of eth transaction hashes to filecoin message Cids.
This will also enable the RealTimeFilterAPI and HistoricFilterAPI by default, but they can be disabled by config options above.`,
		},
		{
			Name: "EthTxHashMappingLifetimeDays",
			Type: "int",

			Comment: `EthTxHashMappingLifetimeDays the transaction hash lookup database will delete mappings that have been stored for more than x days
Set to 0 to keep all mappings`,
		},
		{
			Name: "Events",
			Type: "Events",

			Comment: ``,
		},
	},
	"FullNode": []DocField{
		{
			Name: "Client",
			Type: "Client",

			Comment: ``,
		},
		{
			Name: "Wallet",
			Type: "Wallet",

			Comment: ``,
		},
		{
			Name: "Fees",
			Type: "FeeConfig",

			Comment: ``,
		},
		{
			Name: "Chainstore",
			Type: "Chainstore",

			Comment: ``,
		},
		{
			Name: "Cluster",
			Type: "UserRaftConfig",

			Comment: ``,
		},
		{
			Name: "Fevm",
			Type: "FevmConfig",

			Comment: ``,
		},
		{
			Name: "Index",
			Type: "IndexConfig",

			Comment: ``,
		},
	},
	"IndexConfig": []DocField{
		{
			Name: "EnableMsgIndex",
			Type: "bool",

			Comment: `EnableMsgIndex enables indexing of messages on chain.`,
		},
	},
	"IndexProviderConfig": []DocField{
		{
			Name: "Enable",
			Type: "bool",

			Comment: `Enable set whether to enable indexing announcement to the network and expose endpoints that
allow indexer nodes to process announcements. Enabled by default.`,
		},
		{
			Name: "EntriesCacheCapacity",
			Type: "int",

			Comment: `EntriesCacheCapacity sets the maximum capacity to use for caching the indexing advertisement
entries. Defaults to 1024 if not specified. The cache is evicted using LRU policy. The
maximum storage used by the cache is a factor of EntriesCacheCapacity, EntriesChunkSize and
the length of multihashes being advertised. For example, advertising 128-bit long multihashes
with the default EntriesCacheCapacity, and EntriesChunkSize means the cache size can grow to
256MiB when full.`,
		},
		{
			Name: "EntriesChunkSize",
			Type: "int",

			Comment: `EntriesChunkSize sets the maximum number of multihashes to include in a single entries chunk.
Defaults to 16384 if not specified. Note that chunks are chained together for indexing
advertisements that include more multihashes than the configured EntriesChunkSize.`,
		},
		{
			Name: "TopicName",
			Type: "string",

			Comment: `TopicName sets the topic name on which the changes to the advertised content are announced.
If not explicitly specified, the topic name is automatically inferred from the network name
in following format: '/indexer/ingest/<network-name>'
Defaults to empty, which implies the topic name is inferred from network name.`,
		},
		{
			Name: "PurgeCacheOnStart",
			Type: "bool",

			Comment: `PurgeCacheOnStart sets whether to clear any cached entries chunks when the provider engine
starts. By default, the cache is rehydrated from previously cached entries stored in
datastore if any is present.`,
		},
	},
	"Libp2p": []DocField{
		{
			Name: "ListenAddresses",
			Type: "[]string",

			Comment: `Binding address for the libp2p host - 0 means random port.
Format: multiaddress; see https://multiformats.io/multiaddr/`,
		},
		{
			Name: "AnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to explicitally announce to other peers. If not specified,
all interface addresses are announced
Format: multiaddress`,
		},
		{
			Name: "NoAnnounceAddresses",
			Type: "[]string",

			Comment: `Addresses to not announce
Format: multiaddress`,
		},
		{
			Name: "BootstrapPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "ProtectedPeers",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DisableNatPortMap",
			Type: "bool",

			Comment: `When not disabled (default), lotus asks NAT devices (e.g., routers), to
open up an external port and forward it to the port lotus is running on.
When this works (i.e., when your router supports NAT port forwarding),
it makes the local lotus node accessible from the public internet`,
		},
		{
			Name: "ConnMgrLow",
			Type: "uint",

			Comment: `ConnMgrLow is the number of connections that the basic connection manager
will trim down to.`,
		},
		{
			Name: "ConnMgrHigh",
			Type: "uint",

			Comment: `ConnMgrHigh is the number of connections that, when exceeded, will trigger
a connection GC operation. Note: protected/recently formed connections don't
count towards this limit.`,
		},
		{
			Name: "ConnMgrGrace",
			Type: "Duration",

			Comment: `ConnMgrGrace is a time duration that new connections are immune from being
closed by the connection manager.`,
		},
	},
	"Logging": []DocField{
		{
			Name: "SubsystemLevels",
			Type: "map[string]string",

			Comment: `SubsystemLevels specify per-subsystem log levels`,
		},
	},
	"MinerAddressConfig": []DocField{
		{
			Name: "PreCommitControl",
			Type: "[]string",

			Comment: `Addresses to send PreCommit messages from`,
		},
		{
			Name: "CommitControl",
			Type: "[]string",

			Comment: `Addresses to send Commit messages from`,
		},
		{
			Name: "TerminateControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DealPublishControl",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "DisableOwnerFallback",
			Type: "bool",

			Comment: `DisableOwnerFallback disables usage of the owner address for messages
sent automatically`,
		},
		{
			Name: "DisableWorkerFallback",
			Type: "bool",

			Comment: `DisableWorkerFallback disables usage of the worker address for messages
sent automatically, if control addresses are configured.
A control address that doesn't have enough funds will still be chosen
over the worker address if this flag is set.`,
		},
	},
	"MinerFeeConfig": []DocField{
		{
			Name: "MaxPreCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxCommitGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxPreCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: `maxBatchFee = maxBase + maxPerSector * nSectors`,
		},
		{
			Name: "MaxCommitBatchGasFee",
			Type: "BatchFeeConfig",

			Comment: ``,
		},
		{
			Name: "MaxTerminateGasFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxWindowPoStGasFee",
			Type: "types.FIL",

			Comment: `WindowPoSt is a high-value operation, so the default fee should be high.`,
		},
		{
			Name: "MaxPublishDealsFee",
			Type: "types.FIL",

			Comment: ``,
		},
		{
			Name: "MaxMarketBalanceAddFee",
			Type: "types.FIL",

			Comment: ``,
		},
	},
	"MinerSubsystemConfig": []DocField{
		{
			Name: "EnableMining",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableSealing",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableSectorStorage",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "EnableMarkets",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "SealerApiInfo",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "SectorIndexApiInfo",
			Type: "string",

			Comment: ``,
		},
	},
	"ProvingConfig": []DocField{
		{
			Name: "ParallelCheckLimit",
			Type: "int",

			Comment: `Maximum number of sector checks to run in parallel. (0 = unlimited)

WARNING: Setting this value too high may make the node crash by running out of stack
WARNING: Setting this value too low may make sector challenge reading much slower, resulting in failed PoSt due
to late submission.

After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Name: "SingleCheckTimeout",
			Type: "Duration",

			Comment: `Maximum amount of time a proving pre-check can take for a sector. If the check times out the sector will be skipped

WARNING: Setting this value too low risks in sectors being skipped even though they are accessible, just reading the
test challenge took longer than this timeout
WARNING: Setting this value too high risks missing PoSt deadline in case IO operations related to this sector are
blocked (e.g. in case of disconnected NFS mount)`,
		},
		{
			Name: "PartitionCheckTimeout",
			Type: "Duration",

			Comment: `Maximum amount of time a proving pre-check can take for an entire partition. If the check times out, sectors in
the partition which didn't get checked on time will be skipped

WARNING: Setting this value too low risks in sectors being skipped even though they are accessible, just reading the
test challenge took longer than this timeout
WARNING: Setting this value too high risks missing PoSt deadline in case IO operations related to this partition are
blocked or slow`,
		},
		{
			Name: "DisableBuiltinWindowPoSt",
			Type: "bool",

			Comment: `Disable Window PoSt computation on the lotus-miner process even if no window PoSt workers are present.

WARNING: If no windowPoSt workers are connected, window PoSt WILL FAIL resulting in faulty sectors which will need
to be recovered. Before enabling this option, make sure your PoSt workers work correctly.

After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Name: "DisableBuiltinWinningPoSt",
			Type: "bool",

			Comment: `Disable Winning PoSt computation on the lotus-miner process even if no winning PoSt workers are present.

WARNING: If no WinningPoSt workers are connected, Winning PoSt WILL FAIL resulting in lost block rewards.
Before enabling this option, make sure your PoSt workers work correctly.`,
		},
		{
			Name: "DisableWDPoStPreChecks",
			Type: "bool",

			Comment: `Disable WindowPoSt provable sector readability checks.

In normal operation, when preparing to compute WindowPoSt, lotus-miner will perform a round of reading challenges
from all sectors to confirm that those sectors can be proven. Challenges read in this process are discarded, as
we're only interested in checking that sector data can be read.

When using builtin proof computation (no PoSt workers, and DisableBuiltinWindowPoSt is set to false), this process
can save a lot of time and compute resources in the case that some sectors are not readable - this is caused by
the builtin logic not skipping snark computation when some sectors need to be skipped.

When using PoSt workers, this process is mostly redundant, with PoSt workers challenges will be read once, and
if challenges for some sectors aren't readable, those sectors will just get skipped.

Disabling sector pre-checks will slightly reduce IO load when proving sectors, possibly resulting in shorter
time to produce window PoSt. In setups with good IO capabilities the effect of this option on proving time should
be negligible.

NOTE: It likely is a bad idea to disable sector pre-checks in setups with no PoSt workers.

NOTE: Even when this option is enabled, recovering sectors will be checked before recovery declaration message is
sent to the chain

After changing this option, confirm that the new value works in your setup by invoking
'lotus-miner proving compute window-post 0'`,
		},
		{
			Name: "MaxPartitionsPerPoStMessage",
			Type: "int",

			Comment: `Maximum number of partitions to prove in a single SubmitWindowPoSt messace. 0 = network limit (10 in nv16)

A single partition may contain up to 2349 32GiB sectors, or 2300 64GiB sectors.

The maximum number of sectors which can be proven in a single PoSt message is 25000 in network version 16, which
means that a single message can prove at most 10 partitions

Note that setting this value lower may result in less efficient gas use - more messages will be sent,
to prove each deadline, resulting in more total gas use (but each message will have lower gas limit)

Setting this value above the network limit has no effect`,
		},
		{
			Name: "MaxPartitionsPerRecoveryMessage",
			Type: "int",

			Comment: `In some cases when submitting DeclareFaultsRecovered messages,
there may be too many recoveries to fit in a BlockGasLimit.
In those cases it may be necessary to set this value to something low (eg 1);
Note that setting this value lower may result in less efficient gas use - more messages will be sent than needed,
resulting in more total gas use (but each message will have lower gas limit)`,
		},
		{
			Name: "SingleRecoveringPartitionPerPostMessage",
			Type: "bool",

			Comment: `Enable single partition per PoSt Message for partitions containing recovery sectors

In cases when submitting PoSt messages which contain recovering sectors, the default network limit may still be
too high to fit in the block gas limit. In those cases, it becomes useful to only house the single partition
with recovering sectors in the post message

Note that setting this value lower may result in less efficient gas use - more messages will be sent,
to prove each deadline, resulting in more total gas use (but each message will have lower gas limit)`,
		},
	},
	"Pubsub": []DocField{
		{
			Name: "Bootstrapper",
			Type: "bool",

			Comment: `Run the node in bootstrap-node mode`,
		},
		{
			Name: "DirectPeers",
			Type: "[]string",

			Comment: `DirectPeers specifies peers with direct peering agreements. These peers are
connected outside of the mesh, with all (valid) message unconditionally
forwarded to them. The router will maintain open connections to these peers.
Note that the peering agreement should be reciprocal with direct peers
symmetrically configured at both ends.
Type: Array of multiaddress peerinfo strings, must include peerid (/p2p/12D3K...`,
		},
		{
			Name: "IPColocationWhitelist",
			Type: "[]string",

			Comment: ``,
		},
		{
			Name: "RemoteTracer",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "JsonTracer",
			Type: "string",

			Comment: `Path to file that will be used to output tracer content in JSON format.
If present tracer will save data to defined file.
Format: file path`,
		},
		{
			Name: "ElasticSearchTracer",
			Type: "string",

			Comment: `Connection string for elasticsearch instance.
If present tracer will save data to elasticsearch.
Format: https://<username>:<password>@<elasticsearch_url>:<port>/`,
		},
		{
			Name: "ElasticSearchIndex",
			Type: "string",

			Comment: `Name of elasticsearch index that will be used to save tracer data.
This property is used only if ElasticSearchTracer propery is set.`,
		},
		{
			Name: "TracerSourceAuth",
			Type: "string",

			Comment: `Auth token that will be passed with logs to elasticsearch - used for weighted peers score.`,
		},
	},
	"RetrievalPricing": []DocField{
		{
			Name: "Strategy",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "Default",
			Type: "*RetrievalPricingDefault",

			Comment: ``,
		},
		{
			Name: "External",
			Type: "*RetrievalPricingExternal",

			Comment: ``,
		},
	},
	"RetrievalPricingDefault": []DocField{
		{
			Name: "VerifiedDealsFreeTransfer",
			Type: "bool",

			Comment: `VerifiedDealsFreeTransfer configures zero fees for data transfer for a retrieval deal
of a payloadCid that belongs to a verified storage deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "default".
default value is true`,
		},
	},
	"RetrievalPricingExternal": []DocField{
		{
			Name: "Path",
			Type: "string",

			Comment: `Path of the external script that will be run to price a retrieval deal.
This parameter is ONLY applicable if the retrieval pricing policy strategy has been configured to "external".`,
		},
	},
	"SealerConfig": []DocField{
		{
			Name: "ParallelFetchLimit",
			Type: "int",

			Comment: ``,
		},
		{
			Name: "AllowSectorDownload",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowAddPiece",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowPreCommit1",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowPreCommit2",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowCommit",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowUnseal",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowReplicaUpdate",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowProveReplicaUpdate2",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "AllowRegenSectorKey",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "LocalWorkerName",
			Type: "string",

			Comment: `LocalWorkerName specifies a custom name for the builtin worker.
If set to an empty string (default) os hostname will be used`,
		},
		{
			Name: "Assigner",
			Type: "string",

			Comment: `Assigner specifies the worker assigner to use when scheduling tasks.
"utilization" (default) - assign tasks to workers with lowest utilization.
"spread" - assign tasks to as many distinct workers as possible.`,
		},
		{
			Name: "DisallowRemoteFinalize",
			Type: "bool",

			Comment: `DisallowRemoteFinalize when set to true will force all Finalize tasks to
run on workers with local access to both long-term storage and the sealing
path containing the sector.
--
WARNING: Only set this if all workers have access to long-term storage
paths. If this flag is enabled, and there are workers without long-term
storage access, sectors will not be moved from them, and Finalize tasks
will appear to be stuck.
--
If you see stuck Finalize tasks after enabling this setting, check
'lotus-miner sealing sched-diag' and 'lotus-miner storage find [sector num]'`,
		},
		{
			Name: "ResourceFiltering",
			Type: "ResourceFilteringStrategy",

			Comment: `ResourceFiltering instructs the system which resource filtering strategy
to use when evaluating tasks against this worker. An empty value defaults
to "hardware".`,
		},
	},
	"SealingConfig": []DocField{
		{
			Name: "MaxWaitDealsSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be waiting for more deals to be packed in it before it begins sealing at any given time.
If the miner is accepting multiple deals in parallel, up to MaxWaitDealsSectors of new sectors will be created.
If more than MaxWaitDealsSectors deals are accepted in parallel, only MaxWaitDealsSectors deals will be processed in parallel
Note that setting this number too high in relation to deal ingestion rate may result in poor sector packing efficiency
0 = no limit`,
		},
		{
			Name: "MaxSealingSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when creating new CC sectors (0 = unlimited)`,
		},
		{
			Name: "MaxSealingSectorsForDeals",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when creating new sectors with deals (0 = unlimited)`,
		},
		{
			Name: "PreferNewSectorsForDeals",
			Type: "bool",

			Comment: `Prefer creating new sectors even if there are sectors Available for upgrading.
This setting combined with MaxUpgradingSectors set to a value higher than MaxSealingSectorsForDeals makes it
possible to use fast sector upgrades to handle high volumes of storage deals, while still using the simple sealing
flow when the volume of storage deals is lower.`,
		},
		{
			Name: "MaxUpgradingSectors",
			Type: "uint64",

			Comment: `Upper bound on how many sectors can be sealing+upgrading at the same time when upgrading CC sectors with deals (0 = MaxSealingSectorsForDeals)`,
		},
		{
			Name: "MinUpgradeSectorExpiration",
			Type: "uint64",

			Comment: `When set to a non-zero value, minimum number of epochs until sector expiration required for sectors to be considered
for upgrades (0 = DealMinDuration = 180 days = 518400 epochs)

Note that if all deals waiting in the input queue have lifetimes longer than this value, upgrade sectors will be
required to have expiration of at least the soonest-ending deal`,
		},
		{
			Name: "MinTargetUpgradeSectorExpiration",
			Type: "uint64",

			Comment: `When set to a non-zero value, minimum number of epochs until sector expiration above which upgrade candidates will
be selected based on lowest initial pledge.

Target sector expiration is calculated by looking at the input deal queue, sorting it by deal expiration, and
selecting N deals from the queue up to sector size. The target expiration will be Nth deal end epoch, or in case
where there weren't enough deals to fill a sector, DealMaxDuration (540 days = 1555200 epochs)

Setting this to a high value (for example to maximum deal duration - 1555200) will disable selection based on
initial pledge - upgrade sectors will always be chosen based on longest expiration`,
		},
		{
			Name: "CommittedCapacitySectorLifetime",
			Type: "Duration",

			Comment: `CommittedCapacitySectorLifetime is the duration a Committed Capacity (CC) sector will
live before it must be extended or converted into sector containing deals before it is
terminated. Value must be between 180-540 days inclusive`,
		},
		{
			Name: "WaitDealsDelay",
			Type: "Duration",

			Comment: `Period of time that a newly created sector will wait for more deals to be packed in to before it starts to seal.
Sectors which are fully filled will start sealing immediately`,
		},
		{
			Name: "AlwaysKeepUnsealedCopy",
			Type: "bool",

			Comment: `Whether to keep unsealed copies of deal data regardless of whether the client requested that. This lets the miner
avoid the relatively high cost of unsealing the data later, at the cost of more storage space`,
		},
		{
			Name: "FinalizeEarly",
			Type: "bool",

			Comment: `Run sector finalization before submitting sector proof to the chain`,
		},
		{
			Name: "MakeNewSectorForDeals",
			Type: "bool",

			Comment: `Whether new sectors are created to pack incoming deals
When this is set to false no new sectors will be created for sealing incoming deals
This is useful for forcing all deals to be assigned as snap deals to sectors marked for upgrade`,
		},
		{
			Name: "MakeCCSectorsAvailable",
			Type: "bool",

			Comment: `After sealing CC sectors, make them available for upgrading with deals`,
		},
		{
			Name: "CollateralFromMinerBalance",
			Type: "bool",

			Comment: `Whether to use available miner balance for sector collateral instead of sending it with each message`,
		},
		{
			Name: "AvailableBalanceBuffer",
			Type: "types.FIL",

			Comment: `Minimum available balance to keep in the miner actor before sending it with messages`,
		},
		{
			Name: "DisableCollateralFallback",
			Type: "bool",

			Comment: `Don't send collateral with messages even if there is no available balance in the miner actor`,
		},
		{
			Name: "BatchPreCommits",
			Type: "bool",

			Comment: `enable / disable precommit batching (takes effect after nv13)`,
		},
		{
			Name: "MaxPreCommitBatch",
			Type: "int",

			Comment: `maximum precommit batch size - batches will be sent immediately above this size`,
		},
		{
			Name: "PreCommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Name: "PreCommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deal in batch would start expiring`,
		},
		{
			Name: "AggregateCommits",
			Type: "bool",

			Comment: `enable / disable commit aggregation (takes effect after nv13)`,
		},
		{
			Name: "MinCommitBatch",
			Type: "int",

			Comment: `minimum batched commit size - batches above this size will eventually be sent on a timeout`,
		},
		{
			Name: "MaxCommitBatch",
			Type: "int",

			Comment: `maximum batched commit size - batches will be sent immediately above this size`,
		},
		{
			Name: "CommitBatchWait",
			Type: "Duration",

			Comment: `how long to wait before submitting a batch after crossing the minimum batch size`,
		},
		{
			Name: "CommitBatchSlack",
			Type: "Duration",

			Comment: `time buffer for forceful batch submission before sectors/deals in batch would start expiring`,
		},
		{
			Name: "BatchPreCommitAboveBaseFee",
			Type: "types.FIL",

			Comment: `network BaseFee below which to stop doing precommit batching, instead
sending precommit messages to the chain individually`,
		},
		{
			Name: "AggregateAboveBaseFee",
			Type: "types.FIL",

			Comment: `network BaseFee below which to stop doing commit aggregation, instead
submitting proofs to the chain individually`,
		},
		{
			Name: "TerminateBatchMax",
			Type: "uint64",

			Comment: ``,
		},
		{
			Name: "TerminateBatchMin",
			Type: "uint64",

			Comment: ``,
		},
		{
			Name: "TerminateBatchWait",
			Type: "Duration",

			Comment: ``,
		},
	},
	"Splitstore": []DocField{
		{
			Name: "ColdStoreType",
			Type: "string",

			Comment: `ColdStoreType specifies the type of the coldstore.
It can be "messages" (default) to store only messages, "universal" to store all chain state or "discard" for discarding cold blocks.`,
		},
		{
			Name: "HotStoreType",
			Type: "string",

			Comment: `HotStoreType specifies the type of the hotstore.
Only currently supported value is "badger".`,
		},
		{
			Name: "MarkSetType",
			Type: "string",

			Comment: `MarkSetType specifies the type of the markset.
It can be "map" for in memory marking or "badger" (default) for on-disk marking.`,
		},
		{
			Name: "HotStoreMessageRetention",
			Type: "uint64",

			Comment: `HotStoreMessageRetention specifies the retention policy for messages, in finalities beyond
the compaction boundary; default is 0.`,
		},
		{
			Name: "HotStoreFullGCFrequency",
			Type: "uint64",

			Comment: `HotStoreFullGCFrequency specifies how often to perform a full (moving) GC on the hotstore.
A value of 0 disables, while a value 1 will do full GC in every compaction.
Default is 20 (about once a week).`,
		},
		{
			Name: "HotStoreMaxSpaceTarget",
			Type: "uint64",

			Comment: `HotStoreMaxSpaceTarget sets a target max disk size for the hotstore. Splitstore GC
will run moving GC if disk utilization gets within a threshold (150 GB) of the target.
Splitstore GC will NOT run moving GC if the total size of the move would get
within 50 GB of the target, and instead will run a more aggressive online GC.
If both HotStoreFullGCFrequency and HotStoreMaxSpaceTarget are set then splitstore
GC will trigger moving GC if either configuration condition is met.
A reasonable minimum is 2x fully GCed hotstore size + 50 G buffer.
At this minimum size moving GC happens every time, any smaller and moving GC won't
be able to run. In spring 2023 this minimum is ~550 GB.`,
		},
		{
			Name: "HotStoreMaxSpaceThreshold",
			Type: "uint64",

			Comment: `When HotStoreMaxSpaceTarget is set Moving GC will be triggered when total moving size
exceeds HotstoreMaxSpaceTarget - HotstoreMaxSpaceThreshold`,
		},
		{
			Name: "HotstoreMaxSpaceSafetyBuffer",
			Type: "uint64",

			Comment: `Safety buffer to prevent moving GC from overflowing disk when HotStoreMaxSpaceTarget
is set.  Moving GC will not occur when total moving size exceeds
HotstoreMaxSpaceTarget - HotstoreMaxSpaceSafetyBuffer`,
		},
	},
	"StorageMiner": []DocField{
		{
			Name: "Subsystems",
			Type: "MinerSubsystemConfig",

			Comment: ``,
		},
		{
			Name: "Dealmaking",
			Type: "DealmakingConfig",

			Comment: ``,
		},
		{
			Name: "IndexProvider",
			Type: "IndexProviderConfig",

			Comment: ``,
		},
		{
			Name: "Proving",
			Type: "ProvingConfig",

			Comment: ``,
		},
		{
			Name: "Sealing",
			Type: "SealingConfig",

			Comment: ``,
		},
		{
			Name: "Storage",
			Type: "SealerConfig",

			Comment: ``,
		},
		{
			Name: "Fees",
			Type: "MinerFeeConfig",

			Comment: ``,
		},
		{
			Name: "Addresses",
			Type: "MinerAddressConfig",

			Comment: ``,
		},
		{
			Name: "DAGStore",
			Type: "DAGStoreConfig",

			Comment: ``,
		},
	},
	"UserRaftConfig": []DocField{
		{
			Name: "ClusterModeEnabled",
			Type: "bool",

			Comment: `EXPERIMENTAL. config to enabled node cluster with raft consensus`,
		},
		{
			Name: "DataFolder",
			Type: "string",

			Comment: `A folder to store Raft's data.`,
		},
		{
			Name: "InitPeersetMultiAddr",
			Type: "[]string",

			Comment: `InitPeersetMultiAddr provides the list of initial cluster peers for new Raft
peers (with no prior state). It is ignored when Raft was already
initialized or when starting in staging mode.`,
		},
		{
			Name: "WaitForLeaderTimeout",
			Type: "Duration",

			Comment: `LeaderTimeout specifies how long to wait for a leader before
failing an operation.`,
		},
		{
			Name: "NetworkTimeout",
			Type: "Duration",

			Comment: `NetworkTimeout specifies how long before a Raft network
operation is timed out`,
		},
		{
			Name: "CommitRetries",
			Type: "int",

			Comment: `CommitRetries specifies how many times we retry a failed commit until
we give up.`,
		},
		{
			Name: "CommitRetryDelay",
			Type: "Duration",

			Comment: `How long to wait between retries`,
		},
		{
			Name: "BackupsRotate",
			Type: "int",

			Comment: `BackupsRotate specifies the maximum number of Raft's DataFolder
copies that we keep as backups (renaming) after cleanup.`,
		},
		{
			Name: "Tracing",
			Type: "bool",

			Comment: `Tracing enables propagation of contexts across binary boundaries.`,
		},
	},
	"Wallet": []DocField{
		{
			Name: "RemoteBackend",
			Type: "string",

			Comment: ``,
		},
		{
			Name: "EnableLedger",
			Type: "bool",

			Comment: ``,
		},
		{
			Name: "DisableLocal",
			Type: "bool",

			Comment: ``,
		},
	},
}
