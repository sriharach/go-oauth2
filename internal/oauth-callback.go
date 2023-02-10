package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type loginCallbackBody struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
type GoogleIdToken struct {
	Issuer          string `json:"iss"`
	Subject         string `json:"sub"`
	Audience        string `json:"aud"`
	Expiry          int    `json:"exp"`
	IssuedAt        int    `json:"iat"`
	AtHash          string `json:"at_hash"`
	Hd              string `json:"hd"`
	AuthorizedParty string `json:"azp"`
	Picture         string `json:"picture"`
	Locale          string `json:"locale"`
	Email           string `json:"email"`
	EmailVerified   bool   `json:"email_verified"`
	Name            string `json:"name"`
	FamilyName      string `json:"family_name"`
	GivenName       string `json:"given_name"`
}

func LoginCallback(c *fiber.Ctx) error {
	// request := new(loginCallbackBody)

	// if err := c.BodyParser(request); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	// }

	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Code is not data.")
	}

	config := Oauth2()
	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("tok", tok)
	// id := tok.Extra("id_token").(string)
	// parts := strings.Split(id, ".")
	// payload, err := base64.RawURLEncoding.DecodeString(parts[1])

	// fmt.Println("payload", payload)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
	fmt.Println("resp", resp.Body)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println("content", content)

	var info GoogleIdToken
	json.Unmarshal(content, &info)
	return c.JSON(info)
}
