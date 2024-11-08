package types

import (
	abci "github.com/cometbft/cometbft/abci/types"
)

// InitChainer initializes application state at genesis
type InitChainer func(ctx Context, req abci.RequestInitChain) abci.ResponseInitChain

// BeginBlocker runs code before the transactions in a block
//
// Note: applications which set create_empty_blocks=false will not have regular block timing and should use
// e.g. BFT timestamps rather than block height for any periodic BeginBlock logic
type BeginBlocker func(Context) (BeginBlock, error)

// EndBlocker runs code after the transactions in a block and return updates to the validator set
//
// Note: applications which set create_empty_blocks=false will not have regular block timing and should use
// e.g. BFT timestamps rather than block height for any periodic EndBlock logic
type EndBlocker func(Context) (EndBlock, error)

// EndBlock defines a type which contains endblock events and validator set updates
type EndBlock struct {
	ValidatorUpdates []abci.ValidatorUpdate
	Events           []abci.Event
}

// PeerFilter responds to p2p filtering queries from Tendermint
type PeerFilter func(info string) abci.ResponseQuery

// ProcessProposalHandler defines a function type alias for processing a proposer
type ProcessProposalHandler func(Context, abci.RequestProcessProposal) abci.ResponseProcessProposal

// PrepareProposalHandler defines a function type alias for preparing a proposal
type PrepareProposalHandler func(Context, abci.RequestPrepareProposal) abci.ResponsePrepareProposal

type ResponsePreBlock struct {
	ConsensusParamsChanged bool
}

// PreBlocker runs code before the `BeginBlocker` and defines a function type alias for executing logic right
// before FinalizeBlock is called (but after its context has been set up). It is
// intended to allow applications to perform computation on vote extensions and
// persist their results in state.
//
// Note: returning an error will make FinalizeBlock fail.
type PreBlocker func(Context, *abci.RequestFinalizeBlock) (*ResponsePreBlock, error)

// BeginBlock defines a type which contains beginBlock events
type BeginBlock struct {
	Events []abci.Event
}
