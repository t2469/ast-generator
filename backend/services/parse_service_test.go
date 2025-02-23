package services

import (
	"testing"
)

var languageTests = []struct {
	language string
	code     string
}{
	{"bash", "echo Hello"},
	{"c", "int main() { return 0; }"},
	{"cpp", "#include <iostream>\nint main() { std::cout << \"Hello\"; return 0; }"},
	{"css", "body { margin: 0; }"},
	{"dockerfile", "FROM alpine"},
	{"go", "package main\nfunc main() {}"},
	{"html", "<!DOCTYPE html><html></html>"},
	{"java", "public class Main { public static void main(String[] args) {} }"},
	{"javascript", "console.log('Hello');"},
	{"kotlin", "fun main() { println(\"Hello\") }"},
	{"php", "<?php echo 'Hello'; ?>"},
	{"python", "print('Hello')"},
	{"ruby", "puts 'Hello'"},
	{"rust", "fn main() { println!(\"Hello\"); }"},
	{"sql", "SELECT 1;"},
	{"yaml", "key: value"},
}

func TestParseCode_AllLanguages(t *testing.T) {
	for _, tt := range languageTests {
		t.Run(tt.language, func(t *testing.T) {
			ast, err := ParseCode(tt.language, tt.code)
			if err != nil {
				t.Errorf("ParseCode(%q) returned error: %v", tt.language, err)
				return
			}
			if ast.Type == "" {
				t.Errorf("ParseCode(%q) returned empty AST type", tt.language)
			}
		})
	}
}

func TestParseCode_UnsupportedLanguage(t *testing.T) {
	_, err := ParseCode("unsupported", "some code")
	if err == nil {
		t.Error("Expected error for unsupported language, but got none")
	}
}
