package buffer

// Special Struct Just for Patch Operation
type PatchOp struct {
	Id           string          `json:"id,omitempty"`
	Type         string          `json:"type,omitempty"`
	Additional   string          `json:"additional,omitempty"`
	CreateParent bool            `json:"create_parent,omitempty"`
	Patch        []*PatchRequest `json:"patch,omitempty"`
}

type PatchRequest struct {
	Op    string      `json:"op,omitempty"`
	Path  string      `json:"path,omitempty"`
	Value interface{} `json:"value"`
}
