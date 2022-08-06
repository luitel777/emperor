package main

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

// https://echo.labstack.com/cookbook/file-upload/

func upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create("data/books/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return err
}
