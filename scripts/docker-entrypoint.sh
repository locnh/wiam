#!/bin/sh

# Main program
/whoami

# Option to download browscap.ini file
if [ -z $FULLBC ]; then
    query="Full_BrowsCapINI"
else
    query="BrowsCapINI"
fi

# Download to browscap.ini
wget https://browscap.org/stream?q=${query} -O browscap.ini
