package downloader

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type job struct {
	url      string
	filename string
}

type Downloader struct {
	threadNum int
	jobs      chan *job
	waitGroup sync.WaitGroup
}

func New(threadNum int) *Downloader {
	downloader := &Downloader{threadNum, make(chan *job, 0), sync.WaitGroup{}}
	for i := 0; i < threadNum; i++ {
		downloader.waitGroup.Add(1)
		go func() {
			defer downloader.waitGroup.Done()
			for j := range downloader.jobs {
				if err := download(j.url, j.filename); err != nil {
					log.Printf("Failed to download %s: %v", j.url, err)
				}
			}
		}()
	}
	return downloader
}

func (d *Downloader) AddJob(url, filename string) {
	d.jobs <- &job{url, filename}
}

func (d *Downloader) Close() {
	close(d.jobs)
	d.waitGroup.Wait()
}

func download(url, filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}
	return nil
}
