create table Telegraph
(
  in_time     varchar(50)   null,
  title       varchar(1000) null,
  title_trans varchar(1000) null,
  href        varchar(1000) null
);

create table telegraph_detail
(
  in_time varchar(50)   null,
  title   varchar(1000) null,
  cotent  text          null,
  href    varchar(1000) null
);

