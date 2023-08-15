package keyhandling

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/theQRL/go-qrllib/common"
	"github.com/theQRL/go-qrllib/dilithium"
	"golang.org/x/crypto/sha3"
)

type Keystore struct {
	Crypto      *KeystoreCrypto
	Description string
	PubKey      string
	Path        string
	UUID        string
	Version     string
}

func (k *Keystore) Save(fileFolder string) error {
	if err := os.WriteFile(fileFolder, k.ToJSON(), 0644); err != nil {
		return err
	}
	
	return nil
}

func Encrypt(seed [common.SeedSize]uint8, password, path string, salt, aesIV []byte) (*Keystore, error) {
	if salt == nil {
		salt = make([]uint8, 256)
		if _, err := io.ReadFull(rand.Reader, salt); err != nil {
			return nil, err
		}
	}
	if aesIV == nil {
		aesIV = make([]uint8, 128)
		if _, err := io.ReadFull(rand.Reader, aesIV); err != nil {
			return nil, err
		}
	}

	decryptionKey, err := passwordToDecryptionKey(password, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(decryptionKey[:16])
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(seed))

	stream := cipher.NewCTR(block, aesIV)
	stream.XORKeyStream(cipherText[aes.BlockSize:], seed[:])

	d, err := dilithium.NewDilithiumFromSeed(seed)
	if err != nil {
		return nil, err
	}
	pk := d.GetPK()
	return &Keystore{
		UUID:   uuid.New().String(),
		Crypto: NewKeystoreCrypto(salt, aesIV, cipherText, decryptionKey[16:]),
		PubKey: hex.EncodeToString(pk[:]),
		Path:   path,
	}, nil
}

func passwordToDecryptionKey(password string, salt []byte) ([32]byte, error) {
	h := sha3.NewShake256()
	if _, err := h.Write([]byte(password)); err != nil {
		return [32]byte{}, fmt.Errorf("shake256 hash write failed %v", err)
	}

	if _, err := h.Write([]byte(salt)); err != nil {
		return [32]byte{}, fmt.Errorf("shake256 hash write failed %v", err)
	}

	var decryptionKey [32]uint8
	_, err := h.Read(decryptionKey[:])
	return decryptionKey, err
}
