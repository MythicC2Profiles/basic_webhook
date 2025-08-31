#!/bin/bash
if [[ "$EUID" -ne 0 ]]; then
   echo "This script must be run as root"
   exit 1
fi

WEBHOOKPATH="/home/kali/source/basic_webhook/C2_Profiles/webhook"  # Full path to Payload_Type/haunt, see my example

cd "/opt/Mythic" || exit 1
sudo /opt/Mythic/mythic-cli remove basic_webhook
sudo rm -rf /opt/Mythic/InstalledServices/basic_webhook
sudo cp -r "$WEBHOOKPATH" /opt/Mythic/InstalledServices
sudo ./mythic-cli add webhook
sudo ./mythic-cli build webhook