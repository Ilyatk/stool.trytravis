package ftp

import (
	"github.com/pkg/errors"
	"time"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"net/textproto"

	"cord.stool/updater"

	"github.com/jlaffaye/ftp"
)

// UploadToFTP upload all files from sourceDir recursive
// User, password and releative ftp path should be set in ftpUrl. 
// ftp://ftpuser:ftppass@ftp.protocol.local:21/cordtest/
func UploadToFTP(ftpUrl, sourceDir string) (cerr error) {
	fullSourceDir, cerr := filepath.Abs(sourceDir)
	if cerr != nil {
		return
	}

	u, err := url.Parse(ftpUrl)

	if err != nil {
		return errors.Wrap(err, "UploadToFTP: Invalid ftp url")
	}

	if !strings.EqualFold(u.Scheme, "ftp") {
		return errors.New("UploadToFTP: Invalid url scheme ")
	}

	conn, err := ftp.DialTimeout(u.Host, time.Second * 10) // TODO add timeout to config
	if err != nil {
		return errors.Wrap(err, "UploadToFTP: Failed to coonect ftp")
	}

	if u.User != nil {
		pass, _ := u.User.Password()
		err = conn.Login(u.User.Username(), pass)
		if err != nil {
			return errors.Wrap(err, "UploadToFTP: Cannnot login to ftp server")
		}
	}

	defer conn.Quit()
	defer conn.Logout()

	ftpRoot := u.Path
	conn.RemoveDirRecur(ftpRoot)

	stopCh := make(chan struct{})
	defer func() {
		select {
		case stopCh <- struct{}{}:
		default:
		}

		close(stopCh)
	}()

	f, e := updater.EnumFilesRecursive(fullSourceDir, stopCh)
	var relativePath string
	var file *os.File

	for path := range f {
		relativePath, cerr = filepath.Rel(fullSourceDir, path)
		if cerr != nil {
			return
		}

		ftpPath := filepath.Join(ftpRoot, relativePath)

		ftpDir,_ := filepath.Split(ftpPath)
		cerr = mkdirRecursive(conn, ftpDir)
		
		if cerr != nil {
			return
		}

		file, cerr = os.Open(path)
		if cerr != nil {
			return
		}

		defer file.Close()
		
		cerr = conn.Stor(filepath.ToSlash(ftpPath), file)

		if cerr != nil {
			return
		}
	}

	cerr = <-e
	if cerr != nil {
		return
	}

	return nil
}

func mkdirRecursive(conn *ftp.ServerConn, dir string) (err error) {
	dir = filepath.ToSlash(dir)
	dirs := strings.Split(dir, "/")

	tmpDir := ""

	for _, d := range dirs {
		if d == "" {
			continue;
		}

		tmpDir += d + "/"

		terr := conn.MakeDir(tmpDir)
		
		if terr != nil {
			code := terr.(*textproto.Error).Code
			if code == 550 {
				continue;
			}

			return terr
		}
	}

	return nil
}
