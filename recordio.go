package recordio

import "errors"

type Mode uint32

const (
	modeBase Mode = 0xe0319fd0 // fnv1a32("RecordIO")

	// No compression.
	NoCompression = modeBase + iota
	// Compression with Snappy.
	Snappy
	// Compression with zlib.
	Zlib

	// End of mode.
	EndOfMode

	recHeaderSize = 12
	maxRecordSize = 1024 * 1024 * 1024 // 1024MB
)

// String implements fmt.Stringer interface.
func (m Mode) String() string {
	switch m {
	case NoCompression:
		return "nocompression"
	case Snappy:
		return "snappy"
	case Zlib:
		return "zlib"
	default:
		return "unknown"
	}
}

var (
	// ErrInvalidHeader means that a record header is not valid.
	ErrInvalidHeader = errors.New("invalid record header")

	// ErrInvalidMode means that a (compression) mode is not valid.
	ErrInvalidMode = errors.New("invalid mode")

	// ErrRecordCorrupted means that a record is corrupted.
	ErrRecordCorrupted = errors.New("record corrupted")

	// ErrRecordTooLarge means that a record is too large.
	ErrRecordTooLarge = errors.New("record too large")
)
