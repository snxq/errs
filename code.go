package errs

// Code define err code.
type Code int

// Int32 returns int32.
func (e Code) Int32() int32 { return int32(e) }

// Int64 returns
func (e Code) Int64() int64 { return int64(e) }
