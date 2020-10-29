package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	emails, err := generateEmails(cfg.GenerateEmailCount, cfg.DuplicatePercentage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	start := time.Now()
	dedupedEmails := dedupe(emails)
	elapsed := time.Since(start)
	log.Printf("Took %s to remove %v duplicate emails from %v emails", elapsed, len(emails)-len(dedupedEmails), len(emails))
}

func dedupe(emails []string) []string {
	keys := make(map[string]bool)
	deDupedEmails := []string{}
	for i, email := range emails {
		if !isValidEmail(email) {
			remove(emails, i)
			continue
		} else {
			if _, value := keys[email]; !value {
				keys[email] = true
				deDupedEmails = append(deDupedEmails, email)
			}
		}
	}
	return deDupedEmails
}

func isValidEmail(email string) bool {
	isValid, _ := regexp.MatchString("[^@]+@[^@]+\\.[^@]+", email)
	return isValid
}

// This is slow, but we want to preserve the order of the original list
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
