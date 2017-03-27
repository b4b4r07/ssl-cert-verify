#!/bin/bash

# within two weeks
danger="${1:-14}"
site="zplug.sh"

if ! type ssl-cert-verify &>/dev/null; then
    echo "ssl-cert-verify: not found in PATH" >&2
    exit 1
fi

if [[ ! $danger =~ ^[0-9]+$ ]]; then
    echo "$danger: not integer" >&2
    exit 1
fi

result="$(ssl-cert-verify "$site:443")"
days="$(perl -pe 's/^.*in ([0-9]+) days.*$/$1/' <<<"$result")"

if (( $days > $danger )); then
    # no prob
    exit 0
fi

echo "$result" >&2
exit 1
