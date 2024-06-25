package types

type Value interface {
    Stringify() string
}

type String struct {
    Value string
}

func (s String) Stringify() string {
    return "\"" + s.Value + "\""
}

type Number struct {
    Value string
}

func (n Number) Stringify() string {
    return n.Value
}

type Bool struct {
    Value bool
}

func (b Bool) Stringify() string {
    if (b.Value) {
        return "true"
    }
    return "false"
}

type Object struct {
    Fields map[string]Value
}

type Null struct {
}

func (n Null) Stringify() string {
    return ""
}

func (o Object) Stringify() string {
    stringified := "{"
    fields := []string{}
    for k, v := range o.Fields {
        str := "\""
        str += k
        str += "\""
        str += ":"
        str += v.Stringify()
        fields = append(fields, str)
    }

    for i, v := range fields {
            stringified += v
        if !(i == len(fields) - 1) {
            stringified += ","
        }
    }

    return stringified + "}"
}
