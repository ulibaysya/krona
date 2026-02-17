package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/ulibaysya/krona/internal/daemon"
)

func main() {
	if len(os.Args) == 1 {
		commandHelp([]string{})
		os.Exit(1)
	}
	switch os.Args[1] {
	case "version":
		fmt.Printf("%s: %s\n", Name, Version)
	case "run":
		commandRun(os.Args[2:])
	case "help":
		commandHelp(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "%s: unknown command: %v\nrun %v help\n", Name, os.Args[1], Name)
		os.Exit(1)
	}
}

func commandRun(args []string) {
	// 1. Provided path
	// 2. Environment var
	// 3. /etc/.../conf.yaml
	var configPath string
	if len(args) != 0 {
		configPath = args[0]
	} else if tmp, ok := os.LookupEnv("KRONA_CONF"); ok {
		configPath = tmp
	} else {
		configPath = path.Join("/etc", Name, "conf.yaml")
	}
	fmt.Println(configPath)

	daemon, err := daemon.New(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: error while initializing daemon: %v\n", Name, err)
		os.Exit(1)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := daemon.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "%v: error while running daemon: %v\n", Name, err)
			os.Exit(1)
		}
	}()
	<-signalChan

	time.Sleep(time.Second * 1)

	if err := daemon.Shutdown(); err != nil {
		fmt.Fprintf(os.Stderr, "%v: error while shutdown: %v\n", Name, err)
		os.Exit(1)
	}
	// defer func() {
	// }()
}

func commandHelp(args []string) {
	if len(args) == 0 {
		fmt.Printf("Usage:\n"+"\t%s\trun [config_path]\n"+"\t%s\tversion\n"+"\t%s\thelp [command]"+"\n", Name, Name, Name)
		return
	}

	switch args[0] {
	case "version":
		fmt.Println("TODO")
	case "run":
		fmt.Println("TODO")
	default:
		fmt.Fprintf(os.Stderr, "%s: unknown topic: %v\n", Name, args[0])
		os.Exit(1)
	}
}
