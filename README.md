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
    Description=Simple mysqlbackup system written on Go by Rishat Sultanov
    
    [Service]
    Type=simple
    Restart=always
    RestartSec=5s
    ExecStart=/home/vagrant/mysql-backup/main
    
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
    ● mysqlbackup.service - mysqlbackup
       Loaded: loaded (/lib/systemd/system/mysqlbackup.service; enabled; vendor preset: enabled)
       Active: active (running) since Wed 2018-12-06 21:00:01 UTC; 12min ago
     Main PID: 24735 (acj)
        Tasks: 3 (limit: 4915)
       Memory: 1.8M
          CPU: 20ms
       CGroup: /system.slice/mysqlbackup.service
               └─24735 /home/vagrant/mysql-backup
    
    Dec 06 21:00:01 serenity systemd[1]: Started mysqlbackup.
    
    ```
## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Rishats/ywpti/tags). 

## Authors

* **Rishat Sultanov** - [Rishats](https://github.com/Rishats)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
