<script lang="ts">
	import { type Question, type OpenQuestion, Mode } from 'src/types/question'
	import { selected } from 'src/store/system'

	import QuestionCardScreen from './QuestionCardScreen.svelte'

	let questions: Question = {
		mode: Mode.TOPIC,
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
			{
				title: 'Infrastructure & Database',
				cards: [
					{
						id: 17,
						score: 100,
						opened: false,
					},
					{
						id: 18,
						score: 200,
						opened: false,
					},
					{
						id: 19,
						score: 300,
						opened: false,
					},
					{
						id: 20,
						score: 400,
						opened: false,
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
	{#if $selected >= 0}
		<QuestionCardScreen {openCard} question={questions.topics[$selected]} />
	{:else}
		<div class="h-[55px] px-10 mt-10 mb-12 space-y-1.5">
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
					$selected = i
					document.body.scrollIntoView()
				}}
			>
				<div
					class=" w-[310px] h-[105px] mb-8 title-box rounded-[16px] flex flex-col justify-center"
				>
					<h1
						class=" p-8 text-xl font-semibold text-center topic-color-{i +
							1}"
					>
						{question.title}
					</h1>
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
	.topic-color-5 {
		background: -webkit-linear-gradient(0deg, #0fbc88, #aff4ee);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
</style>
