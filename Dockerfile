FROM debian:8.7

RUN /bin/echo "deb http://ftp.ru.debian.org/debian/ jessie main" > /etc/apt/sources.list \
  && /bin/echo "deb http://mirror.yandex.ru/debian jessie main" >> /etc/apt/sources.list

RUN apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates locales mysql-client \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8
EXPOSE 1323
ADD migrations /migrations
COPY statserver /
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh 

ENTRYPOINT [ "/entrypoint.sh" ]