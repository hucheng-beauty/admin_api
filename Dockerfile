FROM test

MAINTAINER hucheng_beauty@163.com

COPY admin-api /root/admin-api

RUN chmod 775 /root/admin-api

WORKDIR /root