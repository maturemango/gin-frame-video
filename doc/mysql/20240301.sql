/* 视频表添加观看数量字段 */
alter table gf_video add watch int not null default 0 comment '观看(播放)数量';