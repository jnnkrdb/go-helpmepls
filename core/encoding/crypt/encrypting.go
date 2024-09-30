// Please ensure, that in this package is no logging activated.
// There wont be console-outputs or other prints, you have to handle the "error", if any, by
// yourself.
//
// The gocrypt package contains a private variable, this variable stores the default encryption
// passphrase. Make sure to change the default encryption passphrase if you dont want to specify your
// own encryption passphrase everytime the encrypt/decrypt functions are called.
//
// Happy Encrypting!
package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	random "math/rand"
)

// the default encryption passphrase, nothing special
//
// can be changed
var _DEFAULT_Passphrase string = "DEFAULT-ENCRYPTION-PASSPHRASE"

// return the Default Encryption Passphrase
func GetDefaultPassphrase() string {

	return _DEFAULT_Passphrase
}

// set another Default Encryption Passphrase, but remember, if changed, the strings, which
// were encrypted with the old value cant be decrypted anymore
//
// Parameters:
//   - `passphrase` : string > contains the new encryption-passphrase
func SetDefaultPassphrase(passphrase string) {

	_DEFAULT_Passphrase = passphrase
}

// create a random passphrase for de- and encryption
//
// Parameters:
//   - `length` : int > length of the generated passphrase
func CreateRandomPassphrase(length int) string {

	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, length)

	for i := range b {

		b[i] = characters[random.Intn(len(characters))]
	}

	return string(b)
}

// ------------------------------------------------------------------
// encryption and decryption functions are defined here

// uses the encryption with the default passphrase, returns the encrypted version
// of "text" or an error
//
// Parameters:
//   - `text` : string > contains the text to encrypt
func EncryptWithDefault(text string) (string, error) {

	return Encrypt(_DEFAULT_Passphrase, text)
}

// uses the encryption with the given passphrase, returns the encrypted version
// of "text" or an error
//
// Parameters:
//   - `passphrase` : string > contains the string, which is used to encrypt the value of `text`
//   - `text` : string > contains the text to encrypt
func Encrypt(passphrase, text string) (string, error) {

	plaintext := []byte(text)

	if block, err := aes.NewCipher([]byte(passphrase)); err != nil {

		return "", err

	} else {

		ciphertext := make([]byte, aes.BlockSize+len(plaintext))

		iv := ciphertext[:aes.BlockSize]

		if _, err := io.ReadFull(rand.Reader, iv); err != nil {

			return "", err

		} else {

			stream := cipher.NewCFBEncrypter(block, iv)

			stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

			return base64.URLEncoding.EncodeToString(ciphertext), nil
		}
	}
}

// uses the decryption with the default passphrase, returns the decrypted version
// of "text" or an error
//
// Parameters:
//   - `text` : string > contains the text to decrypt
func DecryptWithDefault(text string) (string, error) {

	return Decrypt(_DEFAULT_Passphrase, text)
}

// uses the decryption with the given passphrase, returns the decrypted version
// of "text" or an error
//
// Parameters:
//   - `passphrase` : string > contains the string, which is used to decrypt the value of `text`
//   - `text` : string > contains the text to decrypt
func Decrypt(_passphrase, text string) (string, error) {

	if ciphertext, err := base64.URLEncoding.DecodeString(text); err != nil {

		return "", err

	} else {

		if block, err := aes.NewCipher([]byte(_passphrase)); err != nil {

			return "", err

		} else {

			if len(ciphertext) < aes.BlockSize {
				return "", errors.New("ciphertext to short")
			}

			iv := ciphertext[:aes.BlockSize]
			ciphertext = ciphertext[aes.BlockSize:]

			stream := cipher.NewCFBDecrypter(block, iv)

			stream.XORKeyStream(ciphertext, ciphertext)

			return string(ciphertext), nil
		}
	}
}
