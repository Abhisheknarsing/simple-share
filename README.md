# 🚀 Simple Share

> **Instantly share files across your local network.**
> 
> *A lightweight, zero-config HTTP file server written in Go. Perfect for transferring photos, documents, or code between your computer and phone/tablet on the same Wi-Fi.*

## ✨ Features

- **Zero Configuration**: Just run it, and it works.
- **Auto-Discovery**: Automatically detects and displays your machine's Local LAN IP.
- **QR Code Ready** (Future plan): Easy mobile access via IP.
- **Logging**: Real-time access logs showing who accessed what.
- **Cross-Platform**: Compiles to a single binary for macOS, Linux, Windows.

---

## 📦 Installation

### Prerequisites
You need **Go (Golang)** installed (Version 1.20+ recommended).

### Build & Install
```bash
# Clone the repo
git clone https://github.com/Abhisheknarsing/simple-share.git
cd simple-share

# Install globally
go install
```
Or simply run it directly:
```bash
go run main.go
```

---

## 🚀 Usage

### Quick Start
Share the **current directory** on the default port (**8080**):
```bash
simple-share
```
**Output:**
```
🚀 Simple File Share Started
📂 Sharing Directory: .
Local:   http://localhost:8080
Network: http://192.168.1.5:8080
```

### Custom Directory & Port
Share a specific folder on a custom port:
```bash
simple-share -d /Users/me/Photos -p 3000
```
Then open `http://<YOUR-IP>:3000` on your phone to browse/download photos.

---

## 🛡️ Security Note
This tool serves **unencrypted HTTP** and is intended for use on **trusted local networks** (Home/Office Wi-Fi). Do not expose this port to the public internet unless you know what you are doing.
