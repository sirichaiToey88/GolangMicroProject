package customer

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"project/models/users"
	"project/other"
	"project/schemas"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
)

var User = []users.T_user_template{}

// func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
// 	rUser := []users.T_user_template{}
// 	db, err := schemas.Connectdb()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query(`SELECT id, nick_name, cust_no, user_name, email,mobile FROM t_user_template limit 50`)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var users users.T_user_template
// 		err := rows.Scan(&users.Id, &users.Nick_name, &users.Cust_no, &users.User_name, &users.Email, &users.Mobile)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
// 			return
// 		}
// 		rUser = append(User, users)
// 		w.WriteHeader(http.StatusOK)
// 		jsonResponse, err := json.Marshal(rUser)
// 		w.Write(jsonResponse)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Println(err)
// 		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
// 		return
// 	}
// }

// func Register(w http.ResponseWriter, r *http.Request) {
// 	var register struct {
// 		User_name string `json:"user_name"`
// 		Email     string `json:"email"`
// 		Mobile    string `json:"mobile"`
// 	}

// 	// Parse the request body to extract the data to be inserted
// 	err = json.NewDecoder(r.Body).Decode(&register)
// 	if err != nil {
// 		log.Println("Failed to parse request body:", err)
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	db, err := schemas.Connectdb()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Prepare the INSERT statement
// 	// stmt, err := db.Prepare("INSERT INTO t_user_template(id,create_date,modify_date,user_name,register_channels,nick_name,user_id,cust_no,mobile,email,user_pic,acct_msg_status,is_locked,potential_user_lv,potential_user_type,potential_user_desc,login_password,update_pwd_time,figure_password,gesture_password,passcode,is_figure_open,is_face_recog_open,is_gesture_login_open,is_hide_gesture_login,is_passcode_open,device_token,device_no,user_barcode,is_email_login,email_auth,language,need_change_pwd,my_referral_code,catch_recover_time,fail_potential_time,esign_kid,esign_ref_number,is_notification,is_email_noti,is_sms_noti,is_push_noti,acct_mask,del,space1,space2,space3,space4,space5,space6,customerStatus,ndid_open_time,reg_pdf_par,channel_sign,edc_open_time,edc_qr_data,cust_journey_reference,channel_info_sign,kyc_risk_level,ocr_toggle_status,test) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
// 	stmt, err := db.Prepare("INSERT INTO t_user_template(id,user_name,email,mobile,create_date,modify_date) VALUES(?,?,?,?,?,?);")
// 	if err != nil {
// 		log.Println("Failed to prepare statement:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer stmt.Close()

// 	// Execute the INSERT statement
// 	UUID := other.GenerateUUID()
// 	Createdate := other.DateTime()
// 	Modifydate := other.DateTime()
// 	result, err := stmt.Exec(UUID, register.Email, register.Mobile, register.User_name, Createdate, Modifydate)
// 	if err != nil {
// 		log.Println("Failed to execute statement:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Get the number of rows affected by the INSERT
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Println("Failed to get rows affected:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the response indicating the number of rows affected
// 	response := struct {
// 		RowsAffected int64 `json:"rows_affected"`
// 	}{
// 		RowsAffected: rowsAffected,
// 	}

// 	// Set the Content-Type header to JSON
// 	w.Header().Set("Content-Type", "application/json")

// 	// Convert the response data to JSON format
// 	jsonResponse, err := json.Marshal(response)
// 	if err != nil {
// 		log.Println("Failed to marshal JSON:", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Write the response data to the client
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonResponse)

// }

func GetAllCustomerEcho(e echo.Context) error {
	rUser := []users.User_data{}
	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, nick_name, cust_no, user_name, email,mobile FROM t_user_template limit 10`)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, rows)
	}
	defer rows.Close()

	for rows.Next() {
		var users users.User_data
		rows.Scan(&users.Id, &users.Nick_name, &users.Cust_no, &users.User_name, &users.Email, &users.Mobile)
		// if err != nil {
		// 	log.Println(err)
		// 	return e.JSON(http.StatusInternalServerError, rows)
		// }
		rUser = append(rUser, users)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, rows)
	}
	return e.JSON(http.StatusOK, rUser)
}

// var data_user = []Datauser{}

type Datauser struct {
	ID        uuid.UUID
	ID_       int64  `json:"id"`
	User_name string `json:"user_name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	User_id   string `json:"user_id"`
}

