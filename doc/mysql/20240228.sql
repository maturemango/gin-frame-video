/* 创建用户视频三连操作表 gf_video_triplet */
create table gf_video_triplet (
  id bigint not null auto_increment,
  created_time datetime not null default current_timestamp,
  updated_time datetime not null default current_timestamp on update current_timestamp,
  video_no varchar(50) not null comment '视频编号',
  user_id bigint not null comment '用户id',
  is_upvote int not null default 0 comment '用户是否点赞: 0:否 1:是',
  is_disagree int not null default 0 comment '用户是否点踩: 0:否 1:是',
  is_coins int not null default 0 comment '用户是否投币: 0:否 1:是',
  is_collect int not null default 0 comment '用户是否收藏: 0:否 1:是',
  is_watch int not null default 0 comment '用户是否观看: 0:否 1:是',
  primary key(id)
)

/* 视频表gf_video添加字段 */
alter table gf_video add coins int not null default 0 comment '投币数量';
alter table gf_video add collect int not null default 0 comment '收藏数量';