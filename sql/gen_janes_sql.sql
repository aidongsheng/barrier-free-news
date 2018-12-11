create table wcc.janes_article
(
  in_time varchar(50)   null,
  title   varchar(1000) null,
  author  varchar(200)  null,
  href    varchar(500)  null,
  constraint janes_article_href_uindex
    unique (href)
);

create table wcc.janes_detail
(
  title   varchar(1000) null,
  content text          null,
  href    varchar(1000) null
);

