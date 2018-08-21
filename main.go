package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/mashup/controllers"
	"github.com/zpatrick/mashup/mashup"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	app := cli.NewApp()
	app.Name = "mashup"
	app.Usage = "Mashup Generator"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "p, port",
			EnvVar: "MASHUP_PORT",
			Value:  80,
		},
		cli.StringFlag{
			Name:   "matrix-file",
			EnvVar: "MASHUP_MATRIX_FILE",
			Value:  "generator/matrix.json",
		},
	}

	app.Action = func(c *cli.Context) error {
		rand.Seed(time.Now().UnixNano())

		path := c.String("matrix-file")
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("Failed to load matrix file '%s': %s", path, err.Error())
		}

		var matrix mashup.Matrix
		if err := json.Unmarshal(bytes, &matrix); err != nil {
			return fmt.Errorf("Failed to unmarshal matrix: %s", err.Error())
		}

		generator := mashup.NewMatrixGenerator(matrix)
		rootController := controllers.NewRootController(generator)
		app := fireball.NewApp(rootController.Routes())
		http.Handle("/", app)

		fs := http.FileServer(http.Dir("static"))
		http.Handle("/static/", http.StripPrefix("/static", fs))

		addr := fmt.Sprintf("0.0.0.0:%d", c.Int("port"))
		log.Printf("[INFO] Listening on %s\n", addr)
		return http.ListenAndServe(addr, nil)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
