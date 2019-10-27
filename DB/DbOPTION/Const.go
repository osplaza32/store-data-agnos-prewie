package DbOPTION
type MOTOR string


func (m MOTOR) isMotor() MOTOR {
	return m
}

// The interface which is exported
type Motor interface {
	isMotor() MOTOR
}

const (
	MYSQL   = MOTOR("mysql")
	PG  = MOTOR("postgres")
	SQLITE = MOTOR("sqlite")
	MSSQL = MOTOR("mssql")
	MONGO   = MOTOR("MONGO")
)



