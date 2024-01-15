package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jgkawell/yarr/server"
	"github.com/jgkawell/yarr/storage"
)

var Version string = "0.0"
var GitHash string = "unknown"

var OptList = make([]string, 0)

func opt(envVar, defaultValue string) string {
	OptList = append(OptList, envVar)
	value := os.Getenv(envVar)
	if value != "" {
		return value
	}
	return defaultValue
}

func parseAuthfile(authfile io.Reader) (username, password string, err error) {
	scanner := bufio.NewScanner(authfile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return "", "", fmt.Errorf("wrong syntax (expected `username:password`)")
		}
		username = parts[0]
		password = parts[1]
		break
	}
	return username, password, nil
}

func main() {
	var (
		addr, db, authfile, auth string
		ver, open bool
		err error
	)

	flag.CommandLine.SetOutput(os.Stdout)

	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(out, "\nThe environmental variables, if present, will be used to provide\nthe default values for the params above:")
		fmt.Fprintln(out, " ", strings.Join(OptList, ", "))
	}

	flag.StringVar(&addr, "addr", opt("YARR_ADDR", "127.0.0.1:7070"), "address to run server on")
	flag.StringVar(&authfile, "auth-file", opt("YARR_AUTHFILE", ""), "`path` to a file containing username:password. Takes precedence over --auth (or YARR_AUTH)")
	flag.StringVar(&auth, "auth", opt("YARR_AUTH", ""), "string with username and password in the format `username:password`")
	flag.StringVar(&db, "db", opt("YARR_DB", ""), "database URI")
	flag.BoolVar(&ver, "version", false, "print application version")
	flag.BoolVar(&open, "open", false, "open the server in browser")
	flag.Parse()

	if ver {
		fmt.Printf("v%s (%s)\n", Version, GitHash)
		return
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)

	if db == "" {
		log.Fatal("database URI is required")
	}

	var username, password string
	if authfile != "" {
		f, err := os.Open(authfile)
		if err != nil {
			log.Fatal("Failed to open auth file: ", err)
		}
		defer f.Close()
		username, password, err = parseAuthfile(f)
		if err != nil {
			log.Fatal("Failed to parse auth file: ", err)
		}
	} else if auth != "" {
		username, password, err = parseAuthfile(strings.NewReader(auth))
		if err != nil {
			log.Fatal("Failed to parse auth literal: ", err)
		}
	}

	store, err := storage.New(db)
	if err != nil {
		log.Fatal("Failed to initialise database: ", err)
	}

	srv := server.NewServer(store, addr)

	if username != "" && password != "" {
		srv.Username = username
		srv.Password = password
	}

	log.Printf("starting server at %s", srv.GetAddr())
	srv.Start()
}
