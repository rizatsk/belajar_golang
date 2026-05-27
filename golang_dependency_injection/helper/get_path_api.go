package helper

import (
	"path/filepath"
	"runtime"
	"strings"
)

func GetPathApi() string {
	_, filename, _, _ := runtime.Caller(1)
	path := filepath.Dir(filename)
	cleanPath := filepath.ToSlash(path)

	// Cari posisi "/api/"
	idx := strings.Index(cleanPath, "/api/")
	if idx == -1 {
		return ""
	}

	// Ambil substring mulai dari /api/
	return cleanPath[idx:]
}
