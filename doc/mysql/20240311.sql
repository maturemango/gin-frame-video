/* 用户表添加索引 phone_role */
create unique index phone_role gf_user(phone, role_id);

/* 修改用户表手机号字段 */
alter table gf_user change account phone varchar(20) null comment '用户手机号';