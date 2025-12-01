package scheduler

type Scheduler interface { // polymorphism?
	SelectCandidateNodes()
	Score()
	Pick()
}
