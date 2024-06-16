package data

// Permissions hold different permissions for a user (movies:read, movies:write).
type Permissions []string

// Include check whether the Permissions slice contains a specific
// permission code.
func (p Permissions) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}
