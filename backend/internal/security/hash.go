package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	"crypto/subtle"
	"errors"

	"golang.org/x/crypto/argon2"
)

type Params struct {
	Time       uint32
	Memory     uint32
	Threads    uint8
	KeyLength  uint32
	SaltLength uint32
}

var (
	timeEnv, _       = strconv.ParseUint(os.Getenv("TIME"), 10, 32)
	memoryEnv, _     = strconv.ParseUint(os.Getenv("MEMORY"), 10, 32)
	threadsEnv, _    = strconv.ParseUint(os.Getenv("THREADS"), 10, 8)
	keyLengthEnv, _  = strconv.ParseUint(os.Getenv("KEYLENGTH"), 10, 32)
	saltLengthEnv, _ = strconv.ParseUint(os.Getenv("SALTLENGTH"), 10, 32)
)

var p = &Params{
	Time:       uint32(timeEnv),
	Memory:     uint32(memoryEnv),
	Threads:    uint8(threadsEnv),
	KeyLength:  uint32(keyLengthEnv),
	SaltLength: uint32(saltLengthEnv),
}

// HashPassword génère un hash Argon2id encodé pour stockage.
func HashPassword(password string) (string, error) {
	salt := make([]byte, p.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, p.Time, p.Memory, p.Threads, p.KeyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encoded := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		p.Memory, p.Time, p.Threads, b64Salt, b64Hash,
	)
	return encoded, nil
}

// VerifyPassword compare un mot de passe avec son hash encodé.
func VerifyPassword(encodedHash, password string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("hash Argon2: format invalide")
	}
	var mem, time uint32
	var threads uint8
	fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &time, &threads)

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	calc := argon2.IDKey([]byte(password), salt, time, mem, threads, uint32(len(hash)))
	if subtle.ConstantTimeCompare(calc, hash) == 1 {
		return true, nil
	}
	return false, nil
}
