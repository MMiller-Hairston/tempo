package tempo

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strings"
)

type ListCommand struct{}

func (c *ListCommand) Help() string {
	return "List your captured tasks"
}

func (c *ListCommand) Run(args []string) int {
	fis, _ := os.ReadDir(storage)
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		} else {
			err := filterLines(args, fi)
			if err != nil {
				fmt.Println(err)
				return 1
			}
		}
	}
	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List all tasks organized by date. Specify dates as additional arguments to further filter the list of returned tasks."
}

func readFileLines(file fs.DirEntry) error {
	f, err := os.Open(storage + "/" + file.Name())
	if err != nil {
		return fmt.Errorf("Unable to open file: %v", err)
	}

	defer f.Close()

	fp := strings.Split(f.Name(), "_")
	n := fp[0][strings.LastIndex(fp[0], "/")+1:]
	fmt.Printf("%s\n", n)

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("Error reading file: %v", err)

	}

	return nil
}

func getFileDate(f fs.DirEntry) string {
	fp := strings.Split(f.Name(), "_")
	return fp[0][strings.LastIndex(fp[0], "/")+1:]
}

func filterLines(args []string, fi fs.DirEntry) error {
	if len(args) > 0 {
		n := getFileDate(fi)
		if slices.Contains(args, n) {
			err := readFileLines(fi)
			if err != nil {
				return err
			}
		}
	} else {
		err := readFileLines(fi)
		if err != nil {
			return err
		}
	}

	return nil
}
