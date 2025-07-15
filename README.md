# malleus - In progress
Malleus aims to be a basic server configuration manager. It aims to automate deployment of [Incus](https://github.com/lxc/incus) with OVN.

**NOTE**: This is a toy project for me to learn Golang :D


## Planned Requirements
1. Record server hardware: primarly nics and storage
2. Assign configuration to servers (IPs, partitioning)
3. Generate deployment ISOs (probably Debian preseed)
4. Automation for networking, Incus, OVN deployment

## Optional Goals
1. OS agent for inventorying, configuration deployment features
2. Configuration revision tracking
3. Mirror environment (e.g., preproduction) creation and deployment
4. Cloud integration