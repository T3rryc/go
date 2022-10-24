package io

type SlowWriter struct {
	// speed is the amount of bytes this writer can write on each Write() call
	speed int
	currentData []byte
}

func NewSlowWriter(speed int) *SlowWriter {
	return &SlowWriter{
		speed: speed,
	}
}

func (w *SlowWriter) Write(b []byte) (int, error){
	w.currentData = append(w.currentData, b[O:w.speed]...)
	return w.speed,nil
}
