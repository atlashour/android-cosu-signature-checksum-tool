package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <apk_path>", os.Args[0])
	}

	apkPath := os.Args[1]

	// Execute keytool (Java must be in PATH)
	cmd := exec.Command("keytool", "-printcert", "-jarfile", apkPath)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running keytool (make sure Java is installed and in your PATH): %v", err)
	}

	// Find the line containing SHA256
	lines := strings.Split(string(output), "\n")
	var hexLine string
	for _, line := range lines {
		if strings.Contains(line, "SHA256:") {
			hexLine = strings.TrimSpace(strings.Replace(line, "SHA256:", "", 1))
			break
		}
	}

	if hexLine == "" {
		log.Fatal("SHA256 not found in keytool output.")
	}

	// Remove colons and decode from hex to bytes
	hexStr := strings.ReplaceAll(hexLine, ":", "")
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Fatalf("Error decoding SHA256 hex string: %v", err)
	}

	// Convert to base64
	b64 := base64.StdEncoding.EncodeToString(bytes)

	// Convert to URL-safe base64 (no padding)
	b64url := strings.TrimRight(b64, "=")
	b64url = strings.ReplaceAll(b64url, "+", "-")
	b64url = strings.ReplaceAll(b64url, "/", "_")

	fmt.Println("✔️ Final checksum for PROVISIONING_DEVICE_ADMIN_SIGNATURE_CHECKSUM:")
	fmt.Println(b64url)
}
