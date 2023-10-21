module url-shortener/pkg/cache-management

go 1.19

require (
	github.com/go-redis/redis/v8 v8.11.5
	url-shortener/pkg/config v0.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace url-shortener/pkg/config => ./../config