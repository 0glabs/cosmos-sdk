package baseapp

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type state struct {
	ms  sdk.CacheMultiStore
	mtx sync.RWMutex
	ctx sdk.Context
}

// CacheMultiStore calls and returns a CacheMultiStore on the state's underling
// CacheMultiStore.
func (st *state) CacheMultiStore() sdk.CacheMultiStore {
	return st.ms.CacheMultiStore()
}

// Context returns the Context of the state.
func (st *state) Context() sdk.Context {
	return st.ctx
}

// SetContext updates the state's context to the context provided.
func (st *state) SetContext(ctx sdk.Context) {
	st.mtx.Lock()
	defer st.mtx.Unlock()
	st.ctx = ctx
}
