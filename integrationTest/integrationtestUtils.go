package integrationtest

import (
	"aeperez24/basego/config"
	"aeperez24/basego/dto"
	"aeperez24/basego/handler"
	"aeperez24/basego/model"
	"aeperez24/basego/port"
	"aeperez24/basego/services"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func RunTestWithIntegrationServerGin(testFunc func(port string)) {
	config.LoadViperConfig("../envs/", "isolation")
	DBConfig := config.BuildDBConfig()
	server, port := createTestServerGin(DBConfig)
	createUserForTest(DBConfig)
	defer deleteUsersForTests(DBConfig)
	go server.Start()
	testFunc(port)
	server.Stop()

}

func createTestServerGin(DBConfig config.MongoCofig) (port.Server, string) {
	port := "11082"
	serverConfig := handler.BuildServerConfigGin(port, "testKey", DBConfig)
	return handler.NewGinServer(serverConfig), port
}

func createUserForTest(dbConfig config.MongoCofig) []interface{} {

	user1 := model.User{Username: "user1", Password: "pass1", IDDocument: "account1Number"}
	user2 := model.User{Username: "user2", Password: "pass2", IDDocument: "account2Number"}
	collection := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.USER_COLLECTION)
	resultID1, error1 := collection.InsertOne(context.TODO(), user1)
	resultID2, error2 := collection.InsertOne(context.TODO(), user2)
	if error1 != nil {
		println(error1)
		panic(error1)
	}
	if error2 != nil {
		println(error1)
		panic(error1)
	}
	result := []interface{}{
		resultID1.InsertedID, resultID2.InsertedID,
	}
	return result

}

func deleteUsersForTests(dbConfig config.MongoCofig) {
	dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.USER_COLLECTION).DeleteMany(context.TODO(), nil)

}

func GetJWTTokenForUser1() string {
	tokenService := services.NewTokenService("testKey")
	res, _ := tokenService.CreateToken(dto.BasicUserDto{
		Username:   "user1",
		IDDocument: "account1Number",
	})
	return res
}

func GetJWTTokenForUser2() string {
	tokenService := services.NewTokenService("testKey")
	res, _ := tokenService.CreateToken(dto.BasicUserDto{
		Username:   "user2",
		IDDocument: "account2Number",
	})
	return res
}

func ExecuteHttpPostCall(url string, bodyInterface interface{}, headers map[string]string) ([]byte, *http.Response, error) {
	body, _ := json.Marshal(bodyInterface)
	postBuffer := bytes.NewBuffer(body)

	req, _ := http.NewRequest("POST", url, postBuffer)
	for name, value := range headers {
		req.Header.Add(name, value)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, _ := client.Do((req))
	bodyresp, err := ioutil.ReadAll(resp.Body)
	return bodyresp, resp, err
}
