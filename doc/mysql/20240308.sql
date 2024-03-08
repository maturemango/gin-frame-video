/* 创建gf_role中角色数据 */
insert into gf_role (name, description) values ('超级管理员', '权限最高管理者');
insert into gf_role (name, description) values ('管理员', '普通管理者');
insert into gf_role (name, description) values ('普通用户', '没有管理权限');

/* 新建角色表 gf_role */
create table gf_role (
    id bigint not null auto_increment,
    created_time datetime not null default current_timestamp,
    updated_time datetime not null default current_timestamp on update current_timestamp,
    name varchar(20) not null comment '角色名称',
    description varchar(50) not null comment '角色描述',
    status int not null default 1 comment '角色状态: 0:不启用 1:启用',
    primary key(id)
)

/* 用户表gf_user添加角色字段 */
alter table gf_user add role_id int not null comment '角色id';