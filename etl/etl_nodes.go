package etl

type JobNode[I interface{}, O interface{}] struct {
	Key     string
	Parents []string
	Read    func(ch chan O, args I)
	Write   func(ch chan O)
}

func (node JobNode[I, O]) Run(args I) {
	ch := make(chan O, 10)
	if nil != node.Read {
		go node.Read(ch, args)
	}
	if nil != node.Write {
		node.Write(ch)
	}
}
