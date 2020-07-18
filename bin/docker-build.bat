cd "%~dp0/.."

docker build -t ericflo/rclone-backup:latest .
docker push ericflo/rclone-backup:latest