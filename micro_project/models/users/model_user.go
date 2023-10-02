package users

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type T_user_template struct {
	Id                     string    `json:"id"`
	Create_date            time.Time `json:"create_date"`
	Modify_date            time.Time `json:"modify_date"`
	User_name              string    `json:"user_name"`
	Register_channels      string    `json:"register_channels"`
	Nick_name              string    `json:"nick_name"`
	User_id                string    `json:"user_id"`
	Cust_no                string    `json:"cust_no"`
	Mobile                 string    `json:"mobile"`
	Email                  string    `json:"email"`
	User_pic               string    `json:"user_pic"`
	Acct_msg_status        string    `json:"acct_msg_status"`
	Is_locked              string    `json:"is_locked"`
	Potential_user_lv      string    `json:"potential_user_lv"`
	Potential_user_type    string    `json:"potential_user_type"`
	Potential_user_desc    string    `json:"potential_user_desc"`
	Login_password         string    `json:"login_password"`
	Update_pwd_time        time.Time `json:"update_pwd_time"`
	Figure_password        string    `json:"figure_password"`
	Gesture_password       string    `json:"gesture_password"`
	Passcode               string    `json:"passcode"`
	Is_figure_open         string    `json:"is_figure_open"`
	Is_face_recog_open     string    `json:"is_face_recog_open"`
	Is_gesture_login_open  string    `json:"is_gesture_login_open"`
	Is_hide_gesture_login  string    `json:"is_hide_gesture_login"`
	Is_passcode_open       string    `json:"is_passcode_open"`
	Device_token           string    `json:"device_token"`
	Device_no              string    `json:"device_no"`
	User_barcode           string    `json:"user_barcode"`
	Is_email_login         string    `json:"is_email_login"`
	Email_auth             string    `json:"email_auth"`
	Language               string    `json:"language"`
	Need_change_pwd        string    `json:"need_change_pwd"`
	My_referral_code       string    `json:"my_referral_code"`
	Catch_recover_time     time.Time `json:"catch_recover_time"`
	Fail_potential_time    time.Time `json:"fail_potential_time"`
	Esign_kid              string    `json:"esign_kid"`
	Esign_ref_number       string    `json:"esign_ref_number"`
	Is_notification        string    `json:"is_notification"`
	Is_email_noti          string    `json:"is_email_noti"`
	Is_sms_noti            string    `json:"is_sms_noti"`
	Is_push_noti           string    `json:"is_push_noti"`
	Acct_mask              string    `json:"acct_mask"`
	Del                    string    `json:"del"`
	Space1                 string    `json:"space1"`
	Space2                 string    `json:"space2"`
	Space3                 string    `json:"space3"`
	Space4                 string    `json:"space4"`
	Space5                 string    `json:"space5"`
	Space6                 string    `json:"space6"`
	CustomerStatus         string    `json:"customerStatus"`
	Ndid_open_time         time.Time `json:"ndid_open_time"`
	Reg_pdf_par            string    `json:"reg_pdf_par"`
	Channel_sign           string    `json:"channel_sign"`
	Edc_open_time          time.Time `json:"edc_open_time"`
	Edc_qr_data            string    `json:"edc_qr_data"`
	Cust_journey_reference string    `json:"cust_journey_reference"`
	Channel_info_sign      string    `json:"channel_info_sign"`
	Kyc_risk_level         string    `json:"kyc_risk_level"`
	Ocr_toggle_status      string    `json:"ocr_toggle_status"`
	Test                   string    `json:"test"`
}

type User_data struct {
	Id          string    `json:"id"`
	Create_date time.Time `json:"create_date"`
	Modify_date time.Time `json:"modify_date"`
	User_name   string    `json:"user_name"`
	Nick_name   string    `json:"nick_name"`
	User_id     string    `json:"user_id"`
	Cust_no     string    `json:"cust_no"`
	Mobile      string    `json:"mobile"`
	Email       string    `json:"email"`
	User_pic    string    `json:"user_pic"`
	Is_locked   string    `json:"is_locked"`
	Space1      string    `json:"space1"`
	Space2      string    `json:"space2"`
	Space3      string    `json:"space3"`
	Space4      string    `json:"space4"`
	Space5      string    `json:"space5"`
	Space6      string    `json:"space6"`
}

type User_login struct {
	Id          string    `json:"id"`
	Create_date time.Time `json:"create_date"`
	Modify_date time.Time `json:"modify_date"`
	User_name   string    `json:"user_name"`
	User_id     string    `json:"user_id"`
	Mobile      string    `json:"mobile"`
	Email       string    `json:"email"`
}
