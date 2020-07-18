FROM golang:1.14-alpine

ENV CURRENT_RCLONE_VERSION "v1.52.2"
ENV CURRENT_RCLONE_FLAVOR "linux-amd64"

RUN wget https://downloads.rclone.org/${CURRENT_RCLONE_VERSION}/rclone-${CURRENT_RCLONE_VERSION}-${CURRENT_RCLONE_FLAVOR}.zip
RUN unzip rclone-${CURRENT_RCLONE_VERSION}-${CURRENT_RCLONE_FLAVOR}.zip
RUN mv rclone-${CURRENT_RCLONE_VERSION}-${CURRENT_RCLONE_FLAVOR}/rclone /usr/local/bin/

COPY . /usr/local/src/rclone-backup
WORKDIR /usr/local/src/rclone-backup

RUN go build -o /usr/local/bin/rclone-backup github.com/ericflo/rclone-backup