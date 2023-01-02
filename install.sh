#!/bin/bash

# Manage rpot privileges
if ! [ $(id -u) = 0 ]; then
   echo "The script need to be run as root." >&2
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
echo -n "Create folders... "
sudo -u $REAL_USER mkdir -p $INSTALL_DIR
echo "OK"
echo -n "Copy files... "
sudo -u $REAL_USER cp docker-compose.yml $INSTALL_DIR
sudo -u $REAL_USER cp Dockerfile $INSTALL_DIR
sudo -u $REAL_USER cp -r src/ $INSTALL_DIR
cp daemon/launch.sh $BIN_DIR/remotelab.sh
sed "s~WORKING_DIR~$SUDO_HOME/.remotelab~g" daemon/remotelab.service > $SERVICE_DIR/remotelab.service
echo "OK"

# Install 
echo -n "Build docker image... "
cd $INSTALL_DIR
sudo -u $REAL_USER docker compose build &>/dev/null
RET_VAL=$?
if [ $RET_VAL -ne 0 ]; then
    echo "ERROR"
    exit $RET_VAL
fi
echo "OK"
echo -n "Enable systemctl service... "
chmod +x $BIN_DIR/remotelab.sh
systemctl enable --now remotelab.service
echo "OK"

echo "Installation successful"