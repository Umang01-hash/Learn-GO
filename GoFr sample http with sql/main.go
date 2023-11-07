package main

import (
	"fmt"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {
		rows, err := ctx.DB().Query("SELECT name from students;")
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		responseText := "Hello from Go Web Server!"

		for rows.Next() {
			var name string

			// Scan copies the columns from the matched row into the values pointed at by dest
			err := rows.Scan(&name)
			if err != nil {
				return nil, err
			}

			responseText += fmt.Sprintf("User: %s", name)
		}

		return responseText, err
	})

	app.Start()
}
