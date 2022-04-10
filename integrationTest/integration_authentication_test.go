package integrationtest

import (
	"aeperez24/basego/config"
	"aeperez24/basego/dto"
	"aeperez24/basego/model"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAuthUser(t *testing.T) {
	RunTestWithIntegrationServerGin(func(port string) {
		username := "username_for_integration_testing"
		document := "document_for_integration_testing"
		password := "pass"
		dbConfig := config.BuildDBConfig()

		collectionUser := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.USER_COLLECTION)
		defer collectionUser.DeleteOne(context.TODO(), bson.M{"username": username})

		userdto := dto.UserWithPasswordDto{
			BasicUserDto: dto.BasicUserDto{
				Username:   username,
				IDDocument: document,
			},
			Password: password,
		}
		apiSignUp := fmt.Sprintf("http://localhost:%s/user/signup", port)
		apiSignIn := fmt.Sprintf("http://localhost:%s/user/signin", port)

		ExecuteHttpPostCall(apiSignUp, userdto, nil)
		bodyresp, resp, err := ExecuteHttpPostCall(apiSignIn, userdto, nil)
		assert.Equal(t, 200, resp.StatusCode, "status wrong")
		assert.Nil(t, err, "error should be nil")
		println(string(bodyresp))

	})
}
