FROM mariadb:focal

ARG database
ARG password

RUN echo ${database}
RUN echo ${password}

MAINTAINER Felipe Moura

ENV MYSQL_DATABASE=${database} \
    MYSQL_ROOT_PASSWORD=${password}

EXPOSE 3306
