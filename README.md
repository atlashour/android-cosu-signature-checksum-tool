# android-cosu-signature-checksum-tool

🔐 A utility to generate valid `SIGNATURE_CHECKSUM` values for Android COSU provisioning. Converts certificate SHA-256 into proper base64url format accepted by Android provisioning QR.

This tool helps generate correct `PROVISIONING_DEVICE_ADMIN_SIGNATURE_CHECKSUM` values for Android COSU (Corporate Owned Single Use) provisioning.  
It extracts the SHA-256 fingerprint from your APK's signing certificate, encodes it to a URL-safe base64 string **without padding** — the exact format Android expects but never documents.

---

## ✅ Why use SIGNATURE_CHECKSUM?

Using the certificate's SHA-256 as a signature checksum allows Android to verify the APK's origin without hard-locking it to a specific build.  
This is the recommended method for provisioning production devices, as it:

- Supports future updates signed with the same keystore
- Avoids breaking provisioning due to minor changes in the APK
- Prevents unnecessary QR regeneration

---

## ⚙️ Requirements

- ✅ Your `.apk` must be **signed** with a custom keystore (release build)
- ✅ You need access to the `keytool` binary

### 📦 Keytool path for Android Studio users (Windows):

Add this to your system `PATH` environment variable:

```
C:\Program Files\Android\Android Studio\jbr\bin
```

Or run it explicitly from that folder if needed.

---

## 🚀 How to use

### Option 1: Using PowerShell (Windows)

Use the provided `apk_signature_checksum.ps1` script:

```powershell
.\apk_signature_checksum.ps1 -ApkPath "C:\path\to\your\app-release.apk"
```

### Option 2: Using Go (Cross-platform)

Use the Go version:

```bash
go run apk_signature_checksum.go path/to/app-release.apk
```

> Requires Go installed. You can compile it with:
> `go build -o cosu-checksum.exe apk_signature_checksum.go`

---

## 📤 Output

You will get a clean, valid checksum like this:

```
✔️ Final checksum for PROVISIONING_DEVICE_ADMIN_SIGNATURE_CHECKSUM:
GQfdGZbRF9hzPKYzJB0y6xgSrplOWAK-W0KL0r4Ud0v
```

This is safe to use in your provisioning QR JSON like:

```json
{
  "android.app.extra.PROVISIONING_DEVICE_ADMIN_SIGNATURE_CHECKSUM": "GQfdGZbRF9hzPKYzJB0y6xgSrplOWAK-W0KL0r4Ud0v"
}
```

---

## 🧠 What makes this tool special?

Android expects a **very specific encoding**:

- SHA-256 of the **certificate** (not the APK itself)
- Encoded as **base64url** (not standard base64)
- **No padding** (`=` must be removed)

Failing to meet this exact format results in **silent provisioning failure** with the dreaded and vague error:

> `"Can't set up device. Contact your IT admin for help."`

You will get **no logs, no stacktrace, and no explanation**.

---

## 👾 Contribute

This tool was built to save developers from losing days on undocumented Android behavior.  
PRs, issues, or usage reports are welcome!

---

With ❤️ by **[Atlashour](https://github.com/atlashour)**  
✨ Special thanks to **ChatGPT-4o** — extremely supportive and accurate.

---

### 🔍 Keywords for searchability (SEO)

- Android COSU provisioning tool
- PROVISIONING_DEVICE_ADMIN_SIGNATURE_CHECKSUM
- Can't set up device Contact your IT admin for help
- Android base64url signature hash
- Silent provisioning failure
- COSU QR code Android 13 / 14 / 15
```
