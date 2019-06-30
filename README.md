# MySQL Backup

Simple system which use mysqldump for dump mysql db and send info in to Telegram via [Horn](https://github.com/requilence/integram)

### Installing
1) Clone project
    ```
    git clone https://github.com/Rishats/mysql-backup.git
    ```
2) Change folder
    ```
    cd mysql-backup
    ```
3) Create .env file from .env.example
    ```
     cp .env.example .env
    ```

4) Configure your .env
    ```APP_ENV=production-or-other
       MYSQL_HOST=127.0.0.1
       MYSQL_PORT=3306
       MYSQL_DB=mydb
       MYSQL_USER=mydbuser
       MYSQL_PASSWORD=pwd
       BACKUP_DIR=/home/vagrant/backups/mysql/
       INTEGRAM_WEBHOOK_URI=your-uri
       SENTRY_DSN=your-dsn
       ```

### Running

Via go native:

Build for linux
```
env GOOS=linux GOARCH=amd64 go build main.go
```

### Creating a Service for Systemd
1) On Ubuntu VPS the following was sufficient to create a service after the go app was placed in home folder: /home/vagrant/mysql-backup
    ```
    touch /lib/systemd/system/mysqlbackup.service
    ```
2) Inserted the following into the file through vim

    ```
    vim /lib/systemd/system/mysqlbackup.service
    ```
    ```
    [Unit]
    Description=Simple mysql-backup system written on Go by Rishat Sultanov
    
    [Service]
    Type=simple
    Restart=always
    RestartSec=5s
    WorkingDirectory=/home/vagrant/code/go/mysql-backup
    ExecStart=/home/vagrant/code/go/mysql-backup/go_build_main_go_linux
    
    [Install]
    WantedBy=multi-user.target
    ```

3) This allows you to start your binary/service/mysqlbackup with:
    ```
    service mysqlbackup start
    ```
4) To enable it on boot, type: (optional)
    ```
    service mysqlbackup enable
    ```
5) Don’t forget to check if everything’s cool through: (optional)
    ```
    service mysqlbackup status
    ```
    Example output:
    ```
    ● mysqlbackup.service - Simple mysql-backup system written on Go by Rishat Sultanov
       Loaded: loaded (/lib/systemd/system/mysqlbackup.service; disabled; vendor preset: enabled)
       Active: active (running) since Sun 2019-06-30 08:58:00 UTC; 1min 30s ago
     Main PID: 6418 (go_build_main_g)
        Tasks: 4
       Memory: 12.9M
          CPU: 154ms
       CGroup: /system.slice/mysqlbackup.service
               └─6418 /home/vagrant/code/go/mysql-backup/go_build_main_go_linux
    
    Jun 30 08:58:00 homestead systemd[1]: mysqlbackup.service: Service hold-off time over, scheduling restart.
    Jun 30 08:58:00 homestead systemd[1]: Stopped Simple mysql-backup system written on Go by Rishat Sultanov.
    Jun 30 08:58:00 homestead systemd[1]: Started Simple mysql-backup system written on Go by Rishat Sultanov.
    Jun 30 08:58:00 homestead go_build_main_go_linux[6418]: Output: 0
    Jun 30 08:58:01 homestead go_build_main_go_linux[6418]: &{200 OK 200 HTTP/2.0 2 0 map[Content-Length:[0] Content-Type:[text/plain; charset=utf-8] Date:[Su
    Jun 30 08:58:01 homestead go_build_main_go_linux[6418]: 2019-07-01 02:00:00 +0000 UTC
    
    ```
## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Rishats/ywpti/tags). 

## Authors

* **Rishat Sultanov** - [Rishats](https://github.com/Rishats)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
