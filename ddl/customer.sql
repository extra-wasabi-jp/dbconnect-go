create table customer (
    customer_id integer primary key autoincrement not null
    , customer_cd varchar(50) not null
    , eff_beg_dt varchar(8) not null
    , eff_end_dt varchar(8) not null
    , pwd varchar(255) not null
    , pwd_lst_chg_dt varchar(8) not null
    , customer_nm varchar(100) not null
    , version_no integer not null
    , created_at timestamp not null
    , created_by varchar(100) not null
    , updated_at timestamp not null
    , updated_by varchar(100) not null
    , is_actived varchar(1) not null
);
