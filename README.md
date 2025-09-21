# TN-Container
A minimal container runtime implementation in Go that demonstrates the core concepts behind containerization technology. This project creates isolated environments using Linux namespaces, cgroups, and chroot, similar to how Docker and other container runtimes work under the hood.
🎯 What This Project Does
This simple container implementation provides:

Process isolation using Linux namespaces (PID, UTS, Mount, Network)
Filesystem isolation using chroot to create separate root filesystems
Resource limiting using cgroups (optional)
Basic container lifecycle management (create, start, stop)

🔧 How It Works
The container implementation uses several Linux kernel features:

Namespaces: Create isolated views of system resources

PID namespace: Isolated process tree
UTS namespace: Separate hostname and domain name
Mount namespace: Independent filesystem mount points
Network namespace: Isolated network stack


chroot: Changes the apparent root directory for the container process
cgroups: Limit and monitor resource usage (CPU, memory, etc.)
