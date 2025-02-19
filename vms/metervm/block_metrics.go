// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

type blockMetrics struct {
	buildBlock,
	buildBlockErr,
	parseBlock,
	parseBlockErr,
	getBlock,
	getBlockErr,
	setPreference,
	lastAccepted,
	verify,
	verifyErr,
	accept,
	reject,
	// Batched metrics
	getAncestors,
	batchedParseBlock,
	// Height metrics
	verifyHeightIndex,
	getBlockIDAtHeight,
	// State sync metrics
	stateSyncEnabled,
	getOngoingSyncStateSummary,
	getLastStateSummary,
	parseStateSummary,
	parseStateSummaryErr,
	getStateSummary,
	getStateSummaryErr metric.Averager
}

func (m *blockMetrics) Initialize(
	supportsBatchedFetching bool,
	supportsHeightIndexing bool,
	supportsStateSync bool,
	namespace string,
	reg prometheus.Registerer,
) error {
	errs := wrappers.Errs{}
	m.buildBlock = newAverager(namespace, "build_block", reg, &errs)
	m.buildBlockErr = newAverager(namespace, "build_block_err", reg, &errs)
	m.parseBlock = newAverager(namespace, "parse_block", reg, &errs)
	m.parseBlockErr = newAverager(namespace, "parse_block_err", reg, &errs)
	m.getBlock = newAverager(namespace, "get_block", reg, &errs)
	m.getBlockErr = newAverager(namespace, "get_block_err", reg, &errs)
	m.setPreference = newAverager(namespace, "set_preference", reg, &errs)
	m.lastAccepted = newAverager(namespace, "last_accepted", reg, &errs)
	m.verify = newAverager(namespace, "verify", reg, &errs)
	m.verifyErr = newAverager(namespace, "verify_err", reg, &errs)
	m.accept = newAverager(namespace, "accept", reg, &errs)
	m.reject = newAverager(namespace, "reject", reg, &errs)

	if supportsBatchedFetching {
		m.getAncestors = newAverager(namespace, "get_ancestors", reg, &errs)
		m.batchedParseBlock = newAverager(namespace, "batched_parse_block", reg, &errs)
	}
	if supportsHeightIndexing {
		m.verifyHeightIndex = newAverager(namespace, "verify_height_index", reg, &errs)
		m.getBlockIDAtHeight = newAverager(namespace, "get_block_id_at_height", reg, &errs)
	}
	if supportsStateSync {
		m.stateSyncEnabled = newAverager(namespace, "state_sync_enabled", reg, &errs)
		m.getOngoingSyncStateSummary = newAverager(namespace, "get_ongoing_state_sync_summary", reg, &errs)
		m.getLastStateSummary = newAverager(namespace, "get_last_state_summary", reg, &errs)
		m.parseStateSummary = newAverager(namespace, "parse_state_summary", reg, &errs)
		m.parseStateSummaryErr = newAverager(namespace, "parse_state_summary_err", reg, &errs)
		m.getStateSummary = newAverager(namespace, "get_state_summary", reg, &errs)
		m.getStateSummaryErr = newAverager(namespace, "get_state_summary_err", reg, &errs)
	}
	return errs.Err
}
