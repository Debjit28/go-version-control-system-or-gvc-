# go-version-control

A lightweight Git-compatible version control system built from scratch in Go. This project implements core Git internals — including object storage, blob hashing, and content inspection — giving you a ground-up understanding of how Git works under the hood.

---

## 📁 Project Structure

```
go-version-control/
├── main.go                  # Entry point & command dispatcher
├── commands/
│   ├── cat_file.go          # Implements the cat-file command
│   └── hash_object.go       # Implements the hash-object command
└── objects/
    └── blob.go              # Core blob read/write logic
```

---

## 🚀 How It Works

### Initialization (`init`)

Running `mygit init` sets up a `.git` directory structure that mirrors a real Git repository:

```
.git/
├── HEAD           # Points to refs/heads/main
├── objects/       # Stores all compressed object data
└── refs/          # Stores branch and tag references
```

The `HEAD` file is initialized to point to `refs/heads/main`, just like a freshly initialized Git repository.

---

### Object Storage — Blobs

The core of the system is the **blob object**, which is how file contents are stored. The pipeline for storing a file is:

1. **Read** the file contents from disk
2. **Prepend a header** in the format `blob <size>\0` (null-terminated)
3. **Hash** the full data (header + content) using **SHA3-256** — a stronger alternative to Git's SHA1, producing a 64-character hex digest
4. **Compress** the data using **zlib**
5. **Write** the compressed result to `.git/objects/<first-2-chars>/<remaining-62-chars>`

> **Why SHA3-256?** Unlike Git's original SHA1 (or even SHA256), SHA3 uses a fundamentally different sponge construction, offering significantly higher collision resistance.

---

## 🛠️ Commands

### `mygit init`

Initializes a new repository in the current directory.

```bash
$ mygit init
Initialized git directory
```

---

### `mygit hash-object -w <file>`

Reads a file, creates a blob object, stores it in `.git/objects/`, and prints the SHA3-256 hash.

```bash
$ mygit hash-object -w hello.txt
a3f1c8d...   # 64-character SHA3-256 hash
```

**Flags:**
| Flag | Description |
|------|-------------|
| `-w` | Write the object to the object store (required) |

**Internally:**
- Reads the file at `<file>`
- Builds the blob: `blob <len>\0<content>`
- Hashes with SHA3-256
- Compresses with zlib and saves to `.git/objects/`

---

### `mygit cat-file -p <hash>`

Reads a stored blob object by its hash and prints the original file content.

```bash
$ mygit cat-file -p a3f1c8d...
Hello, world!
```

**Flags:**
| Flag | Description |
|------|-------------|
| `-p` | Pretty-print the object's content |

**Internally:**
- Resolves the object path from the hash: `.git/objects/<hash[:2]>/<hash[2:]>`
- Decompresses using zlib
- Strips the blob header (everything up to and including the null byte `\0`)
- Prints the raw file content

---

## ⚙️ Technical Details

| Feature | Detail |
|---|---|
| Language | Go |
| Hashing | SHA3-256 (64-char hex digest) |
| Compression | zlib |
| Object format | `blob <size>\0<content>` |
| Storage path | `.git/objects/<2-char prefix>/<62-char suffix>` |

---

## 🔧 Building & Running

```bash
# Build the project
go build -o mygit .

# Initialize a repo
./mygit init

# Store a file as a blob
./mygit hash-object -w myfile.txt

# Read back the blob
./mygit cat-file -p <hash>
```

---

## 🗺️ Roadmap

Planned additions to expand git compatibility:

- [ ] `ls-tree` — list tree objects
- [ ] `write-tree` — write the current index as a tree object
- [ ] `commit-tree` — create a commit object
- [ ] `clone` — clone a remote repository
- [ ] Switch to SHA256 for full Git compatibility

---
