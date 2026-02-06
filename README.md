<div align="center">

# ⚡ Thunder-TUI
**Blazing Fast HTTP Client | Terminal Based | Keyboard Driven**

<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&pause=1000&color=F7B93E&center=true&vCenter=true&width=435&lines=Test+APIs+via+SSH;No+Electron+bloat;JSON+Syntax+Highlighting;Vim+Keybindings+Supported" alt="Typing SVG" />

---

### ⚡ Quick Insights
[![Go Version](https://img.shields.io/github/go-mod/go-version/meklasdev/thunder-tui?style=for-the-badge&color=00ADD8&labelColor=1a1b26)](https://github.com/meklasdev/thunder-tui)
[![License](https://img.shields.io/github/license/meklasdev/thunder-tui?style=for-the-badge&color=bb9af7&labelColor=1a1b26)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge&color=7dcfff&labelColor=1a1b26)](http://makeapullrequest.com)

</div>

---

<table border="0">
  <tr>
    <td width="60%" valign="top">
      <h3>🛠️ The Arsenal</h3>
      <table>
        <tr><td><b>Core</b></td><td><img src="https://img.shields.io/badge/-Go-00ADD8?style=flat&logo=go&logoColor=white"/> <img src="https://img.shields.io/badge/-Bubbletea-F7B93E?style=flat"/> <img src="https://img.shields.io/badge/-Lipgloss-FF006E?style=flat"/></td></tr>
        <tr><td><b>Format</b></td><td><img src="https://img.shields.io/badge/-YAML-CB171E?style=flat&logo=yaml&logoColor=white"/> <img src="https://img.shields.io/badge/-JSON-000000?style=flat&logo=json&logoColor=white"/></td></tr>
        <tr><td><b>Platform</b></td><td><img src="https://img.shields.io/badge/-Linux-FCC624?style=flat&logo=linux&logoColor=black"/> <img src="https://img.shields.io/badge/-macOS-000000?style=flat&logo=apple&logoColor=white"/></td></tr>
      </table>
    </td>
    <td width="40%" valign="top" align="center">
      <h3>🚀 Quick Start</h3>
      <code>go install github.com/meklasdev/thunder-tui@latest</code>
      <br/><br/>
      <code>thunder-tui run collection.yaml</code>
    </td>
  </tr>
</table>

---

## 🎯 The Problem

Postman is **slow** (10+ second startup), **bloated** (Electron), and **cloud-dependent**.

When you're SSH'd into a server or working in a minimal environment, you need something **fast**, **lightweight**, and **terminal-native**.

---

## ✨ The Solution

**Thunder-TUI** is a blazing-fast HTTP client that runs entirely in your terminal. Built with Go and Bubbletea, it provides:

✅ **Instant Startup** - No Electron bloat  
✅ **SSH-Friendly** - Works over remote connections  
✅ **YAML Collections** - Version-controlled API tests  
✅ **Keyboard Driven** - Vim-style navigation  
✅ **Beautiful TUI** - Cyberpunk aesthetic with Lipgloss  
✅ **Zero Dependencies** - Single binary, no runtime required

---

## 🚀 Installation

### Via Go Install (Recommended)

```bash
go install github.com/meklasdev/thunder-tui/cmd/thunder-tui@latest
```

### From Source

```bash
git clone https://github.com/meklasdev/thunder-tui.git
cd thunder-tui
go build -o thunder-tui ./cmd/thunder-tui
```

### Binary Releases

Download pre-built binaries from the [Releases](https://github.com/meklasdev/thunder-tui/releases) page.

---

## 📖 Usage

### 1️⃣ Create a Collection

Create a `collection.yaml` file:

```yaml
requests:
  - name: "Get GitHub User"
    method: GET
    url: "https://api.github.com/users/meklasdev"
    headers:
      Accept: "application/vnd.github.v3+json"

  - name: "Create Post"
    method: POST
    url: "https://jsonplaceholder.typicode.com/posts"
    headers:
      Content-Type: "application/json"
    body: |
      {
        "title": "Hello World",
        "body": "This is a test",
        "userId": 1
      }
```

### 2️⃣ Run Thunder-TUI

```bash
thunder-tui run collection.yaml
```

### 3️⃣ Navigate & Test

- **↑/↓** or **j/k** - Navigate requests
- **Enter** - Send request
- **Tab** - Switch between panels
- **q** - Quit

---

## 🎨 Features

### 🎯 **Split-Panel Interface**

```
┌─────────────────────┬──────────────────────────┐
│ 📋 Requests         │ 📡 Response              │
│                     │                          │
│ ▶ GET    Get Users  │ Status: 200 OK           │
│   POST   Create     │ Duration: 234ms          │
│   DELETE Remove     │                          │
│                     │ Headers:                 │
│                     │   Content-Type: json     │
│                     │                          │
│                     │ Body:                    │
│                     │ { "users": [...] }       │
└─────────────────────┴──────────────────────────┘
```

### 🌈 **Method Color Coding**

- 🔵 **GET** - Blue
- 🟢 **POST** - Green
- 🟡 **PUT** - Yellow
- 🔴 **DELETE** - Red
- 🟣 **PATCH** - Purple

### ⚡ **Real-time Response**

- Status codes
- Response headers
- JSON body (syntax highlighted)
- Request duration

---

## 🎮 Keyboard Shortcuts

<div align="center">

| Key | Action |
|-----|--------|
| `↑` / `k` | Navigate up |
| `↓` / `j` | Navigate down |
| `Enter` | Send request |
| `Tab` | Switch panels |
| `q` / `Ctrl+C` | Quit |

</div>

---

## 📦 Collection Format

Thunder-TUI uses YAML for collections:

```yaml
requests:
  - name: "Request Name"
    method: GET|POST|PUT|DELETE|PATCH
    url: "https://api.example.com/endpoint"
    headers:
      Header-Name: "Header Value"
    body: |
      Request body (for POST/PUT/PATCH)
```

### Example Collection

See [`example-collection.yaml`](example-collection.yaml) for a complete example.

---

## 🛡️ Why Thunder-TUI?

| Feature | Thunder-TUI | Postman | Insomnia | curl |
|---------|-------------|---------|----------|------|
| **Startup Time** | ⚡ Instant | 🐌 10s+ | 🐌 5s+ | ⚡ Instant |
| **Memory Usage** | 🟢 ~10MB | 🔴 ~500MB | 🟡 ~200MB | 🟢 ~5MB |
| **SSH-Friendly** | ✅ Yes | ❌ No | ❌ No | ✅ Yes |
| **Collections** | ✅ YAML | ✅ Cloud | ✅ JSON | ❌ No |
| **TUI** | ✅ Beautiful | ❌ GUI | ❌ GUI | ❌ CLI |
| **Version Control** | ✅ Git-friendly | ⚠️ Export | ⚠️ Export | ❌ Scripts |

---

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...

# Build
go build -o thunder-tui ./cmd/thunder-tui
```

---

## 🤝 Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push and open a PR

---

## 📜 License

MIT © [meklasdev](https://github.com/meklasdev)

---

<div align="center">

### 🔗 Connect With Me
[<img src="https://img.shields.io/badge/Website-meklas.pl-7AA2F7?style=for-the-badge&logo=google-chrome&logoColor=white" />](https://meklas.pl)
[<img src="https://img.shields.io/badge/GitHub-meklasdev-181717?style=for-the-badge&logo=github&logoColor=white" />](https://github.com/meklasdev)

**[⭐ Star this repo](https://github.com/meklasdev/thunder-tui)** if you find it useful!

Made with ⚡ Go by **[meklasdev](https://github.com/meklasdev)**

</div>
