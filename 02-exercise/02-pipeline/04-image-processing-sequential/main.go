package main

import (
	"errors"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

type walkResult struct {
	Path string
	Err  error
}

type processImageResult struct {
	SrcImagePath   string
	ThumbnailImage *image.NRGBA
	Err            error
}

type saveResult struct {
	ok  bool
	err error
}

// Image processing - sequential
// Input - directory with images.
// output - thumbnail images
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 2 {
		log.Fatal("need to send directory path of images")
	}
	start := time.Now()
	//FIXME
	ch1 := saveThumbnail(processImage(walkFiles(os.Args[1])))

	for c := range ch1 {
		if !c.ok {
			fmt.Println("Error occured: ", c.err)
		}
	}
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

// walfiles - take diretory path as input
// does the file walk
// generates thumbnail images
// saves the image to thumbnail directory.
func walkFiles(root string) <-chan *walkResult {
	out := make(chan *walkResult) // Channel for path variable

	go func() {
		defer close(out)
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			// filter out error
			if err != nil {
				out <- &walkResult{
					Err: err,
				}
			}

			// check if it is file
			if !info.Mode().IsRegular() {
				out <- &walkResult{
					Err: errors.New(fmt.Sprintf("%s is not a file", path)),
				}
			}

			// check if it is image/jpeg
			contentType, _ := getFileContentType(path)
			if contentType != "image/jpeg" {
				out <- &walkResult{
					Err: errors.New(fmt.Sprintf("%s is not of contenttype image/jpeg", path)),
				}
			}

			out <- &walkResult{
				Path: path,
				Err:  nil,
			}

			return nil
		})

		if err != nil {
			out <- &walkResult{
				Err: err,
			}
		}
	}()
	return out
}

// processImage - takes image file as input
// return pointer to thumbnail image in memory.
func processImage(path <-chan *walkResult) <-chan *processImageResult {
	out := make(chan *processImageResult)
	scalefactor := 5
	var wg sync.WaitGroup
	for i := 1; i <= scalefactor; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range path {
				if p.Err != nil {
					out <- &processImageResult{
						Err: p.Err, //Propagating the errors
					}
				} else {
					srcImage, err := imaging.Open(p.Path)
					if err != nil {
						out <- &processImageResult{
							Err: err,
						}
					} else {
						// scale the image to 100px * 100px
						out <- &processImageResult{
							SrcImagePath:   p.Path,
							ThumbnailImage: imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos),
							Err:            nil,
						}
					}

				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)

	}()
	return out

}

// saveThumbnail - save the thumnail image to folder
func saveThumbnail(pir <-chan *processImageResult) <-chan *saveResult {
	out := make(chan *saveResult)

	go func() {
		defer close(out)
		for p := range pir {
			if p.Err != nil {
				out <- &saveResult{
					err: p.Err, //Propagating the errors
				}
			} else {
				filename := filepath.Base(p.SrcImagePath)
				dstImagePath := "thumbnail/" + filename
				err := imaging.Save(p.ThumbnailImage, dstImagePath)
				if err != nil {
					out <- &saveResult{
						ok:  false,
						err: err,
					}
				}
				fmt.Printf("%s -> %s\n", p.SrcImagePath, dstImagePath)
			}
		}
	}()

	return out
}

// getFileContentType - return content type and error status
func getFileContentType(file string) (string, error) {

	out, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
