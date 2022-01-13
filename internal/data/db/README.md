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
``` 
Roles
```
CREATE TABLE IF NOT EXISTS roles (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
```
Role Users 
```
CREATE TABLE IF NOT EXISTS role_users (
  id int(10) NOT NULL AUTO_INCREMENT,
  role_id INT(10),
  user_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
```
Permissions
```
CREATE TABLE IF NOT EXISTS permissions (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY (id)
);
```
Role Permissions
```
CREATE TABLE IF NOT EXISTS role_permissions (
  id int(10) NOT NULL AUTO_INCREMENT,
  role_id INT(10),
  permission_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
```
Usergroups
```
CREATE TABLE IF NOT EXISTS usergroups (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
```
Usergroup Users
```
CREATE TABLE IF NOT EXISTS usergroup_users (
  id int(10) NOT NULL AUTO_INCREMENT,
  usergroup_id INT(10),
  user_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
```
Role Usergroups
```
CREATE TABLE IF NOT EXISTS role_usergroups (
  id int(10) NOT NULL AUTO_INCREMENT,
  role_id INT(10),
  usergroup_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
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
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS role_users (
  id int(10) NOT NULL AUTO_INCREMENT,
  role_id INT(10),
  user_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS permissions (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY (id)
);
CREATE TABLE IF NOT EXISTS role_permissions (
  id int(10) NOT NULL AUTO_INCREMENT,
  role_id INT(10),
  permission_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS usergroups (
  id int(10) NOT NULL AUTO_INCREMENT,
  parent_id INT(10),
  code VARCHAR(255),
  name VARCHAR(255),
  description VARCHAR(255),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS usergroup_users (
  id int(10) NOT NULL AUTO_INCREMENT,
  usergroup_id INT(10),
  user_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS role_usergroups (
  id int(10) NOT NULL AUTO_INCREMENT,
  role_id INT(10),
  usergroup_id INT(10),
  state TINYINT(1) COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
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
desc role_users;
desc permissions;
desc role_permissions;
desc usergroups;
desc usergroup_users;
desc role_usergroups;
```
![用户中心—完整版](https://pic2.zhimg.com/v2-4ffca73291dcf129cd466a4cb604865d_r.jpg "用户中心—完整版")
Output:
```
+-----------+---------------------+-------------------------------------------+------------------+
| Host      | username            | Password                                  | password_expired |
+-----------+---------------------+-------------------------------------------+------------------+
| %         | hfcms_articles_user | *7B5855D2183D391D34685174C1A975CD23979B17 | N                |
| %         | hfcms_users_user    | *CBF9C118375B4DBFCE57493A331215C4F059F702 | N                |
| localhost | mariadb.sys         |                                           | Y                |
| localhost | root                | *27C01464AD101AED1E65AC21152499A396B4CF72 | N                |
| %         | root                | *27C01464AD101AED1E65AC21152499A396B4CF72 | N                |
+-----------+---------------------+-------------------------------------------+------------------+
+-----------------------------------------------------------------------------------------------------------------+
| Grants for hfcms_users_user@%                                                                                   |
+-----------------------------------------------------------------------------------------------------------------+
| GRANT USAGE ON *.* TO `hfcms_users_user`@`%` IDENTIFIED BY PASSWORD '*CBF9C118375B4DBFCE57493A331215C4F059F702' |
| GRANT ALL PRIVILEGES ON `hfcms_users`.* TO `hfcms_users_user`@`%`                                               |
+-----------------------------------------------------------------------------------------------------------------+
describe users;
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
describe roles;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | int(10)      | NO   | PRI | NULL                | auto_increment                |
| parent_id   | int(10)      | YES  |     | NULL                |                               |
| code        | varchar(255) | YES  |     | NULL                |                               |
| name        | varchar(255) | YES  |     | NULL                |                               |
| description | varchar(255) | YES  |     | NULL                |                               |
| state       | tinyint(1)   | YES  |     | NULL                |                               |
| deleted     | tinyint(1)   | YES  |     | 0                   |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
desc role_users;
+-------------+------------+------+-----+---------------------+-------------------------------+
| Field       | Type       | Null | Key | Default             | Extra                         |
+-------------+------------+------+-----+---------------------+-------------------------------+
| id          | int(10)    | NO   | PRI | NULL                | auto_increment                |
| role_id     | int(10)    | YES  |     | NULL                |                               |
| user_id     | int(10)    | YES  |     | NULL                |                               |
| state       | tinyint(1) | YES  |     | NULL                |                               |
| deleted     | tinyint(1) | YES  |     | 0                   |                               |
| update_time | timestamp  | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+------------+------+-----+---------------------+-------------------------------+
desc permissions;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | int(10)      | NO   | PRI | NULL                | auto_increment                |
| parent_id   | int(10)      | YES  |     | NULL                |                               |
| code        | varchar(255) | YES  |     | NULL                |                               |
| name        | varchar(255) | YES  |     | NULL                |                               |
| description | varchar(255) | YES  |     | NULL                |                               |
| state       | tinyint(1)   | YES  |     | NULL                |                               |
| deleted     | tinyint(1)   | YES  |     | 0                   |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
desc role_permissions;
+---------------+------------+------+-----+---------------------+-------------------------------+
| Field         | Type       | Null | Key | Default             | Extra                         |
+---------------+------------+------+-----+---------------------+-------------------------------+
| id            | int(10)    | NO   | PRI | NULL                | auto_increment                |
| role_id       | int(10)    | YES  |     | NULL                |                               |
| permission_id | int(10)    | YES  |     | NULL                |                               |
| state         | tinyint(1) | YES  |     | NULL                |                               |
| deleted       | tinyint(1) | YES  |     | 0                   |                               |
| update_time   | timestamp  | NO   |     | current_timestamp() | on update current_timestamp() |
+---------------+------------+------+-----+---------------------+-------------------------------+
desc usergroups;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | int(10)      | NO   | PRI | NULL                | auto_increment                |
| parent_id   | int(10)      | YES  |     | NULL                |                               |
| code        | varchar(255) | YES  |     | NULL                |                               |
| name        | varchar(255) | YES  |     | NULL                |                               |
| description | varchar(255) | YES  |     | NULL                |                               |
| state       | tinyint(1)   | YES  |     | NULL                |                               |
| deleted     | tinyint(1)   | YES  |     | 0                   |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
desc usergroup_users;
+--------------+------------+------+-----+---------------------+-------------------------------+
| Field        | Type       | Null | Key | Default             | Extra                         |
+--------------+------------+------+-----+---------------------+-------------------------------+
| id           | int(10)    | NO   | PRI | NULL                | auto_increment                |
| usergroup_id | int(10)    | YES  |     | NULL                |                               |
| user_id      | int(10)    | YES  |     | NULL                |                               |
| state        | tinyint(1) | YES  |     | NULL                |                               |
| deleted      | tinyint(1) | YES  |     | 0                   |                               |
| update_time  | timestamp  | NO   |     | current_timestamp() | on update current_timestamp() |
+--------------+------------+------+-----+---------------------+-------------------------------+
desc role_usergroups;
+--------------+------------+------+-----+---------------------+-------------------------------+
| Field        | Type       | Null | Key | Default             | Extra                         |
+--------------+------------+------+-----+---------------------+-------------------------------+
| id           | int(10)    | NO   | PRI | NULL                | auto_increment                |
| role_id      | int(10)    | YES  |     | NULL                |                               |
| usergroup_id | int(10)    | YES  |     | NULL                |                               |
| state        | tinyint(1) | YES  |     | NULL                |                               |
| deleted      | tinyint(1) | YES  |     | 0                   |                               |
| update_time  | timestamp  | NO   |     | current_timestamp() | on update current_timestamp() |
+--------------+------------+------+-----+---------------------+-------------------------------+
```
