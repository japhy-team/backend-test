// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// Defines values for BreedsPetSize.
const (
	Medium BreedsPetSize = "medium"
	Small  BreedsPetSize = "small"
	Tall   BreedsPetSize = "tall"
)

// Defines values for Species.
const (
	Cat Species = "cat"
	Dog Species = "dog"
)

// Breeds defines model for Breeds.
type Breeds struct {
	// AverageFemaleAdultWeight Average weight of the female adult in gramme
	AverageFemaleAdultWeight *int `json:"average_female_adult_weight,omitempty"`

	// AverageMaleAdultWeight Average weight of the male adult in gramme
	AverageMaleAdultWeight *int `json:"average_male_adult_weight,omitempty"`

	// Name Name of the breed. Should be in snake case and should be unique
	Name string `json:"name"`

	// PetSize size of the pet
	PetSize BreedsPetSize `json:"pet_size"`
	Species Species       `json:"species"`
}

// BreedsPetSize size of the pet
type BreedsPetSize string

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Species defines model for Species.
type Species string

// AverageFemaleAdultWeight Average weight of the female adult in gramme
type AverageFemaleAdultWeight = int

// AverageMaleAdultWeight Average weight of the male adult in gramme
type AverageMaleAdultWeight = int

// BreedName defines model for BreedName.
type BreedName = string

// BadRequestError defines model for BadRequestError.
type BadRequestError = Error

// BreedResponse defines model for BreedResponse.
type BreedResponse = Breeds

// BreedsList defines model for BreedsList.
type BreedsList = []Breeds

// InternalServerError defines model for InternalServerError.
type InternalServerError = Error

// ResourceAlreadyExistsError defines model for ResourceAlreadyExistsError.
type ResourceAlreadyExistsError = Error

// ResourceNotFoundError defines model for ResourceNotFoundError.
type ResourceNotFoundError = Error

// Breed defines model for Breed.
type Breed = Breeds

// ListBreedsParams defines parameters for ListBreeds.
type ListBreedsParams struct {
	Species                  *Species                  `form:"species,omitempty" json:"species,omitempty"`
	AverageFemaleAdultWeight *AverageFemaleAdultWeight `form:"average_female_adult_weight,omitempty" json:"average_female_adult_weight,omitempty"`
	AverageMaleAdultWeight   *AverageMaleAdultWeight   `form:"average_male_adult_weight,omitempty" json:"average_male_adult_weight,omitempty"`
}

// CreateOneBreedJSONRequestBody defines body for CreateOneBreed for application/json ContentType.
type CreateOneBreedJSONRequestBody = Breeds

// CreateOrUpdateBreedByNameJSONRequestBody defines body for CreateOrUpdateBreedByName for application/json ContentType.
type CreateOrUpdateBreedByNameJSONRequestBody = Breeds

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List breeds
	// (GET /breeds)
	ListBreeds(w http.ResponseWriter, r *http.Request, params ListBreedsParams)
	// Create one breed
	// (POST /breeds)
	CreateOneBreed(w http.ResponseWriter, r *http.Request)
	// Delete a given breed by its name
	// (DELETE /breeds/name/{breed_name})
	DeleteBreedByName(w http.ResponseWriter, r *http.Request, breedName BreedName)
	// Retrieve a given breed by its name
	// (GET /breeds/name/{breed_name})
	GetBreedByName(w http.ResponseWriter, r *http.Request, breedName BreedName)
	// Update or create one breed
	// (PUT /breeds/name/{breed_name})
	CreateOrUpdateBreedByName(w http.ResponseWriter, r *http.Request, breedName BreedName)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListBreeds operation middleware
func (siw *ServerInterfaceWrapper) ListBreeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListBreedsParams

	// ------------- Optional query parameter "species" -------------

	err = runtime.BindQueryParameter("form", true, false, "species", r.URL.Query(), &params.Species)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "species", Err: err})
		return
	}

	// ------------- Optional query parameter "average_female_adult_weight" -------------

	err = runtime.BindQueryParameter("form", true, false, "average_female_adult_weight", r.URL.Query(), &params.AverageFemaleAdultWeight)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "average_female_adult_weight", Err: err})
		return
	}

	// ------------- Optional query parameter "average_male_adult_weight" -------------

	err = runtime.BindQueryParameter("form", true, false, "average_male_adult_weight", r.URL.Query(), &params.AverageMaleAdultWeight)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "average_male_adult_weight", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListBreeds(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateOneBreed operation middleware
func (siw *ServerInterfaceWrapper) CreateOneBreed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateOneBreed(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteBreedByName operation middleware
func (siw *ServerInterfaceWrapper) DeleteBreedByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "breed_name" -------------
	var breedName BreedName

	err = runtime.BindStyledParameterWithOptions("simple", "breed_name", mux.Vars(r)["breed_name"], &breedName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "breed_name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteBreedByName(w, r, breedName)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetBreedByName operation middleware
func (siw *ServerInterfaceWrapper) GetBreedByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "breed_name" -------------
	var breedName BreedName

	err = runtime.BindStyledParameterWithOptions("simple", "breed_name", mux.Vars(r)["breed_name"], &breedName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "breed_name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBreedByName(w, r, breedName)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateOrUpdateBreedByName operation middleware
func (siw *ServerInterfaceWrapper) CreateOrUpdateBreedByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "breed_name" -------------
	var breedName BreedName

	err = runtime.BindStyledParameterWithOptions("simple", "breed_name", mux.Vars(r)["breed_name"], &breedName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "breed_name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateOrUpdateBreedByName(w, r, breedName)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/breeds", wrapper.ListBreeds).Methods("GET")

	r.HandleFunc(options.BaseURL+"/breeds", wrapper.CreateOneBreed).Methods("POST")

	r.HandleFunc(options.BaseURL+"/breeds/name/{breed_name}", wrapper.DeleteBreedByName).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/breeds/name/{breed_name}", wrapper.GetBreedByName).Methods("GET")

	r.HandleFunc(options.BaseURL+"/breeds/name/{breed_name}", wrapper.CreateOrUpdateBreedByName).Methods("PUT")

	return r
}
