package blockchain

import (
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ethereumSimulator SimulatedEthereumEnvironment
var blockchainTest *blockchainTestType

func TestMain(m *testing.M) {
	ethereumSimulator = GetSimulatedEthereumEnvironment()
	defer ethereumSimulator.Close()

	blockchainTest = func() *blockchainTestType {
		return &blockchainTestType{
			processor: &Processor{
				rawClient: ethereumSimulator.EthClient,
			},
		}
	}()

	err := ethereumSimulator.EthServer.RegisterName("eth", &Service{})
	if err != nil {
		panic(err)
	}

	result := m.Run()

	os.Exit(result)
}

type blockchainTestType struct {
	processor *Processor
}

type Service struct {
}

func (s *Service) BlockNumber() (string, error) {
	return "FF", nil
}

func TestCurrentBlock(t *testing.T) {

	block, err := blockchainTest.processor.CurrentBlock()

	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(255), block)
}
