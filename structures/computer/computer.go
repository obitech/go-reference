package computer

// Capital structs are exported
type Spec struct {
    // exported
    Maker string

    // unexported
    model string

    // exported
    Price int
}
