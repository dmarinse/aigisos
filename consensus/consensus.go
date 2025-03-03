package consensus

import (
	"context"
	"log"

	"github.com/JResearchLabs/aigisos/blockchain"
	"github.com/JResearchLabs/aigisos/chain"
	"github.com/JResearchLabs/aigisos/helper/progress"
	"github.com/JResearchLabs/aigisos/network"
	"github.com/JResearchLabs/aigisos/secrets"
	"github.com/JResearchLabs/aigisos/state"
	"github.com/JResearchLabs/aigisos/txpool"
	"github.com/JResearchLabs/aigisos/types"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	
)

// Consensus is the public interface for consensus mechanism
// Each consensus mechanism must implement this interface in order to be valid
type Consensus interface {
	// VerifyHeader verifies the header is correct
	VerifyHeader(header *types.Header) error

	// ProcessHeaders updates the snapshot based on the verified headers
	ProcessHeaders(headers []*types.Header) error

	// GetBlockCreator retrieves the block creator (or signer) given the block header
	GetBlockCreator(header *types.Header) (types.Address, error)

	// PreStateCommit a hook to be called before finalizing state transition on inserting block
	PreStateCommit(header *types.Header, txn *state.Transition) error

	// GetSyncProgression retrieves the current sync progression, if any
	GetSyncProgression() *progress.Progression

	// Initialize initializes the consensus (e.g. setup data)
	Initialize() error

	// Start starts the consensus and servers
	Start() error

	// Close closes the connection
	Close() error
}

// Config is the configuration for the consensus
type Config struct {
	// Logger to be used by the consensus
	Logger *log.Logger

	// Params are the params of the chain and the consensus
	Params *chain.Params

	// Config defines specific configuration parameters for the consensus
	Config map[string]interface{}

	// Path is the directory path for the consensus protocol tos tore information
	Path string
}

type Params struct {
	Context        context.Context
	Seal           bool
	Config         *Config
	TxPool         *txpool.TxPool
	Network        *network.Server
	Blockchain     *blockchain.Blockchain
	Executor       *state.Executor
	Grpc           *grpc.Server
	Logger         hclog.Logger
	Metrics        *Metrics
	SecretsManager secrets.SecretsManager
	BlockTime      uint64
}

// Factory is the factory function to create a discovery consensus
type Factory func(*Params) (Consensus, error)
