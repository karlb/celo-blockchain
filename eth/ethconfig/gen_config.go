// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethconfig

import (
	"math/big"
	"time"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/consensus/istanbul"
	"github.com/celo-org/celo-blockchain/core"
	"github.com/celo-org/celo-blockchain/eth/downloader"
	"github.com/celo-org/celo-blockchain/miner"
	"github.com/celo-org/celo-blockchain/params"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		EthDiscoveryURLs        []string
		SnapDiscoveryURLs       []string
		NoPruning               bool
		NoPrefetch              bool
		TxLookupLimit           uint64                 `toml:",omitempty"`
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightServ               int                    `toml:",omitempty"`
		LightIngress            int                    `toml:",omitempty"`
		LightEgress             int                    `toml:",omitempty"`
		LightPeers              int                    `toml:",omitempty"`
		LightNoPrune            bool                   `toml:",omitempty"`
		LightNoSyncServe        bool                   `toml:",omitempty"`
		SyncFromCheckpoint      bool                   `toml:",omitempty"`
		GatewayFee              *big.Int               `toml:",omitempty"`
		Validator               common.Address         `toml:",omitempty"`
		TxFeeRecipient          common.Address         `toml:",omitempty"`
		BLSbase                 common.Address         `toml:",omitempty"`
		UltraLightServers       []string               `toml:",omitempty"`
		UltraLightFraction      int                    `toml:",omitempty"`
		UltraLightOnlyAnnounce  bool                   `toml:",omitempty"`
		SkipBcVersionCheck      bool                   `toml:"-"`
		DatabaseHandles         int                    `toml:"-"`
		DatabaseCache           int
		DatabaseFreezer         string
		TrieCleanCache          int
		TrieCleanCacheJournal   string        `toml:",omitempty"`
		TrieCleanCacheRejournal time.Duration `toml:",omitempty"`
		TrieDirtyCache          int
		TrieTimeout             time.Duration
		SnapshotCache           int
		Preimages               bool
		Miner                   miner.Config
		TxPool                  core.TxPoolConfig
		EnablePreimageRecording bool
		Istanbul                istanbul.Config
		DocRoot                 string `toml:"-"`
		RPCGasInflationRate     float64
		RPCGasCap               uint64
		RPCTxFeeCap             float64
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		OverrideEHardfork       *big.Int                       `toml:",omitempty"`
		OverrideV2IstanbulFork  *big.Int                       `toml:",omitempty"`
		MinSyncPeers            int                            `toml:",omitempty"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.EthDiscoveryURLs = c.EthDiscoveryURLs
	enc.SnapDiscoveryURLs = c.SnapDiscoveryURLs
	enc.NoPruning = c.NoPruning
	enc.NoPrefetch = c.NoPrefetch
	enc.TxLookupLimit = c.TxLookupLimit
	enc.Whitelist = c.Whitelist
	enc.LightServ = c.LightServ
	enc.LightIngress = c.LightIngress
	enc.LightEgress = c.LightEgress
	enc.LightPeers = c.LightPeers
	enc.LightNoPrune = c.LightNoPrune
	enc.LightNoSyncServe = c.LightNoSyncServe
	enc.SyncFromCheckpoint = c.SyncFromCheckpoint
	enc.GatewayFee = c.GatewayFee
	enc.Validator = c.Validator
	enc.TxFeeRecipient = c.TxFeeRecipient
	enc.BLSbase = c.BLSbase
	enc.UltraLightServers = c.UltraLightServers
	enc.UltraLightFraction = c.UltraLightFraction
	enc.UltraLightOnlyAnnounce = c.UltraLightOnlyAnnounce
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieCleanCacheJournal = c.TrieCleanCacheJournal
	enc.TrieCleanCacheRejournal = c.TrieCleanCacheRejournal
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.SnapshotCache = c.SnapshotCache
	enc.Preimages = c.Preimages
	enc.Miner = c.Miner
	enc.TxPool = c.TxPool
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.Istanbul = c.Istanbul
	enc.DocRoot = c.DocRoot
	enc.RPCGasInflationRate = c.RPCGasInflationRate
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	enc.Checkpoint = c.Checkpoint
	enc.CheckpointOracle = c.CheckpointOracle
	enc.OverrideEHardfork = c.OverrideEHardfork
	enc.OverrideV2IstanbulFork = c.OverrideV2IstanbulFork
	enc.MinSyncPeers = c.MinSyncPeers
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		EthDiscoveryURLs        []string
		SnapDiscoveryURLs       []string
		NoPruning               *bool
		NoPrefetch              *bool
		TxLookupLimit           *uint64                `toml:",omitempty"`
		Whitelist               map[uint64]common.Hash `toml:"-"`
		LightServ               *int                   `toml:",omitempty"`
		LightIngress            *int                   `toml:",omitempty"`
		LightEgress             *int                   `toml:",omitempty"`
		LightPeers              *int                   `toml:",omitempty"`
		LightNoPrune            *bool                  `toml:",omitempty"`
		LightNoSyncServe        *bool                  `toml:",omitempty"`
		SyncFromCheckpoint      *bool                  `toml:",omitempty"`
		GatewayFee              *big.Int               `toml:",omitempty"`
		Validator               *common.Address        `toml:",omitempty"`
		TxFeeRecipient          *common.Address        `toml:",omitempty"`
		BLSbase                 *common.Address        `toml:",omitempty"`
		UltraLightServers       []string               `toml:",omitempty"`
		UltraLightFraction      *int                   `toml:",omitempty"`
		UltraLightOnlyAnnounce  *bool                  `toml:",omitempty"`
		SkipBcVersionCheck      *bool                  `toml:"-"`
		DatabaseHandles         *int                   `toml:"-"`
		DatabaseCache           *int
		DatabaseFreezer         *string
		TrieCleanCache          *int
		TrieCleanCacheJournal   *string        `toml:",omitempty"`
		TrieCleanCacheRejournal *time.Duration `toml:",omitempty"`
		TrieDirtyCache          *int
		TrieTimeout             *time.Duration
		SnapshotCache           *int
		Preimages               *bool
		Miner                   *miner.Config
		TxPool                  *core.TxPoolConfig
		EnablePreimageRecording *bool
		Istanbul                *istanbul.Config
		DocRoot                 *string `toml:"-"`
		RPCGasInflationRate     *float64
		RPCGasCap               *uint64
		RPCTxFeeCap             *float64
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		OverrideEHardfork       *big.Int                       `toml:",omitempty"`
		OverrideV2IstanbulFork  *big.Int                       `toml:",omitempty"`
		MinSyncPeers            *int                           `toml:",omitempty"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.EthDiscoveryURLs != nil {
		c.EthDiscoveryURLs = dec.EthDiscoveryURLs
	}
	if dec.SnapDiscoveryURLs != nil {
		c.SnapDiscoveryURLs = dec.SnapDiscoveryURLs
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.NoPrefetch != nil {
		c.NoPrefetch = *dec.NoPrefetch
	}
	if dec.TxLookupLimit != nil {
		c.TxLookupLimit = *dec.TxLookupLimit
	}
	if dec.Whitelist != nil {
		c.Whitelist = dec.Whitelist
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightIngress != nil {
		c.LightIngress = *dec.LightIngress
	}
	if dec.LightEgress != nil {
		c.LightEgress = *dec.LightEgress
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.LightNoPrune != nil {
		c.LightNoPrune = *dec.LightNoPrune
	}
	if dec.LightNoSyncServe != nil {
		c.LightNoSyncServe = *dec.LightNoSyncServe
	}
	if dec.SyncFromCheckpoint != nil {
		c.SyncFromCheckpoint = *dec.SyncFromCheckpoint
	}
	if dec.GatewayFee != nil {
		c.GatewayFee = dec.GatewayFee
	}
	if dec.Validator != nil {
		c.Validator = *dec.Validator
	}
	if dec.TxFeeRecipient != nil {
		c.TxFeeRecipient = *dec.TxFeeRecipient
	}
	if dec.BLSbase != nil {
		c.BLSbase = *dec.BLSbase
	}
	if dec.UltraLightServers != nil {
		c.UltraLightServers = dec.UltraLightServers
	}
	if dec.UltraLightFraction != nil {
		c.UltraLightFraction = *dec.UltraLightFraction
	}
	if dec.UltraLightOnlyAnnounce != nil {
		c.UltraLightOnlyAnnounce = *dec.UltraLightOnlyAnnounce
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieCleanCacheJournal != nil {
		c.TrieCleanCacheJournal = *dec.TrieCleanCacheJournal
	}
	if dec.TrieCleanCacheRejournal != nil {
		c.TrieCleanCacheRejournal = *dec.TrieCleanCacheRejournal
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.SnapshotCache != nil {
		c.SnapshotCache = *dec.SnapshotCache
	}
	if dec.Preimages != nil {
		c.Preimages = *dec.Preimages
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.Istanbul != nil {
		c.Istanbul = *dec.Istanbul
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.RPCGasInflationRate != nil {
		c.RPCGasInflationRate = *dec.RPCGasInflationRate
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = *dec.RPCGasCap
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	if dec.Checkpoint != nil {
		c.Checkpoint = dec.Checkpoint
	}
	if dec.CheckpointOracle != nil {
		c.CheckpointOracle = dec.CheckpointOracle
	}
	if dec.OverrideEHardfork != nil {
		c.OverrideEHardfork = dec.OverrideEHardfork
	}
	if dec.OverrideV2IstanbulFork != nil {
		c.OverrideV2IstanbulFork = dec.OverrideV2IstanbulFork
	}
	if dec.MinSyncPeers != nil {
		c.MinSyncPeers = *dec.MinSyncPeers
	}
	return nil
}
