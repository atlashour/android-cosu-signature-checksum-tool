# android-cosu-signature-checksum-tool

🔐 A utility to generate valid SIGNATURE_CHECKSUM values for Android COSU provisioning. Converts certificate SHA-256 into proper base64url format accepted by Android provisioning QR.

This tool helps generate correct PROVISIONING_DEVICE_ADMIN_SIGNATURE_CHECKSUM values for Android COSU (Corporate Owned Single Use) provisioning.
It extracts the SHA-256 fingerprint from your APK's signing certificate, encodes it to a URL-safe base64 string without padding — the exact format Android expects but never documents.
