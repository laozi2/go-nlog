package log

import (
	"fmt"
	"net"
)

type UdpWriter struct {
	LocalADDR  string
	RemoteADDR string
	conn       net.Conn
}

func NewUdpWriter(LocalADDR string, RemoteADDR string) *UdpWriter {
	updWriter := new(UdpWriter)
	updWriter.LocalADDR = LocalADDR
	updWriter.RemoteADDR = RemoteADDR
	updWriter.initConn()

	return updWriter
}

// Write implements io.Writer
func (w *UdpWriter) Write(p []byte) (n int, err error) {

	if w.conn == nil {
		return
	}

	return w.conn.Write(p)
}

func (w *UdpWriter) initConn() {

	localAddr, err := net.ResolveUDPAddr("udp4", w.LocalADDR)
	if err != nil {
		fmt.Printf("initConn Resolve local addr %s err: %v\n", w.LocalADDR, err)
		return
	}
	remoteAddr, err := net.ResolveUDPAddr("udp4", w.RemoteADDR)
	if err != nil {
		fmt.Printf("initConn Resolve remote addr %s err: %v\n", w.RemoteADDR, err)
		return
	}

	udpconn, err2 := net.DialUDP("udp", localAddr, remoteAddr)
	if err2 != nil {
		fmt.Printf("initConn DialUDP err: %v\n", err)
		return
	}
	//	defer udpconn.Close()
	w.conn = udpconn
}

func (w *UdpWriter) Close() error {
	if w.conn != nil {
		//	fmt.Printf("UdpWriter Close()\n")
		w.conn.Close()
	}
	return nil
}
