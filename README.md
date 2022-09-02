# Url shortener

Url shortener backend written in golang and using redis as database.

The saved data has one month ttl.

It also uses docker compose to bootstrap and manage the containers.

shortener.ly is an alias for localhost. Add `127.0.0.1 shortener.ly` to your `/etc/hosts` file.