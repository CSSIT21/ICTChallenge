export interface Question {
	mode: Mode
	topics: Topic[]
}

export interface Topic {
	title: string
	cards: Card[]
}

export interface Card {
	id: number
	score: number
	opened: boolean
}

export interface OpenQuestion {
	question_id: number
	question: string
	bonus: boolean
}

export enum Mode {
	RULE = 'rule',
	TOPIC = 'topic',
}
