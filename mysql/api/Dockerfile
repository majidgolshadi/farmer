FROM nginx:latest

ENV DEBIAN_FRONTEND noninteractive

RUN \
    apt-get update -y && \
    apt-get install -y --force-yes git curl php5-fpm php5-cli php5-common php5-curl php5-intl php5-mysql

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN \
    curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN \
    sed -i "s/owner = www-data/owner = nginx/" /etc/php5/fpm/pool.d/www.conf && \
    sed -i "s/group = www-data/group = nginx/" /etc/php5/fpm/pool.d/www.conf && \
    sed -i -e "s/user www-data/user nginx/g" /etc/nginx/nginx.conf && \
    sed -i -e "s/www-data/nginx/g" /etc/php5/fpm/pool.d/www.conf && \
    sed -i -e "s/;clear_env = no/clear_env = no/g" /etc/php5/fpm/pool.d/www.conf && \
    sed -i -e "s/DAEMON_ARGS=\"/DAEMON_ARGS=\"--allow-to-run-as-root /g" /etc/init.d/php5-fpm

ADD ./nginx.conf /etc/nginx/conf.d/default.conf
ADD ./web /var/www
ADD ./start.sh /start.sh

EXPOSE 80

CMD ["/start.sh"]