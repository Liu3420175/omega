package auth



func ForceBytes(s interface{}) []byte{
	/*
	    force to decode bytes
	 */
	switch s.(type) {
	case []byte:
		return s.([]byte)
	case byte:
		return []byte{s.(byte)}
	case string:
		return []byte(s.(string))
	default:
		return []byte{}
	}
}