package stakingdeposit

import "github.com/cyyber/staking-deposit-cli/network"

func GenerateKeys(validatorStartIndex, numValidators uint,
	folder, chain, keystorePassword, executionAddress string) {
	chainSettings, ok := network.GetConfig().ChainSettings[chain]
	if !ok {

	}
	credentials, err := NewCredentialsFromMnemonic(seed, numValidators, amounts, chainSettings, validatorStartIndex, executionAddress)
	if err != nil {

	}
	keystoreFileFolders := credentials.ExportKeystores(keystorePassword, folder)
	depositFile, err := credentials.ExportDepositDataJSON(folder)
	if !credentials.VerifyKeystores(keystoreFileFolders, keystorePassword) {
		panic("failed to verify the keystores")
	}
	if !VerifyDepositDataJSON(depositFile, credentials.credentials) {
		panic("failed to verify the deposit data JSON files")
	}
}
