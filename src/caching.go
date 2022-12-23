package ralphred

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"time"
)


func requestAndCache(url string, cacheFile string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	cacheDir := filepath.Dir(cacheFile)
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.MkdirAll(cacheDir, 0700)
	}

	os.WriteFile(cacheFile, respDump, 0700)

	return resp, nil
}

func requestFromCache(cacheFile string) (*http.Response, error) {
	cacheData, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	cacheDataReader := bufio.NewReader(bytes.NewReader(cacheData))
	return http.ReadResponse(cacheDataReader, nil)
}

func getCacheFile(url string) (string, error) {
	cacheKey := hashString(sha1.New(), url)

	homeDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	cacheFile := filepath.Join(homeDir, "ralphred", cacheKey)
	return cacheFile, nil
}

func isCached(cacheFile string, ttl int) (bool, error) {
	cacheFileInfo, err := os.Stat(cacheFile)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	cacheAge := time.Now().UTC().Sub(cacheFileInfo.ModTime())
	return int(cacheAge.Seconds()) < ttl, nil
}

func cachedRequest(url string, ttl int) (*http.Response, error) {
	cacheFile, err := getCacheFile(url)
	if err != nil {
		return nil, err
	}

	cached, err := isCached(cacheFile, ttl)
	if err != nil {
		return nil, err
	}

	if cached {
		log.Printf("Loading %s from cache", url)
		return requestFromCache(cacheFile)
	} else {
		log.Printf("Making request for %s and caching it", url)
		return requestAndCache(url, cacheFile)
	}
}
