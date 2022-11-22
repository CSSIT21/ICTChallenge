<script lang="ts">
	import { type Question, Mode } from 'src/types/question'
	import { selected } from 'src/store/system'
	import { ArtWS } from 'src/store/websocket'
	import { onDestroy } from 'svelte'

	import QuestionCardScreen from './QuestionCardScreen.svelte'

	let questions: Question = { mode: Mode.TOPIC, topics: [] }

	const client = ArtWS.connect(
		'ws://ictc-int.sit.kmutt.ac.th:3000/ws/student?token=XrO3ole8bS83OQ3p',
		{
			log: true, // Log for console.warning
			reconnect: true, // Reconnect on close
			reconnectInterval: 1000, // Reconnect interval in ms
			name: 'client', // Name for the client
		}
	)

	const unsubscribeclient = client.subscribe('st/turn', (payload) => {
		questions = { mode: Mode.TOPIC, topics: payload.topics }
	})

	onDestroy(() => {
		unsubscribeclient()
	})
</script>

<div class="flex justify-center items-center flex-col">
	{#if $selected >= 0}
		<QuestionCardScreen question={questions.topics[$selected]} />
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
