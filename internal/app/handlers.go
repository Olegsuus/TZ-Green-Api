package app

import (
	"TZ-GREEN-API_/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
	"path"
)

const (
	MethodGetSettings      = "getSettings"
	MethodGetStateInstance = "getStateInstance"
	MethodSendMessage      = "sendMessage"
	MethodSendFileByUrl    = "sendFileByUrl"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func constructURL(apiUrl, idInstance, apiTokenInstance, method string) string {
	return fmt.Sprintf("%s/waInstance%s/%s/%s", apiUrl, idInstance, method, apiTokenInstance)
}

func (a *App) getSettings(c echo.Context) error {
	params := models.Params{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	url := constructURL(a.Config.Api.ApiUrl, params.IDInstance, params.ApiTokenInstance, MethodGetSettings)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (a *App) getStateInstance(c echo.Context) error {
	params := models.Params{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	url := constructURL(a.Config.Api.ApiUrl, params.IDInstance, params.ApiTokenInstance, MethodGetStateInstance)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (a *App) sendMessage(c echo.Context) error {
	params := models.Params{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	url := constructURL(a.Config.Api.ApiUrl, params.IDInstance, params.ApiTokenInstance, MethodSendMessage)

	reqBody, err := json.Marshal(map[string]string{
		"chatId":  params.PhoneNumber,
		"message": params.Message,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (a *App) sendFileByUrl(c echo.Context) error {
	params := models.ParamsFileUrl{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	url := constructURL(a.Config.Api.ApiUrl, params.IDInstance, params.ApiTokenInstance, MethodSendFileByUrl)

	reqBody, err := json.Marshal(map[string]string{
		"chatId":   params.ChatId,
		"urlFile":  params.UrlFile,
		"fileName": path.Base(params.UrlFile),
		"caption":  params.Caption,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
