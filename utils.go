package M00N

import (
	"math/big"
	"unsafe"

	"github.com/NginProject/ngind/common"
)

var (
	maxUint256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))
)

func bytesToHash(in unsafe.Pointer) common.Hash {
	return *(*common.Hash)(in)
}
