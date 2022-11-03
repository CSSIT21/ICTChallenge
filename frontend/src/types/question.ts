export interface Question {
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
