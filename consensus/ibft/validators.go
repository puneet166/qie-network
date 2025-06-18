package ibft

import (
	// "math"
 "fmt"
	"github.com/0xPolygon/polygon-edge/types"
	"github.com/0xPolygon/polygon-edge/validators"
)

func CalcMaxFaultyNodes(s validators.Validators) int {
	    fmt.Println("CalcMaxFaultyNodes---------------!",(s.Len() - 1) / 3)
	// Keeping original calculation but not used in quorum override below
	return (s.Len() - 1) / 3
}

type QuorumImplementation func(validators.Validators) int

// LegacyQuorumSize overrides quorum to 2 if 2 or more validators exist,
// allowing block production with just 2 validators.
func LegacyQuorumSize(set validators.Validators) int {
	
	if set.Len() >= 2 {
		return 2
	}
	return set.Len()
}

// OptimalQuorumSize overrides quorum to 2 if 2 or more validators exist,
// similar to LegacyQuorumSize.
func OptimalQuorumSize(set validators.Validators) int {
	if set.Len() >= 2 {
		return 2
	}
	return set.Len()
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
