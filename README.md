# vebra-api
Package for querying the Vebra v10 API, implemented in Golang

## Example Usage

``` go
package main

import(
	"log"
)

func main() {

	// We need a way of storing our current session token,
	// so we'll use the FileTokenStorage
	ts := &vebra.FileTokenStorage{}
	ts.SetFileName(tokenFile)

	// Next, craete our client
	api := vebra.Create(datafeedID, "user", "password", ts)

	// Set the way we want to access the data. I've
	// included 3 methods:
	// * NewRemoteFileGetter - queries the Vebra API
	// * NewRemoteFileGetterLocalWriter - queries the Vebra API but qrites the results to file locally
	// * NewLocalFileGetter - Reads in the files written out by NewRemoteFileGetterLocalWriter. Good for testing.

	api.GetDataFunc = vebra.NewRemoteFileGetter()

	// Return a summary of the available branches
	branches, err := api.GetBranches()

	checkErr(err)

	for _, branch := range branches.Branches {
		// Return the full details of a branch
		fullBranch, err := api.GetBranch(branch)
		checkErr(err)

		// The the branch ID. We need this for getting properties etc.
		id, err := branch.GetClientID()

		checkErr(err)

		props, err := api.GetPropertyList(id)


		for _, propSum := range props.Properties {
			property, err := api.GetProperty(id, propSum.PropID)

			if err != nil {
				log.Printf("ERR: Error getting property: [%s]\n", err)
				continue
			}

			// Do stuff with the property here...
		}
	}
}

func checkErr(err) {
	if err != nil {
		panic(err)
	}
}
```
