#!/bin/sh

# Manage rpot privileges
if ! [ $(id -u) = 0 ]; then
   echo "[ERROR] The script need to be run as root." >&2
   exit 1
fi

if [ $SUDO_USER ]; then
    REAL_USER=$SUDO_USER
else
    REAL_USER=$(whoami)
fi

# Variables
SUDO_HOME=/home/$SUDO_USER
INSTALL_DIR=$SUDO_HOME/.remotelab
BIN_DIR=/usr/bin
SERVICE_DIR=/etc/systemd/system

# Copy installation files to installation location
sudo -u $REAL_USER mkdir -p $INSTALL_DIR
sudo -u $REAL_USER cp docker-compose.yml $INSTALL_DIR
sudo -u $REAL_USER cp Dockerfile $INSTALL_DIR
sudo -u $REAL_USER cp -r src/ $INSTALL_DIR
cp daemon/launch.sh $BIN_DIR/remotelab.sh
sed "s~WORKING_DIR~$SUDO_HOME/.remotelab~g" daemon/remotelab.service > $SERVICE_DIR/remotelab.service

# Install 
cd $INSTALL_DIR
sudo -u $REAL_USER docker compose build
chmod +x $BIN_DIR/remotelab.sh
systemctl enable --now remotelab.service