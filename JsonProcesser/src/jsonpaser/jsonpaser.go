package main

import (
	"downloader"
	"fashionjson"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var outDir = flag.String("o", ".", "The folder where the processed data store")
var dlImg = flag.Bool("d", false, "If download image files while processing json")
var threadNum = flag.Int("tn", 80, "Number of threads for download image files")

func main() {
	flag.Parse()
	if _, err := os.Stat(*outDir); os.IsNotExist(err) {
		log.Fatal(err)
	}
	*outDir = filepath.ToSlash(*outDir)
	if strings.HasSuffix(*outDir, "/") {
		*outDir = (*outDir)[:len(*outDir)-1]
	}
	log.Println("Save processed data to ", *outDir)
	var dl *downloader.Downloader
	if *dlImg == true {
		dl = downloader.New(*threadNum)
		defer dl.Close()
	}
	for _, filename := range flag.Args() {
		fashionjson.Process(filename, *outDir, dl)
	}
}
