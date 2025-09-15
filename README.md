# 📁 File Uploader

**Universal file uploader for old browsers and mobile devices**

A simple, lightweight web server that allows you to upload files through any browser, including old browsers like IE9+, Safari 5+, and mobile devices running iOS 9.3.5+ or Android 4.0+.

## 🚀 Features

- ✅ **Universal compatibility** - Works with old browsers and mobile devices
- ✅ **No dependencies** - Single executable file, ready to run
- ✅ **Multiple file upload** - Upload multiple files at once
- ✅ **Automatic IP detection** - Shows your local IP for easy access
- ✅ **Progress tracking** - Real-time upload progress
- ✅ **Any file type** - Upload any file type (50MB limit)
- ✅ **Local network access** - Access from any device on the same network
- ✅ **Cross-platform** - Windows, macOS, and Linux support

## 📥 Quick Download & Run

### Latest Release: [v1.2.0](https://github.com/Yakwilik/photo-uploader/releases/latest)

**Download the appropriate file for your system:**

| Platform | Architecture | File |
|----------|-------------|------|
| 🪟 Windows | 64-bit | `file-uploader-windows-amd64.exe` |
| 🪟 Windows | 32-bit | `file-uploader-windows-386.exe` |
| 🍎 macOS | Intel | `file-uploader-darwin-amd64` |
| 🍎 macOS | Apple Silicon (M1/M2) | `file-uploader-darwin-arm64` |
| 🐧 Linux | 64-bit | `file-uploader-linux-amd64` |
| 🐧 Linux | 32-bit | `file-uploader-linux-386` |
| 🐧 Linux | ARM64 | `file-uploader-linux-arm64` |
| 🐧 Linux | ARM | `file-uploader-linux-arm` |

## 🔒 Important: Unblock Downloaded Files

Modern operating systems block files downloaded from the internet for security. You need to unblock them first:

### 🪟 Windows
1. Right-click the downloaded `.exe` file
2. Select 'Properties'
3. Check 'Unblock' at the bottom and click 'OK'
4. Or run in PowerShell: `Unblock-File file-uploader-windows-amd64.exe`

### 🍎 macOS
1. Try to run: `./file-uploader-darwin-amd64`
2. If blocked, run: `xattr -d com.apple.quarantine file-uploader-darwin-amd64`
3. Or: Right-click → 'Open' → 'Open' again to bypass Gatekeeper

### 🐧 Linux
1. Usually no blocking, but if needed: `chmod +x file-uploader-linux-amd64`
2. Some distros may show warning - click 'Trust and Launch' or 'Execute'

## 💻 Installation & Usage

### 🪟 Windows
```bash
# Download file-uploader-windows-amd64.exe
# Double-click to run, or run from command line:
file-uploader-windows-amd64.exe
```

### 🍎 macOS
```bash
# Download file-uploader-darwin-amd64 (Intel) or file-uploader-darwin-arm64 (Apple Silicon)
chmod +x file-uploader-darwin-amd64
./file-uploader-darwin-amd64
```

### 🐧 Linux
```bash
# Download appropriate file for your architecture
chmod +x file-uploader-linux-amd64
./file-uploader-linux-amd64
```

## 🚀 What happens after running?

1. **Server starts** on port 8080 (default)
2. **IP address is displayed** in the terminal
3. **Web interface opens** automatically in your browser
4. **Upload files** by dragging & dropping or clicking to select
5. **Access from any device** on the same network using the IP address

**Example output:**
```
File Uploader v1.2.0 starting...
Server running on: http://192.168.1.100:8080
Open this URL in any browser to upload files!
```

## 🔧 Usage

1. **Run the binary** on your computer
2. **The server will start** and show your IP address
3. **Open the URL** in any browser on the same network
4. **Upload files** through the web interface

## ❓ Troubleshooting & Tips

### 🔧 Common Issues

#### 🚫 Security/Blocking Issues
- **Windows: 'Windows protected your PC'**: Click 'More info' → 'Run anyway'
- **Windows: 'Unblock' checkbox**: Right-click file → Properties → Check 'Unblock' → OK
- **macOS: 'Cannot be opened'**: Right-click → 'Open' → 'Open' again
- **macOS: Quarantine error**: Run `xattr -d com.apple.quarantine file-uploader-darwin-amd64`
- **Linux: Permission denied**: Run `chmod +x file-uploader-linux-amd64`

#### ⚙️ Runtime Issues
- **Port 8080 busy**: The app will automatically find an available port
- **Can't access from phone**: Make sure both devices are on the same WiFi network
- **Firewall blocking**: Allow the application through Windows/macOS firewall

### 💡 Pro Tips
- Use **Intel Macs**: Download `darwin-amd64` version
- Use **Apple Silicon Macs** (M1/M2): Download `darwin-arm64` version
- **Upload large files**: Drag & drop works better than clicking 'Choose Files'
- **Multiple files**: Hold Ctrl/Cmd while selecting multiple files

### 🔧 Alternative Launch Methods
- **Windows (Command Prompt)**: Open CMD as Administrator, navigate to file, run directly
- **macOS (Terminal)**: Open Terminal, `cd` to download folder, `./file-uploader-darwin-amd64`
- **Linux (Terminal)**: Open terminal, `cd` to download folder, `./file-uploader-linux-amd64`

### 🛑 To stop the server: Press `Ctrl+C` in the terminal

## 📋 System Requirements

- Any modern operating system (Windows, macOS, Linux)
- Network access for file sharing
- Go 1.21+ (for building from source)

## 🛠️ Building from Source

If you want to build the application from source:

```bash
# Clone the repository
git clone https://github.com/Yakwilik/photo-uploader.git
cd photo-uploader

# Build for your platform
go build -o file-uploader ./cmd/photo-uploader

# Or build for specific platforms
GOOS=windows GOARCH=amd64 go build -o file-uploader-windows-amd64.exe ./cmd/photo-uploader
GOOS=darwin GOARCH=arm64 go build -o file-uploader-darwin-arm64 ./cmd/photo-uploader
GOOS=linux GOARCH=amd64 go build -o file-uploader-linux-amd64 ./cmd/photo-uploader
```

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📞 Support

If you encounter any issues or have questions:

1. Check the [Troubleshooting](#-troubleshooting--tips) section above
2. Look at the [Releases](https://github.com/Yakwilik/photo-uploader/releases) for the latest version
3. Open an [Issue](https://github.com/Yakwilik/photo-uploader/issues) on GitHub

---

**Made with ❤️ for easy file sharing across all devices and browsers**