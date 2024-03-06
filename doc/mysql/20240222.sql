/* 新建视频弹幕表 gf_video_caption */
create table gf_video_caption (
    id bigint not null auto_increment,
    created_time datetime not null default current_timestamp,
    updated_time datetime not null default current_timestamp on update current_timestamp,
    video_no varchar(20) not null comment '视频编号',
    user_id bigint not null comment '用户id',
    send_time datetime not null comment '弹幕发送时间',
    caption varchar(150) not null comment '弹幕内容',
    second int not null comment '视频时间节点',
    primary key(id)
)