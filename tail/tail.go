package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func tail(name string, num int) error {
	var count = 0
	var err error

	path, err := getFilePath(name)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_RDONLY, 0700)
	if err != nil {
		return err
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	if fi.Size() > 0 {
		count, err = lineCounter(f, -1)
		if err != nil {
			return err
		}
	}

	if count == 0 {
		return nil
	}

	start := 0

	if num <= 0 || num > count {
		start = count
	} else {
		start = count - num
	}

	//fmt.Println(count, start)
	f.Seek(0, os.SEEK_SET)
	lineCounter(f, start)

	return nil
}

func getFilePath(name string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", dir, name), nil
}

func lineCounter(r io.Reader, n int) (int, error) {
	count := 0
	buf := make([]byte, 32*1024)
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count + 1, nil

		case err != nil:
			return count, err
		}

		if n == -2 {
			fmt.Println(string(buf))
		}

		if n >= 0 && count >= n {
			//fmt.Println(count, n)
			splitStrFromTrunk(buf[:c], count-n)
			n = -2
		}
	}
}

func splitStrFromTrunk(buf []byte, n int) {

	n = bytes.Count(buf, []byte{'\n'})-n

	var newBuf []byte
	//var tmp []byte

	for _, c := range buf {
		if c == '\n' {
			//fmt.Println(string(tmp))
			n--
		}
		//tmp = append(tmp, c)

		if n <= 0 {
			newBuf = append(newBuf, c)
		}
	}

	fmt.Println(string(newBuf))
}
