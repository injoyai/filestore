package main

import (
	"fmt"
	"github.com/injoyai/conv"
	"github.com/injoyai/conv/cfg/v2"
	"github.com/injoyai/goutil/frame/in/v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	dir := cfg.GetString("dir", "./resource/")
	port := cfg.GetInt("port", 8080)
	enableUpload := cfg.GetBool("enableUpload", false)

	http.ListenAndServe(fmt.Sprintf(":%d", port), in.Recover(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filename := filepath.Join(dir, r.URL.Path)
		fullDir, name := filepath.Split(filename)

		switch r.Method {
		case http.MethodGet:
			f, err := os.Open(filename)
			if err != nil {
				in.Return(http.StatusForbidden, nil)
			}
			defer f.Close()

			info, err := f.Stat()
			if err != nil {
				in.Return(http.StatusForbidden, nil)
			}

			w.Header().Set("Content-Disposition", "attachment; filename="+name)
			w.Header().Set("Content-Length", conv.String(info.Size()))
			io.Copy(w, f)

		case http.MethodPost:
			if !enableUpload {
				in.Return(http.StatusForbidden, nil)
			}

			uploadFile, _, err := r.FormFile("file")
			in.CheckErr(err)
			defer uploadFile.Close()

			err = os.MkdirAll(fullDir, os.ModePerm)
			in.CheckErr(err)

			localFile, err := os.Create(filename + ".uploading")
			in.CheckErr(err)
			defer localFile.Close()

			_, err = io.Copy(localFile, uploadFile)
			in.CheckErr(err)

			err = os.Rename(filename+".uploading", filename)
			in.CheckErr(err)

			in.Succ(nil)

		default:
			in.Return(http.StatusForbidden, nil)

		}

	})))

}
