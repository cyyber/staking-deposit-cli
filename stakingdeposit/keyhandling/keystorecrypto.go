package keyhandling

import "crypto/sha256"

type KeystoreCrypto struct {
	kdf      *KeystoreModule
	checksum *KeystoreModule
	cipher   *KeystoreModule
}

func NewKeystoreCryptoFromJSON() *KeystoreCrypto {
	return nil
}

func NewKeystoreCrypto(salt, aesIV, cipherText, partialDecryptionKey []uint8) *KeystoreCrypto {
	var keyAndCipherText []byte
	keyAndCipherText = append(keyAndCipherText, partialDecryptionKey...)
	keyAndCipherText = append(keyAndCipherText, cipherText...)
	checksum := sha256.Sum256(keyAndCipherText)

	return &KeystoreCrypto{
		kdf: &KeystoreModule{
			params: map[string]interface{}{"salt": salt},
		},
		cipher: &KeystoreModule{
			params:  map[string]interface{}{"iv": aesIV},
			message: cipherText,
		},
		checksum: &KeystoreModule{
			message: checksum[:],
		},
	}
}
