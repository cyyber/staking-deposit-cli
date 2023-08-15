package newmnemonic

import (
	"github.com/sirupsen/logrus"
	"github.com/theQRL/go-qrllib/dilithium"
	"github.com/urfave/cli/v2"
)

var (
	newMnemonicFlags = struct {
		ValidatorStartIndex uint64
		NumValidators       uint64
		Folder              string
		ChainName           string
		KeystorePassword    string
		ExecutionAddress    string
	}{}
	log = logrus.WithField("prefix", "deposit")
)
var Commands = []*cli.Command{
	{
		Name:    "new-mnemonic",
		Aliases: []string{"new-mnemonic"},
		Usage:   "",
		Action: func(cliCtx *cli.Context) error {
			if err := cliActionNewMnemonic(cliCtx); err != nil {
				log.WithError(err).Fatal("Could not generate new mnemonic")
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:        "validator-start-index",
				Usage:       "",
				Destination: &newMnemonicFlags.ValidatorStartIndex,
				Required:    true,
			},
			&cli.Uint64Flag{
				Name:        "num-validators",
				Usage:       "",
				Destination: &newMnemonicFlags.NumValidators,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "folder",
				Usage:       "",
				Destination: &newMnemonicFlags.Folder,
				Value:       "betanet",
			},
			&cli.StringFlag{
				Name:        "chain-name",
				Usage:       "",
				Destination: &newMnemonicFlags.ChainName,
				Value:       "betanet",
			},
			&cli.StringFlag{
				Name:        "keystore-password",
				Usage:       "",
				Destination: &newMnemonicFlags.KeystorePassword,
				Value:       "betanet",
			},
			&cli.StringFlag{
				Name:        "execution-address",
				Usage:       "",
				Destination: &newMnemonicFlags.ExecutionAddress,
				Value:       "betanet",
			},
		},
	},
}

func cliActionNewMnemonic(cliCtx *cli.Context) error {
	z := dilithium.New()
	return nil
}
