package main

import (
  "os"
  "log"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  // define flags
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "save",
      Value:       "no",
      Usage:       "Should save to database (yes/no)",
    },
  }
  // define action
  app.Action = func(c *cli.Context) error {
    var args []string
    if c.NArg() > 0 {
      // Fetch arguments in a array
      args = c.Args()
    }
    // check the flag value
    if c.String("save") == "no" {
      log.Println("Skipping saving to the database")
    } else {
      log.Println("Saving to the database", args)
    }
    return nil
  }

  app.Run(os.Args)
}