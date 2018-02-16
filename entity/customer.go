package entity

import (
    "time"
)

type Customer struct {
    Customer_id int64
    Customer_cd string
    Eff_beg_dt string
    Eff_end_dt string
    Pwd string
    Pwd_lst_chg_dt string
    Customer_nm string
    Version_no int16
    Created_at time.Time
    Created_by string
    Updated_at time.Time
    Updated_by string
    Is_actived string
}
