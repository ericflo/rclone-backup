# rclone-backup

This is a utility that uses rclone to periodically back up a folder on disk to
a cloud storage provider.

## Usage

Invoke the binary and pass in arguments that specify which folder to back up,
and how many copies should be stored:

```
# rclone-backup --source /usr/local/my-app/data --days 7
```

## Environment variables

The most important environment variable to set is `BACKUP_TARGET`, which
determines the target directory on the cloud storage provider. Set it to
something like `bucketname/projectname`, for example:

```
export BACKUP_TARGET="rclone-backups-maxint/siacdn-pro"
```

In order for rclone to know how to connect to your cloud storage provider,
you'll need to set up environment variables to configure a `backup` remote. 
Here's how that would look for B2:

```
export RCLONE_CONFIG_BACKUP_TYPE=b2
export RCLONE_CONFIG_BACKUP_ACCOUNT=XXX
export RCLONE_CONFIG_BACKUP_KEY=XXX
```

Here's how it might look for S3:

```
export RCLONE_CONFIG_BACKUP_TYPE=s3
export RCLONE_CONFIG_BACKUP_ACCESS_KEY_ID=XXX
export RCLONE_CONFIG_BACKUP_SECRET_ACCESS_KEY=XXX
```