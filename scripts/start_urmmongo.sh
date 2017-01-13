#!/bin/bash
LOGDIR="/var/log/mhealth-urm-mongo"
USER="orange"
if [ ! -d "$LOGDIR" ]; then
	echo "Create log directory"
    sudo mkdir -p $LOGDIR
fi
    sudo chown -R $USER $LOGDIR
echo "Urmmongo process started."
./urmmongo &
