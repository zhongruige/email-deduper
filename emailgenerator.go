package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	defaultEmailCount  = 100000
	defaultDupePercent = 0.5
)

func generateEmails(emailCount int, dupePercentage float32) ([]string, error) {
	var err error
	if emailCount < 0 {
		return []string{}, errors.New("emailCount must be greater than 0")
	}
	if dupePercentage > 0.9 || dupePercentage < 0 {
		return []string{}, errors.New("dupePercentage must be a value between 0.0 and 0.9")
	}
	generateEmailCount := defaultEmailCount
	if emailCount > 0 {
		generateEmailCount = emailCount
	}
	dupeCount := 0
	if dupePercentage > 0 && dupePercentage < 1 {
		dupeCount = int(float32(generateEmailCount) * dupePercentage)
	} else {
		dupeCount = int(float32(generateEmailCount) * defaultDupePercent)
	}

	var emails []string
	dupeCounter := 0
	counter := 0
	for len(emails) < generateEmailCount {
		emails = append(emails, fmt.Sprint(counter, "_email@gomail.com"))
		for dupeCounter < dupeCount {
			emails = append(emails, fmt.Sprint(counter, "_email@gomail.com"))
			dupeCounter++
		}
		counter++
	}

	// Shuffle our emails before returning them
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(emails), func(i, j int) { emails[i], emails[j] = emails[j], emails[i] })
	return emails, err
}
