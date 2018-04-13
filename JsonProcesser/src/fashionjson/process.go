package fashionjson

import (
	"downloader"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type TrainInfo struct {
	Url         string
	DateCreated string
	Version     string
	Description string
	Year        string
}

type ImageInfo struct {
	Url     string
	ImageId string
}

type Annotation struct {
	LabelId []string
	ImageId string
}

type License struct {
	Url  string
	Name string
	Id   string
}

func Process(jsonfilename, outDir string, downloader *downloader.Downloader) {
	if !strings.HasSuffix(jsonfilename, ".json") {
		log.Fatalln(jsonfilename, "is not a json file!")
	}
	jsonfile, err := os.Open(jsonfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonfile.Close()

	pureName := regexp.MustCompile("^.*/").ReplaceAllString(filepath.ToSlash(jsonfilename), "")
	pureName = strings.Replace(pureName, ".json", "", 1)

	dec := json.NewDecoder(jsonfile)
	passToken(dec)
	for dec.More() {
		switch objName := getObjName(dec); objName {
		case "info":
			var info TrainInfo
			if err := dec.Decode(&info); err != nil {
				log.Fatal(err)
			}
			log.Printf("%T: %+v\n", info, info)
		case "images":
			imgFolder := filepath.FromSlash(fmt.Sprintf("%s/%s", outDir, pureName))
			if _, err := os.Stat(imgFolder); os.IsNotExist(err) {
				if err = os.Mkdir(imgFolder, 0755); err != nil {
					log.Fatal(err)
				}
			}
			passToken(dec)
			var info ImageInfo
			for dec.More() {
				if err := dec.Decode(&info); err != nil {
					log.Println(err)
				} else if downloader != nil {
					filename := fmt.Sprintf("%s/%s.png", imgFolder, info.ImageId)
					downloader.AddJob(info.Url, filepath.FromSlash(filename))
				}
			}
			passToken(dec)
		case "annotations":
			passToken(dec)
			csvfile, err := os.Create(fmt.Sprintf("%s/%s_labels.csv", outDir, pureName))
			if err != nil {
				log.Fatalln(err)
			}
			defer csvfile.Close()
			csvWriter := csv.NewWriter(csvfile)
			csvWriter.Write([]string{"filename", "labels"})
			var info Annotation
			for dec.More() {
				if err := dec.Decode(&info); err != nil {
					log.Fatal(err)
				}
				imgName := fmt.Sprintf("%s.png", info.ImageId)
				imgLabels := fmt.Sprintf("%v", info.LabelId)
				csvWriter.Write([]string{imgName, strings.Trim(imgLabels, "[]")})
			}
			csvWriter.Flush()
			passToken(dec)
		case "license":
			var info License
			if err := dec.Decode(&info); err != nil {
				log.Fatal(err)
			}
			log.Printf("%T: %+v\n", info, info)

		default:
			log.Fatal("Unknown obj type: ", objName)
		}
	}
}

func passToken(dec *json.Decoder) {
	if t, err := dec.Token(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Pass Token: ", t)
	}
	return
}

func getObjName(dec *json.Decoder) string {
	if t, err := dec.Token(); err == nil {
		switch t.(type) {
		case string:
			log.Println(t)
			return fmt.Sprintf("%v", t)
		default:
			log.Fatalf("Expect string got %T:%v", t, t)
		}
	} else {
		log.Fatal(err)
	}
	return ""
}
