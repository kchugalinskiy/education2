package strstrstr

func Substring(s, sub string) string {
	if len(s) < len(sub) {
		return ""
	}
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return s[i:]
		}
	}
	return ""
}
