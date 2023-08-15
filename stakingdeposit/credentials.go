package stakingdeposit

import (
	"github.com/cyyber/staking-deposit-cli/network"
)

type Credentials struct {
	credentials []*Credential
}

func NewCredentialsFromMnemonic(seed string, numKeys uint, amounts []uint,
	chainSettings *network.ChainSetting, startIndex uint, hexZondWithdrawalAddress string) (*Credentials, error) {
	credentials := &Credentials{
		credentials: make([]*Credential, numKeys),
	}
	for index := startIndex; index < startIndex+numKeys; index++ {
		c, err := NewCredential(seed, index, amounts[index], chainSettings, hexZondWithdrawalAddress)
		if err != nil {
			return nil, err
		}
		credentials.credentials[index] = c
	}
	return credentials, nil
}
