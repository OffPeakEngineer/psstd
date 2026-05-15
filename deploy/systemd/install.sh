#!/usr/bin/env sh
set -eu

install -m 0755 psstd /usr/local/bin/psstd

if ! id psstd >/dev/null 2>&1; then
  useradd --system --home /var/lib/psstd --shell /usr/sbin/nologin psstd
fi

install -d -m 0755 /etc/psstd
install -d -m 0755 -o psstd -g psstd /var/lib/psstd

if [ ! -f /etc/psstd/psstd.env ]; then
  install -m 0644 deploy/systemd/psstd.env /etc/psstd/psstd.env
fi

install -m 0644 deploy/systemd/psstd.service /etc/systemd/system/psstd.service
systemctl daemon-reload
systemctl enable --now psstd
systemctl status psstd --no-pager
