// Code generated by go-swagger; DO NOT EDIT.

package hint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jsam/udebug-cli/models"
)

// GetHintAPIHintRetrieveJSONReader is a Reader for the GetHintAPIHintRetrieveJSON structure.
type GetHintAPIHintRetrieveJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHintAPIHintRetrieveJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHintAPIHintRetrieveJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHintAPIHintRetrieveJSONOK creates a GetHintAPIHintRetrieveJSONOK with default headers values
func NewGetHintAPIHintRetrieveJSONOK() *GetHintAPIHintRetrieveJSONOK {
	return &GetHintAPIHintRetrieveJSONOK{}
}

/*GetHintAPIHintRetrieveJSONOK handles this case with default header values.

OK
*/
type GetHintAPIHintRetrieveJSONOK struct {
	Payload models.ArrayOfString
}

func (o *GetHintAPIHintRetrieveJSONOK) Error() string {
	return fmt.Sprintf("[GET /hint_api/hint/retrieve.json][%d] getHintApiHintRetrieveJsonOK  %+v", 200, o.Payload)
}

func (o *GetHintAPIHintRetrieveJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}