<script lang="ts">
	import Navbar from 'src/lib/components/Navbar.svelte'
	import QuestionSection from 'src/lib/components/QuestionSection.svelte'

	import type { Question, OpenQuestion } from 'src/types/question'

	let questions: Question = {
		topics: [
			{
				title: 'Web Development',
				cards: [
					{
						id: 1,
						score: 100,
						opened: false,
					},
					{
						id: 2,
						score: 200,
						opened: false,
					},
					{
						id: 3,
						score: 300,
						opened: false,
					},
					{
						id: 4,
						score: 400,
						opened: false,
					},
				],
			},
			{
				title: 'Cloud & Information Technology',
				cards: [
					{
						id: 5,
						score: 100,
						opened: false,
					},
					{
						id: 6,
						score: 200,
						opened: false,
					},
					{
						id: 7,
						score: 300,
						opened: false,
					},
					{
						id: 8,
						score: 400,
						opened: false,
					},
				],
			},
			{
				title: 'Artificial Intelligence & IoT',
				cards: [
					{
						id: 9,
						score: 100,
						opened: false,
					},
					{
						id: 10,
						score: 200,
						opened: false,
					},
					{
						id: 11,
						score: 300,
						opened: false,
					},
					{
						id: 12,
						score: 400,
						opened: false,
					},
				],
			},
			{
				title: 'Data Structure & Algorithm',
				cards: [
					{
						id: 13,
						score: 100,
						opened: false,
					},
					{
						id: 14,
						score: 200,
						opened: false,
					},
					{
						id: 15,
						score: 300,
						opened: false,
					},
					{
						id: 16,
						score: 400,
						opened: false,
					},
				],
			},
		],
	}

	let openQuestion: OpenQuestion = {
		question_id: 2,
		question:
			'Who do you think he/she is my least favorite professor in SIT@CS?',
	}

	function getQuestion(id: number) {
		console.log(id)
	}

	function openCard(cardCol: number, cardIndex: number) {
		questions.topics[cardCol].cards[cardIndex].opened =
			!questions.topics[cardCol].cards[cardIndex].opened
	}
</script>

<main class="bg-[#1B2D51] h-screen w-screen px-24 py-12">
	<Navbar />
	<div class="flex justify-between px-12 mt-3">
		{#each questions.topics as question, i (i)}
			<QuestionSection
				{question}
				index={i}
				{openCard}
				{getQuestion}
				openQuestion={openQuestion.question}
			/>
		{/each}
	</div>
</main>
