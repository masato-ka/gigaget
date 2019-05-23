package config

import (
	"gigaget/util/cache"
	"testing"
)

func TestConfigManagerSaveNormal(t *testing.T) {
	target := cache.CacheManager{}
	target.Append("line1")
	target.Append("line2")

	target.Save()
}

func TestConfigManagerLoadNormal(t *testing.T) {
	target := cache.CacheManager{}
	target.Load()
	target.Append("line3")
	target.Save()
}
