package exoprt_utils

import (
	"archive/zip"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/goccy/go-json"
	"io"
	"os"
	"time"
)

type ExportRedisDataModel struct {
	Key   string      `json:"key"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type ExportUtils struct {
	data []*ExportRedisDataModel
}

func (e *ExportUtils) LoadKeysData(client *redis.Client, keys []string) error {
	for _, key := range keys {
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

func (e *ExportUtils) SaveFile() (string, error) {
	baseDir := "./var/export"
	// 判断./var/export 目录是否存在，没创建则创建
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.Mkdir("./var/export", 0755)
		if err != nil {
			return "", err
		}
	}
	// 将数据保存到JSON文件
	filePath := baseDir + "/export_" + time.Now().Format("2006-01-02 15:04:05") + ".json"

	// 保存到JSON文件
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
	//oldFilePath := filePath
	//// 压缩为zip文件
	//err = CompressFile(filePath, filePath+".zip")
	//if err != nil {
	//	return "", fmt.Errorf("failed to compress file: %v", err)
	//}
	//filePath = filePath + ".zip"
	//// 删除原始文件
	//err = os.Remove(oldFilePath)
	//if err != nil {
	//	return "", fmt.Errorf("failed to remove original file: %v", err)
	//}

	return filePath, nil
}

func CompressFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	// 创建zip.Writer
	zipWriter := zip.NewWriter(dstFile)
	defer zipWriter.Close()

	// 将源文件添加到zip中
	zipFile, err := zipWriter.Create(src)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	_, err = io.Copy(zipFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file to zip: %v", err)
	}
	return nil
}
