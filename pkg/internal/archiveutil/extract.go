package archiveutil

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"barglvojtech.net/govm/pkg/internal/fsutil"
	"barglvojtech.net/x/pkg/errutil"
)

const (
	archivePrefixDirectory = "go/"
)

func Extract(src, dst string) error {
	switch {
	case strings.HasSuffix(src, ".zip"):
		return extractZip(src, dst)
	case strings.HasSuffix(src, ".tar.gz"):
		return extractTarGz(src, dst)
	default:
		return fmt.Errorf("only .zip and .tar.gz extension are allowed, got %s", ext(src))
	}
}

func ext(src string) string {
	base := filepath.Base(src)
	extIndex := strings.IndexRune(src, '.')
	return base[extIndex+1:]
}

func extractZip(src, dst string) (err error) {
	r, e := zip.OpenReader(src)
	if errutil.AssignIfErr(&err, e, errutil.PrefixWithFormatted("could not open %s", src)) {
		return
	}

	for _, f := range r.File {
		path := filepath.Join(dst, f.Name)

		// for directory
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		if dir := filepath.Dir(path); !fsutil.Exists(dir) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		}

		// for file
		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer dstFile.Close()

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}
		defer fileInArchive.Close()

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}
	}

	return nil
}

func extractTarGz(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tar := tar.NewReader(gz)
	for {
		h, err := tar.Next()
		if err == io.EOF {
			break
		}

		path := filepath.Join(dst, strings.TrimPrefix(h.Name, archivePrefixDirectory))

		if h.FileInfo().IsDir() {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if dir := filepath.Dir(path); !fsutil.Exists(dir) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		}

		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.FileMode(h.Mode))
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err := io.CopyN(dstFile, tar, h.Size); err != nil {
			return err
		}
	}

	return nil
}
