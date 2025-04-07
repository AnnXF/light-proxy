//go:build darwin

package cmd

import (
	"fmt"
	"os"
	"strings"
)

const shellConfigPath = "~/.zshrc" // 适用于大多数 macOS 用户，若使用 bash 可改为 ~/.bash_profile

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
	configPath, err := expandPath(shellConfigPath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\nexport %s=\"%s\"", variable, value))
	return err
}

func removeFromEnv(variable string) error {
	configPath, err := expandPath(shellConfigPath)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", configPath, err)
	}
	lines := strings.Split(string(content), "\n")

	var newLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, "export "+variable+"=") {
			continue
		}
		newLines = append(newLines, line)
	}

	newContent := strings.Join(newLines, "\n")
	err = os.WriteFile(configPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to %s: %w", configPath, err)
	}

	return nil
}

func expandPath(path string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "~", home, 1), nil
}
