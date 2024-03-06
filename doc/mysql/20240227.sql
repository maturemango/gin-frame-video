/* 视频表添加列名 */
alter table gf_video add is_del int not null default 0 comment '视频是否删除: 0:否 1:是';
alter table gf_video add upvote int not null default 0 comment '视频点赞数量';
alter table gf_video add disagree int not null default 0 comment '视频点踩数量';


/* 新建视频表 gf_video */
create table gf_video (
    id bigint not null auto_increment,
    created_time datetime not null default current_timestamp,
    updated_time datetime not null default current_timestamp on update current_timestamp,
    video_no varchar(50) not null comment '视频编号',
    user_id bigint not null comment '用户id',
    title varchar(50) not null comment '视频标题',
    introduction text null comment '视频简介',
    upload_time datetime not null comment '视频上传时间',
    primary key(id)
)