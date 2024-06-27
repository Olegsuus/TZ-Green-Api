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
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getSettings(c echo.Context) error {
	params := new(models.Params)
	if err := c.Bind(params); err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.green-api.com/getSettings/%s/%s", params.IDInstance, params.ApiTokenInstance)
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

func getStateInstance(c echo.Context) error {
	params := new(models.Params)
	if err := c.Bind(params); err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.green-api.com/getStateInstance/%s/%s", params.IDInstance, params.ApiTokenInstance)
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

func sendMessage(c echo.Context) error {
	params := new(models.Params)
	if err := c.Bind(params); err != nil {
		return err
	}

	url := "https://api.green-api.com/sendMessage"
	reqBody, err := json.Marshal(map[string]string{
		"idInstance":       params.IDInstance,
		"apiTokenInstance": params.ApiTokenInstance,
		"phoneNumber":      params.PhoneNumber,
		"message":          params.Message,
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

func sendFileByUrl(c echo.Context) error {
	params := new(models.Params)
	if err := c.Bind(params); err != nil {
		return err
	}

	url := "https://api.green-api.com/sendFileByUrl"
	reqBody, err := json.Marshal(map[string]string{
		"idInstance":       params.IDInstance,
		"apiTokenInstance": params.ApiTokenInstance,
		"phoneNumberFile":  params.PhoneNumber,
		"fileUrl":          params.FileUrl,
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
