/*
 * Minio Cloud Storage, (C) 2015, 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"crypto/rand"
	"encoding/base64"
)

const (
	accessKeyMinLen = 5
	accessKeyMaxLen = 20
	secretKeyMinLen = 8
	secretKeyMaxLen = 40

	alphaNumericTable    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaNumericTableLen = byte(len(alphaNumericTable))
)

func mustGetAccessKey() string {
	keyBytes := make([]byte, accessKeyMaxLen)
	if _, err := rand.Read(keyBytes); err != nil {
		panic(err)
	}

	for i := 0; i < accessKeyMaxLen; i++ {
		keyBytes[i] = alphaNumericTable[keyBytes[i]%alphaNumericTableLen]
	}

	return string(keyBytes)
}

func mustGetSecretKey() string {
	keyBytes := make([]byte, secretKeyMaxLen)
	if _, err := rand.Read(keyBytes); err != nil {
		panic(err)
	}

	return string([]byte(base64.StdEncoding.EncodeToString(keyBytes))[:secretKeyMaxLen])
}

// isAccessKeyValid - validate access key for right length.
func isAccessKeyValid(accessKey string) bool {
	return len(accessKey) >= accessKeyMinLen && len(accessKey) <= accessKeyMaxLen
}

// isSecretKeyValid - validate secret key for right length.
func isSecretKeyValid(secretKey string) bool {
	return len(secretKey) >= secretKeyMinLen && len(secretKey) <= secretKeyMaxLen
}

// credential container for access and secret keys.
type credential struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func newCredential() credential {
	return credential{mustGetAccessKey(), mustGetSecretKey()}
}

func getCredential(accessKey, secretKey string) (credential, error) {
	if !isAccessKeyValid(accessKey) {
		return credential{}, errInvalidAccessKeyLength
	}

	if !isSecretKeyValid(secretKey) {
		return credential{}, errInvalidSecretKeyLength
	}

	return credential{accessKey, secretKey}, nil
}
