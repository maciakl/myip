# myip

A simple command-line tool to print the default local outbound IP address of the machine. It also provides an option to print the public IP address by making a request to an external service (ipify.org).

## Usage

    Usage: myip [options]
    Options:
      -p, --public     Print the public IP address instead (requires internet access)
      -v, --version    Print version information and exit
      -h, --help       Print this message and exit

## Installation

### All platforms:

Install via `go`:

    go install github.com/maciakl/myip@latest

### Linux

On Linux (requires `wget` & `unzip`, installs to `/usr/local/bin`):

    p="myip" && wget -qN "https://github.com/maciakl/${p}/releases/latest/download/${p}_lin.zip" && unzip -oq ${p}_lin.zip && rm -f ${p}_lin.zip && chmod +x ${p} && sudo mv ${p} /usr/local/bin

To uninstall, simply delete it:

    rm -f /usr/local/bin/myip

### Windows
 
On Windows, this tool is also distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

Next simply run:
 
    scoop install myip
