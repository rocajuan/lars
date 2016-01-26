package lars

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestContext(t *testing.T) {

	l := New()
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	defaultContext := NewContext(l)

	var varParams []Param

	// Parameter
	param1 := Param{
		Key:   "userID",
		Value: "507f191e810c19729de860ea",
	}

	varParams = append(varParams, param1)

	//store
	storeMap := store{
		"User":        "Alice",
		"Information": []string{"Alice", "Bob", "40.712784", "-74.005941"},
	}

	defaultContext.params = varParams
	defaultContext.store = storeMap
	defaultContext.request = r

	c := defaultContext.UnderlyingContext()

	//Request
	NotEqual(t, c.Request(), nil)

	//Response
	NotEqual(t, c.Response(), nil)

	//Parameters by ID
	bsonValue, ok := c.P(0)
	Equal(t, true, ok)
	Equal(t, "507f191e810c19729de860ea", bsonValue)

	//Paramter by name
	bsonValue, ok = c.Param("userID")
	Equal(t, true, ok)
	Equal(t, "507f191e810c19729de860ea", bsonValue)

	//Store
	c.Set("publicKey", "U|ydN3SX)B(hI8SV1R;(")

	value, exists := c.Get("publicKey")

	//Get
	Equal(t, true, exists)
	Equal(t, "U|ydN3SX)B(hI8SV1R;(", value)

	value, exists = c.Get("User")
	Equal(t, true, exists)
	Equal(t, "Alice", value)

	value, exists = c.Get("UserName")
	NotEqual(t, true, exists)
	NotEqual(t, "Alice", value)

	value, exists = c.Get("Information")
	Equal(t, true, exists)
	vString := value.([]string)

	Equal(t, "Alice", vString[0])
	Equal(t, "Bob", vString[1])
	Equal(t, "40.712784", vString[2])
	Equal(t, "-74.005941", vString[3])

	// Reset
	c.Reset(w, r)

	//Request
	NotEqual(t, c.Request(), nil)

	//Response
	NotEqual(t, c.Response(), nil)

	//Param
	Equal(t, len(c.Params()), 0)

	//Set
	Equal(t, c.store, nil)

	// Index
	Equal(t, c.index, -1)

	// Handlers
	Equal(t, c.handlers, nil)

}