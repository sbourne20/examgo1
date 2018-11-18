package main

import (
	"database/sql"
	"fmt"
	"time"
)

type dateType time.Time

type user struct {
	Tglrecord                   time.Time      `json:"Tgl_record"`
	Idpenyelenggara             int            `json:"Id_penyelenggara"`
	Idpengguna                  int            `json:"Id_pengguna"`
	Namapengguna                string         `json:"Nama_pengguna"`
	Jenispengguna               int            `json:"Jenis_pengguna"`
	Idbadanhukum                sql.NullString `json:"Id_badan_hukum"`
	Jenisbadanhukum             sql.NullInt64  `json:"Jenis_badan_hukum"`
	Namabadanhukum              sql.NullString `json:"Nama_badan_hukum"`
	Idktp                       string         `json:"Idktp"`
	Tempatlahir                 string         `json:"Id_ktp"`
	Tanggallahir                time.Time      `json:"Tanggal_lahir"`
	Usia                        int            `json:"Usia"`
	Jeniskelamin                string         `json:"Jenis_kelamin"`
	Pendidikan                  int            `json:"Pendidikan"`
	Pekerjaan                   int            `json:"Pekerjaan"`
	Bidangpekerjaan             string         `json:"Bidang_pekerjaan"`
	Nomorrekening               string         `json:"Nomor_rekening"`
	Namabank                    string         `json:"nama_bank"`
	Idmodusermember             int            `json:"id_mod_user_member"`
	Peringkatpengguna           string         `json:"peringkat_pengguna"`
	Peringkatpenggunapersentase int            `json:"peringkat_pengguna_persentase"`
	Skoring                     int            `json:"skoring"`
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("SELECT * FROM user2")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user

		err = rows.Scan(&u.Tglrecord,
			&u.Idpenyelenggara,
			&u.Idpengguna,
			&u.Namapengguna,
			&u.Jenispengguna,
			&u.Idbadanhukum,
			&u.Jenisbadanhukum,
			&u.Namabadanhukum,
			&u.Idktp,
			&u.Tempatlahir,
			&u.Tanggallahir,
			&u.Usia,
			&u.Jeniskelamin,
			&u.Pendidikan,
			&u.Pekerjaan,
			&u.Bidangpekerjaan,
			&u.Nomorrekening,
			&u.Namabank,
			&u.Idmodusermember,
			&u.Peringkatpengguna,
			&u.Peringkatpenggunapersentase,
			&u.Skoring)

		if err != nil {
			panic(err.Error())
		}
		users = append(users, u)
	}

	return users, nil
}
