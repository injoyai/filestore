package main

import (
	"fmt"
	"github.com/injoyai/conv"
	"github.com/injoyai/conv/cfg/v2"
	"github.com/injoyai/goutil/frame/in/v3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	cfg.Init(
		cfg.WithYaml("./config/config.yaml"),
		cfg.WithFlag(
			&cfg.Flag{Name: "port", Default: "8080", Usage: "服务的端口"},
			&cfg.Flag{Name: "dir", Default: "./resource/", Usage: "存储的目录"},
			&cfg.Flag{Name: "enableDownload", Default: "true", Usage: "启用下载"},
			&cfg.Flag{Name: "enableUpload", Default: "false", Usage: "启用上传"},
			&cfg.Flag{Name: "enableDelete", Default: "false", Usage: "启用删除"},
		),
	)
}

func main() {

	port := cfg.GetInt("port", 8080)
	dir := cfg.GetString("dir", "./resource/")
	enableDownload := cfg.GetBool("enableDownload", true)
	enableUpload := cfg.GetBool("enableUpload")
	enableDelete := cfg.GetBool("enableDelete")

	log.Println("Server listen on:", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), in.Recover(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filename := filepath.Join(dir, r.URL.Path)
		fullDir, name := filepath.Split(filename)

		switch r.Method {
		case http.MethodGet:
			if !enableDownload {
				in.Return(http.StatusForbidden, nil)
			}

			f, err := os.Open(filename)
			if err != nil {
				in.Return(http.StatusForbidden, nil)
			}
			defer f.Close()

			info, err := f.Stat()
			if err != nil {
				in.Return(http.StatusForbidden, nil)
			}

			if ls := r.URL.Query()["show"]; len(ls) == 0 {
				w.Header().Set("Content-Disposition", "attachment; filename="+name)
				w.Header().Set("Content-Length", conv.String(info.Size()))
			}
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

		case http.MethodDelete:
			if !enableDelete {
				in.Return(http.StatusForbidden, nil)
			}

			err := os.Remove(filename)
			in.CheckErr(err)

			in.Succ(nil)

		default:
			in.Return(http.StatusForbidden, nil)

		}

	})))

}
