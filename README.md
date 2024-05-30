<h1 style="font-family:monospace">GO.DO</h1>
<div style="padding-bottom:20px">
    <img src="https://img.shields.io/badge/go-1.22.0-blue" />
    <img src="https://img.shields.io/badge/redis-7.2.4-red" />
    <img src="https://img.shields.io/badge/license-GPL%20v3-grey" />
</div>

<!-- DESCRIPTION -->
<h2 style="font-family:monospace">Description</h2>
<p style="font-family:monospace">
Go + Redis implementation of a todo list web application.<br>
Lightning fast, lightweight, and designed with excellent memory management.
</p>



<!-- INSTALLATION -->
<h2 style="font-family:monospace">Installation</h2>

<h4 style="font-family:monospace">Option 1: Docker</h4>

```bash
# 1) Install project
git clone https://github.com/ttamre/go.do.git
cd go.do

# 2) Build docker image
docker build -t go.do .
```

<h4 style="font-family:monospace">Option 2: Makefile</h4>

```bash
# 1) Install redis
# https://redis.io/downloads/

# 2) Run redis server in background
redis-server --port 5000 &

# 3) Install project
git clone https://github.com/ttamre/go.do.git
cd go.do

# 4) Build project
make deps       # get dependencies
make build      # build binaries
```

<!-- USAGE -->
<h2 style="font-family:monospace">Usage</h2>

```bash
docker run -p 5000:5000 go.do   # Run docker image
./bin/godo                      # Run binary
```

<!-- LICENSE -->
<h2 style="font-family:monospace">License</h2>
<p style="font-family:monospace">This project is licensed under the GNU v3 General Public License. For more information, see the <a href="LICENSE">LICENSE</a></p>

<img src="demo.png">