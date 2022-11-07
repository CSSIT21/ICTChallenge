<script lang="ts">
	import { onMount } from 'svelte'
	import type { Topic } from 'src/types/question'

	import QuestionCard from '../../lib/components/QuestionCard.svelte'
	import FlippedCard from './FlippedCard.svelte'

	import iconsTopic1 from '../../assets/images/icons-topic1.png'
	import iconsTopic2 from '../../assets/images/icons-topic2.png'
	import iconsTopic3 from '../../assets/images/icons-topic3.png'
	import iconsTopic4 from '../../assets/images/icons-topic4.png'

	export let question: Topic
	export let selectedId: number

	let icon: string = ''
	const widthCard: string = '224px'
	const heightCard: string = '124px'
	const widthImg: string = '60.3px'
	const textSize: string = '27px'
	const iconSize: string = '80px'

	onMount(() => {
		if (selectedId == 0) {
			icon = iconsTopic1
		} else if (selectedId == 1) {
			icon = iconsTopic2
		} else if (selectedId == 2) {
			icon = iconsTopic3
		} else if (selectedId == 3) {
			icon = iconsTopic4
		}
	})
</script>

<div class="mt-11">
	<div class="h-[592px] w-[390px] flex flex-col justify-between items-center">
		{#each question.cards as card, i (i)}
			{#if card.opened == true}
				<QuestionCard
					img={icon}
					score={card.score}
					textColor={'text-color-' + (selectedId + 1)}
					cardIndex={i}
					{widthCard}
					{heightCard}
					{widthImg}
					{textSize}
					cardId={card.id}
				/>
			{:else}
				<FlippedCard {widthCard} {heightCard} {iconSize} />
			{/if}
		{/each}
	</div>
</div>
