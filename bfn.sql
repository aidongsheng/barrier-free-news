create table article
(
  in_time varchar(50)   null,
  title   varchar(1000) null,
  author  varchar(100)  null,
  content text          null,
  href    varchar(1000) null
);

create table article_list
(
  al_time     varchar(50)   null,
  al_title    varchar(1000) null,
  al_ti_trans varchar(1000) null,
  al_href     varchar(1000) null
);

create table wcclog
(
  udid        varchar(100) null,
  record_date varchar(100) null,
  log_text    varchar(100) null,
  city_name   varchar(100) null
);

