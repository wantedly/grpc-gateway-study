#!/bin/sh

set -eu
set -o pipefail

if [ -d /host_ssh ]; then
  cp -r /host_ssh /root/.ssh
  chown -R root:root /root/.ssh
fi

"$@"
