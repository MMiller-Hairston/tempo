package tempo

import (
	"fmt"
	"os"
	"strings"
)

type CleanupCommand struct{}

func (c *CleanupCommand) Help() string {
	return "Cleanup"
}

func (c *CleanupCommand) Run(args []string) int {
	t := getToday()
	d := formatDate(t)

	cnt := 0
	fis, _ := os.ReadDir(storage)
	for _, f := range fis {
		if f.IsDir() {
			continue
		} else {
			s := strings.Split(f.Name(), "_")

			if s[0] < d {
				err := os.Remove(storage + "/" + s[0] + file_ext)
				if err != nil {
					fmt.Printf("Unable to delete file %s: %v\n", d+file_ext, err)
					return 1
				}
				cnt += 1
			}
		}
	}
	fmt.Printf("%d files removed\n", cnt)
	return 0
}

func (c *CleanupCommand) Synopsis() string {
	return "Cleanup all previous caught tasks that occurred before today."
}
