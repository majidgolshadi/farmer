#!/bin/bash

cd /var/www
composer update
php5-fpm --allow-to-run-as-root --nodaemonize --fpm-config /etc/php5/fpm/php-fpm.conf &
nginx -g "daemon off;"