/* 创建系统操作日志表 gf_log */
create table gf_log (
  id bigint not null auto_increment,
  created_time datetime not null default current_timestamp,
  updated_time datetime not null default current_timestamp on update current_timestamp,
  user_id bigint not null comment '用户id',
  addr varchar(50) not null comment 'ip地址',
  log_level varchar(20) not null comment '日志等级',
  operate_time datetime not null comment '操作时间',
  operate_desc varchar(50) not null comment '操作描述',
  detail varchar(255) not null comment '操作详情',
  primary key(id)
)