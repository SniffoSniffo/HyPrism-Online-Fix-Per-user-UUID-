// internal/game/client_replace.go
package game

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"HyPrism/internal/assets"
)

// ReplaceHytaleClient replaces the installed HytaleClient with a bundled custom one.
func ReplaceHytaleClient(gameDir string) error {
	clientDir := filepath.Join(gameDir, "Client")
	if _, err := os.Stat(clientDir); os.IsNotExist(err) {
		return fmt.Errorf("client directory not found: %s", clientDir)
	}

	var targetPath string
	var data []byte

	switch runtime.GOOS {
	case "windows":
		targetPath = filepath.Join(clientDir, "HytaleClient.exe")
		data = assets.WindowsClient
	default:
		// Skip replacement on non-Windows for now (unless you add Unix support)
		return nil
	}

	// Safety check: ensure embedded data isn't empty
	if len(data) == 0 {
		return fmt.Errorf("embedded custom client is empty")
	}

	// Write the bundled client
	if err := os.WriteFile(targetPath, data, 0755); err != nil {
		return fmt.Errorf("failed to write custom client: %w", err)
	}

	fmt.Printf("✅ Replaced HytaleClient at %s\n", targetPath)
	return nil
}