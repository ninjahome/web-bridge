package main

import (
	"encoding/json"
	"fmt"
	"github.com/hopwesley/fdlimit"
	"github.com/ninjahome/web-bridge/server"
	"github.com/ninjahome/web-bridge/util"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	PidFileName    = "pid"
	ConfigFIleName = "config.json"
)

var (
	param     = &startParam{}
	Version   string
	Commit    string
	BuildTime string
)

type startParam struct {
	version bool
	config  string
}

var rootCmd = &cobra.Command{
	Use: "web-bridge",

	Short: "web-bridge",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

func init() {
	flags := rootCmd.Flags()
	flags.BoolVarP(&param.version, "version",
		"v", false, "web-bridge -v")
	flags.StringVarP(&param.config, "conf",
		"c", ConfigFIleName, "web-bridge -c config.json")
}
func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func mainRun(_ *cobra.Command, _ []string) {

	if param.version {
		fmt.Println("\n==================================================")
		fmt.Printf("Version:\t%s\n", Version)
		fmt.Printf("Build:\t\t%s\n", BuildTime)
		fmt.Printf("Commit:\t\t%s\n", Commit)
		fmt.Println("==================================================")
		return
	}

	if err := fdlimit.MaxIt(); err != nil {
		panic(err)
	}

	cf := initConfig()
	var basisSrv = server.NewMainService(cf)
	go func() {
		basisSrv.Start()
	}()

	waitShutdownSignal()
}

func initConfig() *server.Conf {
	cf := new(server.Conf)

	bts, err := os.ReadFile(ConfigFIleName)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(bts, &cf); err != nil {
		panic(err)
	}
	util.SetLogLevel(cf.Log)
	fmt.Println(cf.String())
	return cf
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>service start at pid(%s)<<<<<<<<<<\n", pid)
	if err := os.WriteFile(PidFileName, []byte(pid), 0644); err != nil {
		fmt.Print("failed to write running pid", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)
	sig := <-sigCh
	fmt.Printf("\n>>>>>>>>>>service finished(%s)<<<<<<<<<<\n", sig)
}
