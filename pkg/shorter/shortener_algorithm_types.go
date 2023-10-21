package shortener

type Algorithm int

const (
	CRC32 Algorithm = iota
	MD5
	SHA1
	SHA256
)

type Shortener struct {
	Algorithm
}
