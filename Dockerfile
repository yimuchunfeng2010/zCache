FROM  ubuntu:latest
MAINTAINER  yuanjun.zeng
RUN mkdir /data
ADD ./zCache /data/
EXPOSE "8005"
CMD  /data/zCache