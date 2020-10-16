create database `demo` default character set utf8mb4 collate utf8mb4_unicode_ci;

use demo;

-- 消息表
DROP TABLE IF EXISTS `msg`;
create table msg(
`id` bigint primary key auto_increment,
`uid` varchar(32) not null,
`nickname` varchar(24) default '',
`peerid` varchar(32) not null,
`ctype` varchar(32) default '' comment '1',
`msg_type` int default 0,
`content` varchar(255) default '',
`pic` varchar(255),
`status` int default 0 comment '1 发送成功 -1失败',
`create_time` bigint default 0
);
insert into msg(uid, nickname,peerid, ctype, msg_type, content, status, create_time)values('100','Mr Peng','101',1,1,'the weather is very good today!',1,1602726388);
insert into msg(uid, nickname,peerid, ctype, msg_type, content, status, create_time)values('100','Ms Wang','102',1,1,'it is very funny!',1,1602726388);
insert into msg(uid, nickname,peerid, ctype, msg_type, content, status, create_time)values('100','测试群','103',2,1,'do you want fishing?',1,1602726388);

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
insert into user(uid, nickname, phone, gender, create_time, login_time)values('100','Mr Peng','18320944165',1,1602726388,1602726388);

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
`job` varchar(16) default '' comment '职业'
);


-- 图片资源表
DROP TABLE IF EXISTS `imgs`;
create table imgs (
`id` bigint primary key auto_increment,
`path` varchar(255) not null
);

-- 动态表
DROP TABLE IF EXISTS `dynamics`;
CREATE  TABLE dynamics(
`id` bigint primary key auto_increment,
`did` varchar(32) not null default '',
`title` varchar(255) default '',
`uid` varchar(32) not null,
`nickname` varchar(16) default '',
`avatar` varchar(255) default '',
`likenum` int default 0,
`location` varchar(128) default '',
`lat` float default 0.0,
`lng` float default 0.0,
`createtime` int default 0,
`resimg` json,
`gender` int default 0,
`description` varchar(255) default ''
);

insert into dynamics(title, uid, nickname, avatar, likenum, commentid, location, lat, lng, createtime, resimg, gender, description)
values ('今天天气真好噢，大家一起出来玩~~','100','Mr Peng','',59,'','厦门市湖里区',0.0,0.0,1602663648,NULL,1,'就在那个体育场');

-- 评论表
DROP TABLE IF EXISTS `comments`;
create table comments(
`id` bigint primary key auto_increment,
`did` varchar(32) not null,
`content` varchar(255) default '',
`uid` varchar(32) not null,
`ateuid` varchar(32) default '',
`nickname` varchar(32) default '',
`atenickname` varchar(32) default '',
`createtime` int default 0
);
insert into comments(did, content, uid, nickname,createtime)values('asdfg','这是留言','100','Mr Peng',1602663648);
insert into comments(did, content, uid, ateuid, nickname, atenickname,createtime) values ('asdfg','this is reply msg','101','100','MS Wang','Mr Peng',1602673648);