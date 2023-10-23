package handlers

import (
	"errors"
	"fmt"
	"time"
	cacheManagement "url-shortener/pkg/cache"
	shortener "url-shortener/pkg/shorter"
)

type Input struct {
	Url    string
	Domain string
	IP     string
	Time   time.Time
}

const (
	SRC  string = "SRC"
	DEST string = "DEST"
)

func CreateShortUrl(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL is empty")
	}

	cache := cacheManagement.Instance()
	shortUrlCached, err := cache.Get(SRC + url)
	if err != nil {
		return "", fmt.Errorf("Failed to get short URL from cache: %v", err)
	}
	if shortUrlCached != nil {
		return *shortUrlCached, nil
	}

	return generateShortURL(url, cache)
}

func generateShortURL(url string, cache cacheManagement.Management) (string, error) {
	shortenerGenerator := shortener.Shortener{
		Algorithm: shortener.CRC32,
	}

	for {
		dest, err := shortenerGenerator.Generate(url)
		if err != nil {
			return "", fmt.Errorf("failed to generate short URL: %v", err)
		}

		shortUrl, err := cache.Get(DEST + dest)
		if err != nil {
			return "", fmt.Errorf("failed to get short URL from cache: %v", err)
		}

		if shortUrl != nil {
			continue
		}

		err = cache.Sets(map[string]string{
			SRC + url:   DEST + dest,
			DEST + dest: SRC + url,
		}, 0)
		if err != nil {
			return "", fmt.Errorf("Failed to set short URL in cache: %v", err)
		}

		return dest, nil
	}
}

func View() string {
	return ""
}
