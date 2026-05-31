package exoprt_utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/goccy/go-json"
)

// ExportRedisDataModel describes one exported Redis key snapshot.
type ExportRedisDataModel struct {
	Key   string      `json:"key"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// ExportUtils collects and persists exported key data.
type ExportUtils struct {
	data []*ExportRedisDataModel
}

// LoadKeysData fetches values for given keys and appends them to export buffer.
//
// Supported Redis types: string, hash, list, set, zset.
// Unsupported types are skipped.
func (e *ExportUtils) LoadKeysData(client *redis.Client, keys []string) error {
	for _, key := range keys {
		// Detect key type first, then use corresponding fetch command.
		keyType, err := client.Type(key).Result()
		if err != nil {
			continue
		}
		var value interface{}
		switch keyType {
		case "string":
			value, err = client.Get(key).Result()
		case "hash":
			value, err = client.HGetAll(key).Result()
		case "list":
			value, err = client.LRange(key, 0, -1).Result()
		case "set":
			value, err = client.SMembers(key).Result()
		case "zset":
			value, err = client.ZRangeWithScores(key, 0, -1).Result()
		default:
			continue
		}

		if err != nil {
			return fmt.Errorf("failed to fetch value for key %s: %v", key, err)
		}

		// Save normalized snapshot record.
		e.data = append(
			e.data, &ExportRedisDataModel{
				Key:   key,
				Type:  keyType,
				Value: value,
			},
		)
	}
	return nil
}

// SaveFile writes current export buffer to a formatted JSON file under var/export.
// It creates target directory recursively when needed.
func (e *ExportUtils) SaveFile() (string, error) {
	baseDir := "./var/export"
	// Create directory tree to avoid parent-directory missing errors.
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", err
	}
	// Compose output filename with timestamp.
	filePath := baseDir + "/export_" + time.Now().Format("2006-01-02 15:04:05") + ".json"

	// Create and write JSON output.
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(e.data); err != nil {
		return "", fmt.Errorf("failed to encode data to JSON: %v", err)
	}

	// The optional zip-compression path is intentionally kept disabled.
	//oldFilePath := filePath
	//// Compress to zip file.
	//err = CompressFile(filePath, filePath+".zip")
	//if err != nil {
	//	return "", fmt.Errorf("failed to compress file: %v", err)
	//}
	//filePath = filePath + ".zip"
	//// Remove original file.
	//err = os.Remove(oldFilePath)
	//if err != nil {
	//	return "", fmt.Errorf("failed to remove original file: %v", err)
	//}

	return filePath, nil
}

// CompressFile compresses a source file into a destination zip file.
func CompressFile(src, dst string) error {
	// Open source file.
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	// Create destination file.
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	// Create ZIP writer on destination file.
	zipWriter := zip.NewWriter(dstFile)
	defer zipWriter.Close()

	// Add source file entry into ZIP archive.
	zipFile, err := zipWriter.Create(src)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	// Stream source content into zip entry.
	_, err = io.Copy(zipFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file to zip: %v", err)
	}
	return nil
}
