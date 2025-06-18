# shll 🧠💻

`shll` is a natural-language-powered shell assistant. Just tell it what you want in plain English, and it suggests real, runnable shell commands — with AI.

```bash
shll install the lts version of node using nvm
```

➡️ Suggests:
```
1. nvm install --lts
2. nvm install lts/*
```

Arrow through options, pick one, and optionally run it.

---

## ✨ Features

- 🔍 Turn natural language into command-line instructions
- 🤖 Powered by OpenAI (or optional local LLMs via Ollama)
- 🧭 Arrow-key interactive interface with `promptui`
- 💥 Optional command execution (with confirmation)
- 📦 Cross-platform static binaries (Linux, macOS, Windows)
- 🧃 Homebrew, APT, and RPM packaging support via Goreleaser

---

## 🧪 Quick Start

### 1. Install (Coming Soon)

**Homebrew:**
```bash
brew install yourname/tap/shll
```

**Debian/Ubuntu:**
```bash
sudo apt install shll
```

**RHEL/Fedora/CentOS:**
```bash
sudo yum install shll
```

Or download the binary from [Releases](https://github.com/YOUR_USERNAME/shll/releases)

---

### 2. Set API Key

```bash
export OPENAI_API_KEY=sk-xxxx...
```

---

### 3. Use It

```bash
shll create a zip of the current folder and exclude node_modules
```

---

## 🧠 Powered By

- [OpenAI GPT-4](https://platform.openai.com/)
- [`go-openai`](https://github.com/sashabaranov/go-openai)
- [`promptui`](https://github.com/manifoldco/promptui)
- [`goreleaser`](https://goreleaser.com/)

---

## 🔒 Privacy & Safety

No commands are run without your confirmation. You are responsible for what you choose to execute.

---

## 🛠 Local Development

```bash
git clone https://github.com/YOUR_USERNAME/shll
cd shll
go build -o shll ./cmd/shll
./shll "list open ports on this machine"
```
> **Note:** Some shells (such as `zsh` or when using special characters) may require your prompt to be enclosed in quotes:
> ```
> shll "find all .txt files in this directory and count lines"
> ```

---

## 📝 Shell Alias: Use shll Without Quotes

To use `shll` without needing quotes around your prompt, add this alias or function to your shell config (e.g., `~/.zshrc` or `~/.bashrc`):

```sh
shll() {
  /full/path/to/shll "$*"
}
```

Replace `/full/path/to/shll` with the actual path to your compiled binary (e.g., `~/shll/bin/shll`).

After adding this, restart your terminal or run `source ~/.zshrc` (or `source ~/.bashrc`). Now you can run:

```
shll find all .txt files in this directory and count lines
```

No quotes needed—the entire prompt will be passed as a single argument.

---

## 📦 Releases

We use [Goreleaser](https://goreleaser.com/) to generate:
- `.tar.gz` binaries for all major OSes
- `.deb` and `.rpm` packages
- Homebrew formula for Mac

---

## 🧰 TODO / Roadmap

- [ ] Support local LLMs via Ollama
- [ ] Save command history & favorites
- [ ] Add plugins & context awareness
- [ ] Dry-run / safe-mode by default
- [ ] Explain command mode

---

## 💬 License

Don't Be a Dick Public License (DBAD) — [read here](https://github.com/philsturgeon/dbad/blob/main/LICENSE.md).

---
