package ssession

import "smessage"

type IEventQueue interface {

}

type SEventQueue struct {

}

func (self *SEventQueue) PostMessage(msg smessage.SMessage,session SSession){

}