package fileUtils

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressDirToZip(sourceDir string, destZipPath string) (bool, error) {
	// 创建 ZIP 文件
	zipFile, err := os.Create(destZipPath)
	if err != nil {
		return false, fmt.Errorf("failed to create ZIP file: %v", err)
	}
	defer zipFile.Close()

	// 创建 ZIP 编写器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历文件夹中的文件并将它们添加到 ZIP 文件中
	err = filepath.Walk(sourceDir, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access file %s: %v", filePath, err)
		}

		if fileInfo.IsDir() {
			return nil // 跳过目录
		}

		// 相对于给定目录的路径
		relativePath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return fmt.Errorf("failed to get relative path for file %s: %v", filePath, err)
		}

		// 将文件添加到 ZIP 归档文件中
		zipFileEntry, err := zipWriter.Create(relativePath)
		if err != nil {
			return fmt.Errorf("failed to create ZIP file entry for file %s: %v", filePath, err)
		}

		fileToCompress, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", filePath, err)
		}
		defer fileToCompress.Close()

		// 将文件内容写入 ZIP 归档文件中
		_, err = io.Copy(zipFileEntry, fileToCompress)
		if err != nil {
			return fmt.Errorf("failed to compress file %s: %v", filePath, err)
		}

		return nil
	})

	if err != nil {
		return false, fmt.Errorf("failed to compress directory to ZIP: %v", err)
	}

	fmt.Printf("Successfully compressed directory %s to %s\n", sourceDir, destZipPath)
	return true, nil
}

func FilesToZip(files []string, zipName string) (bool, error) {
	// 检查 zip 文件名是否存在
	_, err := os.Stat(zipName)
	if err == nil {
		return false, errors.New("zip 文件已存在")
	}

	// 创建 zip 文件
	zipFile, err := os.Create(zipName)
	if err != nil {
		return false, err
	}
	defer zipFile.Close()

	// 创建 zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, path := range files {
		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println("获取文件信息失败：", err)
			return false, err
		}

		if !fileInfo.IsDir() {
			// 打开要压缩的文件
			fileToZip, err := os.Open(path)
			if err != nil {
				return false, err
			}
			defer fileToZip.Close()

			// 创建 ZIP 归档文件中的文件条目
			zipFileEntry, err := zipWriter.Create(filepath.Base(path))
			if err != nil {
				return false, err
			}

			// 将文件内容写入 ZIP 归档文件中的文件条目
			_, err = io.Copy(zipFileEntry, fileToZip)
			if err != nil {
				fmt.Println(err)
				return false, err
			}
			fmt.Println("成功创建 archive.zip")
			return true, nil
		}
	}

	// 如果没有找到要压缩的文件，则返回错误消息
	return false, errors.New("不支持压缩文件夹或没有找到要压缩的文件")
}