func RegisterEcho(e echo.Context) error {
	// Parse the request body to extract the data to be inserted
	body := &Datauser{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the INSERT statement
	// stmt, err := db.Prepare("INSERT INTO t_user_template(id,create_date,modify_date,user_name,register_channels,nick_name,user_id,cust_no,mobile,email,user_pic,acct_msg_status,is_locked,potential_user_lv,potential_user_type,potential_user_desc,login_password,update_pwd_time,figure_password,gesture_password,passcode,is_figure_open,is_face_recog_open,is_gesture_login_open,is_hide_gesture_login,is_passcode_open,device_token,device_no,user_barcode,is_email_login,email_auth,language,need_change_pwd,my_referral_code,catch_recover_time,fail_potential_time,esign_kid,esign_ref_number,is_notification,is_email_noti,is_sms_noti,is_push_noti,acct_mask,del,space1,space2,space3,space4,space5,space6,customerStatus,ndid_open_time,reg_pdf_par,channel_sign,edc_open_time,edc_qr_data,cust_journey_reference,channel_info_sign,kyc_risk_level,ocr_toggle_status,test) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	stmt, err := db.Prepare("INSERT INTO t_user_template(id,user_name,email,mobile,create_date,modify_date) VALUES(?,?,?,?,?,?);")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	// Execute the INSERT statement
	UUID := other.GenerateUUID()
	Createdate := other.DateTime()
	Modifydate := other.DateTime()
	result, err := stmt.Exec(UUID, body.Email, body.Mobile, body.User_name, Createdate, Modifydate)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	} else {
		id, _ := result.LastInsertId()
		body.ID_ = id
		return e.JSON(http.StatusCreated, body)
	}
}

func UpdateCustInfo(e echo.Context) error {
	dataUser := []Datauser{}
	body := &Datauser{}

	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, user_name, email,mobile,user_id FROM t_user_template WHERE user_id=?`, body.User_id)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, rows)
	}
	defer rows.Close()

	for rows.Next() {
		var m Datauser
		if err := rows.Scan(&m.ID, &m.User_name, &m.Email, &m.Mobile, &m.User_id); err != nil {
			dataUser = append(dataUser, m)
		}
		dataUser = append(dataUser, m)
	}
	if err != nil {
		log.Fatal(err)
	}

	var UserName string
	var Email string
	var Mobile string
	for _, data := range dataUser {
		UserName = data.User_name
		Email = data.Email
		Mobile = data.Mobile
	}

	update, err := db.Prepare(`UPDATE t_user_template SET user_name=?, mobile=?, email=? WHERE user_id=?`)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	defer update.Close()

	// fmt.Println("JSON Username :", body.User_name)
	// fmt.Println("JSON User_id :", body.User_id)
	// fmt.Println("JSON Mobile :", body.Mobile)
	// fmt.Println("JSON Email :", body.Email)

	var dataUpdate Datauser
	if body.Email == "" {
		dataUpdate.User_name = UserName
	} else {
		dataUpdate.User_name = body.User_name
	}

	if body.Email == "" {
		dataUpdate.Email = Email
	} else {
		dataUpdate.Email = body.Email
	}

	if body.Mobile == "" {
		dataUpdate.Mobile = Mobile
	} else {
		dataUpdate.Mobile = body.Mobile
	}

	fmt.Println("JSON Username :", dataUpdate.User_name)
	fmt.Println("JSON User_id :", body.User_id)
	fmt.Println("JSON Mobile :", dataUpdate.Mobile)
	fmt.Println("JSON Email :", dataUpdate.Email)

	_, err = update.Exec(dataUpdate.User_name, dataUpdate.Mobile, dataUpdate.Email, body.User_id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, body)
}

func GetUserHandleByID(e echo.Context) error {
	body := &Datauser{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow(`SELECT id, user_name, email,mobile,user_id FROM t_user_template WHERE user_id = ?`, body.User_id)

	var rawData Datauser
	err1 := row.Scan(&rawData.ID, &rawData.User_name, &rawData.Email, &rawData.Mobile, &rawData.User_id)

	switch err1 {
	case sql.ErrNoRows:
		return e.JSON(http.StatusNotFound, map[string]string{"message": "not found"})
	case nil:
		return e.JSON(http.StatusOK, rawData)

	default:
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
}
