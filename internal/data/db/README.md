# 1. MariaDB
Refer: https://hub.docker.com/\_/mariadb  
https://zhuanlan.zhihu.com/p/97035035
## 1.1. docker
```
docker pull mariadb
docker network create hfcms-mariadb
docker run --detach \
--publish 3306:3306 \
--env MARIADB_ROOT_PASSWORD=my-secret-pw \
--network hfcms-mariadb \
--name hfcms-mariadb \
mariadb:latest

docker exec -it hfcms-mariadb mysql -u root -pmy-secret-pw
```
MARIADB\_ROOT\_PASSWORD=`my-secret-pw`

## 1.2 mariadb

1.Change root password:
```
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '[newpassword]';
```
2.Create tables:
Database
```
DROP database hfcms_users;
CREATE database hfcms_users;
CREATE USER 'hfcms_users_user'@'%' IDENTIFIED BY 'hfcms_users_user_passwd';
GRANT ALL PRIVILEGES ON hfcms_users.* TO 'hfcms_users_user'@'%';
FLUSH PRIVILEGES;
USE hfcms_users;
```
Users
```
CREATE TABLE IF NOT EXISTS users (
  id int(10) NOT NULL AUTO_INCREMENT,
  username VARCHAR(255),
  password VARCHAR(255),
  realname VARCHAR(255),
  nickname VARCHAR(255),
  avatar_url VARCHAR(255),
  phone VARCHAR(11),
  user_ip INT(4) UNSIGNED,
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY (id)
);
desc users;
+-------------+-----------------+------+-----+---------------------+-------------------------------+
| Field       | Type            | Null | Key | Default             | Extra                         |
+-------------+-----------------+------+-----+---------------------+-------------------------------+
| id          | int(10)         | NO   | PRI | NULL                | auto_increment                |
| username    | varchar(255)    | YES  |     | NULL                |                               |
| password    | varchar(255)    | YES  |     | NULL                |                               |
| realname    | varchar(255)    | YES  |     | NULL                |                               |
| nickname    | varchar(255)    | YES  |     | NULL                |                               |
| avatar_url  | varchar(255)    | YES  |     | NULL                |                               |
| phone       | varchar(11)     | YES  |     | NULL                |                               |
| user_ip     | int(4) unsigned | YES  |     | NULL                |                               |
| state       | tinyint(1)      | YES  |     | NULL                |                               |
| deleted     | tinyint(1)      | YES  |     | 0                   |                               |
| create_time | timestamp       | NO   |     | current_timestamp() |                               |
| update_time | timestamp       | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+-----------------+------+-----+---------------------+-------------------------------+
``` 
Roles
```
CREATE TABLE IF NOT EXISTS roles (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
);
desc roles;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | int(10)      | NO   | PRI | NULL                | auto_increment                |
| parent_id   | int(10)      | YES  |     | NULL                |                               |
| code        | varchar(255) | YES  |     | NULL                |                               |
| name        | varchar(255) | YES  |     | NULL                |                               |
| description | varchar(255) | YES  |     | NULL                |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
```
Permissions
```
CREATE TABLE IF NOT EXISTS roles (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY (id)
);
```

SQL All in one
```
DROP database hfcms_users;
CREATE database hfcms_users;
CREATE USER 'hfcms_users_user'@'%' IDENTIFIED BY 'hfcms_users_user_passwd';
GRANT ALL PRIVILEGES ON hfcms_users.* TO 'hfcms_users_user'@'%';
FLUSH PRIVILEGES;
USE hfcms_users;

CREATE TABLE IF NOT EXISTS users (
  id int(10) NOT NULL AUTO_INCREMENT,
  username VARCHAR(255),
  password VARCHAR(255),
  realname VARCHAR(255),
  nickname VARCHAR(255),
  avatar_url VARCHAR(255),
  phone VARCHAR(11),
  user_ip INT(4) UNSIGNED,
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY (id)
);
CREATE TABLE IF NOT EXISTS roles (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY (id)
);
```
Validate
```
select host,
       user as username,
       password,
       password_expired
from mysql.user
order by user;
SHOW GRANTS FOR hfcms_users_user;
describe users;
describe roles;
```
