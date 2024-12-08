//go:build windows

package openuem_utils

import (
	"encoding/binary"
	"fmt"

	"github.com/danieljoos/wincred"
)

func GetJWTKey() (string, error) {
	pass, err := wincred.GetGenericCredential("OpenUEM JWT Key")
	if err != nil {
		return "", fmt.Errorf("could not read password from Windows Credential Manager")
	}

	decodedPass := UTF16BytesToString(pass.CredentialBlob, binary.LittleEndian)
	return decodedPass, nil
}
