// aBasic code is from: https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
// aI just created the benchmark function
// Command: go test uuid_test.go -bench=.
package benchmarks

import (
	"math/rand"
	"testing"
	"time"

	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func BenchmarkGenShortUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortuuid.New()
		//fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)
	}
}

func BenchmarkGenUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = guuid.New()
		//fmt.Printf("github.com/google/uuid:         %s\n", id.String())
	}
}

func BenchmarkGenXid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = xid.New()
		//fmt.Printf("github.com/rs/xid:              %s\n", id.String())
	}
}

func BenchmarkGenKsuid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ksuid.New()
		//fmt.Printf("github.com/segmentio/ksuid:     %s\n", id.String())
	}
}

func BenchmarkGenBetterGUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = betterguid.New()
		//fmt.Printf("github.com/kjk/betterguid:      %s\n", id)
	}
}

func BenchmarkGenUlid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := time.Now().UTC()
		entropy := rand.New(rand.NewSource(t.UnixNano()))
		_ = ulid.MustNew(ulid.Timestamp(t), entropy)
		//fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
	}
}

func BenchmarkGenSonyflake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		flake := sonyflake.NewSonyflake(sonyflake.Settings{})
		_, err := flake.NextID()
		if err != nil {
			b.Errorf("flake.NextID() failed with %s\n", err)
		}
		// Note: this is base16, could shorten by encoding as base62 string
		//fmt.Printf("github.com/sony/sonyflake:      %x\n", id)
	}
}

func BenchmarkGenSid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sid.Id()
		//fmt.Printf("github.com/chilts/sid:          %s\n", id)
	}
}

func BenchmarkGenUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := uuid.NewV4()
		if err != nil {
			b.Errorf("uuid.NewV4() failed with %s\n", err)
		}
		//fmt.Printf("github.com/satori/go.uuid:      %s\n", id)
	}
}
