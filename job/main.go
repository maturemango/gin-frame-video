package job

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gin-frame",
	Short: "background job",
	Long: "gin-frame background job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("job start")
		job := &BackgroundJob{}
		go gracefulShutdown(job.Stop)
		job.Run()
	},
}

func main() {
	// rootCmd.AddCommand()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	if flag := rootCmd.Flags().Lookup("help"); flag != nil && flag.Changed == true {
		os.Exit(0)
	}
}

func gracefulShutdown(f func()) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit
	log.Println("graceful Shutdown Job ")
	f()
	log.Println("Job is done")
}