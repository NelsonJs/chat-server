create database `demo` default character set utf8mb4 collate utf8mb4_unicode_ci;

use demo;

-- 消息表
DROP TABLE IF EXISTS `msg`;
create table msg(
`id` bigint primary key auto_increment,
`send_id` varchar(100) not null,
`receive_id` varchar(100) not null,
`group_id` varchar(100) default '',
`msg_type` int default 0,
`content` varchar(255) default '',
`status` int default 0 comment '0 发送成功 -1失败',
`create_time` bigint default 0
);

-- 用户表
DROP TABLE IF EXISTS `user`;
create table user(
`uid` bigint not null primary key auto_increment,
`nick_name` varchar(255) default '',
`phone` varchar(24) default '',
`gender` int default 0,
`years_old` int default 0,
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
`uid` varchar(255) not null,
`friend_id` varchar(255) not null,
`group` varchar(24) default '' comment '分组'
);

-- 活动表
DROP TABLE IF EXISTS `active`;
create table active(
`id` bigint primary key auto_increment,
`uid` bigint default 0,
`title` varchar(255) default '',
`description` varchar(255) default '',
`img` varchar(255) default '',
`gender` int default 0 comment '0-不限 1-男 2-女',
`begin` bigint default 0,
`loc` varchar(255) default '',
`lng` float default 0,
`lat` float default 0,
`peopel_num` int default 0,
`people_total_num` int default 0,
`like` int default 0,
`comment_num` bigint default 0,
`comment_id` bigint default 0
);

-- 动态表
DROP TABLE IF EXISTS `dynamic`;
create table dynamic(
`id` bigint primary key auto_increment,
`uid` bigint default 0,
`title` varchar(255) default '',
`img` varchar(255) default '',
`like` int default 0,
`comment_id` int default 0,
`loc` varchar(255) default '',
`lat` float default 0,
`lng` float default 0,
`time` bigint default 0,
`res_img` json default null
);

-- 个人介绍表
DROP TABLE IF EXISTS `intro`;
create table intro(
`id` bigint primary key auto_increment,
`uid` bigint not null,
`nickname` varchar(16) default '',
`img_path` varchar(255) default '',
`gender` int default 0,
`years_old` int default 0,
`habit` varchar(255) default '' comment '爱好',
`jiguan` varchar(255) default '' comment '籍贯',
`curlocal` varchar(255) default '' comment '当前所在地',
`xueli` varchar(30) default '' comment '学历',
`job` varchar(16) default '' comment '职业',
`shengao` int default 0 comment '身高',
`tizhong` int default 0 comment '体重',
`love_word` varchar(255) default '' comment '恋爱宣言'
);

-- 图片资源表
DROP TABLE IF EXISTS `imgs`;
create table imgs (
`id` bigint primary key auto_increment,
`path` varchar(255) not null
);
