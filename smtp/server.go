package smtp

//var registeredHandlers = stateHandlers{
//	StateInit: onInit,
//	StateHelo: onHelo,
//	StateMail: onMail,
//	StateData: onData,
//}

//type Server struct {
//	Addr    string
//	Handler Handler
//}
//
//func (s *Server) ListenAndServe() error {
//	addr := s.Addr
//	ln, err := net.Listen("tcp", addr)
//	if err != nil {
//		return err
//	}
//	s.Serve(ln.(*net.TCPListener))
//	return nil
//}
//
//func (s *Server) Serve(l net.Listener) {
//	var err error
//	var rw net.Conn
//	log.Printf("SMTP Server listening on %s\n", s.Addr)
//	for {
//		rw, err = l.Accept()
//		if err != nil {
//			log.Println(fmt.Sprintf("Error accepting connection: %s", err))
//		}
//		//c := s.newConn(rw)
//		//go c.serve()
//	}
//}
//
//type conn struct {
//	server *Server
//	rwc    net.Conn
//}
//
//func (s *Server) newConn(rwc net.Conn) *conn {
//	return &conn{
//		server: s,
//		rwc:    rwc,
//	}
//}

//func (conn *conn) serve() {
//	defer func () {
//		err := conn.rwc.Close()
//		if err != nil {
//			log.Println(fmt.Sprintf("Error closing connection %s", err))
//		}
//	}()
//
//	ssm := NewSmtpStateMachine()
//	r, wr := conn.rwc, conn.rwc
//
//	var next State
//	var tests State
//	for !ssm.IsFinished() {
//		tests = ssm.State()
//		_, next = registeredHandlers[tests](r, wr)
//		if ssm.Cannot(next) {
//			log.Fatalf("Cannot move to state %s from %s", string(next), ssm.State())
//		}
//
//		ssm.S.SetState(string(next))
//	}
//	_ = onFini(r, wr)
//
//	go conn.server.Handler.HandleMessage(stubMessage("message"))
//}
//
//type stateHandler func (reader io.Reader, writer io.Writer) (error, State)
//
//type stateHandlers map[State]stateHandler

//func onInit(reader io.Reader, writer io.Writer) (error, State) {
//	err := writeResponse(writer, StateInit)
//	if err != nil {
//		return err, ""
//	}
//
//	bufReader := bufio.NewReader(reader)
//	input, err := bufReader.ReadString('\n')
//	if err != nil {
//		return err, ""
//	}
//
//	log.Println(string(input))
//
//	return nil, StateHelo
//}
//
//func onHelo(reader io.Reader, writer io.Writer) (error, State) {
//	err := writeResponse(writer, StateHelo)
//	if err != nil {
//		return err, ""
//	}
//
//	bufReader := bufio.NewReader(reader)
//	input, err := bufReader.ReadString('\n')
//	if err != nil {
//		return err, ""
//	}
//
//	log.Println(string(input))
//
//	return nil, StateMail
//}
//
//func onMail(reader io.Reader, writer io.Writer) (error, State) {
//	err := writeResponse(writer, StateMail)
//	if err != nil {
//		return err, ""
//	}
//
//	bufReader := bufio.NewReader(reader)
//	input, err := bufReader.ReadString('\n')
//	if err != nil {
//		return err, ""
//	}
//
//	log.Println(string(input))
//
//	return nil, StateData
//}
//
//func onData(reader io.Reader, writer io.Writer) (error, State) {
//	err := writeResponse(writer, StateData)
//	if err != nil {
//		return err, ""
//	}
//
//	bufReader := bufio.NewReader(reader)
//	input, err := bufReader.ReadString('\n')
//	if err != nil {
//		return err, ""
//	}
//
//	log.Println(string(input))
//
//	return nil, StateComp
//}

//func onFini(reader io.Reader, writer io.Writer) error {
//	err := writeResponse(writer, StateComp)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//func writeResponse(writer io.Writer, state State) error {
//	var err error
//
//	protocolText := stateResponses[state]
//
//	_, err = writer.Write((protocolText).bytes())
//	if err != nil {
//		return err
//	}
//
//	_, err = writer.Write([]byte("\n"))
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
