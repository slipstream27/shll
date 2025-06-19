# shll 🧠💻

`shll` is a natural-language-powered shell assistant. Just tell it what you want in plain English, and it suggests real, runnable shell commands — with AI.

```bash
shll install the lts version of node using nvm
```

Example output:
```
Available commands:
1: nvm install --lts
2: nvm install lts/*
3: curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
m: More results

Select a command by number to copy to clipboard (add '?' to explain, e.g. 1?):
```
- Add `?` after a number (e.g. `1?`) to get an explanation for that command.
- Enter `m` for more suggestions.
- Enter `q` to exit at any prompt.

---

## Quick Start

1. Clone the repository:
   ```sh
   git clone https://github.com/YOUR_USERNAME/shll.git
   cd shll
   ```
2. Build the binary:
   ```sh
   make build
   ```
   The binary will be in the `bin/` directory.

   > **Note:** Official packages for Homebrew, APT, and YUM are coming soon.

3. Set your OpenAI API key:
   ```sh
   export OPENAI_API_KEY=sk-xxxx...
   ```
4. Run shll:
   
   Once built, you can run shll from your terminal:
   ```sh
   ./bin/shll find all .txt files in this directory and count lines
   ```
   > Some shells (such as `zsh` or when using special characters) may require your prompt to be enclosed in quotes:
   > ```
   > shll "find all .txt files in this directory and count lines"
   > ```
   > 
   > **Tip:** To use `shll` without needing quotes, add this to your `~/.zshrc` or `~/.bashrc`:
   > ```sh
   > shll() {
   >   /full/path/to/bin/shll "$*"
   > }
   > ```
   > Replace `/full/path/to/bin/shll` with the actual path to your compiled binary. After adding this, restart your terminal or run `source ~/.zshrc` (or `source ~/.bashrc`).
