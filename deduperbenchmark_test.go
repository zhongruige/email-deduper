package main

import "testing"

func BenchmarkDeDupe(b *testing.B) {
	emails, _ := generateEmails(100000, 0.5)
	for n := 0; n < b.N; n++ {
		dedupe(emails)
	}
}
