package fileHandle

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadFile(path string, url string) (err error) {
	p := path + "/tempFile.txt"
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		newpath := filepath.Join(".", path)
		if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
			return err
		}
		out, err := os.Create(newpath + "/tempFile.txt")
		if err != nil {
			return err
		}
		defer out.Close()

		// Get the data
		timeout := time.Duration(5) * time.Second
		transport := &http.Transport{
			ResponseHeaderTimeout: timeout,
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, timeout)
			},
			DisableKeepAlives: true,
		}
		client := &http.Client{
			Transport: transport,
		}
		resp, err := client.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Check server response
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status: %s", resp.Status)
		}

		// Writer the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
