#!/usr/bin/env sh
set -eu

install -m 0755 pulsed /usr/local/bin/pulsed

if ! id pulsed >/dev/null 2>&1; then
  useradd --system --home /var/lib/pulsed --shell /usr/sbin/nologin pulsed
fi

install -d -m 0755 /etc/pulsed
install -d -m 0755 -o pulsed -g pulsed /var/lib/pulsed

if [ ! -f /etc/pulsed/pulsed.env ]; then
  install -m 0644 deploy/systemd/pulsed.env /etc/pulsed/pulsed.env
fi

install -m 0644 deploy/systemd/pulsed.service /etc/systemd/system/pulsed.service
systemctl daemon-reload
systemctl enable --now pulsed
systemctl status pulsed --no-pager
