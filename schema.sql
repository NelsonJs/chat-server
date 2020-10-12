create database `demo` default character set utf8mb4 collate utf8mb4_unicode_ci;

use demo;

-- 消息表
DROP TABLE IF EXISTS `msg`;
create table msg(
`id` bigint primary key auto_increment,
`uid` varchar(32) not null,
`peerid` varchar(32) not null,
`ctype` varchar(32) default '',
`msg_type` int default 0,
`content` varchar(255) default '',
`pic` varchar(255),
`status` int default 0 comment '1 发送成功 -1失败',
`create_time` bigint default 0
);

-- 注册用户表
DROP TABLE IF EXISTS `user`;
create table user(
`id`  bigint primary key auto_increment,
`uid` varchar(32) not null,
`nickname` varchar(255) default '',
`phone` varchar(24) default '',
`gender` int default 0,
`avatar` varchar(255) default '',
`create_time` bigint default 0 comment '注册账号时间',
`login_time` bigint default 0 comment '登录时间',
`logout_time` bigint default 0 comment '登出时间',
`status` int default 0 comment '1-封禁'
);

-- 好友表
DROP TABLE IF EXISTS `friends`;
create table friends(
`id` bigint primary key auto_increment,
`uid` varchar(32) not null,
`friend_id` varchar(255) not null,
`group` varchar(24) default '' comment '分组'
);

-- 群组表
DROP TABLE IF EXISTS `cgroup`;
CREATE TABLE cgroup(
    `id` bigint primary key auto_increment,
    `groupid` varchar(32) default '',
    `name` varchar(24) default '',
    `intro` varchar(255) default '',
    `avatar` varchar(255) default '',
    `ownerid` varchar(32) default '',
    `helpers` json,
    `members` json,
    `grouptype` int default 0,
    `status` int default 0,
    `apply` int default 0,
    `max` int default 0,
    `maxhelper` int default 0
);


-- 个人介绍表
DROP TABLE IF EXISTS `intro`;
create table intro(
`id` bigint primary key auto_increment,
`uid` varchar(32) not null,
`nickname` varchar(16) default '',
`avatar` varchar(255) default '',
`gender` int default 0,
`years_old` int default 0,
`habit` varchar(255) default '' comment '爱好',
`jiguan` varchar(255) default '' comment '籍贯',
`curlocal` varchar(255) default '' comment '当前所在地',
`xueli` varchar(30) default '' comment '学历',
`job` varchar(16) default '' comment '职业',
);

-- 图片资源表
DROP TABLE IF EXISTS `imgs`;
create table imgs (
`id` bigint primary key auto_increment,
`path` varchar(255) not null
);
