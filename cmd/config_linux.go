//bo:build linux

package cmd

import (
	"fmt"
	"os"
	"strings"
)

func configClear() {
	if err := removeFromEnv("http_proxy"); err != nil {
		println("[sys] proxy config clear failed", err.Error())
		return
	}
	if err := removeFromEnv("https_proxy"); err != nil {
		println("[sys] proxy config clear failed", err.Error())
		return
	}
	println("[sys] proxy config clear successfully")
}

func configSet(domain string) {
	if err := appendToEnv("http_proxy", domain); err != nil {
		println("[sys] proxy config set failed", err.Error())
		return
	}
	if err := appendToEnv("https_proxy", domain); err != nil {
		println("[sys] proxy config set failed", err.Error())
		return
	}
	println("[sys] proxy config set successfully")
}

func appendToEnv(variable, value string) error {
	file, err := os.OpenFile("/etc/environment", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\n%s=\"%s\"", variable, value))
	return err
}

func removeFromEnv(variable string) error {
	envPath := "/etc/environment"
	content, err := os.ReadFile(envPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", envPath, err)
	}
	lines := strings.Split(string(content), "\n")

	var newLines []string
	for _, line := range lines {
		if strings.Contains(line, variable+"=") {
			continue
		}
		newLines = append(newLines, line)
	}

	newContent := strings.Join(newLines, "\n")
	err = os.WriteFile(envPath, []byte(newContent), 0755)
	if err != nil {
		return fmt.Errorf("failed to write to %s: %w", envPath, err)
	}

	return nil
}
