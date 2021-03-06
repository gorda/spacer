package formatter

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"spacer/scanner"
	"sync"
	"time"
)

func Start(pathname string) {
	fmt.Println("start formatter...")
	now := time.Now()
	files := scanner.ScanAllFile(pathname, []string{".md"})
	wg := sync.WaitGroup{}
	for _, v := range files {
		wg.Add(1)
		go format(v, &wg)
	}
	wg.Wait()
	cost := time.Now().Sub(now).Milliseconds()
	fmt.Printf("done! affect %v files, cost %v ms\n", len(files), cost)
}

func format(filename string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		if r := recover(); r != nil {
			fmt.Printf("format file fail. filename: %v, panic: %v", filename, r)
			return
		}
	}()
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file fail. filename: %v, err: %v", filename, err.Error())
		return
	}

	new := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		new.WriteString(Handling(sc.Text()))
	}
	err = ioutil.WriteFile(filename, new.Bytes(), 0644)
	if err != nil {
		fmt.Printf("write file fail. filename: %v, err: %v", filename, err.Error())
		return
	}
}
