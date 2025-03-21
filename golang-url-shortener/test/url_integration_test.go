package test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/matheusapostulo/url-shortener/cmd/url/application"
	"github.com/matheusapostulo/url-shortener/cmd/url/dependencies"
	"github.com/stretchr/testify/require"
)

var (
	url_base   = "http://localhost:8080/api/v1"
	app_json   = "application/json"
	ErrSendReq = errors.New("error while sending request to server")
	ErrReadRes = errors.New("error while reading the body of the response")
)

// func startMySQLContainer() (string, string, func(), error) {
// 	ctx := context.Background()
// 	req := testcontainers.ContainerRequest{
// 		Image:        "mysql:latest",
// 		Env:          map[string]string{"MYSQL_ROOT_PASSWORD": "root", "MYSQL_DATABASE": "urls"},
// 		ExposedPorts: []string{"3306/tcp"},
// 		Cmd:          []string{"--init-file", "mysql-init/init.sql"},
// 		WaitingFor:   wait.ForLog("port: 3306  MySQL Community Server - GPL"),
// 	}

// 	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: req,
// 		Started:          true,
// 	})
// 	if err != nil {
// 		return "", "", nil, err
// 	}

// 	host, err := mysqlContainer.Host(ctx)
// 	if err != nil {
// 		return "", "", nil, err
// 	}

// 	port, err := mysqlContainer.MappedPort(ctx, "3306")
// 	if err != nil {
// 		return "", "", nil, err
// 	}

// 	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/urls", host, port.Port())

// 	stopContainer := func() {
// 		mysqlContainer.Terminate(ctx)
// 	}

// 	fmt.Println("MySQL container running on", host, port.Port())

// 	return dsn, host, stopContainer, nil
// }

func init() {
	go func() {
		dependencies, err := dependencies.BuildDependencies()
		if err != nil {
			panic(err)
		}
		defer dependencies.DB.Close()
		defer dependencies.CacheClient.Close()

		application.Run(*dependencies)
	}()
}
func TestIntegrationCreateURL(t *testing.T) {
	bodyReq := `{
		"long_url": "http://www.google.com"
	}`
	resp, err := http.Post(url_base+"/shorten", app_json, strings.NewReader(bodyReq))
	if err != nil {
		panic(ErrSendReq)
	}
	defer resp.Body.Close()

	bodyRes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(ErrReadRes)
	}

	var urlRes struct {
		ShortURL string `json:"short_url"`
	}
	err = json.Unmarshal(bodyRes, &urlRes)

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.NotEmpty(t, urlRes.ShortURL)
}
