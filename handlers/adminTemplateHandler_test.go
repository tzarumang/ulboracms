package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
	sr "github.com/Ulbora/ulboracms/services"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
)

func TestCmsHandler_AdminTemplateList(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminTemplateList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestCmsHandler_AdminTemplateListNotLoggedIn(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = false
	s.Save(r, w)
	h.AdminTemplateList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminAddTemplate(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminAddTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestCmsHandler_AdminAddTemplateNotLoggedIn(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = false
	s.Save(r, w)
	h.AdminAddTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminUploadTemplate(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"
	ci.TemplateFilePath = "../services/testDownloads"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()

	file, err := os.Open("../services/testUploads/hestia.tar.gz")
	if err != nil {
		fmt.Println("file test template open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("template fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("tempFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("template upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminUploadTemplate(w, r)
	fmt.Println("code in upload: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminUploadTemplateFail(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"
	ci.TemplateFilePath = "../services/testDownloads"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()

	file, err := os.Open("../services/testUploads/hestia.tar.gz")
	if err != nil {
		fmt.Println("file test template open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("template fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("tempFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("template upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminUploadTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestCmsHandler_AdminUploadTemplateNotLoggedIn(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"
	ci.TemplateFilePath = "../services/testDownloads"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()

	file, err := os.Open("../services/testUploads/hestia.tar.gz")
	if err != nil {
		fmt.Println("file test template open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("template fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("tempFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("template upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = false
	s.Save(r, w)
	h.AdminUploadTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminActivateTemplate(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	ch.ActiveTemplateLocation = "../services/testDownloads"

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	vars := map[string]string{
		"name": "hestia",
	}
	r = mux.SetURLVars(r, vars)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminActivateTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminActivateTemplateNotLoggedIn(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	vars := map[string]string{
		"name": "hestia",
	}
	r = mux.SetURLVars(r, vars)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = false
	s.Save(r, w)
	h.AdminActivateTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminDeleteTemplate(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	vars := map[string]string{
		"name": "hestia",
	}
	r = mux.SetURLVars(r, vars)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminDeleteTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminActivateTemplate2(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	ch.ActiveTemplateLocation = "../services/testDownloads"

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	vars := map[string]string{
		"name": "temp2",
	}
	r = mux.SetURLVars(r, vars)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminActivateTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminDeleteTemplate2(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"
	ci.TemplateFilePath = "../services/testDownloads"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	vars := map[string]string{
		"name": "hestia",
	}
	r = mux.SetURLVars(r, vars)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Save(r, w)
	h.AdminDeleteTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestCmsHandler_AdminDeleteTemplateNotLoggedIn(t *testing.T) {
	var ch CmsHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ci sr.CmsService
	ci.TemplateStorePath = "../services/testBackup/templateStore"
	ci.TemplateFilePath = "../services/testDownloads"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../services/testBackup/templateStore"
	ci.TemplateStore = ds.GetNew()
	ch.Service = ci.GetNew()

	h := ch.GetNew()
	r, _ := http.NewRequest("GET", "/test", nil)
	vars := map[string]string{
		"name": "hestia",
	}
	r = mux.SetURLVars(r, vars)
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s, suc := ch.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = false
	s.Save(r, w)
	h.AdminDeleteTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
