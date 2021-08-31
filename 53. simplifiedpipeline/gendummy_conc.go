package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

const totalFile = 3000
const contentLength = 1000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()
	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	//WaitGroup to control the workers
	wg := new(sync.WaitGroup)

	//allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		//dispatch N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				//listen to channel chanIn for incoming jobs
				for job := range chanIn {
					//do the jobs
					filePath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

					//construct the job's result, and send it to chanOut
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				//if chanIn closed, and the remaining jobs are finished, only then we mark the worker as completed
				wg.Done()
			}(workerIndex)

		}
	}()
	//wait until chanIn closed and then all workers are done, because right after that we need to close chanOut channel
	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	//pipeline1: job distribution
	chanFileIndex := generateFileIndexes()

	//pipeline2: the main logic (creating files)
	createFilesWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFilesWorker)

	//track and print output
	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating files %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
