package shortenerConfig

import shortener "url-shortener/pkg/shorter"

type ShortenerConfig struct {
	ID         string              `json:"id"`
	AccountId  string              `json:"account-id"`
	ExpireTime int                 `json:"expire-time"`
	Algorithm  shortener.Algorithm `json:"algorithm"`
}
