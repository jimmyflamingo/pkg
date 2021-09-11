package tbdiags

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Error-69]
	_ = x[Warning-87]
}

const (
	_Severity_name_0 = "Error"
	_Severity_name_1 = "Warning"
)

func (i Severity) String() string {
	switch {
	case i == 69:
		return _Severity_name_0
	case i == 87:
		return _Severity_name_1
	default:
		return "Severity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
