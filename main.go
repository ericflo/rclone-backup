package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ericflo/rclone-backup/backup"
)

var source string
var days int
var period int

func main() {
	flag.StringVar(&source, "source", "", "source directory to back up")
	flag.IntVar(&days, "days", 14, "number of days to keep each backup")
	flag.IntVar(&period, "period", 60*5, "seconds between checking whether to backup")
	flag.Parse()

	if source == "" {
		flag.Usage()
		log.Fatalln("Must provide at least `source` flags to rclone-backup.")
	}

	target := os.Getenv("BACKUP_TARGET")
	if target == "" {
		log.Fatalln("Must provide a backup target via the BACKUP_TARGET environment variable.")
	}

	for {
		if err := backup.RunBackup(source, target); err != nil {
			log.Println("Got an error during a backup run:", err)
		}
		if err := backup.RemoveExpiredBackups(target, days); err != nil {
			log.Println("Got an error while removing expired backups:", err)
		}
		time.Sleep(time.Second * time.Duration(period))
	}
}
