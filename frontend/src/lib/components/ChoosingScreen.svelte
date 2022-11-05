<script lang="ts">
	import type { Question } from 'src/types/question'

	import QuestionCardScreen from './QuestionCardScreen.svelte'

	export let handleSelected: (selected: boolean) => void
	export let selected: boolean

	let selectedId: number = 1
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

<div class="flex justify-center items-center flex-col">
	{#if selected == true}
		<QuestionCardScreen
			question={questions.topics[selectedId]}
			{selectedId}
			{openCard}
		/>
	{:else}
		<div class="h-[55px] px-11 mt-12 mb-8 space-y-1.5">
			<h1 class="font-semibold text-lg text-[#4279E8]">
				Question Categories
			</h1>
			<p class="text-xs text-white">
				Tap on the category that you want to see the questions
			</p>
		</div>

		{#each questions.topics as question, i (i)}
			<button
				on:click={() => {
					selectedId = i
					console.log(selectedId)

					handleSelected(selected)
				}}
			>
				<div
					class=" w-[310px] h-[105px] mb-8 title-box rounded-[16px] flex flex-col justify-center"
				>
					<p
						class=" p-2 text-xl font-semibold text-center topic-color-{i +
							1}"
					>
						{question.title}
					</p>
				</div>
			</button>
		{/each}
	{/if}
</div>

<style>
	.title-box {
		background-color: rgba(0, 0, 0, 0.15);
	}
	.topic-color-1 {
		background: -webkit-linear-gradient(0deg, #48a3fa, #5ad2f3);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topic-color-2 {
		background: -webkit-linear-gradient(0deg, #ac77e7, #ead7fe);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topic-color-3 {
		background: -webkit-linear-gradient(0deg, #ffb34c, #ffe155);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topic-color-4 {
		background: -webkit-linear-gradient(0deg, #fc6cb9, #ffc7e5);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
</style>
