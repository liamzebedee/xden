package XDen

import (

)

/*
A Document is an object that has the following traits:
 - High persistence
 - Overturn
 - Load balancing
Documents are identified by the hash of their name (abitrarily chosen)
*/
type Document struct {
	Name string
	ID Key
	Data []byte
}

/*
An update is:
 - Short lived
 - Quickly propogated
Updates are identified by the hash of their data
*/
type Update struct {
	ID Key
	Data []byte
}

type DENSDB interface {
	/*
	PUT stores a document-like object in the overlay, associating it with a key (calculated as
	the hash of the name). 
	The name presented by a PUT will be persistent, however data will not (due to overturn)
	Returns any error encountered and the document put
	*/
	PUT(document Document)	(error, *Document)
	
	/*
	GET tries to retrieve a document associated with a name.
	Returns any error encountered and the document retrieved (if any)
	*/
	GET(name string) (error, Document)
	
	/*
	PUSH propogates an update-like object in the overlay
	*/
	PUSH(update Update)		(error)
	
	
	PULL()					(error)
}