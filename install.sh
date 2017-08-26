#!/bin/bash

if [[ $EUID -ne 0 ]]
then
    echo "This script needs to be run as root."
else
    mkdir -p /var/openbox-spawn
    curl --output /var/openbox-spawn/wrapper --url "https://raw.githubusercontent.com/roypur/openbox-spawn/master/bin/wrapper"
    curl --output /var/openbox-spawn/spawn --url "https://raw.githubusercontent.com/roypur/openbox-spawn/master/bin/spawn"
    curl --output /usr/share/xsessions/openbox-spawn.desktop --url "https://raw.githubusercontent.com/roypur/openbox-spawn/master/openbox-spawn.desktop"
    chmod --recursive 755 /var/openbox-spawn
    if [ ! -f /var/openbox-spawn/config.json ]; then
        curl --output /var/openbox-spawn/config.json --url "https://raw.githubusercontent.com/roypur/openbox-spawn/master/config.json"
    else
        echo "Config file exists. Skipping config file"
    fi
fi
