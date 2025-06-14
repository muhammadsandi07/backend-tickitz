package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashConfig struct {
	Threads uint8
	Time    uint32
	Memory  uint32
	Keylen  uint32
	Saltlen uint32
}

func InitHashConfig() *HashConfig {
	return &HashConfig{}

}
func (h *HashConfig) UseConfig(time, memory, keylen, saltlen uint32, threads uint8) {
	h.Time = time
	h.Threads = threads
	h.Keylen = keylen
	h.Memory = memory
	h.Saltlen = saltlen
}

func (h *HashConfig) UseDefaultConfig() {
	h.Threads = 2
	h.Time = 3
	h.Memory = 64 * 1024
	h.Keylen = 32
	h.Saltlen = 16
}

func (h *HashConfig) GenSalt() ([]byte, error) {
	salt := make([]byte, h.Saltlen)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func (h *HashConfig) GenHashedPassword(password string) (string, error) {
	// hash = password + salt +config
	salt, err := h.GenSalt()

	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, h.Time, h.Memory, h.Threads, h.Keylen)
	// dalam penulisan hash ada format

	version := argon2.Version
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)
	hashedPwd := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", version, h.Memory, h.Time, h.Threads, base64Salt, base64Hash)
	return hashedPwd, nil

}
func (h *HashConfig) CompareHashAndPassword(hashedPass string, password string) (bool, error) {

	salt, hash, err := h.decodeHash(hashedPass)

	if err != nil {
		return false, err
	}
	newHash := argon2.IDKey([]byte(password), salt, h.Time, h.Memory, h.Threads, h.Keylen)

	if subtle.ConstantTimeCompare(hash, newHash) == 0 {
		return false, err
	}
	return true, nil
}
func (h *HashConfig) decodeHash(hashedPass string) (salt []byte, hash []byte, err error) {
	values := strings.Split(hashedPass, "$")

	if len(values) != 6 {
		return nil, nil, errors.New("invalid 1 format")
	}
	if values[1] != "argon2id" {
		return nil, nil, errors.New("invalid hash type")
	}
	var version int
	if _, err := fmt.Sscanf(values[2], "v=%d", &version); err != nil {
		return nil, nil, errors.New("Invalid version format")
	}
	if version != argon2.Version {
		return nil, nil, errors.New("invalid argon2 version")
	}

	if _, err := fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &h.Memory, &h.Time, &h.Threads); err != nil {
		return nil, nil, errors.New("invalid N format")
	}

	salt, err = base64.RawStdEncoding.DecodeString(values[4])
	if err != nil {
		return nil, nil, err
	}
	h.Saltlen = uint32(len(salt))
	hash, err = base64.RawStdEncoding.DecodeString(values[5])
	if err != nil {
		return nil, nil, err
	}
	h.Keylen = uint32(len(hash))
	return salt, hash, nil
}



// student
// student id
// name
// email


// course
// course code
// title
// credit hourse

// student_course
// student_id
// course_code
// semester
// grade 


// users          
// id pk 	   
// name
// address		


// menu   
// id pk 
// menu name
// price

// users_meals
// order_id
// menu_id
// pk (order_id, menu_id)



// orders
// id pk
// user id fk
// status
// total price




