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
    Version int16
    Reg_user_nm string
    Reg_tmst time.Time
    Mod_user_nm string
    Mod_tmst time.Time
    Is_active string
}
