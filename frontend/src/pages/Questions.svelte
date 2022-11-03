<script lang="ts">
	import Navbar from 'src/lib/components/Navbar.svelte'
	import QuestionSection from 'src/lib/components/QuestionSection.svelte'

	import type { Question } from 'src/types/question'

	let questions: Question = {
		topics: [
			{
				title: 'Web Development',
				cards: [
					{
						id: 1,
						score: 100,
						opened: true,
					},
					{
						id: 2,
						score: 200,
						opened: true,
					},
					{
						id: 3,
						score: 300,
						opened: true,
					},
					{
						id: 4,
						score: 400,
						opened: true,
					},
				],
			},
			{
				title: 'Cloud & Information Technology',
				cards: [
					{
						id: 5,
						score: 100,
						opened: true,
					},
					{
						id: 6,
						score: 200,
						opened: true,
					},
					{
						id: 7,
						score: 300,
						opened: true,
					},
					{
						id: 8,
						score: 400,
						opened: true,
					},
				],
			},
			{
				title: 'Artificial Intelligence & IoT',
				cards: [
					{
						id: 9,
						score: 100,
						opened: true,
					},
					{
						id: 10,
						score: 200,
						opened: true,
					},
					{
						id: 11,
						score: 300,
						opened: true,
					},
					{
						id: 12,
						score: 400,
						opened: true,
					},
				],
			},
			{
				title: 'Data Structure & Algorithm',
				cards: [
					{
						id: 13,
						score: 100,
						opened: true,
					},
					{
						id: 14,
						score: 200,
						opened: true,
					},
					{
						id: 15,
						score: 300,
						opened: true,
					},
					{
						id: 16,
						score: 400,
						opened: true,
					},
				],
			},
		],
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
			<QuestionSection {question} index={i} {openCard} />
		{/each}
	</div>
</main>
