package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"time"

	"github.com/caarlos0/env/v8"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func main() {
	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateValuer, url.URL{})

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalf("%+v\n", err)
	}

	registerFn := func(ut ut.Translator) error {
		return ut.Add("http_url", `invalid HTTP address: '{0}' in field '{1}'`, true) // see universal-translator for details
	}
	translationFn := func(ut ut.Translator, fe validator.FieldError) string {
		value := fmt.Sprintf("%s", fe.Value())
		field := fe.StructNamespace()
		t, _ := ut.T("http_url", value, field)

		return t
	}

	if err := validate.RegisterTranslation("http_url", trans, registerFn, translationFn); err != nil {
		log.Fatalf("%+v\n", err)
	}

	cfg := config{}
	opts := env.Options{
		RequiredIfNoDef:       true,
		UseFieldNameByDefault: true,
	}

	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err := validate.Struct(cfg); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}

		os.Exit(1)
	}

	log.Printf("%+v\n", cfg)
}

type config struct {
	Home         string
	Port         int
	IsProduction bool
	Hosts        []*url.URL `validate:"dive,http_url"`
	Host         *url.URL   `validate:"http_url"`
	Duration     time.Duration
	TempFolder   string
}

func ValidateValuer(field reflect.Value) any {
	if valuer, ok := field.Interface().(url.URL); ok {
		return valuer.String()
		// handle the error how you want
	}

	return nil
}
