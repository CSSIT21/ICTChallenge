<script lang="ts">
	import { onMount } from 'svelte'

	import iconsTopic1 from '../../assets/images/icons-topic1.png'
	import iconsTopic2 from '../../assets/images/icons-topic2.png'
	import iconsTopic3 from '../../assets/images/icons-topic3.png'
	import iconsTopic4 from '../../assets/images/icons-topic4.png'

	import QuestionCard from 'src/lib/components/QuestionCard.svelte'
	import type { Topic } from 'src/types/question'
	import FlippedCard from './FlippedCard.svelte'

	export let question: Topic
	export let index: number
	let icon: string = ''
	let topicColor: string = ''
	export let openCard: (cardCol: number, cardIndex: number) => void

	onMount(() => {
		if (index == 0) {
			icon = iconsTopic1
			topicColor = 'topicColor1'
		} else if (index == 1) {
			icon = iconsTopic2
			topicColor = 'topicColor2'
		} else if (index == 2) {
			icon = iconsTopic3
			topicColor = 'topicColor3'
		} else if (index == 3) {
			icon = iconsTopic4
			topicColor = 'topicColor4'
		}
	})
</script>

<div class="w-[324px] flex flex-col items-center">
	<div class="flex flex-col justify-center h-[150px]">
		<p class="py-2 text-4xl font-semibold text-center {topicColor}">
			{question.title}
		</p>
	</div>
	<div class="h-[772px] flex flex-col justify-between items-center">
		{#each question.cards as card, i (i)}
			{#if card.opened == true}
				<QuestionCard
					img={icon}
					score={card.score}
					textColor={'textColor' + (index + 1)}
					{openCard}
					cardCol={index}
					cardIndex={i}
				/>
			{:else}
				<FlippedCard />
			{/if}
		{/each}
	</div>
</div>

<style>
	.topicColor1 {
		background: -webkit-linear-gradient(0deg, #48a3fa, #5ad2f3);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topicColor2 {
		background: -webkit-linear-gradient(0deg, #ac77e7, #ead7fe);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topicColor3 {
		background: -webkit-linear-gradient(0deg, #ffb34c, #ffe155);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.topicColor4 {
		background: -webkit-linear-gradient(0deg, #fc6cb9, #ffc7e5);
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
</style>
