package concurrent

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// 入口函数
func WalkDirectory(dirs ...string) string {
	if len(dirs) == 0 {
		dirs = []string{"."}
	}
	// 初始化变量，channel用于完成Size传递，WaitGroup用于等待调度
	fileSizeCh := make(chan int64, 1)
	wg := &sync.WaitGroup{}

	// 启动多个协程统计信息，取决于len(dirs)
	for _, dir := range dirs {
		wg.Add(1)
		// 并发统计目录信息
		go walkDir(dir, fileSizeCh, wg)
	}

	// 启动累计运算的协程
	go func(wg *sync.WaitGroup) {
		// 等待统计工作完成
		wg.Wait()
		close(fileSizeCh)
	}(wg)

	// range方式从fileSizeCh获取文件大小
	fileNumCh := make(chan int64, 1)
	fileTotalSizeCh := make(chan int64, 1)
	go func(fileSizeCh <-chan int64, fileNumCh, fileTotalSizeCh chan<- int64) {
		// 统计文件数和文件整体大小
		var fileNum, sizeTotal int64
		for fileSize := range fileSizeCh {
			fileNum++
			sizeTotal += fileSize
		}
		fileNumCh <- fileNum
		fileTotalSizeCh <- sizeTotal
	}(fileSizeCh, fileNumCh, fileTotalSizeCh)

	// 整理返回值
	result := fmt.Sprintf("%d files, %.2f MB\n", <-fileNumCh, float64(<-fileTotalSizeCh)/1e6)
	return result
}

// 遍历并统计某个特定目录的信息
// 核心功能
func walkDir(dir string, fileSizeCh chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()

	// 读取dir下的全部文件
	for _, file := range fileInfo(dir) {
		if file.IsDir() {
			subDir := filepath.Join(dir, file.Name())
			wg.Add(1)
			go walkDir(subDir, fileSizeCh, wg)
		} else {
			fileSizeCh <- file.Size()
		}
	}

	// 根据dir下的文件信息
	// 如果是目录，递归获取信息
	// 如果不是，就是文件，统计文件大小，放入channel

}

// 获取某个目录下文件信息
func fileInfo(dir string) []fs.FileInfo {
	// 读取目录的全部文件
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Println("WalkDir error: ", err)
		return []fs.FileInfo{}
	}

	// 获取文件的文件信息
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		if info, err := entry.Info(); err == nil {
			infos = append(infos, info)
		}
	}

	return infos
}
