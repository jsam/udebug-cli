// Code generated by go-swagger; DO NOT EDIT.

package input

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jsam/udebug-cli/models"
)

// GetInputAPIInputRetrieveJSONReader is a Reader for the GetInputAPIInputRetrieveJSON structure.
type GetInputAPIInputRetrieveJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInputAPIInputRetrieveJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInputAPIInputRetrieveJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInputAPIInputRetrieveJSONOK creates a GetInputAPIInputRetrieveJSONOK with default headers values
func NewGetInputAPIInputRetrieveJSONOK() *GetInputAPIInputRetrieveJSONOK {
	return &GetInputAPIInputRetrieveJSONOK{}
}

/*GetInputAPIInputRetrieveJSONOK handles this case with default header values.

OK
*/
type GetInputAPIInputRetrieveJSONOK struct {
	Payload models.ArrayOfString
}

func (o *GetInputAPIInputRetrieveJSONOK) Error() string {
	return fmt.Sprintf("[GET /input_api/input/retrieve.json][%d] getInputApiInputRetrieveJsonOK  %+v", 200, o.Payload)
}

func (o *GetInputAPIInputRetrieveJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}