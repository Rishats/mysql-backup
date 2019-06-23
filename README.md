# MySQL Backup

Simple system which use mysqldump for dump mysql db and send info in to Telegram via [Horn](https://github.com/requilence/integram)

### Installing
1 Clone project
```
git clone https://github.com/Rishats/mysql-backup.git
```
2 Change folder
```
cd mysql-backup
```
3 Create .env file from .env.example
```
 cp .env.example .env
```

4 Configure your .env
```APP_ENV=production-or-other
   MYSQL_HOST=127.0.0.1
   MYSQL_PORT=3306
   MYSQL_DB=mydb
   MYSQL_USER=mydbuser
   MYSQL_PASSWORD=pwd
   BACKUP_DIR=~/backups/mysql/
   INTEGRAM_WEBHOOK_URI=your-uri
   SENTRY_DSN=your-dsn
   ```

### Running

Via go native:

```
go run main.go
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Rishats/ywpti/tags). 

## Authors

* **Rishat Sultanov** - [Rishats](https://github.com/Rishats)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
