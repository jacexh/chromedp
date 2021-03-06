package domstorage

// AUTOGENERATED. DO NOT EDIT.

// StorageID dOM Storage identifier.
type StorageID struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin for the storage.
	IsLocalStorage bool   `json:"isLocalStorage"` // Whether the storage is local storage (not session storage).
}

// Item dOM Storage item.
type Item []string
