gapi-go is a Go client for the G Adventures REST API. It wraps the necessities to fetch, update, and delete resources.

Usage
===

First, set your `ApiKey`

    import (
        gapi "github.com/bartek/gapi-go/gapigo"
    )

    func main() {
        gapi.ApiKey = "your_apikey_replaced"
    }

To access a resource, first, define the parameters for this query:

    var req = gapi.ApiRequest{
        Url: "/tours/21346/",
    }

You are able to update `req` as you see fit, technically (but maybe you
shouldn't!)

Then, make a request to fetch it when you're ready.

    status, err := req.Fetch()

The result is stored in the `Result` key. This is unmarshalled json:

    m := req.Result.(map[string]interface{})
    for k, v := range m {
        // .. do stuff
    }

See the result in raw form

    log.Println(req.RawRext)

