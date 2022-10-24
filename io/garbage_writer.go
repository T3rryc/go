package io

import "time"

type GarbageWriter struct {
	writtenData string
}

func NewGarbageWriter() *GarbageWriter{ 
	return &GarbageWriter
}

func (w *GarbageWriter) Write(b []byte)(int, error){
	for _,r := range b {
		w.writtenData
	}
}