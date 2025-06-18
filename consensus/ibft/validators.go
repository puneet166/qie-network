package ibft

import (
	// "math"
 "fmt"
	"github.com/0xPolygon/polygon-edge/types"
	"github.com/0xPolygon/polygon-edge/validators"
)

func CalcMaxFaultyNodes(s validators.Validators) int {
	fmt.Println("FAKE faulty nodes: returning 0 regardless of validator count")
	return 0
}

type QuorumImplementation func(validators.Validators) int

// LegacyQuorumSize overrides quorum to 2 if 2 or more validators exist,
// allowing block production with just 2 validators.
func LegacyQuorumSize(set validators.Validators) int {
	fmt.Println("FORCING legacy quorum size to 2")
	return 2
}

// OptimalQuorumSize overrides quorum to 2 if 2 or more validators exist,
// similar to LegacyQuorumSize.
func OptimalQuorumSize(set validators.Validators) int {
	fmt.Println("FORCING quorum size to 2 no matter what")
	return 2
}

func CalcProposer(
	validators validators.Validators,
	round uint64,
	lastProposer types.Address,
) validators.Validator {
	var seed uint64

	if lastProposer == types.ZeroAddress {
		seed = round
	} else {
		offset := int64(0)

		if index := validators.Index(lastProposer); index != -1 {
			offset = index
		}

		seed = uint64(offset) + round + 1
	}

	pick := seed % uint64(validators.Len())

	return validators.At(pick)
}
