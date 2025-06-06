package ea

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"

	"github.com/solaire/genie/internal/utils"

	"golang.org/x/crypto/sha3"
)

type EaManifest struct {
	InstallInfos []InstallInfo `json:"installInfos"`
}

type InstallInfo struct {
	InstallPath string `json:"baseInstallPath"`
	InstallReg  string `json:"installCheck"`
}

// https://erri120.github.io/posts/2023-01-18/
// https://github.com/erri120/GameFinder/wiki/EA-Desktop
// Decrypt the manifest file and parse the json information.
// The 'IS' file is encrypted using AES with a key size of 256
//
//	  bits in the Cipher Block Chaining (CBC) mode that requires an
//	Initialization Vector (IV) of 128 bits.
func decryptManifest(path string) (*EaManifest, error) {
	iv := createIV()
	key := createDecryptKey()
	// path = `C:\Users\Kacper\Downloads\IS_erri120.encrypted`

	// NOTE: Decryption on the example works, but the real IS does not
	//  checked against the C# lib and the IV, key is the same (and it failed too).
	decrypted, err := decrypt(path, iv, key)
	if err != nil {
		return nil, err
	}

	var manifest EaManifest
	if err := json.Unmarshal(decrypted, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

func createIV() []byte {
	hash := sha3.Sum256([]byte("allUsersGenericIdIS"))
	return hash[:16]
}

func createDecryptKey() []byte {
	m, _ := utils.GetHardwareInfo()
	sha1Sum := sha1.Sum([]byte(m))
	sha1Hex := hex.EncodeToString(sha1Sum[:])
	combined := "allUsersGenericIdIS" + sha1Hex
	key := sha3.Sum256([]byte(combined))
	return key[:]
}

func decrypt(file string, iv, key []byte) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if len(data) <= 64 {
		return nil, errors.New("file too short to skip header")
	}

	// logger.Printf("IV: %s\n", hex.EncodeToString(iv[:]))
	// logger.Printf("KEY: %s\n", hex.EncodeToString(key[:]))
	// logger.Printf("DATA: %x\n", data)

	cipertext := data[64:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(cipertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(cipertext))
	mode.CryptBlocks(plaintext, cipertext)

	// Strip PKCS#7 padding
	plaintext, err = pkcs7Unpad(plaintext, aes.BlockSize)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func pkcs7Unpad(data []byte, blocksize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blocksize != 0 {
		return nil, errors.New("invalid PKCS#7 data")
	}

	pad_len := int(data[len(data)-1])
	if pad_len == 0 || pad_len > blocksize {
		return nil, errors.New("invalid padding length")
	}

	for i := 0; i < pad_len; i++ {
		if data[len(data)-1-i] != byte(pad_len) {
			return nil, errors.New("invalid PKCS#7 padding")
		}
	}
	return data[:len(data)-pad_len], nil
}
