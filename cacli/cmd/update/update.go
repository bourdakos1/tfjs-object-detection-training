package update

import (
	"fmt"
	"os/exec"

	"github.com/cloud-annotations/training/cacli/e"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println(buildDownloadURL())

	out, err := exec.Command("which", "cacli").Output()
	if err != nil {
		e.Exit(err)
	}

	output := string(out)
	fmt.Println(output)

	// os.Rename("/usr/local/bin/cacli", "/usr/local/bin/cacli-old")
	// os.Remove("/usr/local/bin/cacli-old")
	// os.Rename("./cacli", "/usr/local/bin/cacli")
}
