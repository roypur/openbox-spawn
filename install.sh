#!/bin/bash
sudo mkdir -p /var/openbox-spawn
sudo wget -O /var/openbox-spawn/wrapper "https://raw.githubusercontent.com/roypur/openbox-spawn/master/bin/wrapper"
sudo wget -O /var/openbox-spawn/spawn "https://raw.githubusercontent.com/roypur/openbox-spawn/master/bin/spawn"
sudo wget -O /usr/share/xsessions/openbox-spawn.desktop "https://raw.githubusercontent.com/roypur/openbox-spawn/master/openbox-spawn.desktop"
sudo chmod --recursive 755 /var/openbox-spawn
if [ ! -f /var/openbox-spawn/config.json ]; then
    sudo wget -O /var/openbox-spawn/config.json "https://raw.githubusercontent.com/roypur/openbox-spawn/master/config.json"
else
    echo "Config file exists. Skipping config file"
fi
