/* 用户表修改用户名字段的属性 */
alter table gf_user modify user_name varchar(50) null comment '用户名';

/* 用户表添加手机号字段 */
alter table gf_user add account varchar(50) not null comment '用户账号';

/* 新建用户表 gf_user */
create table gf_user (
    id bigint not null auto_increment,
    created_time datetime not null default current_timestamp,
    updated_time datetime not null default current_timestamp on update current_timestamp,
    user_name varchar(50) not null comment '用户名',
    password varchar(255) not null comment '用户密码',
    primary key(id)
)