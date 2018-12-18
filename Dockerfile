FROM  ubuntu:latest
MAINTAINER  yuanjun.zeng
RUN mkdir /data
ADD ./ZCache /data/
EXPOSE "8005"
CMD  /data/ZCache