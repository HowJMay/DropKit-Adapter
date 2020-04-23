package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/DropKit/DropKit-Adapter/constants"
	"github.com/DropKit/DropKit-Adapter/package/crypto/account"
	"github.com/DropKit/DropKit-Adapter/package/crypto/transaction"
	"github.com/DropKit/DropKit-Adapter/package/parser"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/spf13/viper"
)

func SQLCreate(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	_, tableAddress := account.GenerateWallet()

	services.AddMetaTable(tableName, tableAddress, metaTableAddress, callerPriavteKey)
	services.GrantAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)
	services.Exec(sqlCommand)

	aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)

	defer r.Body.Close()
	response := constants.Exec_Response{"200", aduitTransactionHash}

	services.ResponseWithJson(w, http.StatusOK, response)
}

func SQLInsert(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		services.Exec(sqlCommand)
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		defer r.Body.Close()
		response := constants.Exec_Response{"200", aduitTransactionHash}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		defer r.Body.Close()
		response := constants.Exec_Response{"401", "NULL"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}

func SQLSelect(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		metadata, _ := services.Query(sqlCommand)
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		defer r.Body.Close()
		response := constants.Query_Response{"200", metadata, aduitTransactionHash}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		defer r.Body.Close()
		response := constants.Query_Response{"401", "NULL", "NULL"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}

func SQLUpdate(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		services.Exec(sqlCommand)
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		defer r.Body.Close()
		response := constants.Exec_Response{"200", aduitTransactionHash}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		defer r.Body.Close()
		response := constants.Exec_Response{"401", "NULL"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}

func SQLDelete(w http.ResponseWriter, r *http.Request) {
	metaTableAddress := viper.GetString(`DROPKIT.METATABLE`)
	authorityAddr := viper.GetString(`DROPKIT.AUTHORITY`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var newStatement constants.DB_Statement
	_ = json.Unmarshal(body, &newStatement)
	sqlCommand := newStatement.Statement
	callerPriavteKey := newStatement.PrivateKey
	callerAddress := account.PrivateKeyToPublicKey(callerPriavteKey)

	tableName := parser.GetTableName(sqlCommand)
	tableAddress := services.GetMetaTable(tableName, metaTableAddress)
	authority := services.VerifyAuthority(authorityAddr, callerPriavteKey, tableName, callerAddress)

	switch authority {
	case true:
		services.Exec(sqlCommand)
		aduitTransactionHash := transaction.SendRawTransaction(tableAddress, sqlCommand, 0, callerPriavteKey)
		defer r.Body.Close()
		response := constants.Exec_Response{"200", aduitTransactionHash}
		services.ResponseWithJson(w, http.StatusOK, response)
	case false:
		defer r.Body.Close()
		response := constants.Exec_Response{"401", "NULL"}
		services.ResponseWithJson(w, http.StatusUnauthorized, response)
	}
}
