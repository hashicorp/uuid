package uuid

import (
	"crypto/rand"
	"fmt"
)

// GenerateUUID is used to generate a random UUID
func GenerateUUID() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v", err))
	}
	buf[6] = (buf[6] & 0x0F) | 0x40 // Version: https://tools.ietf.org/html/rfc4122#section-4.1.3
	buf[8] = (buf[8] & 0x3F) | 0x80 // Variant: https://tools.ietf.org/html/rfc4122#section-4.1.1

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}
