package update

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/cloud-annotations/training/cacli/version"
)

func buildDownloadURL() string {
	baseDownloadURL := "https://github.com/cloud-annotations/training/releases/download"
	version := version.BuildVersion()
	return fmt.Sprintf("%s/v%s/cacli_%s", baseDownloadURL, version, determinePlatform())
}

func determinePlatform() string {
	arch := runtime.GOARCH

	switch runtime.GOOS {
	case "windows":
		if arch == "386" {
			return "windows_i386.exe"
		}
		return "windows_x86_64.exe"

	case "linux":
		switch arch {
		case "arm64":
			return "linux_arm64"
		case "arm":
			return "linux_armv6"
		case "386":
			return "linux_i386"
		default:
			return "linux_x86_64"
		}
	case "darwin":
		if arch == "386" {
			return "darwin_i386"
		}
		return "darwin_x86_64"
	}
	return "unknown"
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
