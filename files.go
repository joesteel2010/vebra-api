package api

import (
	"os"
	"io"
	"net/http"
	"github.com/nfnt/resize"
	"github.com/disintegration/imaging"
	"sync"
	"path/filepath"
	"strconv"
	"image"
)

func CreatePropertyFilesManifest(property *Property) *propertyFilesManifest {
	urls := make([]string, len(property.Files))
	for _, file := range property.Files {
		urls = append(urls, file.Url)
	}
	return &propertyFilesManifest{
		strconv.Itoa(int(property.ID)),
		urls,
	}
}

func CreatePropertyFilesManifestFromChangedPropertySummaries(property *ChangedFilesSummaries) []*propertyFilesManifest {
	urls := make(map[string][]string, len(property.Files))

	for _, file := range property.Files {
		if file.Deleted {
			continue
		}
		propertyID := strconv.Itoa(file.FilePropId)
		urls[propertyID] = append(urls[propertyID], file.Url)
	}

	manifests := make([]*propertyFilesManifest, len(urls))
	for key := range urls {
		manifests = append(manifests, &propertyFilesManifest{
			key,
			urls[key],
		})
	}
	return manifests
}

type propertyFilesManifest struct {
	propertyID string
	urls []string
}

type fileDownloader struct {
	wg                    sync.WaitGroup
	propertyChan          chan *propertyFilesManifest
	interpolationFunction resize.InterpolationFunction
	outputDirectory       string
	thumbNailFilePrefix   string
	maxHeight             uint
	maxWidth              uint
}

func FileDownloader(chanSize int) *fileDownloader {
	fileDownloader := &fileDownloader{
		sync.WaitGroup{},
		make(chan *propertyFilesManifest, chanSize),
		resize.Lanczos3,
		"files",
		"tn",
		355,
		0,
	}
	go fileDownloader.listenForImages()
	return fileDownloader
}

func (fileDownloader *fileDownloader) SetInterpolationFunction(interpolationFunction resize.InterpolationFunction) {
	fileDownloader.interpolationFunction = interpolationFunction
}

func (fileDownloader *fileDownloader) SetOutputDirectory(outputDirectory string) {
	fileDownloader.outputDirectory = outputDirectory
}

func (fileDownloader *fileDownloader) SetMaxHeight(maxHeight uint) {
	fileDownloader.maxHeight = maxHeight
}

func (fileDownloader *fileDownloader) SetMaxWidth(maxWidth uint) {
	fileDownloader.maxWidth = maxWidth
}

func (fileDownloader *fileDownloader) Download(property *propertyFilesManifest) {
	fileDownloader.propertyChan <- property
	fileDownloader.wg.Add(1)
}

func (fileDownloader *fileDownloader) Wait() {
	fileDownloader.wg.Wait()
}

func (fileDownloader *fileDownloader) listenForImages() {
	for {
		property := <-fileDownloader.propertyChan
		for _, file := range property.urls {
			dir := fileDownloader.outputDirectory + "/" + property.propertyID + "/"
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.MkdirAll(dir, os.ModePerm)
			}
			rawFileName := filepath.Base(file)
			outFile := dir + rawFileName
			fileDownloader.downloadFile(file, outFile)
			thumbNameFile := fileDownloader.thumbNailFilePrefix + rawFileName
			fileDownloader.convertImage(outFile, dir+thumbNameFile)
		}
		fileDownloader.wg.Done()
	}
}

func (fileDownloader *fileDownloader) downloadFile(source string, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func (fileDownloader *fileDownloader) convertImage(source string, dest string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}

	imgMetaData, _, err := image.DecodeConfig(file)

	if ! (fileDownloader.maxHeight < 1) && imgMetaData.Height < int(fileDownloader.maxHeight) {
		return Copy(source, dest)
	}

	if ! (fileDownloader.maxWidth < 1) && imgMetaData.Height < int(fileDownloader.maxWidth) {
		return Copy(source, dest)
	}

	// decode jpeg into image.Image
	img, err := imaging.Decode(file)
	if err != nil {
		return err
	}
	file.Close()

	// resize to specified height/width using given interpolationFunction resampling
	// and preserve aspect ratio
	m := resize.Resize(fileDownloader.maxWidth, fileDownloader.maxHeight, img, fileDownloader.interpolationFunction)

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	return imaging.Save(m, dest)
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}


type fileDeleter struct {
	wg                    sync.WaitGroup
	propertyChan          chan *Property
	outputDirectory       string
	thumbNailFilePrefix   string
}

func FileDeleter(chanSize int) *fileDeleter {
	fileDeleter := &fileDeleter{
		sync.WaitGroup{},
		make(chan *Property, chanSize),
		"files",
		"tn",
	}
	go fileDeleter.listenForImages()
	return fileDeleter
}

func (fileDeleter *fileDeleter) SetOutputDirectory(outputDirectory string) {
	fileDeleter.outputDirectory = outputDirectory
}

func (fileDeleter *fileDeleter) Delete(property *Property) {
	fileDeleter.propertyChan <- property
	fileDeleter.wg.Add(1)
}

func (fileDeleter *fileDeleter) Wait() {
	fileDeleter.wg.Wait()
}

func (fileDeleter *fileDeleter) listenForImages() {
	for {
		property := <- fileDeleter.propertyChan
		for _, file := range property.Files {
			dir := fileDeleter.outputDirectory + "/" + strconv.Itoa(int(property.ID)) + "/"
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.MkdirAll(dir, os.ModePerm)
			}
			rawFileName := filepath.Base(file.Url)
			outFile := dir + rawFileName
			if _, err := os.Stat(outFile); err == nil {
				os.Remove(outFile)
			}
			thumbNameFile := fileDeleter.thumbNailFilePrefix + rawFileName
			if _, err := os.Stat(thumbNameFile); err == nil {
				os.Remove(thumbNameFile)
			}
		}
		fileDeleter.wg.Done()
	}
}