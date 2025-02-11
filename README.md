# HiveShake
This toolkit will help test system resilience by injecting failures like process termination, CPU stress, network disruptions, and disk space exhaustion.

### **🚀 Chaos Engineering Toolkit in Golang**  
This toolkit will help test system resilience by injecting failures like process termination, CPU stress, network disruptions, and disk space exhaustion.  

---

## **🛠️ Architecture Breakdown**  

### **1️⃣ Core Components**  
| Component         | Description |
|------------------|-------------|
| **CLI Interface** | Handles user commands and triggers experiments. |
| **Fault Injection Engine** | Manages different chaos experiments (CPU stress, network latency, process killing, etc.). |
| **Observability Module** | Logs experiment results and provides metrics via Prometheus/Grafana. |
| **API Server (Optional)** | Exposes REST/gRPC endpoints for remote execution of chaos tests. |
| **Kubernetes Controller (Optional)** | Runs chaos tests inside Kubernetes clusters. |

---

### **2️⃣ High-Level Flow**  
1️⃣ **User triggers chaos experiment via CLI or API.**  
2️⃣ **Fault Injection Engine executes the selected failure mode.**  
3️⃣ **Observability Module logs experiment details.**  
4️⃣ **System metrics are collected and exposed via Prometheus.**  
5️⃣ **(Optional) Kubernetes integration for testing containerized workloads.**  

---

### **3️⃣ Fault Injection Modules**  
| Failure Type | Method in Golang | Notes |
|-------------|------------------|------|
| **Process Killing** | `syscall.Kill(pid, syscall.SIGKILL)` | Kills a process by PID or name. |
| **CPU Stress** | `for {}` inside Goroutines | Spins up busy loops to consume CPU. |
| **Memory Leak Simulation** | `make([]byte, size)` | Allocates large memory blocks. |
| **Network Latency** | `tc qdisc add` via `exec.Command` | Uses Linux `tc` to add delay/loss. |
| **Disk Fill** | `f.Write(make([]byte, size))` | Writes large files to simulate full disk. |

---

## **📂 Folder Structure**  
```
hiveshake/
│── cmd/                 # CLI commands
│── main.go              # Entry point
│── go.mod               # Dependencies
```

---

🔥 **Next Steps:**  
- Add **Disk Space Filling**  
- Implement **Observability (Prometheus)**  
- Support **Kubernetes chaos testing**  

Would you like me to expand any specific feature next? 🚀