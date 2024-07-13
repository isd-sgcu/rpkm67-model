package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"regexp"
	"time"

	"github.com/google/uuid"
)

func hash(bv []byte) string {
	h := sha256.New()
	h.Write(bv)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GenGroupToken(id *uuid.UUID) string {
	hashed := hash([]byte(time.Now().String() + id.String()))
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return re.ReplaceAllString(hashed, "")
}
