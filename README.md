# Serpent

A basic cryptolocker and data exfiltration demonstration tool for educational and authorized security testing purposes.

## ⚠️ LEGAL DISCLAIMER

**THIS SOFTWARE IS PROVIDED FOR EDUCATIONAL AND AUTHORIZED SECURITY TESTING PURPOSES ONLY.**

This repository contains a demonstration model of ransomware-like behavior (file encryption) and data exfiltration capabilities. It is intended **EXCLUSIVELY** for:

- Educational purposes in cybersecurity courses
- Authorized penetration testing and red team exercises
- Purple team training and security awareness demonstrations
- Controlled laboratory environments with explicit permission

### Prohibited Uses

**UNAUTHORIZED USE OF THIS SOFTWARE IS STRICTLY PROHIBITED AND ILLEGAL.**

You may NOT use this software to:
- Encrypt files on systems you do not own or have explicit authorization to test
- Exfiltrate data from systems without proper authorization
- Cause harm, damage, or disruption to any computer systems or networks
- Violate any local, state, national, or international laws

### Legal Consequences

Unauthorized use of this software may violate:
- Computer Fraud and Abuse Act (CFAA) in the United States
- Computer Misuse Act in the United Kingdom
- Similar cybercrime legislation in other jurisdictions

Violators may face severe criminal penalties including imprisonment and substantial fines.

### No Responsibility

**THE AUTHORS AND CONTRIBUTORS OF THIS SOFTWARE DISCLAIM ALL RESPONSIBILITY AND LIABILITY FOR ANY MISUSE, DAMAGE, OR ILLEGAL ACTIVITIES CONDUCTED WITH THIS SOFTWARE.**

By downloading, accessing, or using this software, you acknowledge that:
1. You will use it only for lawful, authorized purposes
2. You accept full responsibility for your actions
3. The authors bear no liability for any consequences of misuse
4. You will comply with all applicable laws and regulations

## Features

- **File Encryption**: AES-256-GCM encryption of files in the current directory
- **Data Exfiltration**: HTTPS-based file transmission to a specified IP address
- **Cross-Platform**: Builds for Windows (PE) and Unix-like systems
- **Configurable**: Command-line flags for target IP configuration
- **Safety Features**: 
  - Excludes `.go` and `.mod` source files from encryption
  - Excludes the executable itself from encryption and transmission

## Architecture

```
Serpent/
├── src/
│   ├── serpent.go          # Main application entry point
│   ├── serpent_test.go     # Test suite
│   ├── Makefile            # Build configuration
│   ├── go.mod              # Go module definition
│   └── pkg/
│       ├── encrypt/        # Encryption and file operations
│       └── https/          # HTTPS file transmission
├── build/                  # Compiled executables
├── test/                   # Test files
└── val/                    # Validation directory
```

## Building

```bash
cd src
make build
```

This will create:
- `build/serpent` - Native executable (macOS/Linux)
- `build/serpent.exe` - Windows PE executable

## Usage

**⚠️ USE ONLY IN AUTHORIZED, CONTROLLED ENVIRONMENTS ⚠️**

```bash
# Run with default IP (127.0.0.1)
./serpent

# Specify target IP address
./serpent -i <ip_address>
```

The program will:
1. Display ASCII art banner
2. Attempt to send all files in the current directory to the specified IP via HTTPS
3. Encrypt all files in the current directory (excluding source files and the executable)

## Testing

```bash
cd src
go test
```

## Requirements

- Go 1.25.5 or higher
- HTTPS server for receiving exfiltrated files (for testing purposes)

## Security Considerations

This tool demonstrates:
- **Attack Vectors**: How ransomware encrypts files and exfiltrates data
- **Detection Opportunities**: File system monitoring, network traffic analysis
- **Mitigation Strategies**: Backups, endpoint protection, network segmentation

Use this knowledge to **defend** against real threats, not to create them.

## License

This software is provided "as is" without warranty of any kind. See the disclaimer above for terms of use.

## Ethical Use Agreement

By using this software, you agree to:
- Obtain proper authorization before testing on any system
- Use it only for defensive security purposes
- Not distribute it to unauthorized parties
- Report any vulnerabilities responsibly

---

**Remember: With great power comes great responsibility. Use your skills ethically and legally.**
