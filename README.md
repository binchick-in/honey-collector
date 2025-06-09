# 🍯 Honey Collector

**Honey Collector** is a lightweight, scalable HTTP honeypot written in Go. It captures and logs HTTP requests—especially those from malicious actors—for security monitoring, threat intelligence, and research. Honey Collector acts as a decoy web server, attracting attackers and recording their activities for later analysis.

---

## 🧐 What is Honey Collector?

Honey Collector is a tool for running one or more fake HTTP servers ("honeypots") on your network. These servers are designed to attract and log unwanted or suspicious traffic, such as scans, exploit attempts, or automated attacks. By capturing this data, you can:

- Detect threats and attacks before they reach production systems
- Analyze attack patterns and payloads
- Gather threat intelligence (malicious IPs, user agents, payloads)
- Support security research and compliance efforts

---

## 🚨 Why Use Honey Collector?

- **Early Threat Detection:** Identify scanning and attack attempts in real time.
- **Attack Pattern Analysis:** Understand how attackers probe and exploit web services.
- **Threat Intelligence:** Collect indicators of compromise (IOCs) for blacklisting or further investigation.
- **Security Research:** Study malware, bots, and automated tools in a controlled environment.
- **Forensics & Compliance:** Maintain detailed logs for incident response and regulatory requirements.

---

## ⚙️ How to Use Honey Collector

### Prerequisites

- Go 1.18 or later
- Google Cloud Platform account with Pub/Sub API enabled (for cloud publishing)
- GCP service account with Pub/Sub Publisher permissions

### ⚡ Quick Start

1. **Clone and build:**
    ```sh
    git clone <repository-url>
    cd honey-collector
    make build
    ```

2. **Set up GCP credentials:**
    ```sh
    # Download your service account JSON key
    export GCP_CREDS=$(cat path/to/service-account-key.json)
    ```

3. **Run the honeypot:**
    ```sh
    ./honey -ports "8080,8443,80" -provider google
    ```

---

### 🛠️ Configuration Options

#### Command Line Flags

- `-ports`  
  Comma-separated list of ports to listen on (required).  
  Example: `-ports "80,443,8080"`

- `-provider`  
  Backend provider:  
  - `"google"` (Google Cloud Pub/Sub, default if using GCP)
  - `"sql"` (SQL backend, not yet implemented)

- `-response`  
  Custom response text for all HTTP requests (default: `\( ^ o ^)/`).

#### Environment Variables

- `GCP_CREDS`  
  Google Cloud service account JSON credentials (required for Google provider).

---

### 💡 Usage Examples

**Basic single-port honeypot:**
```sh
GCP_CREDS=$(cat creds.json) ./honey -ports "8080" --provider google
```

**Multi-port honeypot with custom response:**
```sh
GCP_CREDS=$(cat creds.json) ./honey \
  -ports "80,443,8080,8443" \
  -provider google \
  -response "Server temporarily unavailable"
```

---

### 📦 Data Format

Captured requests are serialized as JSON with the following structure:

```json
{
  "host": "example.com",
  "method": "POST",
  "path": "/admin/login",
  "remote_address": "192.168.1.100",
  "user_agent": "Mozilla/5.0...",
  "query_params": {"id": "123"},
  "headers": {"Content-Type": "application/json"},
  "body": "username=admin&password=123456",
  "time": 1640995200
}
```

---

## 🔒 Deployment & Security Considerations

- Deploy on isolated network segments or cloud sandboxes.
- Use dedicated service accounts with minimal permissions.
- Monitor honeypot logs for operational security.
- Regularly rotate credentials.
- Do **not** expose honeypots on production networks unless you understand the risks.

---

## 🛠️ Troubleshooting

- **"Failed to create Honey Provider"**  
  Check GCP credentials and permissions.

- **"Server error: listen tcp :8080: bind: address already in use"**  
  Port already in use; choose different ports.

- **"Publish error"**  
  Verify Pub/Sub topic exists and service account has publisher permissions.

---

## ⚠️ Current Limitations

- SQL backend provider is not yet implemented.
- No built-in rate limiting or DDoS protection.
- Basic logging only; no structured log levels.
- No configuration file support (CLI flags only).

---

**Deploy responsibly. Honey Collector is strictly for research purposes and fun!**
