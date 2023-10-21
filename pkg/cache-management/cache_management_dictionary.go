package cacheManagement

import (
	"sync"
)

// DictionaryCacheManager is an in-memory dictionary cache implementation.
type DictionaryCacheManager struct {
	cache map[string]string
	mu    sync.RWMutex
}

// Set is a method of the DictionaryCacheManager to set a key-value pair in the in-memory dictionary cache.
func (client *DictionaryCacheManager) Set(key string, value interface{}, expire int) error {
	client.mu.Lock()
	defer client.mu.Unlock()
	client.cache[key] = value.(string)
	return nil
}

// Get is a method of the DictionaryCacheManager to retrieve a value from the in-memory dictionary cache.
func (client *DictionaryCacheManager) Get(key string) (*string, error) {
	client.mu.RLock()
	defer client.mu.RUnlock()
	value, found := client.cache[key]
	if !found {
		return nil, nil
	}
	return &value, nil
}

// Delete is a method of the DictionaryCacheManager to delete a key from the in-memory dictionary cache.
func (client *DictionaryCacheManager) Delete(key string) error {
	client.mu.Lock()
	defer client.mu.Unlock()
	delete(client.cache, key)
	return nil
}

func (client *DictionaryCacheManager) SetAndAdd(key, value string) error {
	// TODO implement it
	return nil
}

func (client *DictionaryCacheManager) Sets(values map[string]string, expire int) error {
	//TODO implement me
	panic("implement me")
}

func (client *DictionaryCacheManager) SetString(key string, value string, expire int) error {
	//TODO implement me
	panic("implement me")
}
