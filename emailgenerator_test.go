package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateNegativeEmails(t *testing.T) {
	_, err := generateEmails(-1, 0)
	assert.Error(t, err, "Received expected error when -1 email count")
}

func TestGenerateNegativeDuplicatePercentage(t *testing.T) {
	_, err := generateEmails(0, -1)
	assert.Error(t, err, "Received expected error when -1 email count")
}

func TestGenerateGreaterDuplicatePercentage(t *testing.T) {
	_, err := generateEmails(0, 2)
	assert.Error(t, err, "Received expected error when -1 email count")
}

func TestGenerateDefaultEmails(t *testing.T) {
	emails, _ := generateEmails(0, 0)
	assert.EqualValues(t, 100000, len(emails), "Email count should match 100000")
	dupeCount := GetDuplicateCount(emails)
	assert.EqualValues(t, dupeCount, 50000, "Duplicate count should be 50000")
}

func TestGenerateSpecifiedEmailCount(t *testing.T) {
	emails, _ := generateEmails(380, 0)
	assert.EqualValues(t, 380, len(emails), "Email count should match 380")
	dupeCount := GetDuplicateCount(emails)
	assert.EqualValues(t, dupeCount, 190, "Duplicate count should be 190")
}

func TestGenerateSpecifiedDuplicatePercent(t *testing.T) {
	emails, _ := generateEmails(0, 0)
	assert.EqualValues(t, 100000, len(emails), "Email count should match 100000")
	dupeCount := GetDuplicateCount(emails)
	assert.EqualValues(t, 50000, dupeCount, "Duplicate count should be 50000")
}

func TestGenerateSpecificEmailAndDuplicateCounts(t *testing.T) {
	emails, _ := generateEmails(1128, 0.3)
	assert.EqualValues(t, 1128, len(emails), "Email count should match 1128")
	dupeCount := GetDuplicateCount(emails)
	assert.EqualValues(t, 338, dupeCount, "Duplicate count should be 338")
}

// Helper function to get duplicate count
func GetDuplicateCount(emails []string) int {
	keys := make(map[string]bool)
	deDupedEmails := []string{}
	for _, email := range emails {
		if _, value := keys[email]; !value {
			keys[email] = true
			deDupedEmails = append(deDupedEmails, email) // this returns all of the non dupes
		}
	}
	return len(emails) - len(deDupedEmails)
}
