package backup

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const backupPrefix = "backup:/"

func makeFullTarget(target string, daysAgo int) string {
	now := time.Now().UTC().AddDate(0, 0, -daysAgo)
	trimmedTarget := strings.Trim(target, "/")
	return fmt.Sprintf("%s%s/%s", backupPrefix, trimmedTarget, now.Format("2006-01-02"))
}

func targetExists(fullTarget string, print bool) (bool, error) {
	if print {
		log.Println("Checking whether", strings.TrimPrefix(fullTarget, backupPrefix), "exists")
	}
	cmd := exec.Command("rclone", "lsf", fullTarget, "--log-level", "ERROR")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("Could not run rclone lsf: %v %s", err, string(stdoutStderr))
	}
	exists := strings.TrimSpace(string(stdoutStderr)) != ""
	return exists, nil
}

func backupToTarget(source, fullTarget string) error {
	dest := strings.TrimPrefix(fullTarget, backupPrefix)
	log.Println("Backing up from", source, "to", dest)
	cmd := exec.Command("rclone", "copy", source, fullTarget, "-v", "--local-no-check-updated", "--no-check-dest")
	stdoutStderr, err := cmd.CombinedOutput()
	if stdoutStderr != nil && len(stdoutStderr) > 0 {
		log.Println("Backup output:", string(stdoutStderr))
	}
	log.Println("Finished backing up to", dest)
	return err
}

// RunBackup runs a single backup job
func RunBackup(source, target string) error {
	if _, err := os.Stat(source); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Source directory does not exist")
		}
		return fmt.Errorf("Could not check source directory: %v", err)
	}
	fullTarget := makeFullTarget(target, 0)
	exists, err := targetExists(fullTarget, true)
	if err != nil {
		return fmt.Errorf("Could not check whether target exists: %v", err)
	}
	if exists {
		return nil
	}
	return backupToTarget(source, fullTarget)
}

// RemoveExpiredBackups removes any expired backups past the day window
func RemoveExpiredBackups(target string, days int) error {
	fullTarget := makeFullTarget(target, days+1)
	if exists, err := targetExists(fullTarget, false); err != nil {
		return err
	} else if !exists {
		return nil
	}
	dest := strings.TrimPrefix(fullTarget, backupPrefix)
	log.Println("Removing expired backup at:", dest)
	cmd := exec.Command("rclone", "purge", fullTarget, "-v")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Could not remove directory: %v", err)
	}
	if stdoutStderr != nil && len(stdoutStderr) > 0 {
		log.Println("Removal output:", string(stdoutStderr))
	}
	log.Println("Finished removing", dest)
	return nil
}
