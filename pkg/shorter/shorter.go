package shortener

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash/crc32"
	"url-shortener/pkg/log"
)

var (
	logger = logHandler.Logger()
)

func (shortener *Shortener) ShortURLCRC32(src string) string {
	logger.Info("")
	hash := crc32.ChecksumIEEE([]byte(src))
	shortURL := fmt.Sprintf("%x", hash)
	return shortURL
}

func (shortener *Shortener) ShortURLMD5(src string) string {
	hash := md5.Sum([]byte(src))
	dest := fmt.Sprintf("%x", hash)
	return dest
}

func (shortener *Shortener) ShortURLSHA1(src string) string {
	hash := sha1.Sum([]byte(src))
	dest := fmt.Sprintf("%x", hash)
	return dest
}

func (shortener *Shortener) ShortURLSHA256(src string) string {
	hash := sha256.Sum256([]byte(src))
	dest := fmt.Sprintf("%x", hash)
	return dest
}

func (shortener *Shortener) Generate(src string) (string, error) {
	var dest string
	if src == "" {
		return "", nil
	} else {
		switch shortener.Algorithm {
		case SHA1:
			dest = shortener.ShortURLSHA1(src)
		case CRC32:
			dest = shortener.ShortURLCRC32(src)
		case MD5:
			dest = shortener.ShortURLMD5(src)
		case SHA256:
			dest = shortener.ShortURLSHA256(src)
		}
	}
	return dest, nil
}
