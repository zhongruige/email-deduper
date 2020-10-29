package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidEmail_MissingAt(t *testing.T) {
	badEmail := "emailgomail.com"
	r := isValidEmail(badEmail)
	assert.False(t, r, "Expected validEmail to be false with missing @")
}

func TestCheckValidEmail_MissingDot(t *testing.T) {
	badEmail := "email@gomailcom"
	r := isValidEmail(badEmail)
	assert.False(t, r, "Expected validEmail to be false with missing .")
}

func TestCheckValidEmail_NoDelimiters(t *testing.T) {
	badEmail := "emailgomailcom"
	r := isValidEmail(badEmail)
	assert.False(t, r, "Expected validEmail to be false with no delimiters")
}

func TestCheckValidEmail_GoodEmail(t *testing.T) {
	goodEmail := "email@gomail.com"
	r := isValidEmail(goodEmail)
	assert.True(t, r, "Expected validEmail to be true for valid email")
}

func TestCheckValidEmail_Dotnames(t *testing.T) {
	goodEmail := "my.email@gomail.com"
	r := isValidEmail(goodEmail)
	assert.True(t, r, "Expected validEmail to be true for valid email with dot names")
}

func TestDeDupeMaintainsOriginalOrder(t *testing.T) {
	emails := []string{"3_email@gomail.com", "4_email@gomail.com", "0_email@gomail.com", "0_email@gomail.com", "0_email@gomail.com", "2_email@gomail.com", "0_email@gomail.com", "1_email@gomail.com", "0_email@gomail.com", "0_email@gomail.com"}
	expectedResult := []string{"3_email@gomail.com", "4_email@gomail.com", "0_email@gomail.com", "2_email@gomail.com", "1_email@gomail.com"}
	dedupedEmails := dedupe(emails)
	assert.EqualValues(t, expectedResult, dedupedEmails, "Dedupe should maintain the same ordering")
}

func TestDeDupeMaintsOriginalOrder_RemovingInvalidEmails(t *testing.T) {
	emails := []string{"invalidemail", "0_email@gomail.com", "0_email@gomail.com", "0_email@gomail.com", "2_email@gomail.com", "0_email@gomail.com", "1_email@gomail.com", "0_email@gomail.com", "0_email@gomail.com"}
	expectedResult := []string{"0_email@gomail.com", "2_email@gomail.com", "1_email@gomail.com"}
	dedupedEmails := dedupe(emails)
	assert.EqualValues(t, expectedResult, dedupedEmails, "Dedupe should maintain the same ordering")
}

func TestDeDupeEmailsDefaultValues(t *testing.T) {
	emails, _ := generateEmails(0, 0)
	dedupedEmails := dedupe(emails)
	assert.Equal(t, 50000, len(dedupedEmails), "Deduped all 50000 duplicates")
}

func TestDeDupeEmailsSpecificEmailAndDuplicateCounts(t *testing.T) {
	emails, _ := generateEmails(1128, 0.3)
	dedupedEmails := dedupe(emails)
	assert.Equal(t, 790, len(dedupedEmails), "Deduped all 338 duplicates")
}

func TestDeDupeEmailsSpecifiedDuplicatePercent(t *testing.T) {
	emails, _ := generateEmails(0, 0.8)
	dedupedEmails := dedupe(emails)
	assert.Equal(t, 20000, len(dedupedEmails), "Deduped all 80000 duplicates")
}

func TestDeDupeInValidEmails(t *testing.T) {
	emails := []string{"valid@email.com", "invalidemail", "real@email.com"}
	dedupedEmails := dedupe(emails)
	assert.Equal(t, 2, len(dedupedEmails), "Deduped emails with an invalid email should return a length of 2")
}
