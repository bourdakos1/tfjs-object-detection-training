package update

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	os.Rename("cacli", "cacli-old")
	os.Remove("./cacli-old")
	fileurl := "https://github.com/cloud-annotations/training/releases/download/v1.2.22/cacli_darwin_x86_64"
	if err := downloadFile("./cacli", fileurl); err != nil {
		panic(err)
	}
	os.Chmod("./cacli", 0700)
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
