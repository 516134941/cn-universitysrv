-- university 大学表
CREATE table university(
    id SERIAL primary key,              --主键
    province text not null DEFAULT '',  --省份
    city text not null DEFAULT '',      --城市
    name text not null DEFAULT '',       --大学名称
    UNIQUE(city,name)
);
CREATE INDEX university_province_index ON university(province);    --添加城市索引
