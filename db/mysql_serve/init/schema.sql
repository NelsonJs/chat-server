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
`pwd` varchar(24) default '',
`gender` int default 0,
`avatar` varchar(255) default '',
`create_time` bigint default 0 comment '注册账号时间',
`login_time` bigint default 0 comment '登录时间',
`logout_time` bigint default 0 comment '登出时间',
`status` int default 0 comment '1-封禁'
);
insert into user(uid, nickname, phone, gender, create_time, login_time)values('100','Mr Peng','18320944165',1,1602726388,1602726388);
insert into user(uid, nickname, phone, gender, create_time, login_time)values('101','Ms Wang','13798554429',2,1602827388,1602826388);
insert into user(uid, nickname, phone, gender, create_time, login_time)values('102','Ms Tong','18779411443',2,1602826388,1602926388);

-- 好友表
DROP TABLE IF EXISTS `friends`;
create table friends(
`id` bigint primary key auto_increment,
`uid` varchar(32) not null,
`friend_id` varchar(255) not null,
`fnickname` varchar(24) default '',
`group` varchar(24) default '' comment '分组'
);
insert into friends(uid, friend_id) values ('100','101');

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
`liked` tinyint(1) default 0 comment '0没有点赞 1点赞了',
`location` varchar(128) default '',
`lat` float default 0.0,
`lng` float default 0.0,
`createtime` int default 0,
`resimg` json,
`gender` int default 0,
`description` varchar(255) default ''
);

insert into dynamics(did,title, uid, nickname, avatar, likenum, liked, location, lat, lng, createtime, resimg, gender, description)
values ('asdfg','今天天气真好噢，大家一起出来玩~~','100','Mr Peng','',59,1,'厦门市湖里区',0.0,0.0,1602663648,NULL,1,'就在那个体育场');

-- 评论表
DROP TABLE IF EXISTS `comments`;
create table comments(
`id` bigint primary key auto_increment,
`did` varchar(32) not null comment '动态的id',
`cid` varchar(32) comment '评论id',
`reply` json comment '[{id:223,uid:100,nickname:mr peng,content:11,replyuid:101,replynickname:mr wang,likenum:43}]',
`content` varchar(255) default '',
`uid` varchar(32) not null comment '评论人的uid',
`nickname` varchar(32) default '' comment '评论人的昵称',
`likenum` int default 0 comment '点赞数量',
`status` int default 0 comment '状态 0正常',
`createtime` int default 0
);
insert into comments(did, cid,content, uid, nickname,createtime)values('asdfg','qqqs','这是留言','100','Mr Peng',1602663648);
insert into comments(did, cid,content, uid, nickname, reply,createtime) values ('asdfg','fegd','this is reply msg','101','MS Wang','[{"id":"feag","uid":"100","nickname":"mr peng","content":"回复了哈","replyuid":"101","replynickname":"mr wang","likenum":"43"}]',1602673648);
insert into comments(did, cid,content, uid, nickname,createtime) values ('asdfg','rtgd','this is reply msg','102','MS Tong',1602683688);
insert into comments(did, cid,content, uid, nickname,createtime) values ('asdfg','geg','this is reply msg','102','MS Tong',1602683688);
insert into comments(did, cid,content, uid, nickname,createtime) values ('asdfg','gewgw','你是真不错呀','102','MS Tong',1602683688);
insert into comments(did, cid,content, uid, nickname,createtime) values ('asdfg','vvda','很好很好','102','MS Tong',1602683688);
insert into comments(did, cid,content, uid, nickname,createtime) values ('asdfg','yrh','哈哈哈','102','MS Tong',1602683688);
insert into comments(did, cid,content, uid, nickname,createtime) values ('asdfg','qreb','不错噢','102','MS Tong',1602683688);
